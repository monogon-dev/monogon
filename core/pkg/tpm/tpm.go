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
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.monogon.dev/source/nexantic.git/core/pkg/logtree"

	"git.monogon.dev/source/nexantic.git/core/pkg/sysfs"

	"github.com/gogo/protobuf/proto"
	tpmpb "github.com/google/go-tpm-tools/proto"
	"github.com/google/go-tpm-tools/tpm2tools"
	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

var (
	// SecureBootPCRs are all PCRs that measure the current Secure Boot configuration.
	// This is what we want if we rely on secure boot to verify boot integrity. The firmware
	// hashes the secure boot policy and custom keys into the PCR.
	//
	// This requires an extra step that provisions the custom keys.
	//
	// Some background: https://mjg59.dreamwidth.org/48897.html?thread=1847297
	// (the initramfs issue mentioned in the article has been solved by integrating
	// it into the kernel binary, and we don't have a shim bootloader)
	//
	// PCR7 alone is not sufficient - it needs to be combined with firmware measurements.
	SecureBootPCRs = []int{7}

	// FirmwarePCRs are alle PCRs that contain the firmware measurements
	// See https://trustedcomputinggroup.org/wp-content/uploads/TCG_EFI_Platform_1_22_Final_-v15.pdf
	FirmwarePCRs = []int{
		0, // platform firmware
		2, // option ROM code
		3, // option ROM configuration and data
	}

	// FullSystemPCRs are all PCRs that contain any measurements up to the currently running EFI payload.
	FullSystemPCRs = []int{
		0, // platform firmware
		1, // host platform configuration
		2, // option ROM code
		3, // option ROM configuration and data
		4, // EFI payload
	}

	// Using FullSystemPCRs is the most secure, but also the most brittle option since updating the EFI
	// binary, updating the platform firmware, changing platform settings or updating the binary
	// would invalidate the sealed data. It's annoying (but possible) to predict values for PCR4,
	// and even more annoying for the firmware PCR (comparison to known values on similar hardware
	// is the only thing that comes to mind).
	//
	// See also: https://github.com/mxre/sealkey (generates PCR4 from EFI image, BSD license)
	//
	// Using only SecureBootPCRs is the easiest and still reasonably secure, if we assume that the
	// platform knows how to take care of itself (i.e. Intel Boot Guard), and that secure boot
	// is implemented properly. It is, however, a much larger amount of code we need to trust.
	//
	// We do not care about PCR 5 (GPT partition table) since modifying it is harmless. All of
	// the boot options and cmdline are hardcoded in the kernel image, and we use no bootloader,
	// so there's no PCR for bootloader configuration or kernel cmdline.
)

var (
	numSRTMPCRs = 16
	srtmPCRs    = tpm2.PCRSelection{Hash: tpm2.AlgSHA256, PCRs: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}
	// TCG Trusted Platform Module Library Level 00 Revision 0.99 Table 6
	tpmGeneratedValue = uint32(0xff544347)
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
	logger logtree.LeveledLogger
	device io.ReadWriteCloser

	// We keep the AK loaded since it's used fairly often and deriving it is expensive
	akHandleCache tpmutil.Handle
	akPublicKey   crypto.PublicKey
}

// Initialize finds and opens the TPM (if any). If there is no TPM available it returns
// ErrNotExists
func Initialize(logger logtree.LeveledLogger) error {
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
		// If this is changed GetMeasurementLog() needs to be updated too
		logger.Warningf("Found more than one TPM, using the first one")
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
	pcrList := []string{}
	for _, pcr := range sealedKey.Pcrs {
		pcrList = append(pcrList, string(pcr))
	}
	tpm.logger.Infof("Attempting to unseal data protected with PCRs %s", strings.Join(pcrList, ","))
	unsealedData, err := srk.Unseal(&sealedKey)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to unseal data")
	}
	return unsealedData, nil
}

