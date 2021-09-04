package main

import (
	"fmt"
	"math"
)

func main() {
	var x int64 // переменная, значение которой будем менять

	x = 255
	fmt.Println("x value before modification (decimal):", x)
	bitOverwrite(&x, 15, 1)
	fmt.Printf("x value after modification (binary): %b\nx value after modification (decimal): %d\n", x, x)

	fmt.Println("--------------------------")
	x = 255
	fmt.Println("x value before modification (decimal):", x)
	bitOverwrite(&x, 4, 0)
	fmt.Printf("x value after modification (decimal): %d", x)

}

func bitOverwrite(x *int64, i, v int) { // Функция имеет параметры: 1:указатель на исходное число (возвращать ничего не будем, модифицируем исходное значение), 2:индекс изменяемого бита (индексация справа налево), 3:на какое значение меняем
	if v == 0 { // рассматривается 3 случая: если i-тый бит надо поменять на 0
		mask := int64(math.MaxInt64 - 1<<i) // смещаем 1-чку на i-ю позицию, используем "-" c операндом math.MaxInt64 слева, чтобы получить инверсию результата выражения 1<<i
		//fmt.Println(1 << i)
		fmt.Printf("mask value: %b\nx value before modification: %b\n", mask, *x)
		*x = *x & mask // выполняем операцию побитового И
		fmt.Printf("x value after modification: %b\n", *x)

	} else if v == 1 { // если i-тый бит надо поменять на 1
		fmt.Printf("x value before modificaton %b\n", *x)
		*x = *x | 1<<i // Справа от | мы берём единицу и с помощью сдвига вправо (умножение на 2^i) переставляем её на нужную позицию.
		// Далее выполняем операцию побитового ИЛИ над нашим значением и результатом сдвига
	} else { // если в аргументе передано что-то отличное от 0 и 1.
		fmt.Println("Error")
		return // Здесь можно было бы вернуть значение ошибки (как положено), но мы конечно же так делать не будем
	}
}
