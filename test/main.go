package main

import "fmt"

//type SwapAdd func() int

//func GenerateRandom() SwapAdd {
//a, b := rand.Intn(10), rand.Intn(20)
//fmt.Println("=================")
//fmt.Printf("generate a = %d, b = %d\n", a, b)
//fmt.Println("=================")
//return func() int {
//a, b = b, a+b
//return a
//}
//}

//func (s SwapAdd) Read(p []byte) (n int, err error) {
//next := s()
//if next > 20 {
//return 0, io.EOF
//}

//fmt.Println("************")
//fmt.Println("next = ", next)
//fmt.Println("************")

//nextStr := strconv.Itoa(next)
//reader := strings.NewReader(nextStr)
//return reader.Read(p)
//}

//func PrintRes(r io.Reader) {
//scanner := bufio.NewScanner(r)
//for scanner.Scan() {
//fmt.Println("###############")
//fmt.Println(scanner.Text())
//fmt.Println("###############")
//}

//}

//func main() {
//s := GenerateRandom()
//PrintRes(s)

//}

type Struct1 struct {
	Struct2 []struct {
		A string
		B string
		C []string
	}
}

func main() {
	a := Struct1{
		Struct2: []struct {
			A string
			B string
			C []string
		}{
			{A: "a1", B: "b1", C: []string{"c1", "c1"}},
			{A: "a2", B: "b2", C: []string{"c2", "c2"}},
			{A: "a3", B: "b3", C: []string{"c3", "c3"}},
		},
	}

	fmt.Println(a.Struct2)
}
