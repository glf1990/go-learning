package main

import "fmt"

func datatype() {
	var a int32 = 1
	var b int64
	b = int64(a)
	fmt.Printf("test b : %d ", b)
	aptr := &a //a 的指针
	fmt.Print(a, aptr)
	var s string
	fmt.Printf("this is the test of default string : %shhhhh , it's length  is %d", s, len(s))
	fmt.Printf("\n\r")
}
func operation() {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3}
	d := [...]int{1, 2, 3, 4}
	fmt.Print(a == b)
	fmt.Print(a == d)
	fmt.Printf("\n\r")

}
func circle() {
	for i := 0; i < 5; i++ {
		if i > 2 {
			fmt.Printf("i is a big data")
			fmt.Printf("\n\r")
		}
		fmt.Print(i)
		fmt.Printf("\n\r")
	}
}
func playarray() {
	var a [3]int
	fmt.Println(a[0], a[1], a[2])
	for idx, e := range a {
		fmt.Println(idx, e)
	}
	b := [...]int{1, 2, 3, 4, 5, 6}
	c := b[3:]
	fmt.Println(c)
	s0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(s0), cap(s0))
	s1 := make([]int, 3, 5)
	s1 = append(s1, 1)
	fmt.Println(s1[0], s1[1], s1[2], s1[3])
	s2 := s0[3:6]
	s3 := s0[4:7]
	fmt.Println(s2, s3)
	s2 = append(s2, 1)
	fmt.Println(s0)
}
func mapplay() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for i, v := range m {
		fmt.Println(i, v)

	}
	fmt.Println(m["one"])
	m1 := map[int]int{}
	m1[2] = 3
	v, ok := m1[3]
	fmt.Println(v, ok)
	m2 := map[int]func(op int) int{}
	m2[1] = func(op int) int { return op }
	m2[2] = func(op int) int { return op * op }
	fmt.Println(m2[1](2), m2[2](2))

	myset := map[int]bool{}
	myset[1] = true
	n := 1
	if myset[n] {
		fmt.Printf("%d is exist\n\r", n)
	} else {
		fmt.Printf("%d is not exist\n\r", n)
	}

}
func abstring() {
	var s string
	fmt.Printf("default string is : %s \n\r", s)
	s = "hello"
	fmt.Printf("s len is : %d \n\r", len(s))
	s = "\xE4\xB8\xA5"
	fmt.Printf("byte s  print result is : %s \n\r", s)
	fmt.Printf("s len is : %d \n\r", len(s))
	s = "中"
	c := []rune(s)
	fmt.Printf("中unicod %x\n\r", c[0])
	fmt.Printf("中 UTF8 %x\n\r", s)
}
func main() {
	datatype()
	operation()
	circle()
	playarray()
	mapplay()
	abstring()
}
