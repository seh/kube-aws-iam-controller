package main

import (
	"strings"
	"testing"
)

func TestValidateExternalIDPrefix(t *testing.T) {
	tests := []struct {
		description   string
		prefix        string
		expectFailure bool
		err           string
	}{
		{
			"empty",
			"",
			false,
			"",
		},
		{
			"short",
			"a",
			false,
			"",
		},
		{
			"long",
			strings.Repeat("a", externalIDPrefixMaxSize),
			false,
			"",
		},
		{
			"too long",
			strings.Repeat("a", externalIDPrefixMaxSize+1),
			true,
			"long",
		},
		{
			"includes tolerated characters",
			"az09+=,.@:\\-",
			false,
			"",
		},
		{
			"includes forward slash at start",
			"/ab",
			true,
			"slash",
		},
		{
			"includes forward slash in middle",
			"a/b",
			true,
			"slash",
		},
		{
			"includes forward slash at end",
			"ab/",
			true,
			"slash",
		},
		{
			"includes other forbidden characters",
			"`~!#$%^&*()_[]{}|",
			true,
			"alphanumeric",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if err := validateExternalIDPrefix(test.prefix); err != nil {
				if !test.expectFailure {
					t.Errorf("got %v, want success", err)
				}
				if len(test.err) > 0 && !strings.Contains(err.Error(), test.err) {
					t.Errorf("substring not found: got %q, want %q", err.Error(), test.err)
				}
			} else if test.expectFailure {
				t.Error("got success, want failure")
			}
		})
	}
}
