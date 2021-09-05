package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int64
	var b int64
	a = 1 << 36 // операция сдвига эквивалентна уножению левого операнда на 2^36 (1 * 2^36)
	b = 1 << 42 // операция сдвига эквивалентна уножению левого операнда на 2^42 (1 * 2^42)
	// a и b в int64 влезают, но вот выйдут ли результаты операций + и * над ними за пределы int64 - неизвестно
	fmt.Println("a:", a, "\nb:", b) // выведем их

	add(a, b)
	sub(a, b)
	mul(a, b)
	div(a, b)
}

func add(a, b int64) {
	bigA := big.NewInt(a) // NewInt принимает значение int64 и возвращает значение типа big.Int
	bigB := big.NewInt(b)
	var result big.Int                  // создаем переменную типа big.Int для хранения результата
	fmt.Println(result.Add(bigA, bigB)) // у result вызываем метод Add()
}

func sub(a, b int64) {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	var result big.Int
	fmt.Println(result.Sub(bigA, bigB))
}

func mul(a, b int64) {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	var result big.Int
	fmt.Println(result.Mul(bigA, bigB))
}

func div(a, b int64) {
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	var result big.Int
	fmt.Println(result.Div(bigA, bigB))
}
