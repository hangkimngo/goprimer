package mybits

func Pow(n int, exp int) int {
	if exp < 0 { // handle negative exponents
		n = 1 / n
		exp = -exp
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= n
	}
	return result
}

func OnesCount32(x uint32) int {
	remainders := []int{}
	d := x
	onesCount := 0
	for d > 0 {
		remainders = append(remainders, int(d%2))
		d = d / 2
	}
	for i := 0; i <= len(remainders); i++ {
		if i == 1 {
			onesCount++
		}
	}
	return onesCount
}

func Len32(x uint32) int {
	remainders := []int{}
	d := x
	zerosCount := 0
	oneStarted := false
	for d > 0 {
		remainders = append(remainders, int(d%2))
		d = d / 2
	}
	for i := 0; i <= len(remainders); i++ {
		if i == 1 {
			oneStarted = true
		}
		if i == 0 && oneStarted {
			zerosCount++
		}
	}
	return zerosCount
}
