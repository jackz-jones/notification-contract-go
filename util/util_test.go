package util

import (
	"testing"
)

func TestBoolToBytes(t *testing.T) {
	okBytes := BoolToBytes(true)
	resolvedOk := BytesToBool(okBytes)
	if true != resolvedOk {
		t.Errorf("origin:%t, resolved:%t", true, resolvedOk)
	} else {
		t.Logf("resolve ok:%t", resolvedOk)
	}
}
