package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct { // структура-счетчик
	count int
	mt    sync.Mutex // для защиты count
}

func (c *Counter) increaser() {
	c.mt.Lock()
	c.count++
	c.mt.Unlock()
}

func main() {
	var c = Counter{ // создаем экземпляр структуры counter
		count: 0,
		mt:    sync.Mutex{},
	}
	c.mt.Lock()
	fmt.Println(c.count)
	c.mt.Unlock()
	for i := 0; i < 25; i++ {
		go c.increaser() // 25 раз вызываем метод в новых горутинах
	}
	time.Sleep(3 * time.Second)
	c.mt.Lock() // в 7-ой задачке объяснял, зачем здесь Lock/Unlock
	fmt.Println(c.count)
	c.mt.Unlock()
}
