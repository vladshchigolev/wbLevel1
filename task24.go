package main

import "fmt"

func main() {
	arr := make([]int, 100) // len == cap
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
