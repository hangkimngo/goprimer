package mybits

import (
	"math/bits"
	"testing"
)

func TestOnesCount32(t *testing.T) {
	tests := []struct {
		name string
		val  uint32
	}{
		{"zero", 0},
		{"one", 1},
		{"seven", 7},
		{"fifteen", 15},
		{"0xFFFFFFFF", 0xFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OnesCount32(tt.val)
			want := bits.OnesCount32(tt.val)
			if want != got {
				t.Errorf("OnesCount32(%v) = %v, want match for %v", tt.val, got, want)
			}
		})
	}
}

func TestLen32(t *testing.T) {
	tests := []struct {
		name string
		val  uint32
	}{
		{"zero", 0},
		{"one", 1},
		{"seven", 7},
		{"eight", 8},
		{"fifteen", 15},
		{"sixteen", 16},
		{"255", 255},
		{"0xFFFFFFFF", 0xFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Len32(tt.val)
			want := bits.Len32(tt.val)
			if want != got {
				t.Errorf("Len32(%v) = %v, want match for %v", tt.val, got, want)
			}
		})
	}
}

func TestRotateLeft32(t *testing.T) {
	tests := []struct {
		name     string
		val      uint32
		rotation int
	}{
		{"zero, k=1", 0, 1},
		{"one, k=1", 1, 1},
		{"0x80000000, k=1", 0x80000000, 1},
		{"0x12345678, k=0", 0x12345678, 0},
		{"0x12345678, k=32", 0x12345678, 32},
		{"0x12345678, k=33", 0x12345678, 33},
		{"0x12345678, k=-1", 0x12345678, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RotateLeft32(tt.val, tt.rotation)
			want := bits.RotateLeft32(tt.val, tt.rotation)
			if want != got {
				t.Errorf("RotateLeft32(%#08x, %d) = %#08x, want match for %#08x", tt.val, tt.rotation, got, want)
			}
		})
	}
}

func TestReverse32(t *testing.T) {
	tests := []struct {
		name string
		val  uint32
	}{
		{"zero", 0},
		{"one", 1},
		{"0x80000000", 0x80000000},
		{"0x12345678", 0x12345678},
		{"0xFFFFFFFF", 0xFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse32(tt.val)
			want := bits.Reverse32(tt.val)
			if want != got {
				t.Errorf("Reverse32(%#08x) = %#08x, want match for %#08x", tt.val, got, want)
			}
		})
	}
}

func TestReverseBytes32(t *testing.T) {
	tests := []struct {
		name string
		val  uint32
	}{
		{"0x00000000", 0x00000000},
		{"0x01020304", 0x01020304},
		{"0x12345678", 0x12345678},
		{"0x000000FF", 0x000000FF},
		{"0xFFFFFFFF", 0xFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseBytes32(tt.val)
			want := bits.ReverseBytes32(tt.val)
			if want != got {
				t.Errorf("ReverseBytes32(%#08x) = %#08x, want match for %#08x", tt.val, got, want)
			}
		})
	}
}
