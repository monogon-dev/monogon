// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tpm

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"smalltown/pkg/sysfs"
	"strconv"
	"sync"

	"github.com/gogo/protobuf/proto"
	tpmpb "github.com/google/go-tpm-tools/proto"
	"github.com/google/go-tpm-tools/tpm2tools"
	"github.com/google/go-tpm/tpm2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
)

var (
	// SecureBootPCRs are all PCRs that measure the current Secure Boot configuration
	SecureBootPCRs = []int{7}

	// FirmwarePCRs are alle PCRs that contain the firmware measurements
	// See https://trustedcomputinggroup.org/wp-content/uploads/TCG_EFI_Platform_1_22_Final_-v15.pdf
	FirmwarePCRs = []int{0, 2, 3}

	// FullSystemPCRs are all PCRs that contain any measurements up to the currently running EFI
	// payload.
	FullSystemPCRs = []int{0, 1, 2, 3, 4}
)

var (
	// ErrNotExists is returned when no TPMs are available in the system
	ErrNotExists = errors.New("no TPMs found")
	// ErrNotInitialized is returned when this package was not initialized successfully
	ErrNotInitialized = errors.New("no TPM was initialized")
)

// Singleton since the TPM is too
var tpm *TPM

// We're serializing all TPM operations since it has a limited number of handles and recovering
// if it runs out is difficult to implement correctly. Might also be marginally more secure.
var lock sync.Mutex

// TPM represents a high-level interface to a connected TPM 2.0
type TPM struct {
	logger *zap.Logger
	device io.ReadWriteCloser
}

// Initialize finds and opens the TPM (if any). If there is no TPM available it returns
// ErrNotExists
func Initialize(logger *zap.Logger) error {
	lock.Lock()
	defer lock.Unlock()
	tpmDir, err := os.Open("/sys/class/tpm")
	if err != nil {
		return errors.Wrap(err, "failed to open sysfs TPM class")
	}
	defer tpmDir.Close()

	tpms, err := tpmDir.Readdirnames(2)
	if err != nil {
		return errors.Wrap(err, "failed to read TPM device class")
	}

	if len(tpms) == 0 {
		return ErrNotExists
	}
	if len(tpms) > 1 {
		logger.Warn("Found more than one TPM, using the first one")
	}
	tpmName := tpms[0]
	ueventData, err := sysfs.ReadUevents(filepath.Join("/sys/class/tpm", tpmName, "uevent"))
	majorDev, err := strconv.Atoi(ueventData["MAJOR"])
	if err != nil {
		return fmt.Errorf("failed to convert uevent: %w", err)
	}
	minorDev, err := strconv.Atoi(ueventData["MINOR"])
	if err != nil {
		return fmt.Errorf("failed to convert uevent: %w", err)
	}
	if err := unix.Mknod("/dev/tpm", 0600|unix.S_IFCHR, int(unix.Mkdev(uint32(majorDev), uint32(minorDev)))); err != nil {
		return errors.Wrap(err, "failed to create TPM device node")
	}
	device, err := tpm2.OpenTPM("/dev/tpm")
	if err != nil {
		return errors.Wrap(err, "failed to open TPM")
	}
	tpm = &TPM{
		device: device,
		logger: logger,
	}
	return nil
}

// GenerateSafeKey uses two sources of randomness (Kernel & TPM) to generate the key
func GenerateSafeKey(size uint16) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, ErrNotInitialized
	}
	encryptionKeyHost := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, encryptionKeyHost); err != nil {
		return []byte{}, errors.Wrap(err, "failed to generate host portion of new key")
	}
	var encryptionKeyTPM []byte
	for i := 48; i > 0; i-- {
		tpmKeyPart, err := tpm2.GetRandom(tpm.device, size-uint16(len(encryptionKeyTPM)))
		if err != nil {
			return []byte{}, errors.Wrap(err, "failed to generate TPM portion of new key")
		}
		encryptionKeyTPM = append(encryptionKeyTPM, tpmKeyPart...)
		if len(encryptionKeyTPM) >= int(size) {
			break
		}
	}

	if len(encryptionKeyTPM) != int(size) {
		return []byte{}, fmt.Errorf("got incorrect amount of TPM randomess: %v, requested %v", len(encryptionKeyTPM), size)
	}

	encryptionKey := make([]byte, size)
	for i := uint16(0); i < size; i++ {
		encryptionKey[i] = encryptionKeyHost[i] ^ encryptionKeyTPM[i]
	}
	return encryptionKey, nil
}

// Seal seals sensitive data and only allows access if the current platform configuration in
// matches the one the data was sealed on.
func Seal(data []byte, pcrs []int) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, ErrNotInitialized
	}
	srk, err := tpm2tools.StorageRootKeyRSA(tpm.device)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to load TPM SRK")
	}
	defer srk.Close()
	sealedKey, err := srk.Seal(pcrs, data)
	sealedKeyRaw, err := proto.Marshal(sealedKey)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to marshal sealed data")
	}
	return sealedKeyRaw, nil
}

// Unseal unseals sensitive data if the current platform configuration allows and sealing constraints
// allow it.
func Unseal(data []byte) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, ErrNotInitialized
	}
	srk, err := tpm2tools.StorageRootKeyRSA(tpm.device)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to load TPM SRK")
	}
	defer srk.Close()

	var sealedKey tpmpb.SealedBytes
	if err := proto.Unmarshal(data, &sealedKey); err != nil {
		return []byte{}, errors.Wrap(err, "failed to decode sealed data")
	}
	// Logging this for auditing purposes
	tpm.logger.Info("Attempting to unseal data protected with PCRs", zap.Int32s("pcrs", sealedKey.Pcrs))
	unsealedData, err := srk.Unseal(&sealedKey)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to unseal data")
	}
	return unsealedData, nil
}
