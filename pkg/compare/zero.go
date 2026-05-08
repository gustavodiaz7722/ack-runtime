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
	"reflect"
)

func IsZeroValue(i interface{}) bool {
	if i == nil {
		return false
	}
	v := reflect.ValueOf(i)
	// Dereference pointer to check if the pointed-to value is zero
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		return v.Elem().IsZero()
	}
	return v.IsZero()
}

// IsNilEqualsZero returns true if 'a' satisfies IsNil and 'b' satisfies IsZeroValue.
func IsNilEqualsZero(a interface{}, b interface{}) bool {
	return IsNil(a) && IsZeroValue(b)
}

func IsZeroEqualsNil(a interface{}, b interface{}) bool {
	return IsZeroValue(a) && IsNil(b)
}
