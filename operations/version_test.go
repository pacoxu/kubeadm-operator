/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package operations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCrossVersions(t *testing.T) {
	var tests = []struct {
		current, target string
		expected        []string
	}{
		{
			current:  "v1.0.0",
			target:   "v1.0.0",
			expected: []string{},
		},
		{
			current:  "v1.0.0",
			target:   "v1.1.0",
			expected: []string{},
		},
		{
			current: "v1.0.0",
			target:  "v1.2.1",
			expected: []string{
				"v1.1.0",
			},
		},
		{
			current: "v1.0.0",
			target:  "v1.5.2",
			expected: []string{
				"v1.1.0",
				"v1.2.0",
				"v1.3.0",
				"v1.4.0",
			},
		},
	}
	for _, rt := range tests {
		t.Run(fmt.Sprintf("%s-%s", rt.current, rt.target), func(t *testing.T) {
			actual := getCrossVersions(rt.current, rt.target)
			if !reflect.DeepEqual(actual, rt.expected) {
				t.Errorf("getCrossVersions(%s, %s) == %v, want %v", rt.current, rt.target, actual, rt.expected)
			}
		})
	}

}
