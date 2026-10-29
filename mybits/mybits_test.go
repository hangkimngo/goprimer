package mybits

import (
	"math/bits"
	"testing"
)

func TestOnesCount32(t *testing.T) {
	onesCount := OnesCount32(5)
	want := bits.OnesCount32(5)
	if want != onesCount {
		t.Errorf("OnesCount32(5) = %v, want match for %v", onesCount, want)
	}
}
