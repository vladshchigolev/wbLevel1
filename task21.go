package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // WaitGroup объяснял во 2-ом задании
	var arr = []int{60, 54, 72, 38, 78, 13, 31, 33, 89, 20}
	for i := range arr {
		wg.Add(1)
		go writer(arr, i, &wg) // каждому воркеру дадим по значению из массива, чтобы выводил в stdout
	}
	wg.Wait()
}

func writer(arr []int, index int, wg *sync.WaitGroup) {
	fmt.Println(arr[index])
	wg.Done()
}
