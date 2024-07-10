package main

import (
	"fmt"
	"math"
)

var DIGITS = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

func ComputeValue(b []byte, base int) int {
	sum := 0
	bx := float64(base)
	bl := len(b)
	blf := float64(bl)
	for i := float64(0); i < blf; i++ {
		sum += int(b[bl-int(i)-1]) * int(math.Pow(bx, i))
	}

	return sum
}

func NumberOfPalindromes(digits, base int) int {
	dimensions := int(float64(digits)/2 + 0.5)
	for d := 0; d < dimensions; d++ {

	}
	return int(math.Pow(float64(base), float64(dimensions)))
}

// func Palindromes(digits, base int) {
// 	dimensions := int(float64(digits)/2 + 0.5)
// 	fmt.Printf("Digi: %d, Dim: %d, Base: %d\n", digits, dimensions, base)
// 	for d := 0; d < dimensions; d++ {

// 	}
// }

func Palindromes(nrDigits, base int) int {
	digits := DIGITS[0:base]
	sum := 0

	for d := 0; d < nrDigits; d++ {

		sum += ComputeValue(digits, base)
		fmt.Println(digits, d)
	}

	return sum
}

func PalindromesOfSize(n, base int) []string {
	digits := make([]byte, n)

	for di := byte(0); di < byte(base); di++ {
		for i := 0; i < int(float64(n/2)+0.5); i++ {
			digits[i] = di
			digits[n-i-1] = di

			fmt.Println(digits, di)
		}
	}

	// p := []string{}

	// cols := int(float64(n)/2 + 0.5)
	// for digit := 0; digit < base; digit++ {
	// 	// number := make([]byte, n)
	// 	for i := 0; i < cols; i++ {

	// 	}
	// }

	// for col := 0; col < cols; col++ {
	// 	for digit := 0; digit < base; digit++ {
	// 		number := make([]byte, n)

	// 	}
	// }

	// for pos := 0; pos < n; pos++ {
	// 	number := make([]byte, pos+1)

	// 	for i := 0; i < pos+1; i++ {
	// 		for j := 0; j < base; j++ {
	// 			number[i] = '0'
	// 			number[pos-i] = '0'
	// 		}
	// 	}

	// 	p = append(p, string(number))
	// }
	return []string{}
}

/*

   n and n+1 is close to equal. Middle 'column' is just doubled

   n = 1

   0
   1
   2

   n = 2
   00
   11
   22

   n = 3
   010
   020
   030
   0k0

   101
   111
   121
   131

   n = 4

   0110
   0220
   0330
   0440

   1001
   1111
   1221
   1331
   1441

   n = 5

   00100
   00200
   00300

   01010
   01110
   01210
   01310
   01410


*/
