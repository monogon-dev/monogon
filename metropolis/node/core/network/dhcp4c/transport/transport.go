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

// Package transport contains Linux-based transports for the DHCP broadcast and
// unicast specifications.
package transport

import (
	"errors"
	"fmt"
	"net"
)

var DeadlineExceededErr = errors.New("deadline exceeded")

func NewInvalidMessageError(internalErr error) error {
	return &InvalidMessageError{internalErr: internalErr}
}

type InvalidMessageError struct {
	internalErr error
}

func (i InvalidMessageError) Error() string {
	return fmt.Sprintf("received invalid packet: %v", i.internalErr.Error())
}

func (i InvalidMessageError) Unwrap() error {
	return i.internalErr
}

func deadlineFromTimeout(err error) error {
	var timeoutErr net.Error
	if errors.As(err, &timeoutErr) && timeoutErr.Timeout() {
		return DeadlineExceededErr
	}
	return err
}
