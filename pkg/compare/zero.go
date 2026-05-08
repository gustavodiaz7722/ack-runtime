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
