package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 5, 10
	wantA, wantB := 10, 5

	gotA, gotB := swap(a, b)
	if gotA != wantA || gotB != wantB {
		t.Errorf("swap(%d, %d) = (%d, %d); want (%d, %d)", a, b, gotA, gotB, wantA, wantB)
	}
	gotX, gotY := swapXor(a, b)
	if gotX != wantA || gotY != wantB {
		t.Errorf("swapXor(%d, %d) = (%d, %d); want (%d, %d)", a, b, gotX, gotY, wantA, wantB)
	}
}
