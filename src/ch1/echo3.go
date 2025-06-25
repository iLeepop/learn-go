package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	b := time.Now()
	echo()
	a := time.Now()
	fmt.Printf("Execution time: %v\n", a.Sub(b))
}
