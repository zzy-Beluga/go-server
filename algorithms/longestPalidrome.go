package algorithms

import "fmt"

func LongestPalindrome(s string) string {
	l := len(s)
	fmt.Println(l)
	p := make([][]bool, l)
	for i := 0; i < l; i++ {
		p[i] = make([]bool, l)
	}
	a := 0
	b := 0
	for j := 0; j < l; j++ {
		p[j][j] = true
		for i := 0; i < j; i++ {
			if j == i+1 && s[i] == s[j] {
				p[i][j] = true
				fmt.Println(p[i][j], i, j)
			} else if j > i+1 {
				p[i][j] = s[i] == s[j] && p[i+1][j-1]
				fmt.Println(p[i][j], i, j)
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if p[i][j] && j-i > b-a {
				a = i
				b = j
				fmt.Println(a, b)
			}
		}
	}

	return s[a : b+1]
}