// Standard AK template for RSA2048 non-duplicatable restricted signing for attestation
var akTemplate = tpm2.Public{
	Type:       tpm2.AlgRSA,
	NameAlg:    tpm2.AlgSHA256,
	Attributes: tpm2.FlagSignerDefault,
	RSAParameters: &tpm2.RSAParams{
		Sign: &tpm2.SigScheme{
			Alg:  tpm2.AlgRSASSA,
			Hash: tpm2.AlgSHA256,
		},
		KeyBits: 2048,
	},
}

func loadAK() error {
	var err error
	// Rationale: The AK is an EK-equivalent key and used only for attestation. Using a non-primary
	// key here would require us to store the wrapped version somewhere, which is inconvenient.
	// This being a primary key in the Endorsement hierarchy means that it can always be recreated
	// and can never be "destroyed". Under our security model this is of no concern since we identify
	// a node by its IK (Identity Key) which we can destroy.
	tpm.akHandleCache, tpm.akPublicKey, err = tpm2.CreatePrimary(tpm.device, tpm2.HandleEndorsement,
		tpm2.PCRSelection{}, "", "", akTemplate)
	return err
}

// Process documented in TCG EK Credential Profile 2.2.1
func loadEK() (tpmutil.Handle, crypto.PublicKey, error) {
	// The EK is a primary key which is supposed to be certified by the manufacturer of the TPM.
	// Its public attributes are standardized in TCG EK Credential Profile 2.0 Table 1. These need
	// to match exactly or we aren't getting the key the manufacturere signed. tpm2tools contains
	// such a template already, so we're using that instead of redoing it ourselves.
	// This ignores the more complicated ways EKs can be specified, the additional stuff you can do
	// is just absolutely crazy (see 2.2.1.2 onward)
	return tpm2.CreatePrimary(tpm.device, tpm2.HandleEndorsement,
		tpm2.PCRSelection{}, "", "", tpm2tools.DefaultEKTemplateRSA())
}

// GetAKPublic gets the TPM2T_PUBLIC of the AK key
func GetAKPublic() ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, ErrNotInitialized
	}
	if tpm.akHandleCache == tpmutil.Handle(0) {
		if err := loadAK(); err != nil {
			return []byte{}, fmt.Errorf("failed to load AK primary key: %w", err)
		}
	}
	public, _, _, err := tpm2.ReadPublic(tpm.device, tpm.akHandleCache)
	if err != nil {
		return []byte{}, err
	}
	return public.Encode()
}

// TCG TPM v2.0 Provisioning Guidance v1.0 7.8 Table 2 and
// TCG EK Credential Profile v2.1 2.2.1.4 de-facto Standard for Windows
// These are both non-normative and reference Windows 10 documentation that's no longer available :(
// But in practice this is what people are using, so if it's normative or not doesn't really matter
const ekCertHandle = 0x01c00002

// GetEKPublic gets the public key and (if available) Certificate of the EK
func GetEKPublic() ([]byte, []byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, []byte{}, ErrNotInitialized
	}
	ekHandle, publicRaw, err := loadEK()
	if err != nil {
		return []byte{}, []byte{}, fmt.Errorf("failed to load EK primary key: %w", err)
	}
	defer tpm2.FlushContext(tpm.device, ekHandle)
	// Don't question the use of HandleOwner, that's the Standard™
	ekCertRaw, err := tpm2.NVReadEx(tpm.device, ekCertHandle, tpm2.HandleOwner, "", 0)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	publicKey, err := x509.MarshalPKIXPublicKey(publicRaw)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	return publicKey, ekCertRaw, nil
}

