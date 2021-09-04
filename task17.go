package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 42
	myStr := "Hello, Go!"
	myBool := true
	myCh := make(chan int, 5)
	showType(myInt)
	showType(myStr)
	showType(myBool)
	showType(myCh)
}

// интерфейс создан, чтобы поддерживаться несколькими типами
// параметр data примет значение любого типа, т к его тип - пустой интерфейс
func showType(val interface{}) { // пустой интерфейс, принимает ВСЁ, т. к. в нём НЕ ОПРЕДЕЛЕНО МЕТОДОВ, которые необходимо иметь ТИПУ для того, чтобы поддерживать его
	fmt.Println("Value:", val, "\nType:", reflect.TypeOf(val))
}
