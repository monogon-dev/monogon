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

package tpm2

import (
	"context"
	"errors"
	"fmt"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/integrity"
	"git.monogon.dev/source/nexantic.git/core/pkg/tpm"
)

type TPM2Agent struct {
}

func (a *TPM2Agent) Initialize(newNode api.NewNodeInfo, enrolment api.EnrolmentConfig) error {
	nmsConn, err := integrity.DialNMS(enrolment)
	nmsClient := api.NewNodeManagementServiceClient(nmsConn)
	ekPub, ekCert, err := tpm.GetEKPublic()
	if err != nil {
		return fmt.Errorf("failed to generate EK: %w", err)
	}

	akPub, err := tpm.GetAKPublic()
	if err != nil {
		return fmt.Errorf("failed to generate AK: %w", err)
	}

	registerSession, err := nmsClient.NewTPM2NodeRegister(context.Background())
	if err != nil {
		return fmt.Errorf("failed to open registration session: %w", err)
	}
	defer registerSession.CloseSend()
	if err := registerSession.Send(&api.TPM2FlowRequest{
		Stage: &api.TPM2FlowRequest_Register{
			Register: &api.TPM2RegisterRequest{
				AkPublic: akPub,
				EkPubkey: ekPub,
				EkCert:   ekCert,
			},
		},
	}); err != nil {
		return fmt.Errorf("failed to send registration: %w", err)
	}

	res1, err := registerSession.Recv()
	if err != nil {
		return fmt.Errorf("failed to receive attest request: %w", err)
	}
	attestReqContainer, ok := res1.Stage.(*api.TPM2FlowResponse_AttestRequest)
	if !ok {
		return errors.New("protocol violation: after RegisterRequest expected AttestRequest")
	}
	attestReq := attestReqContainer.AttestRequest
	solution, err := tpm.SolveAKChallenge(attestReq.AkChallenge, attestReq.AkChallengeSecret)
	if err != nil {
		return fmt.Errorf("failed to solve AK challenge: %w", err)
	}
	pcrs, err := tpm.GetPCRs()
	if err != nil {
		return fmt.Errorf("failed to get SRTM PCRs: %w", err)
	}
	quote, quoteSig, err := tpm.AttestPlatform(attestReq.QuoteNonce)
	if err != nil {
		return fmt.Errorf("failed Quote operation: %w", err)
	}
	if err := registerSession.Send(&api.TPM2FlowRequest{
		Stage: &api.TPM2FlowRequest_AttestResponse{
			AttestResponse: &api.TPM2AttestResponse{
				AkChallengeSolution: solution,
				Pcrs:                pcrs,
				Quote:               quote,
				QuoteSignature:      quoteSig,
			},
		},
	}); err != nil {
		return fmt.Errorf("failed to send AttestResponse: %w", err)
	}
	if err := registerSession.Send(&api.TPM2FlowRequest{
		Stage: &api.TPM2FlowRequest_NewNodeInfo{
			NewNodeInfo: &newNode,
		},
	}); err != nil {
		return fmt.Errorf("failed to send NewNodeInfo: %w", err)
	}
	return nil
}

// Unlock attests the node state to the remote NMS and asks it for the global unlock key
func (a *TPM2Agent) Unlock(enrolment api.EnrolmentConfig) ([]byte, error) {
	nmsConn, err := integrity.DialNMS(enrolment)
	if err != nil {
		return []byte{}, err
	}
	nmsClient := api.NewNodeManagementServiceClient(nmsConn)
	unlockClient, err := nmsClient.TPM2Unlock(context.Background())
	if err != nil {
		return []byte{}, err
	}
	defer unlockClient.CloseSend()
	unlockInitContainer, err := unlockClient.Recv()
	if err != nil {
		return []byte{}, err
	}
	unlockInitVariant, ok := unlockInitContainer.Stage.(*api.TPM2UnlockFlowResponse_UnlockInit)
	if !ok {
		return []byte{}, errors.New("TPM2Unlock protocol violation")
	}
	unlockInit := unlockInitVariant.UnlockInit
	quote, sig, err := tpm.AttestPlatform(unlockInit.Nonce)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to attest platform: %w", err)
	}
	pcrs, err := tpm.GetPCRs()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get PCRs from TPM: %w", err)
	}
	if err := unlockClient.Send(&api.TPM2UnlockFlowRequeset{Stage: &api.TPM2UnlockFlowRequeset_UnlockRequest{
		UnlockRequest: &api.TPM2UnlockRequest{
			Pcrs:           pcrs,
			Quote:          quote,
			QuoteSignature: sig,
			NodeId:         enrolment.NodeId,
		},
	}}); err != nil {
		return []byte{}, err
	}
	unlockResponseContainer, err := unlockClient.Recv()
	if err != nil {
		return []byte{}, err
	}
	unlockResponseVariant, ok := unlockResponseContainer.Stage.(*api.TPM2UnlockFlowResponse_UnlockResponse)
	if !ok {
		return []byte{}, errors.New("violated TPM2Unlock protocol")
	}
	unlockResponse := unlockResponseVariant.UnlockResponse

	return unlockResponse.GlobalUnlockKey, nil
}
