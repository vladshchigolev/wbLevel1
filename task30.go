package main

import (
	"fmt"
	"math/rand"
)

// 1. последний элемент среза скопировали на место удаляемого
// 2. последний элем среза (в конце) удалили
func main() {
	n := 10                  // кол-во элементов в срезе (len) будет 20
	arr := make([]int, n)    // создаем его
	for i := 0; i < n; i++ { // набиваем значениями (перезаписываем нулевые)
		arr[i] = rand.Int() % 100
	}
	fmt.Println("исходный срез:", arr)
	arr = deleteElem(arr, 6)
	fmt.Println("результи срез:", arr)
	//arr = deleteElem2(arr, 0)
	//fmt.Println(arr)
}

func deleteElem(arr []int, index int) []int { //
	arr[index] = arr[len(arr)-1] // перезаписываем элемент, который требуется удалить, значением последнего элемента в срезе
	fmt.Println("промежут срез:", arr)
	return arr[:len(arr)-1] // будет создан новый срез,
}
