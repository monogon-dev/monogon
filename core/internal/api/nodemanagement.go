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

package api

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/pkg/tpm"
	"github.com/gogo/protobuf/proto"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const nodesPrefix = "nodes/"
const enrolmentsPrefix = "enrolments/"

func nodeId(idCert []byte) (string, error) {
	// Currently we only identify nodes by ID key
	cert, err := x509.ParseCertificate(idCert)
	if err != nil {
		return "", err
	}
	pubKey, ok := cert.PublicKey.(ed25519.PublicKey)
	if !ok {
		return "", errors.New("invalid node identity certificate")
	}

	return "smalltown-" + base64.RawStdEncoding.EncodeToString([]byte(pubKey)), nil
}

func (s *Server) registerNewNode(node *api.Node) error {
	nodeRaw, err := proto.Marshal(node)
	if err != nil {
		return err
	}

	nodeID, err := nodeId(node.IdCert)
	if err != nil {
		return err
	}

	key := nodesPrefix + nodeID

	// Overwriting nodes is a BadIdea(TM), so make this a Compare-and-Swap
	res, err := s.getStore().Txn(context.Background()).If(
		clientv3.Compare(clientv3.CreateRevision(key), "=", 0),
	).Then(
		clientv3.OpPut(key, string(nodeRaw)),
	).Commit()
	if err != nil {
		return fmt.Errorf("failed to store new node: %w", err)
	}
	if !res.Succeeded {
		s.Logger.Warn("double-registration of node attempted", zap.String("node", nodeID))
	}
	return nil
}

func (s *Server) TPM2BootstrapNode(newNodeInfo *api.NewNodeInfo) (*api.Node, error) {
	akPublic, err := tpm.GetAKPublic()
	if err != nil {
		return nil, err
	}
	ekPubkey, ekCert, err := tpm.GetEKPublic()
	if err != nil {
		return nil, err
	}
	return &api.Node{
		Address: newNodeInfo.Ip,
		Integrity: &api.Node_Tpm2{Tpm2: &api.NodeTPM2{
			AkPub:    akPublic,
			EkCert:   ekCert,
			EkPubkey: ekPubkey,
		}},
		GlobalUnlockKey: newNodeInfo.GlobalUnlockKey,
		IdCert:          newNodeInfo.IdCert,
		State:           api.Node_MASTER,
	}, nil
}

func (s *Server) TPM2Unlock(unlockServer api.NodeManagementService_TPM2UnlockServer) error {
	nonce := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return status.Error(codes.Unavailable, "failed to get randonmess")
	}
	if err := unlockServer.Send(&api.TPM2UnlockFlowResponse{
		Stage: &api.TPM2UnlockFlowResponse_UnlockInit{
			UnlockInit: &api.TPM2UnlockInit{
				Nonce: nonce,
			},
		},
	}); err != nil {
		return err
	}
	unlockReqContainer, err := unlockServer.Recv()
	if err != nil {
		return err
	}
	unlockReqVariant, ok := unlockReqContainer.Stage.(*api.TPM2UnlockFlowRequeset_UnlockRequest)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "protocol violation")
	}
	unlockRequest := unlockReqVariant.UnlockRequest

	store := s.getStore()
	// This is safe, etcd does not do relative paths
	path := nodesPrefix + unlockRequest.NodeId
	nodeRes, err := store.Get(unlockServer.Context(), path)
	if err != nil {
		return status.Error(codes.Unavailable, "consensus request failed")
	}
	if nodeRes.Count == 0 {
		return status.Error(codes.NotFound, "this node does not exist")
	} else if nodeRes.Count > 1 {
		panic("invariant violation: more than one node with the same id")
	}
	nodeRaw := nodeRes.Kvs[0].Value
	var node api.Node
	if err := proto.Unmarshal(nodeRaw, &node); err != nil {
		s.Logger.Error("Failed to decode node", zap.Error(err))
		return status.Error(codes.Internal, "invalid node")
	}

	nodeTPM2, ok := node.Integrity.(*api.Node_Tpm2)
	if !ok {
		return status.Error(codes.InvalidArgument, "node not integrity-protected with TPM2")
	}

	validQuote, err := tpm.VerifyAttestPlatform(nonce, nodeTPM2.Tpm2.AkPub, unlockRequest.Quote, unlockRequest.QuoteSignature)
	if err != nil {
		return status.Error(codes.PermissionDenied, "invalid quote")
	}

	pcrHash := sha256.New()
	for _, pcr := range unlockRequest.Pcrs {
		pcrHash.Write(pcr)
	}
	expectedPCRHash := pcrHash.Sum(nil)

	if !bytes.Equal(validQuote.AttestedQuoteInfo.PCRDigest, expectedPCRHash) {
		return status.Error(codes.InvalidArgument, "the quote's PCR hash does not match the supplied PCRs")
	}

	// TODO: Plug in policy engine to decide if the unlock should actually happen

	return unlockServer.Send(&api.TPM2UnlockFlowResponse{Stage: &api.TPM2UnlockFlowResponse_UnlockResponse{
		UnlockResponse: &api.TPM2UnlockResponse{
			GlobalUnlockKey: node.GlobalUnlockKey,
		},
	}})
}

