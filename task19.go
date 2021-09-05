package main

import "fmt"

var justString string // на уровне пакета объявили justString

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // результат вычисления выражения (операция взятия подстроки) будет присвоен justString
	fmt.Println("cap of underlying array:", cap([]byte(justString)), "bytes")
	// если в кач-ве индекса последнего байта указать индекс, превышающий capacity нижележащего слайса, выход за пределы соответстующего нижележащего массива
	fmt.Println(justString)
}

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ { // в результирующей строке size раз будет повторено "A"

		v += "A" // Оператор "+" создает новую строку путем конкатенации двух строк. Может быть преобразовать исходную строку к слайсу рун, и добавлять каждый раз по руне?
	}
	return v // будет возвращена копия значения переменной v
}

func main() {
	someFunc()
	fmt.Println(justString)
}
