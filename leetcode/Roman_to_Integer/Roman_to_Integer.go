package main

var RomanValueMap = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt(s string) int {
	sum := 0
	for i := range s {
		value := RomanValueMap[s[i]]
		if i != len(s)-1 && value < RomanValueMap[s[i+1]] {
			sum -= value
		} else {
			sum += value
		}
	}

	return sum
}
