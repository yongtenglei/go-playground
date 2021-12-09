package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := len(s); i+1 < j; j-- {
			if isPalindrome(s[i:j]) {
				return s[i:j]
			}
		}
	}
	return string(s[0])
}

func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "cbbd"
	fmt.Println(isPalindrome(s[0:1]))
	fmt.Println(longestPalindrome(s))
}
