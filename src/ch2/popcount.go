package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var c byte
	for i := range 8 {
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}

func main() {
	count := PopCount(12345)
	fmt.Printf("PopCount = %d\n", count)
}
