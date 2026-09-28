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
