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

// This file is adapted from github.com/google/go-tpm/tpm2/credactivation which outputs broken
// challenges for unknown reasons. They use u16 length-delimited outputs for the challenge blobs
// which is incorrect.
// TODO(lorenz): I'll eventually deal with this upstream, but for now just fix it here (it's not that)
// much code after all.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rsa"
	"fmt"
	"io"

	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

const (
	labelIdentity  = "IDENTITY"
	labelStorage   = "STORAGE"
	labelIntegrity = "INTEGRITY"
)

func generateRSA(aik *tpm2.HashValue, pub *rsa.PublicKey, symBlockSize int, secret []byte, rnd io.Reader) ([]byte, []byte, error) {
	newAIKHash, err := aik.Alg.HashConstructor()
	if err != nil {
		return nil, nil, err
	}

	// The seed length should match the keysize used by the EKs symmetric cipher.
	// For typical RSA EKs, this will be 128 bits (16 bytes).
	// Spec: TCG 2.0 EK Credential Profile revision 14, section 2.1.5.1.
	seed := make([]byte, symBlockSize)
	if _, err := io.ReadFull(rnd, seed); err != nil {
		return nil, nil, fmt.Errorf("generating seed: %v", err)
	}

	// Encrypt the seed value using the provided public key.
	// See annex B, section 10.4 of the TPM specification revision 2 part 1.
	label := append([]byte(labelIdentity), 0)
	encSecret, err := rsa.EncryptOAEP(newAIKHash(), rnd, pub, seed, label)
	if err != nil {
		return nil, nil, fmt.Errorf("generating encrypted seed: %v", err)
	}

	// Generate the encrypted credential by convolving the seed with the digest of
	// the AIK, and using the result as the key to encrypt the secret.
	// See section 24.4 of TPM 2.0 specification, part 1.
	aikNameEncoded, err := aik.Encode()
	if err != nil {
		return nil, nil, fmt.Errorf("encoding aikName: %v", err)
	}
	symmetricKey, err := tpm2.KDFa(aik.Alg, seed, labelStorage, aikNameEncoded, nil, len(seed)*8)
	if err != nil {
		return nil, nil, fmt.Errorf("generating symmetric key: %v", err)
	}
	c, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil, nil, fmt.Errorf("symmetric cipher setup: %v", err)
	}
	cv, err := tpmutil.Pack(tpmutil.U16Bytes(secret))
	if err != nil {
		return nil, nil, fmt.Errorf("generating cv (TPM2B_Digest): %v", err)
	}

	// IV is all null bytes. encIdentity represents the encrypted credential.
	encIdentity := make([]byte, len(cv))
	cipher.NewCFBEncrypter(c, make([]byte, len(symmetricKey))).XORKeyStream(encIdentity, cv)

	// Generate the integrity HMAC, which is used to protect the integrity of the
	// encrypted structure.
	// See section 24.5 of the TPM specification revision 2 part 1.
	macKey, err := tpm2.KDFa(aik.Alg, seed, labelIntegrity, nil, nil, newAIKHash().Size()*8)
	if err != nil {
		return nil, nil, fmt.Errorf("generating HMAC key: %v", err)
	}

	mac := hmac.New(newAIKHash, macKey)
	mac.Write(encIdentity)
	mac.Write(aikNameEncoded)
	integrityHMAC := mac.Sum(nil)

	idObject := &tpm2.IDObject{
		IntegrityHMAC: integrityHMAC,
		EncIdentity:   encIdentity,
	}
	id, err := tpmutil.Pack(idObject)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding IDObject: %v", err)
	}

	packedID, err := tpmutil.Pack(id)
	if err != nil {
		return nil, nil, fmt.Errorf("packing id: %v", err)
	}
	packedEncSecret, err := tpmutil.Pack(encSecret)
	if err != nil {
		return nil, nil, fmt.Errorf("packing encSecret: %v", err)
	}

	return packedID, packedEncSecret, nil
}
