package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var myMap = map[int]int{}
	mt := sync.Mutex{}        // mt охраняет myMap
	for i := 0; i <= 1; i++ { // запускаем 20 воркеров
		go writeToMapRand(myMap, &mt) // передаем каждому воркеру карту, указатель на мьютекс
	}
	time.Sleep(4 * time.Second) // main-горутина засыпает на некоторое время, чтобы позволить воркерам отработать
	mt.Lock()                   //
	fmt.Println(myMap)          // main может обратиться к общей переменной, когда ею занимается другая горутина (даже после ожидания main-горутиной 4-х секунд),
	// так что пусть main придерживается общих правил, и вызовет Lock перед доступом, так мы будем уверены, что 19-я строка не выполнится, пока другой горутиной не будет вызван Unlock
	mt.Unlock()

}

func writeToMapRand(aMap map[int]int, mt *sync.Mutex) {
	mt.Lock()
	aMap[rand.Int()%1000] = rand.Int() % 100 // каждый воркер должен писать в карту с уникальным ключом
	mt.Unlock()
}
