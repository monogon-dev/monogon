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
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"go.etcd.io/etcd/clientv3"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
)

const enrolmentPrefix = "enrolments/"

var errNotExists = errors.New("not found")

type EnrolmentStore struct {
	backend clientv3.KV
}

func (s *EnrolmentStore) GetBySecret(ctx context.Context, secret []byte) (*api.EnrolmentConfig, error) {

	res, err := s.backend.Get(ctx, enrolmentPrefix+base64.RawURLEncoding.EncodeToString(secret))
	if err != nil {
		return nil, fmt.Errorf("failed to query consensus: %w", err)
	}
	if res.Count == 0 {
		return nil, errNotExists
	} else if res.Count > 1 {
		panic("more than one value for the same key, bailing")
	}
	rawVal := res.Kvs[0].Value
	var config *api.EnrolmentConfig
	if err := proto.Unmarshal(rawVal, config); err != nil {
		return nil, err
	}
	return config, nil
}
