package main

import (
	"fmt"
	"reflect"
)

// Можно реализовать это с помощью ВЛОЖЕНИЯ и ВСТРАИВАНИЯ.

type Gender string

type Nationality string

type Human struct {
	firstName string
	lastName string
	age uint16
	Gender // Такие поля называют анонимными
	Nationality
}

type Action struct { // Такие типы, как Action, создаются путем композиции одного или нескольких полей, каждое из которых предоставляет несколько полей и/или методов
	Human // Мы встраиваем Human для предоставления полей firstName, lastName, age, Gender и Nationality,
		  // а так же методов, связанных с типом Human
}

func (h *Human) chAge(newAge uint16) {
	h.age = newAge
}

func (a *Action) chName(newFstName, newLstName string) {
	// Именованные типы и указатели на них - единственные типы,
	// которые могут появляться в объявлении получателя
	a.firstName = newFstName // Встраивание позволяет нам использовать синтаксические сокращения
	a.lastName = newLstName
}

func main()  {
	actionInstance := Action{Human{firstName: "Ivan", lastName: "Smirnov", age: 48, Gender: "Male", Nationality: "Russian"}}
	fmt.Println(reflect.TypeOf(actionInstance))
	fmt.Println(actionInstance)
	actionInstance.chName("Sergey", "Petrov")
	fmt.Println(actionInstance)
	actionInstance.chAge(33) // Методы Human доступны для экземпляра Action
	fmt.Println(actionInstance)
}
// Не стоит рассматривать Human как базовый класс, а Action - как производный класс,
// или интерпретировать связь между этими типами, как если бы Action "являлся" Human