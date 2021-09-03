package main

import (
	"fmt"
	"math"
)

func main() {
	var x int64

	x = 255
	fmt.Println("x =", x)
	writeToBit(&x, 1, 0)
	fmt.Println("Записали в 0-ой бит 1, x =", x) // 10 -> 11

	//x = 3
	//fmt.Println("x =", x)
	//writeToBit(&x, 2, 1)
	//fmt.Println("Записали в 2-ой бит 1, x =", x) // 011 -> 111
	//
	//x = 6
	//fmt.Println("x =", x)
	//writeToBit(&x, 1, 0)
	//fmt.Println("Записали в 1-ой бит 0, x =", x) // 110 -> 100

}

// число, в какой бит записываем значение, значение
func writeToBit(x *int64, i, v int) { // Функция имеет параметры: 1:указатель на исходное число (возвращать ничего не будем, модифицируем исходное значение), 2:индекс изменяемого бита (индексация справа налево), 3:на какое значение меняем
	if v == 0 { // рассматривается 3 случая: если i-тый бит надо поменять на 0
		temp := int64(math.MaxInt64 - 1<<i)
		fmt.Printf("temp value: %b\nx value: %b\n", temp, *x)
		*x = *x & temp

	} else if v == 1 { // если i-тый бит надо поменять на 1
		*x = *x | 1<<i // Справа от | мы берём единицу и с помощью сдвига вправо (умножение на 2^i) переставляем её на нужную позицию.
		// Далее выполняем операцию побитового ИЛИ над нашим значением и результатом сдвига
	} else { // если в аргументе передано что-то отличное от 0 и 1.
		fmt.Println("Error")
		return // Здесь можно было бы вернуть значение ошибки (как положено), но мы конечно же так делать не будем
	}
}