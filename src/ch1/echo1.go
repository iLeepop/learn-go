package main

import (
	"fmt"
	"os"
	"time"
)

func echo() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Index of Args: ", i, "Value: ", os.Args[i])
	}
}

func main() {
	b := time.Now()
	echo()
	a := time.Now()
	fmt.Printf("Execution time: %v\n", a.Sub(b))
}
