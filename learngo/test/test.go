package main

import "fmt"

func f(s []int) {
	s = append(s, 0)
	fmt.Printf("1%#v\n", s)
	for i := range s {
		s[i]++
	}
	fmt.Printf("2%#v\n", s)

}

func main() {

	s1 := []int{1, 2}
	f(s1)
	fmt.Printf("%#v\n", s1)
	fmt.Println("==========")
	s2 := s1
	s2 = append(s2, 3)
	fmt.Printf("%#v\n", s2)
	f(s2)
	fmt.Printf("%#v\n", s2)
	//fmt.Printf("%#v\n", s1)
	//fmt.Printf("%#v\n", s2)
	//s2 = append(s2, 3)
	//fmt.Printf("%#v\n", s1)
	//fmt.Printf("%#v\n", s2)
	//f(s1)
	//f(s2)
	//fmt.Printf("%#v\n", s1)
	//fmt.Printf("%#v\n", s2)
	//f(s1)
	//fmt.Printf("%#v\n", s1)
	//f(s1)
	//fmt.Printf("%#v\n", s1)

}
