package main

func NumberToBase_v1(n, base int) string {
	var digs = "0123456789abcdefghijklmnopqrstuvwxyz"
	if n == 0 {
		return string(digs[0])
	}

	digits := []byte{}
	for {
		index := n % base
		digits = append([]byte{digs[index]}, digits...)
		n = n / base // integer division

		if n == 0 {
			break
		}
	}

	return string(digits)
}

func ReverseString_v1(s string) string {
	b := []byte(s)
	out := []byte{}
	for i := len(b) - 1; i >= 0; i-- {
		out = append(out, b[i])
	}

	return string(out)
}

func PalindromeSum_v1(limit int) (int, int) {
	count := 0
	sum := 0
	for i := 0; i <= limit; i++ {
		base22 := NumberToBase_v1(i, 22)

		// Check for palindrome
		base22reversed := ReverseString_v1(base22)
		if base22 == base22reversed {
			count += 1
			sum += i
		}
	}

	return sum, count
}
