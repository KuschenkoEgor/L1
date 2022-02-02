package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
func main() {
	var Str = "snow dog sun"

	arr := strings.Split(Str," ")

	ln := len(arr)

	for i:=0; i<ln/2;i++{
	arr[i],arr[ln-1-i]= arr[ln-1-i],arr[i]
	}
	fmt.Println("Исходная строка:",Str)
	fmt.Println("Новая строка:",strings.Join(arr," "))
}
