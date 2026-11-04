package mybits

func OnesCount32(x uint32) int {
	onesCount := 0
	for x > 0 {
		x &= x - 1
		onesCount++
	}
	return onesCount
}

func Len32(x uint32) int {
	zerosCount := 0
	for x > 0 {
		x >>= 1
		zerosCount++
	}
	return zerosCount
}

func RotateLeft32(x uint32, k int) uint32 {
	k %= 32
	if k < 0 {
		k += 32
	}
	return (x<<k | x>>(32-k))
}

func Reverse32(x uint32) uint32 {
	var reversedBit uint32
	for i := 0; i < 32; i++ {
		mostRight := x & 1
		reversedBit = reversedBit<<1 | mostRight
		// reversedBit |= mostRight << (31 - int32(i))
		x >>= 1
	}
	return reversedBit
}

func ReverseBytes32(x uint32) uint32 {
	byte1 := x >> 24 & 0xff
	byte2 := x >> 16 & 0xff
	byte3 := x >> 8 & 0xff
	byte4 := x & 0xff
	return (byte4 << 24) | (byte3 << 16) | (byte2 << 8) | byte1
}
