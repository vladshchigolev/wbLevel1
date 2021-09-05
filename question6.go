package main

import "fmt"

func main() {
	var a []int
	b := []int{} // GoLand советует не использовать определение пустого слайса, используя литерал

	fmt.Println(a, len(a), cap(a)) // длина и ёмкость одинаковы
	fmt.Println(b, len(b), cap(b))
}
