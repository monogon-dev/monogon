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

package common

import "errors"

type DataPlace uint32

const (
	PlaceESP  DataPlace = 0
	PlaceData           = 1
)

var (
	// ErrNotInitialized will be returned when trying to access a place that's not yet initialized
	ErrNotInitialized = errors.New("This place is not initialized")
	// ErrUnknownPlace will be returned when trying to access a place that's not known
	ErrUnknownPlace = errors.New("This place is not known")
)

type StorageManager interface {
	GetPathInPlace(place DataPlace, path string) (string, error)
}
