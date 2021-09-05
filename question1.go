package main

import (
	"fmt"
	"strings"
)

func main() {
	fstStr := "qpwoeiruty"
	sndStr := "alskdjfhg"
	var resStr string

	// strings.Builder имеет метод Grow(). Если необходимо, Grow увеличит емкость буфера нашего значения builder,
	// чтобы гарантировать место для еще n байтов, после чего в builder можно записать не менее n байтов без
	// дополнительного выделения памяти
	var builder strings.Builder // предназначен для конструирования строки
	// fmt.Println(builder.buf) // приватное поле
	builder.Grow(len(fstStr) + len(sndStr)) // Выделяем память
	builder.WriteString(fstStr)             // добавляет содержимое fstStr к буферу (поле buf) builder
	builder.WriteString(sndStr)             // аналогично
	resStr = builder.String()               // метод возвращает собранное строковое значение
	fmt.Println(fstStr, sndStr, resStr)
}
