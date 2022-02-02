package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	var S string
	var k int //счетчик
	fmt.Println("Введите строку:")
	fmt.Scan(&S)
	ln := utf8.RuneCountInString(S)
	for i:=0; i<ln; i++{
		val := strings.Count(S,string(S[i]))
		if val > 1 {
			fmt.Printf("%v - false",S)
			break
		} else {
			k++
		}
	}
	if k ==ln {
		fmt.Printf("%v - true",S)
	}
}