// MakeAKChallenge generates a challenge for TPM residency and attributes of the AK
func MakeAKChallenge(ekPubKey, akPub []byte, nonce []byte) ([]byte, []byte, error) {
	ekPubKeyData, err := x509.ParsePKIXPublicKey(ekPubKey)
	if err != nil {
		return []byte{}, []byte{}, fmt.Errorf("failed to decode EK pubkey: %w", err)
	}
	akPubData, err := tpm2.DecodePublic(akPub)
	if err != nil {
		return []byte{}, []byte{}, fmt.Errorf("failed to decode AK public part: %w", err)
	}
	// Make sure we're attesting the right attributes (in particular Restricted)
	if !akPubData.MatchesTemplate(akTemplate) {
		return []byte{}, []byte{}, errors.New("the key being challenged is not a valid AK")
	}
	akName, err := akPubData.Name()
	if err != nil {
		return []byte{}, []byte{}, fmt.Errorf("failed to derive AK name: %w", err)
	}
	return generateRSA(akName.Digest, ekPubKeyData.(*rsa.PublicKey), 16, nonce, rand.Reader)
}

// SolveAKChallenge solves a challenge for TPM residency of the AK
func SolveAKChallenge(credBlob, secretChallenge []byte) ([]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, ErrNotInitialized
	}
	if tpm.akHandleCache == tpmutil.Handle(0) {
		if err := loadAK(); err != nil {
			return []byte{}, fmt.Errorf("failed to load AK primary key: %w", err)
		}
	}

	ekHandle, _, err := loadEK()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to load EK: %w", err)
	}
	defer tpm2.FlushContext(tpm.device, ekHandle)

	// This is necessary since the EK requires an endorsement handle policy in its session
	// For us this is stupid because we keep all hierarchies open anyways since a) we cannot safely
	// store secrets on the OS side pre-global unlock and b) it makes no sense in this security model
	// since an uncompromised host OS will not let an untrusted entity attest as itself and a
	// compromised OS can either not pass PCR policy checks or the game's already over (you
	// successfully runtime-exploited a production Smalltown Core)
	endorsementSession, _, err := tpm2.StartAuthSession(
		tpm.device,
		tpm2.HandleNull,
		tpm2.HandleNull,
		make([]byte, 16),
		nil,
		tpm2.SessionPolicy,
		tpm2.AlgNull,
		tpm2.AlgSHA256)
	if err != nil {
		panic(err)
	}
	defer tpm2.FlushContext(tpm.device, endorsementSession)

	_, err = tpm2.PolicySecret(tpm.device, tpm2.HandleEndorsement, tpm2.AuthCommand{Session: tpm2.HandlePasswordSession, Attributes: tpm2.AttrContinueSession}, endorsementSession, nil, nil, nil, 0)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to make a policy secret session: %w", err)
	}

	for {
		solution, err := tpm2.ActivateCredentialUsingAuth(tpm.device, []tpm2.AuthCommand{
			{Session: tpm2.HandlePasswordSession, Attributes: tpm2.AttrContinueSession}, // Use standard no-password authentication
			{Session: endorsementSession, Attributes: tpm2.AttrContinueSession},         // Use a full policy session for the EK
		}, tpm.akHandleCache, ekHandle, credBlob, secretChallenge)
		if warn, ok := err.(tpm2.Warning); ok && warn.Code == tpm2.RCRetry {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		return solution, err
	}
}

// FlushTransientHandles flushes all sessions and non-persistent handles
func FlushTransientHandles() error {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return ErrNotInitialized
	}
	flushHandleTypes := []tpm2.HandleType{tpm2.HandleTypeTransient, tpm2.HandleTypeLoadedSession, tpm2.HandleTypeSavedSession}
	for _, handleType := range flushHandleTypes {
		handles, err := tpm2tools.Handles(tpm.device, handleType)
		if err != nil {
			return err
		}
		for _, handle := range handles {
			if err := tpm2.FlushContext(tpm.device, handle); err != nil {
				return err
			}
		}
	}
	return nil
}

// AttestPlatform performs a PCR quote using the AK and returns the quote and its signature
func AttestPlatform(nonce []byte) ([]byte, []byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return []byte{}, []byte{}, ErrNotInitialized
	}
	if tpm.akHandleCache == tpmutil.Handle(0) {
		if err := loadAK(); err != nil {
			return []byte{}, []byte{}, fmt.Errorf("failed to load AK primary key: %w", err)
		}
	}
	// We only care about SHA256 since SHA1 is weak. This is supported on at least GCE and
	// Intel / AMD fTPM, which is good enough for now. Alg is null because that would just hash the
	// nonce, which is dumb.
	quote, signature, err := tpm2.Quote(tpm.device, tpm.akHandleCache, "", "", nonce, srtmPCRs,
		tpm2.AlgNull)
	if err != nil {
		return []byte{}, []byte{}, fmt.Errorf("failed to quote PCRs: %w", err)
	}
	return quote, signature.RSA.Signature, err
}

