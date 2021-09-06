package main

import "fmt"

// Иногда конкретный тип значения не важен. Вас не интересует,
// с чем вы работаете. Вы просто хотите быть уверены в том, что оно может
// делать то, что нужно вам. Тогда вы сможете вызывать для значения определенные методы.
// Интерфейсы в Go позволяют определять переменные и
// параметры функций, способные хранить ЛЮБОЙ тип при условии, что этот
// тип определяет некоторые методы (поддерживает интерфейс)
// интерфейсы в go декларируют полу-абстрактный тип (экземпляры этого типа всё же ВОЗМОЖНО создать)

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() { // Кроме прочего, есть ещё метод Record
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) { // параметр device, имеет тип интерфейса, => принимает значения любого типа, поддерживающего его
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	songs := []string{"Alberto Balsalm - Aphex Twin", "If Loving You Is Wrong - Against All Logic", "Take Me Home, Country Roads - John Denver"}
	sonyWalkmanFX10 := TapePlayer{Batteries: "Full"}
	playList(sonyWalkmanFX10, songs)
	funai31B := TapeRecorder{}
	playList(funai31B, songs)

}
