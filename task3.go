package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	res := make(chan int, len(arr))
	var sum int

	for _, val := range arr {
		go squarer(val, res)
	}
	for i := 0; i < len(arr); i++ {
		sum += <-res // Операция чтения (результата вычисления квадрата элемента массива) из канала даст нам целочисленный результат, который мы прибавим к sum
	}
	fmt.Println("Sum of squares:", sum)
	fmt.Println("---------------")
	var wg sync.WaitGroup
	for _, v := range arr {
		wg.Add(1)
		go square2(v, &wg)
	}
	wg.Wait()

}

func squarer(origVal int, result chan int) {
	result <- origVal * origVal
}

func square2(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(x * x)
}