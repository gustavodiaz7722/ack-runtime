package compare_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func TestIsZeroValue(t *testing.T) {
	require := require.New(t)

	emptyString := ""
	nonEmptyString := "hello"
	zeroInt := 0
	nonZeroInt := 42
	falseBool := false
	trueBool := true
	var nullPtr *string

	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"nil", nil, false},
		{"nil pointer", nullPtr, false},
		{"pointer to empty string", &emptyString, true},
		{"pointer to non-empty string", &nonEmptyString, false},
		{"pointer to zero int", &zeroInt, true},
		{"pointer to non-zero int", &nonZeroInt, false},
		{"pointer to false", &falseBool, true},
		{"pointer to true", &trueBool, false},
		{"direct zero int", 0, true},
		{"direct non-zero int", 1, false},
		{"direct empty string", "", true},
		{"direct non-empty string", "hello", false},
		{"direct false", false, true},
		{"direct true", true, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected {
				require.True(compare.IsZeroValue(tc.input), tc.name)
			} else {
				require.False(compare.IsZeroValue(tc.input), tc.name)
			}
		})
	}
}

func TestIsNilEqualsZero(t *testing.T) {
	require := require.New(t)

	emptyString := ""
	nonEmptyString := "hello"
	var nullPtr *string

	tests := []struct {
		name     string
		a        interface{}
		b        interface{}
		expected bool
	}{
		{"nil and nil", nil, nil, false},
		{"nil and empty string", nil, "", true},
		{"nil and zero int", nil, 0, true},
		{"nil and false", nil, false, true},
		{"nil and pointer to empty string", nil, &emptyString, true},
		{"nil pointer and empty string", nullPtr, "", true},
		{"nil pointer and pointer to empty string", nullPtr, &emptyString, true},
		{"nil and non-empty string", nil, "hello", false},
		{"nil and non-zero int", nil, 42, false},
		{"nil and true", nil, true, false},
		{"nil and pointer to non-empty string", nil, &nonEmptyString, false},
		{"non-nil and empty string", &emptyString, "", false},
		{"non-nil and pointer to empty string", &nonEmptyString, &emptyString, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected {
				require.True(compare.IsNilEqualsZero(tc.a, tc.b), tc.name)
			} else {
				require.False(compare.IsNilEqualsZero(tc.a, tc.b), tc.name)
			}
		})
	}
}

func TestIsZeroEqualsNil(t *testing.T) {
	require := require.New(t)

	emptyString := ""
	nonEmptyString := "hello"
	var nullPtr *string

	tests := []struct {
		name     string
		a        interface{}
		b        interface{}
		expected bool
	}{
		{"empty string and nil", "", nil, true},
		{"zero int and nil", 0, nil, true},
		{"false and nil", false, nil, true},
		{"pointer to empty string and nil", &emptyString, nil, true},
		{"pointer to empty string and nil pointer", &emptyString, nullPtr, true},
		{"nil and nil", nil, nil, false},
		{"non-empty string and nil", "hello", nil, false},
		{"pointer to non-empty string and nil", &nonEmptyString, nil, false},
		{"empty string and non-nil", "", &emptyString, false},
		{"nil pointer and nil", nullPtr, nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected {
				require.True(compare.IsZeroEqualsNil(tc.a, tc.b), tc.name)
			} else {
				require.False(compare.IsZeroEqualsNil(tc.a, tc.b), tc.name)
			}
		})
	}
}
