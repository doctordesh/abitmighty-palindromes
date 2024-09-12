package main

import (
	"math"
)

func PalindromeGenerator(base int) func(func([]uint8) bool) {
	return func(yield func([]uint8) bool) {
		palindrome_length := 1
		for {
			slots := palindrome_length / 2 // intentional integer division, rounds down
			from := int(math.Pow(float64(base), float64(slots-1)))
			to := int(math.Pow(float64(base), float64(slots))) - 1

			if palindrome_length%2 == 1 { // odd
				for abc := from; abc <= to; abc++ {
					for x := 0; x < base; x++ {
						value := BuildPalindrome(abc, x, palindrome_length, base)
						if !yield(value) {
							return
						}
					}
				}
			} else { // even
				for abc := from; abc <= to; abc++ {
					value := BuildPalindrome(abc, 0, palindrome_length, base)
					if !yield(value) {
						return
					}
				}
			}

			palindrome_length += 1
		}
	}
}

func BuildPalindrome(abc, x, length, base int) []uint8 {
	slots := length / 2 // intentional integer division, rounds down
	palindrome := make([]uint8, length)

	for index := 0; index < slots; index++ {
		palindrome[index] = uint8(abc / int(math.Pow(float64(base), float64(slots-index-1))) % base)
		palindrome[length-index-1] = palindrome[index]
	}

	if length%2 == 1 { // odd
		palindrome[slots] = uint8(x) // x goes in the middle
	}

	return palindrome
}

func PalindromeValue(p []uint8, base int) int {
	sum := 0
	for i := 0; i < len(p); i++ {
		sum += int(p[len(p)-i-1]) * int(math.Pow(float64(base), float64(i)))
	}

	return sum
}

func PalindromeSum_v4(limit int) (int, int) {
	sum := 0
	count := 0
	for palindrome := range PalindromeGenerator(22) {
		value := PalindromeValue(palindrome, 22)
		if value > limit {
			break
		}
		sum += value
		count += 1
	}

	return sum, count
}
