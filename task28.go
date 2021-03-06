package main

import "fmt"

// паттерн "адаптер" позволяет подружить несовместимые объекты.
// Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому
// Чтобы адаптировать архитектуру, рассчитанную на один интерфейс, для типа, поддерживающего другой интерфейс
// мы создадим структуру-адаптер, которая будет:

// 1. Поддерживать тот же интерфейс, который ожидает Human
// 2. Переводить запрос от Human к адаптируемому объекту в форме, которую он ожидает

// человек может открутить гайку, так определено в сигнатуре его метода unscrew(): метод принимает на вход значения типа "Washer", являющегося ИНТЕРФЕЙСОМ,
// это значит, что unscrew() может принять любой тип, который имеет метод unscrew

type Human struct { // есть human, он имеет метод unscrew()
}

func (h *Human) unscrew(ws IntfWasherBy4) { // чел-к может открутить только то, что поддерживается интерфейсом IntfWasherBy4 (то, что МОЖЕТ откручиваться условным "гайковертом на 4")
	fmt.Println("Human unscrew the washer")
	ws.unscrewBy4()
}

//////////////////////////
type IntfWasherBy4 interface {
	unscrewBy4()
}

//////////////////////////
// декларируем тип гаек на 4
type WasherBy4 struct { // для этой гайки не нужен адаптер, она поддерживает интерфейс IntfWasherBy4 - имеет метод unscrewBy4() (её можно открутить Human'у его "гайковертом на 4")
	isScrewed bool
}

func (w *WasherBy4) unscrewBy4() {
	w.isScrewed = false // при вызове метода гайка становится открученной
	fmt.Println("Washer by 4 is unscrewed")
}

/////////////////////////////////////////////////////////////////////////
type WasherBy12 struct { // гайка на 12 (Human уже не сможет её открутить гайковертом на 4, нету метода unscrewBy4() => не поддерживается интерфейс IntfWasherBy4))
	isScrewed bool
}

func (w *WasherBy12) unscrewBy12() { // под эту гайку нужна головка на 12
	w.isScrewed = false
	fmt.Println("Washer by 12 is unscrewed")
}

/////////////////////////////////////////////////////////////////////////
type Adapter12To4 struct { // адаптер для гайки на 12 (тип-адаптер для неподдерживающегося Human'ом типа WasherBy12). Human примет этот тип, т к он имеет метод unscrewBy4()
	washer12 *WasherBy12 // КЛЮЧЕВОЙ МОМЕНТ! В себе имеет ссылку на экземпляр значения типа WasherBy12. В этом обьекте будет как-бы "запрятан" обьект, который Human не принимает
}

func (wa *Adapter12To4) unscrewBy4() {
	fmt.Println("Adapted")
	wa.washer12.unscrewBy12() // переводим запрос к адаптируемому обьекту
}

////////////////////////////////////////////////////////////////////////
func main() {
	human := &Human{}      // создали экземпляр Human
	bw := &WasherBy4{true} // создали экземпляр WasherBy4

	human.unscrew(bw) // открутили поддерживающуюся Human'ом гайку

	w12 := &WasherBy12{true}      // создали неподдерживающуюся Human'ом гайку
	adapter := &Adapter12To4{w12} // передали её адаптеру

	human.unscrew(adapter) // открутили неподдерживающуюся Human'ом гайку через адаптер (добрались до поля неподдерживаемого Human'ом типа и изменили его значение)

}
