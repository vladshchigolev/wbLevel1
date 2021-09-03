package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan bool) // Имеется "сигнальный" канал (можем сделать его с буфером 1, чтобы не было блокировок main-горутины,
	// в случае, если выполнение секции default у select'a функции timer() занимает много времени)
	go timer(stopCh) // Передаем его запускаемой горутине
	time.Sleep(5 * time.Second) // Горутина работает столько, сколько нам надо (горутина main спит)
	stopCh <- true // Проснувшись, горутина main отправляет сигнал по каналу о необходимости завершения "timer"
	// т к вызванная горутина итеративно постоянно проверяет содержимое канала
	ctx := context.Background()
	go timer2(ctx)
	time.Sleep(6 * time.Second)
	ctx.Done()
}

func timer(stopCh chan bool) {
	for {
		select {
		case <-stopCh: // разблокировка этого case произойдет только если станет возможна операция чтения из stopCh
			return // происходит возврат из функции
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Second is gone...")
		}
	}
}

func timer2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("2 seconds is gone...")
		}
	}
}