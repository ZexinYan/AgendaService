package testutil

import (
	"reflect"
	"testing"
)

// ExpectDeepEq ..
func ExpectDeepEq(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Fail: Expected: %v, Actual: %v\n", b, a)
	} else {
		t.Log("Pass")
	}
}

// NoError ..
func NoError(t *testing.T, err error) bool {
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
		return false
	}
	return true
}

// ShouldError ..
func ShouldError(t *testing.T, err error) bool {
	if err == nil {
		t.Error("Should have error here")
		return false
	}
	return true
}
