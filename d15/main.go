package main

import (
	"fmt"
	"math/big"
)

func main() {

	var first *big.Int
	var second *big.Int
	first = big.NewInt(783)
	second = big.NewInt(325)
	// first = big.NewInt(65)
	// second = big.NewInt(8921)

	var firstMultiplier = big.NewInt(16807)
	var secondMultiplier = big.NewInt(48271)
	var maxInt = big.NewInt(2147483647)
	var matches = 0
	var z1 *big.Int = big.NewInt(0)
	var z2 *big.Int = big.NewInt(0)

	for i := 0; i < 40000000; i++ {
		first = first.Mul(first, firstMultiplier)
		second = second.Mul(second, secondMultiplier)

		var m1 *big.Int = big.NewInt(0)
		var m2 *big.Int = big.NewInt(0)
		z1, m1 = big.NewInt(0).DivMod(first, maxInt, m1)
		first = m1

		z2, m2 = big.NewInt(0).DivMod(second, maxInt, m2)
		second = m2

		var firstCutoff = first.Uint64()
		var secondCutoff = second.Uint64()
		firstCutoff = firstCutoff & 0xFFFF
		secondCutoff = secondCutoff & 0xFFFF

		if firstCutoff == secondCutoff {
			matches++
		}
	}

	z1 = z1
	z2 = z2
	fmt.Println(matches)
}