// VerifyAttestPlatform verifies a given attestation. You can rely on all data coming back as being
// from the TPM on which the AK is bound to.
func VerifyAttestPlatform(nonce, akPub, quote, signature []byte) (*tpm2.AttestationData, error) {
	hash := crypto.SHA256.New()
	hash.Write(quote)

	akPubData, err := tpm2.DecodePublic(akPub)
	if err != nil {
		return nil, fmt.Errorf("invalid AK: %w", err)
	}
	akPublicKey, err := akPubData.Key()
	if err != nil {
		return nil, fmt.Errorf("invalid AK: %w", err)
	}
	akRSAKey, ok := akPublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid AK: invalid key type")
	}

	if err := rsa.VerifyPKCS1v15(akRSAKey, crypto.SHA256, hash.Sum(nil), signature); err != nil {
		return nil, err
	}

	quoteData, err := tpm2.DecodeAttestationData(quote)
	if err != nil {
		return nil, err
	}
	// quoteData.Magic works together with the TPM's Restricted key attribute. If this attribute is set
	// (which it needs to be for the AK to be considered valid) the TPM will not sign external data
	// having this prefix with such a key. Only data that originates inside the TPM like quotes and
	// key certifications can have this prefix and sill be signed by a restricted key. This check
	// is thus vital, otherwise somebody can just feed the TPM an arbitrary attestation to sign with
	// its AK and this function will happily accept the forged attestation.
	if quoteData.Magic != tpmGeneratedValue {
		return nil, errors.New("invalid TPM quote: data marker for internal data not set - forged attestation")
	}
	if quoteData.Type != tpm2.TagAttestQuote {
		return nil, errors.New("invalid TPM qoute: not a TPM quote")
	}
	if !bytes.Equal(quoteData.ExtraData, nonce) {
		return nil, errors.New("invalid TPM quote: wrong nonce")
	}

	return quoteData, nil
}

// GetPCRs returns all SRTM PCRs in-order
func GetPCRs() ([][]byte, error) {
	lock.Lock()
	defer lock.Unlock()
	if tpm == nil {
		return [][]byte{}, ErrNotInitialized
	}
	pcrs := make([][]byte, numSRTMPCRs)

	// The TPM can (and most do) return partial results. Let's just retry as many times as we have
	// PCRs since each read should return at least one PCR.
readLoop:
	for i := 0; i < numSRTMPCRs; i++ {
		sel := tpm2.PCRSelection{Hash: tpm2.AlgSHA256}
		for pcrN := 0; pcrN < numSRTMPCRs; pcrN++ {
			if len(pcrs[pcrN]) == 0 {
				sel.PCRs = append(sel.PCRs, pcrN)
			}
		}

		readPCRs, err := tpm2.ReadPCRs(tpm.device, sel)
		if err != nil {
			return nil, fmt.Errorf("failed to read PCRs: %w", err)
		}

		for pcrN, pcr := range readPCRs {
			pcrs[pcrN] = pcr
		}
		for _, pcr := range pcrs {
			// If at least one PCR is still not read, continue
			if len(pcr) == 0 {
				continue readLoop
			}
		}
		break
	}

	return pcrs, nil
}

// GetMeasurmentLog returns the binary log of all data hashed into PCRs. The result can be parsed by eventlog.
// As this library currently doesn't support extending PCRs it just returns the log as supplied by the EFI interface.
func GetMeasurementLog() ([]byte, error) {
	return ioutil.ReadFile("/sys/kernel/security/tpm0/binary_bios_measurements")
}
