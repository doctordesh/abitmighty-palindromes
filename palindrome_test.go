package main

import "testing"

func Benchmark_PalindromeSum_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeSum_v1(1000)
	}
}

func Benchmark_PalindromeSum_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeSum_v2(1000)
	}
}

func Benchmark_PalindromeSum_v3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeSum_v3(1000)
	}
}

func Benchmark_Reverse_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString_v1("lorem ipsum dolor sit amet")
	}
}

func Benchmark_Reverse_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString_v2("lorem ipsum dolor sit amet")
	}
}

func Benchmark_NumberToBase_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberToBase_v1(1_000_000, 22)
	}
}

func Benchmark_NumberToBase_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberToBase_v2(1_000_000, 22)
	}
}

func TestEquality(t *testing.T) {
	limit := 5468885
	n1, c1 := PalindromeSum_v1(limit)
	n2, c2 := PalindromeSum_v2(limit)
	n3, c3 := PalindromeSum_v3(limit)

	if n1 != n2 || n1 != n3 || c1 != c2 || c1 != c3 {
		t.FailNow()
	}
}

func TestIsPalindrome(t *testing.T) {
	type row struct {
		In  string
		Exp bool
	}

	table := []row{
		{"abc", false},
		{"a", true},
		{"abcba", true},
		{"abba", true},
		{"foo", false},
		{"ccc", true},
	}

	for _, r := range table {
		p := IsPalindrome(r.In)
		t.Log(p, r.Exp)
		if p != r.Exp {
			t.Fail()
		}
	}
}

func Benchmark_IsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("aaaaabbbbbcccccdddddcccccbbbbbaaaaa")
	}
}
