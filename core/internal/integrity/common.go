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

package integrity

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
)

// Agent specifices the interface which every integrity agent needs to fulfill
// TODO: This interface is not yet used, we call the TPM2 agent directly.
type Agent interface {
	// Initialize needs to be called once and initializes the systems required to maintain integrity
	// on the given platform.
	// nodeCert is a X.509 DER certificate which identifies the node once it's unlocked. This is
	// required to bind the node certificate (which is only available when the node is unlocked) to
	// the integrity subsystem used to attest said node.
	// Initialize returns the cryptographic identity that it's bound to.
	Initialize(newNode api.NewNodeInfo, enrolment api.EnrolmentConfig) (string, error)

	// Unlock performs all required actions to assure the integrity of the platform and securely retrieves
	// the unlock key.
	Unlock(enrolment api.EnrolmentConfig) ([]byte, error)
}

// DialNMS creates a secure GRPC connection to the NodeManagementService
func DialNMS(enrolment api.EnrolmentConfig) (*grpc.ClientConn, error) {
	var targets []string
	for _, target := range enrolment.MasterIps {
		targets = append(targets, fmt.Sprintf("%v:%v", net.IP(target), common.MasterServicePort))
	}
	cert, err := x509.ParseCertificate(enrolment.MastersCert)
	if err != nil {
		return nil, err
	}
	mastersPool := x509.NewCertPool()
	mastersPool.AddCert(cert)

	secureTransport := &tls.Config{
		InsecureSkipVerify: true,
		// Critical function, please review any changes with care
		// TODO(lorenz): Actively check that this actually provides the security guarantees that we need
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			for _, cert := range rawCerts {
				// X.509 certificates in DER can be compared like this since DER has a unique representation
				// for each certificate.
				if bytes.Equal(cert, enrolment.MastersCert) {
					return nil
				}
			}
			return errors.New("failed to find authorized NMS certificate")
		},
		MinVersion: tls.VersionTLS13,
	}
	secureTransportCreds := credentials.NewTLS(secureTransport)

	return grpc.Dial(strings.Join(targets, ","), grpc.WithTransportCredentials(secureTransportCreds))
}
