package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestA(t *testing.T) {
	if diff := cmp.Diff("a", "b"); diff != "" {
		t.Errorf("diff %s", diff)
	}
}
