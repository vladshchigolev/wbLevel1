package main

import "fmt"

func main() {
	sliceA := make([]int, 3)    // capacity == length (make создает неименованную переменную массива и возвращает его срез
	sliceB := make([]int, 3, 4) // сам массив доступен только через возвращаемый срез)
	sliceC := []int{1, 2, 3}    // С помощью литерала, инициализация значениями
	var sliceD = []int{}        // пустой срез
	var sliceE []int            // пустой срез

	fmt.Println(sliceA, len(sliceA), cap(sliceA))
	fmt.Println(sliceB, len(sliceB), cap(sliceB))
	fmt.Println(sliceC, len(sliceC), cap(sliceC))
	fmt.Println(sliceD, len(sliceD), cap(sliceD))
	fmt.Println(sliceE, len(sliceE), cap(sliceE))

	mapA := make(map[int]int)
	mapB := make(map[int]int, 4)
	var mapC = map[int]int{}
	mapD := map[int]int{}
	fmt.Println(mapA, len(mapA))
	fmt.Println(mapB, len(mapB))
	fmt.Println(mapC, len(mapC))
	fmt.Println(mapD, len(mapD))

}
