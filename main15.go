package main

import (
	"fmt"
	"unicode/utf8"
)

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "ф"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	rs := []rune(v)
	justString := string(rs[:100])
	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("len of justString = ",utf8.RuneCountInString(justString))
}

func main() {
	someFunc()
}
//Негативные последствия будут заключаться в том, что
//если в функциии createHugeString() мы создадим
//строку из русских символов или латинских (с одиноковым колличеством символов),
//то в результате срез строки justString будет различным (срез идет не до символа, а до номера байта)
// размер русского символа - 2 байта, латинского - 1 байт