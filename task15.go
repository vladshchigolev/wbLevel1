package main

import "fmt"

func main() {
	var a = 7
	var b = 4
	// необходимо произвести обмен значениями
	fmt.Println("before swapping", a, b)
	a = a + b // перезаписываем а результатом сложения a и b (11) (сначала будет вычислено выражение, затем выполнена операция присваивания)
	b = a - b // перезаписываем b результатом вычитания b из a (b станет тем, чем изначально была a(7))
	a = a - b // перезаписываем b результатом вычитания b из a (a станет тем, чем изначально была b(4))
	fmt.Println("after swapping", a, b)
}
