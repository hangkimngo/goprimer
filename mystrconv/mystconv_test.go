package mystrconv

import (
	"math"
	"strconv"
	"testing"
)

func TestFormatUint(t *testing.T) {
	tests := []struct {
		name string
		val  uint64
		base int
	}{
		{"0, base = 10", 0, 10},
		{"255, base = 16", 255, 16},
		{"35, base = 36", 35, 36},
		{"1, base = 37", 1, 37},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := FormatUint(tt.val, tt.base)
			if tt.base < 2 || tt.base > 36 {
				if gotOK || got != "" {
					t.Errorf("FormatUint() = (%q, %v), want (\"\", false)", got, gotOK)
				}
				return
			}
			want := strconv.FormatUint(tt.val, tt.base)
			if !gotOK || want != got {
				t.Errorf("FormatUint(%d, %d) = (%q, %v), want (%q, true)", tt.val, tt.base, got, gotOK, want)
			}
		})
	}
}

func TestFormatInt(t *testing.T) {
	tests := []struct {
		name string
		val  int64
		base int
	}{
		{"0, base = 10", 0, 10},
		{"255, base = 16", 255, 16},
		{"35, base = 36", 35, 36},
		{"1, base = 37", 1, 37},
		{"-1, base = 1", -1, 1},
		{"-255, base = 16", -255, 16},
		{"-13, base = 2", -13, 2},
		{"min value -9223372036854775808, base=10", -9223372036854775808, 10},
		{"max value, base 64", int64(math.Pow(2, 64)) - 1, 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := FormatInt(tt.val, tt.base)
			if tt.base < 2 || tt.base > 36 {
				if gotOK || got != "" {
					t.Errorf("FormatInt() = (%q, %v), want (\"\", false)", got, gotOK)
				}
				return
			}

			want := strconv.FormatInt(tt.val, tt.base)
			if !gotOK || want != got {
				t.Errorf("FormatInt(%d, %d) = (%q, %v), want (%q, true)", tt.val, tt.base, got, gotOK, want)
			}
		})
	}
}

func TestParseUint(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		base    int
		bitSize int
	}{
		{"decimal", "42", 10, 64},
		{"hex lowercase", "ff", 16, 64},
		{"hex uppercase", "FF", 16, 64},
		{"base 36", "z", 36, 64},
		{"zero", "0", 10, 64},
		{"empty", "", 10, 64},
		{"negative", "-1", 10, 64},
		{"plus sign", "+1", 10, 64},
		{"invalid base 1", "1", 1, 64},
		{"invalid base 37", "1", 37, 64},
		{"invalid bitSize", "1", 10, 65},
		{"256 exceeds 8 bits", "256", 10, 8},
		{"255 fits 8 bits", "255", 10, 8},
		{"whitespace", "1 ", 10, 64},
		{"underscore", "1_000", 10, 64},
		{"invalid digit", "2a", 10, 64},
		{"bitSize zero", "300", 10, 0},
		{"one bit max", "1", 10, 1},
		{"one bit overflow", "2", 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := ParseUint(tt.s, tt.base, tt.bitSize)

			want, err := strconv.ParseUint(tt.s, tt.base, tt.bitSize)
			wantOK := err == nil

			if gotOK != wantOK {
				t.Errorf("success status = %v, want %v", gotOK, wantOK)
				return
			}

			if gotOK && got != want {
				t.Errorf(
					"ParseUint(%q, %d, %d) = (%d, %v), want (%d, %v)",
					tt.s, tt.base, tt.bitSize,
					got, gotOK, want, wantOK,
				)
			}

			if !gotOK && got != 0 {
				t.Errorf("got %d on error, want 0", got)
			}
		})
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{"positive", "42"},
		{"negative", "-42"},
		{"explicit plus", "+42"},
		{"zero", "0"},
		{"negative zero", "-0"},
		{"positive max int64", "9223372036854775807"},
		{"negative min int64", "-9223372036854775808"},
		{"positive overflow", "9223372036854775808"},
		{"negative overflow", "-9223372036854775809"},
		{"empty", ""},
		{"only minus", "-"},
		{"only plus", "+"},
		{"decimal", "4.2"},
		{"letters", "hello"},
		{"whitespace", " 42"},
		{"trailing whitespace", "42 "},
		{"double sign", "--42"},
		{"sign without digits", "+-"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := Atoi(tt.s)

			want, err := strconv.Atoi(tt.s)
			wantOK := err == nil

			if gotOK != wantOK {
				t.Errorf("success status = %v, want %v",
					gotOK, wantOK)
				return
			}

			if gotOK && got != want {
				t.Errorf("Atoi(%q) = %d, want %d",
					tt.s, got, want)
			}

			if !gotOK && got != 0 {
				t.Errorf("got %d on error, want 0", got)
			}
		})
	}
}
