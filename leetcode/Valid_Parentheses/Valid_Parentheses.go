package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	if len(s) == 0 {
		return true
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' || v == '}' || v == ']' {

			remained := len(stack)
			if remained == 0 {
				return false
			}

			last := stack[remained-1]

			switch v {
			case ')':
				if last != '(' {
					return false
				}
			case '}':
				if last != '{' {
					return false
				}
			case ']':
				if last != '[' {
					return false
				}
			}

			stack = stack[:remained-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()"
	s1 := "()[]{}"
	s2 := "(]"
	s3 := "([)]"
	s4 := "{[]}"
	fmt.Println(isValid(s))
	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
	fmt.Println(isValid(s4))
}
