//package palindromenumber
package main

import (
	"fmt"
	"strconv"
)

// Method1: covert int to string
func isPalindrome1(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.FormatInt(int64(x), 10)
	i := 0
	j := len(s)
	for i <= (i+j)/2 {
		if s[i] != s[j-1] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome(x int) bool {

	s := strconv.FormatInt(int64(x), 10)
	slen := len(s)
	for i, j := 0, slen-1; i < slen/2; i, j = i+1, j-1 {
		if s[i] != s[j-1] {
			return false
		}
	}
	return true
}

// Method 2: without covert int to string
func isPalindrome2(x int) bool {

	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}

	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func main() {
	//x := 123
	x := 123454321
	fmt.Println(isPalindrome1(x))
	fmt.Println(isPalindrome2(x))
}
