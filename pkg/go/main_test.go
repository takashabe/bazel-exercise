package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestA(t *testing.T) {
	if diff := cmp.Diff("a", "a"); diff != "" {
		t.Errorf("diff %s", diff)
	}
}
