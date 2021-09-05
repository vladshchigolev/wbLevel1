package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	fakeSleep(3)
	fmt.Println("End")
}

func fakeSleep(seconds int) { // время в секундах, на которое надо прервать дальнейшее выполнение (хотя на самом деле исполнение горутины не прервется)
	startTime := time.Now().Second() // возвращает текущее время (на момент вызова fakeSleep)
	fmt.Println(startTime)
	for time.Now().Second()-startTime < seconds { // цикл будет крутиться до тех пор, пока разница между временем начала итерации и временем вызова fakeSleep не станет равной
		// промежутку времени, на который была запущена fakeSleep
	}
}
