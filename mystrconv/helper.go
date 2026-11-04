package mystrconv

func digitValue(c byte) (uint64, bool) {
	if c >= '0' && c <= '9' {
		return uint64(c - '0'), true
	}
	if c >= 'a' && c <= 'z' {
		return uint64(c-'a') + 10, true
	}
	if c >= 'A' && c <= 'Z' {
		return uint64(c-'A') + 10, true
	}
	return 0, false
}
