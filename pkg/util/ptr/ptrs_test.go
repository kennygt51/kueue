/*
Copyright The Kubernetes Authors.

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

package ptr

import (
	"testing"
)

func TestValEquals(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		cases := map[string]struct {
			ptr  *string
			want string
			exp  bool
		}{
			"nil pointer": {
				ptr:  nil,
				want: "test",
				exp:  false,
			},
			"equal values": {
				ptr:  stringPtr("test"),
				want: "test",
				exp:  true,
			},
			"different values": {
				ptr:  stringPtr("hello"),
				want: "world",
				exp:  false,
			},
			"empty string match": {
				ptr:  stringPtr(""),
				want: "",
				exp:  true,
			},
			"empty string no match": {
				ptr:  stringPtr(""),
				want: "test",
				exp:  false,
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				result := ValEquals(tc.ptr, tc.want)
				if result != tc.exp {
					t.Errorf("ValEquals(%v, %q) = %v; want %v", tc.ptr, tc.want, result, tc.exp)
				}
			})
		}
	})

	t.Run("Int", func(t *testing.T) {
		cases := map[string]struct {
			ptr  *int
			want int
			exp  bool
		}{
			"nil pointer": {
				ptr:  nil,
				want: 42,
				exp:  false,
			},
			"equal values": {
				ptr:  intPtr(42),
				want: 42,
				exp:  true,
			},
			"different values": {
				ptr:  intPtr(10),
				want: 20,
				exp:  false,
			},
			"zero value match": {
				ptr:  intPtr(0),
				want: 0,
				exp:  true,
			},
			"zero value no match": {
				ptr:  intPtr(0),
				want: 1,
				exp:  false,
			},
			"negative values": {
				ptr:  intPtr(-5),
				want: -5,
				exp:  true,
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				result := ValEquals(tc.ptr, tc.want)
				if result != tc.exp {
					t.Errorf("ValEquals(%v, %d) = %v; want %v", tc.ptr, tc.want, result, tc.exp)
				}
			})
		}
	})

	t.Run("Bool", func(t *testing.T) {
		cases := map[string]struct {
			ptr  *bool
			want bool
			exp  bool
		}{
			"nil pointer": {
				ptr:  nil,
				want: true,
				exp:  false,
			},
			"true equals true": {
				ptr:  boolPtr(true),
				want: true,
				exp:  true,
			},
			"false equals false": {
				ptr:  boolPtr(false),
				want: false,
				exp:  true,
			},
			"true not equals false": {
				ptr:  boolPtr(true),
				want: false,
				exp:  false,
			},
			"false not equals true": {
				ptr:  boolPtr(false),
				want: true,
				exp:  false,
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				result := ValEquals(tc.ptr, tc.want)
				if result != tc.exp {
					t.Errorf("ValEquals(%v, %t) = %v; want %v", tc.ptr, tc.want, result, tc.exp)
				}
			})
		}
	})

	t.Run("CustomStruct", func(t *testing.T) {
		type TestStruct struct {
			Name string
			ID   int
		}

		cases := map[string]struct {
			ptr  *TestStruct
			want TestStruct
			exp  bool
		}{
			"nil pointer": {
				ptr:  nil,
				want: TestStruct{Name: "test", ID: 1},
				exp:  false,
			},
			"equal structs": {
				ptr:  &TestStruct{Name: "test", ID: 1},
				want: TestStruct{Name: "test", ID: 1},
				exp:  true,
			},
			"different name": {
				ptr:  &TestStruct{Name: "test", ID: 1},
				want: TestStruct{Name: "other", ID: 1},
				exp:  false,
			},
			"different id": {
				ptr:  &TestStruct{Name: "test", ID: 1},
				want: TestStruct{Name: "test", ID: 2},
				exp:  false,
			},
			"zero value struct": {
				ptr:  &TestStruct{},
				want: TestStruct{},
				exp:  true,
			},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				result := ValEquals(tc.ptr, tc.want)
				if result != tc.exp {
					t.Errorf("ValEquals(%v, %v) = %v; want %v", tc.ptr, tc.want, result, tc.exp)
				}
			})
		}
	})
}

// Helper functions to create pointers for test cases
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
