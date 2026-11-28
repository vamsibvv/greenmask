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

func TestEmail(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "empty string", value: "", expected: ""},
		{name: "empty local part", value: "@example.com", expected: "@example.com"},
		{name: "empty domain", value: "user@", expected: "use****@"},
		{name: "without domain", value: "alexandra.johnson", expected: "ale****"},
		{name: "local part shorter than 3 characters", value: "ab@example.com", expected: "ab****@example.com"},
		{name: "local part of 3 characters", value: "bob@example.com", expected: "bob****@example.com"},
		{name: "local part of 7 characters", value: "ggw.cha@gmail.com", expected: "ggw****@gmail.com"},
		{name: "long local part", value: "alexandra.johnson@example.com", expected: "ale****@example.com"},
		{name: "several @", value: "john@doe@example.com", expected: "joh****@example.com"},
		{name: "quoted local part with @", value: `"john@doe"@example.com`, expected: `"jo****@example.com`},
		{name: "multibyte", value: "иван.петров@example.com", expected: "ива****@example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Email(tt.value))
		})
	}
}
