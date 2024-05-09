package main

import (
	"math"
	"strings"
)

var digs = "0123456789abcdefghijklmnopqrstuvwxyz"

//	func DigitsInBase(x, base int) int {
//		return int(math.Floor(
//			math.Log(float64(x))/math.Log(float64(base)),
//		)) + 2 // Yes, a 2. It's correct, please continue reading.
//	}
func DigitsInBase(x, base int) int {
	return int(math.Floor(
		math.Log(float64(x))/math.Log(float64(base)),
	)) + 2
}

func NumberToBase_v2(n, base int) string {
	if n == 0 {
		return string(digs[0])
	}

	digits := make([]byte, DigitsInBase(n, base))
	i := 0
	for {
		index := n % base
		digits[len(digits)-1-i] = digs[index]
		n = n / base // integer division

		if n == 0 {
			break
		}

		i = i + 1
	}

	// Because of the potential extra null byte from the 'DigitsInBase'
	return strings.Trim(string(digits), "\x00")
}

func ReverseString_v2(s string) string {
	b := []byte(s)
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}

	return string(b)
}

func PalindromeSum_v2(limit int) (int, int) {
	count := 0
	sum := 0
	for i := 0; i <= limit; i++ {
		base22 := NumberToBase_v2(i, 22)

		// Check for palindrome
		base22reversed := ReverseString_v2(base22)
		if base22 == base22reversed {
			count += 1
			sum += i
		}
	}

	return sum, count
}
