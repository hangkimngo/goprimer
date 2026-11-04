package myhex

const hexChars = "0123456789abcdef"

func Encode(src []byte) []byte {
	dst := make([]byte, len(src)*2)
	for i, b := range src {
		high := b >> 4
		low := b & 0x0F

		dst[2*i] = hexChars[high]
		dst[2*i+1] = hexChars[low]
	}
	return dst
}

func Decode(src []byte) ([]byte, bool) {
	if len(src)%2 != 0 {
		return nil, false
	}
	dst := make([]byte, len(src)/2)
	for i := 0; i < len(src); i = i + 2 {
		highIndex, ok := hexIndex(src[i])
		if !ok {
			return nil, false
		}

		lowIndex, ok := hexIndex(src[i+1])
		if !ok {
			return nil, false
		}

		b := highIndex<<4 | lowIndex
		dst[i/2] = b
	}
	return dst, true
}
