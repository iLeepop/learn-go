package main

import "fmt"

var a [3]int

// 初始化字面值
var q [3]int = [3]int{1, 2, 3}

func main() {
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 使用...代替具体长度 使用字面量初始化
	p := [...]int{1, 2, 3, 4, 5}

	for _, v := range p {
		fmt.Printf("%d\n", v)
	}
}
