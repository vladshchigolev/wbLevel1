package main

import "fmt"

func update(pointer *int) { // функция update в качестве параметра принимает значение типа "указатель на int"
	// НО ЗДЕСЬ ВАЖНЫЙ МОМЕНТ!
	// В Go используется ПЕРЕДАЧА ПО ЗНАЧЕНИЮ, это значит, что при вызове update()
	// переданный аргумент СКОПИРУЕТСЯ в параметр, то есть мы внутри функции будем иметь
	// локальную переменную - указатель, ссылающийся на тот же объект, что и передававшийся
	// оригинал, но изменения, совершенные над этим значением, никак не затронут оригинал
	b := 2 // Объявляем локальную переменную с типом int
	pointer = &b // справа получаем указатель на b (*int), присваиваем pointer'у, являющемуся копией "p"
}

func main() {
	var (
		a = 1
		p = &a // получаем указатель на переменную "a", тип значения p - *int
	)
	fmt.Println(*p) // разыменовываем ссылку (равнозначно fmt.Println(a))
	update(p)
	fmt.Println(*p) // разыменовываем снова и убеждаемся, что p всё так же ссылается на "a"
}
