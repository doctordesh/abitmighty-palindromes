package main

func IsPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func PalindromeSum_v3(limit int) (int, int) {
	count := 0
	sum := 0
	for i := 0; i <= limit; i++ {
		base22 := NumberToBase_v2(i, 22)

		// Check for palindrome
		if IsPalindrome(base22) {
			count += 1
			sum += i
		}
	}

	return sum, count
}
