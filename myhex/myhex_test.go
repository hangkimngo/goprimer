package myhex

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		src  []byte
	}{
		{"empty input", []byte{}},
		{"single byte", []byte{0x00}},
		{"all bits set", []byte{0xFF}},
		{"valid input", []byte{0xDE, 0xAD}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.src)
			want := []byte(hex.EncodeToString(tt.src))
			if !bytes.Equal(got, want) {
				t.Errorf("Encode(%v) = %s, want %s", tt.src, got, want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name   string
		src    []byte
		want   []byte
		wantOK bool
	}{
		{"lowercase", []byte("dead"), []byte{0xDE, 0xAD}, true},
		{"uppercase", []byte("DEAD"), []byte{0xDE, 0xAD}, true},
		{"empty", []byte(""), []byte{}, true},
		{"odd length", []byte("abc"), nil, false},
		{"invalid character", []byte("zz"), nil, false},
		{"partially valid", []byte("dezz"), nil, false},
		{"invalid low nibble", []byte("dz"), nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := Decode(tt.src)

			if gotOK != tt.wantOK {
				t.Errorf("Decode() ok = %v, want %v", gotOK, tt.wantOK)
			}

			if !bytes.Equal(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}

			// Compare successful decoding with Go standard library.
			if tt.wantOK {
				wantStd, err := hex.DecodeString(string(tt.src))
				if err != nil {
					t.Fatalf("standard library DecodeString() error: %v", err)
				}

				if !bytes.Equal(got, wantStd) {
					t.Errorf("got %v, standard library returned %v", got, wantStd)
				}
			}
		})
	}
}

func TestHexIndex(t *testing.T) {
	tests := []struct {
		name   string
		input  byte
		want   byte
		wantOK bool
	}{
		{"zero", '0', 0, true},
		{"nine", '9', 9, true},
		{"lowercase a", 'a', 10, true},
		{"lowercase f", 'f', 15, true},
		{"uppercase A", 'A', 10, true},
		{"uppercase F", 'F', 15, true},
		{"invalid", 'z', 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := hexIndex(tt.input)

			if got != tt.want || ok != tt.wantOK {
				t.Errorf("hexIndex(%q) = (%d, %v), want (%d, %v)",
					tt.input, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}
