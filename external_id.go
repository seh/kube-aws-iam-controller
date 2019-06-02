package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	// externalIDMaxSize is the maximum tolerated length of an IAM external ID.
	externalIDMaxSize = 1224
	// externalIDPrefixMaxSize is the maximum tolerated length of an IAM external ID prefix.
	//
	// Its value comes from the following constraints:
	// The maximum length of a Kubernetes namespace name is 63 characters.
	// The maximum length of a Kubernetes object name is 253 characters.
	// There are three delimiter characters in an external ID.
	externalIDPrefixMaxSize = externalIDMaxSize - 63 - 253 - 3
)

func validateExternalIDPrefix(p string) error {
	if len(p) > externalIDPrefixMaxSize {
		return fmt.Errorf("must not be longer than %d characters", externalIDPrefixMaxSize)
	}
	if strings.IndexByte(p, '/') != -1 {
		return errors.New("must not contain a forward slash character")
	}
	// See the following document for the syntactic constraints for External IDs:
	// https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html#API_AssumeRole_RequestParameters
	// Note that our slightly stricter format, we leave the forward slash out, having confirmed its
	// lack in the prior validation step.
	//
	// Don't bother saving this compiled regular expression as a top-level var, as we only expect
	// this to be used once per run of the program.
	if regexp.MustCompile(`[^\w+=,.@:\\-]`).MatchString(p) {
		return errors.New(`must contain only alphanumeric characters and _+=,.@:\-`)
	}
	return nil
}
