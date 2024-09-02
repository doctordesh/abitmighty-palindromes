package main

import (
	"math"
)

func generateWithLength(length, base int, cb func([]byte) bool) bool {
	// if length < 1 {
	// 	return true
	// }

	// if length == 1 {
	// 	for i := byte(1); i < byte(base); i++ {
	// 		done := cb([]byte{i})
	// 		if done {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	half_length := length / 2 // intentional integer division, rounds down
	from := int(math.Pow(float64(base), float64(half_length-1)))
	to := int(math.Pow(float64(base), float64(half_length)))

	if length%2 == 1 { // odd
		for y := from; y < to; y++ {
			for x := 0; x < base; x++ {
				done := cb(computeExpandedForm(y, x, length, half_length, base))
				if done {
					return true
				}
			}
		}
	} else { // even
		for y := from; y < to; y++ {
			done := cb(computeExpandedForm(y, 0, length, half_length, base))
			if done {
				return true
			}
		}
	}

	return false
}

func computePalindromeValue(p []byte, base int) int {
	sum := 0
	for i := 0; i < len(p); i++ {
		sum += int(p[len(p)-i-1]) * int(math.Pow(float64(base), float64(i)))
	}

	return sum
}

func computeExpandedForm(y, x, length, half_length, base int) []byte {
	number := make([]byte, length)
	for hl := 0; hl < half_length; hl++ {
		number[hl] = byte(y / int(math.Pow(float64(base), float64(half_length-hl-1))) % base)
		number[length-hl-1] = number[hl]
	}
	if length%2 != 0 { // odd
		number[half_length] = byte(x)
	}
	return number
}

func PalindromeSum_v4(limit int) (int, int) {
	sum := 0
	count := 0

	cb := func(palindrome []byte) bool {
		// fmt.Println(palindrome)
		value := computePalindromeValue(palindrome, 22)
		if value >= limit {
			return true
		}

		sum += value
		count += 1
		return false
	}

	digits := 1
	for generateWithLength(digits, 22, cb) == false {
		digits += 1
	}

	return sum, count
}
