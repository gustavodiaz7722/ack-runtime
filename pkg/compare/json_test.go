// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package compare

import (
	"testing"
)

func TestJSONStringEqual(t *testing.T) {
	cases := []struct {
		name    string
		a       string
		b       string
		want    bool
		wantErr bool
	}{
		{
			name: "IdenticalStrings",
			a:    `{"key":"value"}`,
			b:    `{"key":"value"}`,
			want: true,
		},
		{
			name: "DifferentWhitespace",
			a:    `{"key":"value","nested":{"a":1}}`,
			b: `{
				"key": "value",
				"nested": {
					"a": 1
				}
			}`,
			want: true,
		},
		{
			name: "TrailingNewline_Issue2869",
			a:    "{\n  \"secrets-store-csi-driver\": {\n    \"syncSecret\": {\n      \"enabled\": true\n    }\n  }\n}\n",
			b:    "{\n  \"secrets-store-csi-driver\": {\n    \"syncSecret\": {\n      \"enabled\": true\n    }\n  }\n}",
			want: true,
		},
		{
			name: "PrettyVsCompact_Issue2877",
			a:    "{\n  \"notificationOrigin\": [\n    \"PRODUCT\"\n  ],\n  \"productType\": [\n    {\n      \"anything-but\": \"simple\"\n    }\n  ]\n}\n",
			b:    `{"notificationOrigin":["PRODUCT"],"productType":[{"anything-but":"simple"}]}`,
			want: true,
		},
		{
			name: "DifferentKeyOrder",
			a:    `{"b":2,"a":1}`,
			b:    `{"a":1,"b":2}`,
			want: true,
		},
		{
			name: "NestedDifferentKeyOrder",
			a:    `{"outer":{"z":26,"a":1},"list":[1,2,3]}`,
			b:    `{"list":[1,2,3],"outer":{"a":1,"z":26}}`,
			want: true,
		},
		{
			name: "DifferentValues",
			a:    `{"key":"value1"}`,
			b:    `{"key":"value2"}`,
			want: false,
		},
		{
			name: "DifferentStructure",
			a:    `{"key":"value"}`,
			b:    `{"key":{"nested":"value"}}`,
			want: false,
		},
		{
			name: "ExtraKey",
			a:    `{"a":1}`,
			b:    `{"a":1,"b":2}`,
			want: false,
		},
		{
			name: "ArrayOrderMatters",
			a:    `{"items":[1,2,3]}`,
			b:    `{"items":[3,2,1]}`,
			want: false,
		},
		{
			name: "BothEmpty",
			a:    `{}`,
			b:    `{}`,
			want: true,
		},
		{
			name: "EmptyVsNonEmpty",
			a:    `{}`,
			b:    `{"key":"value"}`,
			want: false,
		},
		{
			name: "LeadingAndTrailingWhitespace",
			a:    "  \n\t{\"key\":\"value\"}\n  ",
			b:    `{"key":"value"}`,
			want: true,
		},
		{
			name: "NullValues",
			a:    `{"key":null}`,
			b:    `{"key":null}`,
			want: true,
		},
		{
			name: "NullVsMissing",
			a:    `{"key":null}`,
			b:    `{}`,
			want: false,
		},
		{
			name: "NumericTypes",
			a:    `{"int":1,"float":1.5}`,
			b:    `{"int":1,"float":1.5}`,
			want: true,
		},
		{
			name: "BooleanValues",
			a:    `{"enabled":true,"disabled":false}`,
			b:    `{"disabled":false,"enabled":true}`,
			want: true,
		},
		{
			name:    "InvalidJSONFirst",
			a:       `{invalid`,
			b:       `{"key":"value"}`,
			want:    false,
			wantErr: true,
		},
		{
			name:    "InvalidJSONSecond",
			a:       `{"key":"value"}`,
			b:       `not json`,
			want:    false,
			wantErr: true,
		},
		{
			name:    "EmptyStringFirst",
			a:       ``,
			b:       `{"key":"value"}`,
			want:    false,
			wantErr: true,
		},
		{
			name:    "BothEmptyStrings",
			a:       ``,
			b:       ``,
			want:    false,
			wantErr: true,
		},
		{
			name: "ComplexNestedStructure",
			a: `{
				"secrets-store-csi-driver": {
					"syncSecret": {
						"enabled": true
					},
					"providers": ["aws", "azure"]
				}
			}`,
			b:    `{"secrets-store-csi-driver":{"syncSecret":{"enabled":true},"providers":["aws","azure"]}}`,
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := JSONStringEqual(tc.a, tc.b)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