func (s *Server) NewTPM2NodeRegister(registerServer api.NodeManagementService_NewTPM2NodeRegisterServer) error {
	registerReqContainer, err := registerServer.Recv()
	if err != nil {
		return err
	}
	registerReqVariant, ok := registerReqContainer.Stage.(*api.TPM2FlowRequest_Register)
	if !ok {
		return status.Error(codes.InvalidArgument, "protocol violation")
	}
	registerReq := registerReqVariant.Register

	challengeNonce := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, challengeNonce); err != nil {
		return status.Error(codes.Unavailable, "failed to get randonmess")
	}
	challenge, challengeBlob, err := tpm.MakeAKChallenge(registerReq.EkPubkey, registerReq.AkPublic, challengeNonce)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "failed to challenge AK: %v", err)
	}
	nonce := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return status.Error(codes.Unavailable, "failed to get randonmess")
	}
	if err := registerServer.Send(&api.TPM2FlowResponse{Stage: &api.TPM2FlowResponse_AttestRequest{AttestRequest: &api.TPM2AttestRequest{
		AkChallenge:       challenge,
		AkChallengeSecret: challengeBlob,
		QuoteNonce:        nonce,
	}}}); err != nil {
		return err
	}
	attestationResContainer, err := registerServer.Recv()
	if err != nil {
		return err
	}
	attestResVariant, ok := attestationResContainer.Stage.(*api.TPM2FlowRequest_AttestResponse)
	if !ok {
		return status.Error(codes.InvalidArgument, "protocol violation")
	}
	attestRes := attestResVariant.AttestResponse

	if subtle.ConstantTimeCompare(attestRes.AkChallengeSolution, challengeNonce) != 1 {
		return status.Error(codes.InvalidArgument, "invalid challenge response")
	}

	validQuote, err := tpm.VerifyAttestPlatform(nonce, registerReq.AkPublic, attestRes.Quote, attestRes.QuoteSignature)
	if err != nil {
		return status.Error(codes.PermissionDenied, "invalid quote")
	}

	pcrHash := sha256.New()
	for _, pcr := range attestRes.Pcrs {
		pcrHash.Write(pcr)
	}
	expectedPCRHash := pcrHash.Sum(nil)

	if !bytes.Equal(validQuote.AttestedQuoteInfo.PCRDigest, expectedPCRHash) {
		return status.Error(codes.InvalidArgument, "the quote's PCR hash does not match the supplied PCRs")
	}

	newNodeInfoContainer, err := registerServer.Recv()
	newNodeInfoVariant, ok := newNodeInfoContainer.Stage.(*api.TPM2FlowRequest_NewNodeInfo)
	newNodeInfo := newNodeInfoVariant.NewNodeInfo

	store := s.getStore()
	res, err := store.Get(registerServer.Context(), "enrolments/"+base64.RawURLEncoding.EncodeToString(newNodeInfo.EnrolmentConfig.EnrolmentSecret))
	if err != nil {
		return status.Error(codes.Unavailable, "Consensus unavailable")
	}
	if res.Count == 0 {
		return status.Error(codes.PermissionDenied, "Invalid enrolment secret")
	} else if res.Count > 1 {
		panic("more than one value for the same key, bailing")
	}
	rawVal := res.Kvs[0].Value
	var config api.EnrolmentConfig
	if err := proto.Unmarshal(rawVal, &config); err != nil {
		return err
	}

	// TODO: Plug in policy engine here

	node := api.Node{
		Address: newNodeInfo.Ip,
		Integrity: &api.Node_Tpm2{Tpm2: &api.NodeTPM2{
			AkPub:    registerReq.AkPublic,
			EkCert:   registerReq.EkCert,
			EkPubkey: registerReq.EkPubkey,
		}},
		GlobalUnlockKey: newNodeInfo.GlobalUnlockKey,
		IdCert:          newNodeInfo.IdCert,
		State:           api.Node_UNININITALIZED,
	}

	if err := s.registerNewNode(&node); err != nil {
		s.Logger.Error("failed to register a node", zap.Error(err))
		return status.Error(codes.Internal, "failed to register node")
	}

	return nil
}
