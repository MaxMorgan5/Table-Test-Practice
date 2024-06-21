package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSepStrings(t *testing.T) {
	tests := map[string]struct {
		input, sep  string
		expected    []string
		expectedErr error
	}{
		"wrong sep": {
			input:       "abcd",
			sep:         "h",
			expected:    []string{"abcd"},
			expectedErr: errors.New("string does not contain sep"),
		},
		"no sep": {
			input:       "aerss",
			sep:         "",
			expected:    nil,
			expectedErr: errors.New("must provide a value for sep"),
		},
		"no string": {
			input:       "",
			sep:         "s",
			expected:    nil,
			expectedErr: errors.New("input string cannot be empty"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			result, err := SepStrings(tc.input, tc.sep)
			if err != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err, "unexpected err : %v", err)
				assert.Equal(t, tc.expected, result, "result mismatch")
			}

		})
	}
}
