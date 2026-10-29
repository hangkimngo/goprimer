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
	if k > 0 {
		for i := 1; i <= k; i++ {
			x <<= 1
		}
	} else if k < 0 {
		for i := -1; i <= k; i-- {
			x >>= 1
		}
	}
	return x
}

func Reverse32(x uint32) uint32 {
	var reversedBit uint32
	for i := 0; i < 32; i++ {
		mostRight := x & 1
		reversedBit |= mostRight << (31 - int32(i))
		x >>= 1
	}
	return reversedBit
}

func ReverseBytes32(x uint32) uint32 {
	var reversedBytes uint32
	fourthPartNew := x >> 24
	secondPartNew := x << 8 & (0x00FF0000)
	thirdPartNew := x >> 8 & (0x0000FF00)
	firstPartNew := x << 24
	reversedBytes = firstPartNew | secondPartNew | thirdPartNew | fourthPartNew
	return reversedBytes
}
