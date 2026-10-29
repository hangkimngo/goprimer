package main

import (
	"fmt"
	"goprimer/mybits"
	"math/bits"
)

func main() {
	a := bits.ReverseBytes32(7)
	fmt.Printf("Reverse order of %d: %b \n", 7, a)
	fmt.Printf("Reverse order of %d: %b", 7, mybits.ReverseBytes32(7))
}
