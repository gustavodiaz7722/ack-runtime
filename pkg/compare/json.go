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
	"encoding/json"
	"fmt"
	"reflect"
)

// JSONStringEqual returns true if two strings containing JSON are
// semantically equivalent. It handles differences in whitespace,
// key ordering, and trailing newlines that commonly occur when comparing
// JSON values from Kubernetes spec fields (which may include YAML block
// scalar artifacts) against values returned by AWS APIs.
//
// Returns an error if either string cannot be parsed as valid JSON.
func JSONStringEqual(a, b string) (bool, error) {
	var aJSON, bJSON interface{}

	if err := json.Unmarshal([]byte(a), &aJSON); err != nil {
		return false, fmt.Errorf("failed to parse first JSON string: %w", err)
	}
	if err := json.Unmarshal([]byte(b), &bJSON); err != nil {
		return false, fmt.Errorf("failed to parse second JSON string: %w", err)
	}

	return reflect.DeepEqual(aJSON, bJSON), nil
}
