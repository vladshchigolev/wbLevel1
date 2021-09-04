package main

import "fmt"

func main() {
	var arr1 = [...]int{7, 15, 26, 3, -12, 16, 31}
	var arr2 = [...]int{40, 26, 5, -12, 16, 34, 22}

	var intersection []int    // здесь будем хранить пересечение массивов
	for _, v1 := range arr1 { // этот цикл перебирает первый слайс
		for _, v2 := range arr2 { // сравниваем КАЖДЫЙ элемент 1-го массива с КАЖДЫМ элементом 2-го. (например сначала 7 сравним со всеми эл-тами 2-го, потом 15, и т д)
			if v1 == v2 { // если значение из 1-го совпадает с каким-то значением из 2-го, добавляем его к слайсу с пересечением
				intersection = append(intersection, v1) // перезаписываем result новым слайсом, который вернула ф-ция append()
			}
		}
	}
	fmt.Println("пересечение множеств:", intersection)

	var count = make(map[int]int)
	for _, v1 := range arr1 {
		if count[v1] != 1 {
			count[v1]++
		}
	}
	for _, v2 := range arr2 {
		if count[v2] == 1 {
			count[v2]++
		}
	}
	var res2 []int
	for k, v := range count {
		if v == 2 {
			res2 = append(res2, k)
		}
	}
	fmt.Println(res2)
}
