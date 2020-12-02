package main

import "fmt"

func main() {
	var a struct{}
	b := make([]int, 0)
	c := []int{}
	fmt.Println(&a == nil, &b == nil, &c == nil)
}
