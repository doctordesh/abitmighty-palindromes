package main

import (
	"testing"

	"github.com/doctordesh/check"
)

// func Benchmark_PalindromeSum_v1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PalindromeSum_v1(1000)
// 	}
// }

// func Benchmark_PalindromeSum_v2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PalindromeSum_v2(1000)
// 	}
// }

// func Benchmark_PalindromeSum_v3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PalindromeSum_v3(1000)
// 	}
// }

// func Benchmark_Reverse_v1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ReverseString_v1("lorem ipsum dolor sit amet")
// 	}
// }

// func Benchmark_Reverse_v2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ReverseString_v2("lorem ipsum dolor sit amet")
// 	}
// }

// func Benchmark_NumberToBase_v1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		NumberToBase_v1(1_000_000, 22)
// 	}
// }

// func Benchmark_NumberToBase_v2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		NumberToBase_v2(1_000_000, 22)
// 	}
// }

// func TestEquality(t *testing.T) {
// 	limit := 5468885
// 	n1, c1 := PalindromeSum_v1(limit)
// 	n2, c2 := PalindromeSum_v2(limit)
// 	n3, c3 := PalindromeSum_v3(limit)

// 	if n1 != n2 || n1 != n3 || c1 != c2 || c1 != c3 {
// 		t.FailNow()
// 	}
// }

// func TestIsPalindrome(t *testing.T) {
// 	type row struct {
// 		In  string
// 		Exp bool
// 	}

// 	table := []row{
// 		{"abc", false},
// 		{"a", true},
// 		{"abcba", true},
// 		{"abba", true},
// 		{"foo", false},
// 		{"ccc", true},
// 	}

// 	for _, r := range table {
// 		p := IsPalindrome(r.In)
// 		t.Log(p, r.Exp)
// 		if p != r.Exp {
// 			t.Fail()
// 		}
// 	}
// }

// func Benchmark_IsPalindrome(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		IsPalindrome("aaaaabbbbbcccccdddddcccccbbbbbaaaaa")
// 	}
// }

// func TestVersion4(t *testing.T) {
// 	fmt.Println(PalindromesOfSize(3, 22))
// 	fmt.Println(PalindromesOfSize(4, 22))
// 	fmt.Println(PalindromesOfSize(5, 22))
// }

// func TestLoops(t *testing.T) {
// 	t.Log(PalindromesOfSize(3, 22))
// }

// func TestCompute(t *testing.T) {
// 	type row struct {
// 		input  []byte
// 		base   int
// 		output int
// 	}
// 	table := []row{
// 		{[]byte{3, 3, 3}, 4, 63},
// 		{[]byte{21, 21, 21}, 22, 10647},
// 		{[]byte{1, 0, 0}, 10, 100},
// 		{[]byte{1, 0, 0, 0}, 10, 1000},
// 		{[]byte{1, 0, 0, 0}, 22, 10648},
// 		{[]byte{1, 0, 0, 0}, 2, 8},
// 	}

// 	for _, row := range table {
// 		check.Equals(t, row.output, ComputeValue(row.input, row.base))
// 	}
// }

func TestPal(t *testing.T) {
	check.Equals(t, 16, NumberOfPalindromes(3, 4))
	check.Equals(t, 25, NumberOfPalindromes(3, 5))
	check.Equals(t, 125, NumberOfPalindromes(5, 5))
}

func TestPall(t *testing.T) {
	// t.Log(Palindromes(3, 4))
	// t.Log(Palindromes(3, 5))
	// t.Log(Palindromes(4, 5))
	// t.Log(Palindromes(5, 5))

	b := []byte{0}
	// first digit (odd)
	t.Log(b)
	b[0] = 1
	t.Log(b)
	b[0] = 2
	t.Log(b)
	// 'base' number of loops

	// second digit (even)
	b = append(b, 0)
	b[0] = 0
	t.Log(b)
	b[0] = 1
	b[1] = 1
	t.Log(b)
	b[0] = 2
	b[1] = 2
	t.Log(b)

	// third digit (odd)
	b = append(b, 0)
	b[0] = 0
	b[1] = 0
	t.Log(b)

}

// func BenchmarkLoopsRolled(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		h()
// 	}
// }

// func BenchmarkLoopsUnrolled(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		k()
// 	}
// }
