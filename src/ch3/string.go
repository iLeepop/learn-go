package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	// c := s[len(s)] // panic: index out of range
	// string[indexStart:indexEnd] 当indexStart被忽略时，默认为零 indexEnd被忽略则为len(string)
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[0:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	left := "left"
	temp := left
	left += ", right"

	fmt.Println(left)
	fmt.Println(temp)

	// s[0] = 'L' // compile error: cannot assign to s[0]

	for i, r := range "hello, 世界" {
		fmt.Println("%d\t%q\t%d\n", i, r, r)
	}

}
