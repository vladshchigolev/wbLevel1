package main

import "fmt"

func main() {
	slice := []string{"a", "a"} // объявляем срез строковых значений с помощью литерала
	fmt.Println("емкость среза в переменной slice:", cap(slice))
	func(slice []string) {
		slice = append(slice, "a") // slice теперь ссылается на новый базовый массив (ёмкость выросла в 2 раза)
		fmt.Println("емкость среза в переменной slice после вызова append:", cap(slice))
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice) // та же ситуация, что и в задаче 18
}
