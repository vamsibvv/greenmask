// Copyright 2026 Greenmask
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

package maskers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "empty string", value: "", expected: ""},
		{name: "3 characters", value: "123", expected: "***"},
		{name: "4 characters", value: "1234", expected: "****"},
		{name: "5 characters", value: "12345", expected: "1****"},
		{name: "9 characters", value: "295847284", expected: "29584****"},
		{name: "10 characters", value: "A123456789", expected: "A12345****"},
		{name: "13 characters", value: "1234567890123", expected: "123456789****"},
		{name: "multibyte prefix", value: "ЖЗ123456", expected: "ЖЗ12****"},
		{name: "multibyte masked part", value: "12ЖЗИК", expected: "12****"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ID(tt.value))
		})
	}
}
