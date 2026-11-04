package mystrconv

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func FormatUint(u uint64, base int) (string, bool) {
	if base < 2 || base > 36 {
		return "", false
	}
	if u == 0 {
		return "0", true
	}

	var result string
	for u > 0 {
		remainder := u % uint64(base)
		result = string(digits[remainder]) + result
		u /= uint64(base)
	}
	return result, true
}

func FormatInt(i int64, base int) (string, bool) {
	if base < 2 || base > 36 {
		return "", false
	}

	if i >= 0 {
		u := uint64(i)
		return FormatUint(u, base)
	}

	u := uint64(-(i + 1)) + 1
	result, _ := FormatUint(u, base)
	return "-" + result, true
}

func ParseUint(s string, base int, bitSize int) (uint64, bool) {
	if base < 2 || base > 36 {
		return 0, false
	}

	if bitSize < 0 || bitSize > 64 {
		return 0, false
	}

	if len(s) == 0 || s[0] == '+' || s[0] == '-' {
		return 0, false
	}

	if bitSize == 0 {
		bitSize = 64
	}

	var value uint64
	var max uint64
	if bitSize == 64 {
		max = ^uint64(0)
	} else {
		max = (uint64(1) << bitSize) - 1
	}

	for i := 0; i < len(s); i++ {
		digit, ok := digitValue(s[i])
		if !ok {
			return 0, false
		}

		if digit >= uint64(base) {
			return 0, false
		}

		if value > max/uint64(base) || (value == max/uint64(base) && digit > max%uint64(base)) {
			return 0, false
		}
		value = value*uint64(base) + digit
	}

	return value, true
}

func Atoi(s string) (int, bool) {
	negative := false
	if len(s) > 0 && s[0] == '-' {
		negative = true
		s = s[1:]
	} else if len(s) > 0 && s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0, false
	}

	// int64 has 64 bits including 1 bit for sign, and 63 for expression
	bitSize := 63
	if negative {
		bitSize = 64
	}

	i, ok := ParseUint(s, 10, bitSize) // check max = 2^63 - 1
	if !ok {
		return 0, false
	}

	minInt64 := -1 << 63
	limit := uint64(1) << 63

	if negative == true {
		if i > limit {
			return 0, false
		}

		if i == uint64(1)<<63 {
			return int(minInt64), true
		}

		return int(i) * -1, true
	}
	return int(i), true
}
