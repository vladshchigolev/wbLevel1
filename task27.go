package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println("before swap:", str)
	str = swapString(str)
	fmt.Println("after swap:", str)
}

func swapString(s string) string {
	words := strings.Split(s, " ") // Split разделяет получаемую строку на подстроки, возвращая слайс с ними.
	// в кач-ве сепаратора, по которому делить исходную строку, указываем " " (символ сепаратора).
	fmt.Println(words)                                    // выведем возвращаемый Split слайс
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // идём по слайсу сразу в обе стороны навстречу и меняем пары элементов местами
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") // Join выполняет конкатенацию строковых элементов слайса, для соединения используя, в нашем случае, символ сепаратора
}
