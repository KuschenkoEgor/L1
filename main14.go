package main

import "fmt"
//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel из переменной типа interface{}
func main() {
	var X interface{}
	X = 54
	fmt.Printf("Для перменной %v тип переменной: %T\n ",X,X)
	X = false
	fmt.Printf("Для перменной %v тип переменной: %T\n ",X,X)
	X = 43.1
	fmt.Printf("Для перменной %v тип переменной: %T\n ",X,X)
	X = make(chan int)
	fmt.Printf("Для перменной %v тип переменной: %T\n ",X,X)
	X = "K P A C U B O"
	fmt.Printf("Для перменной %v тип переменной: %T\n ",X,X)
}
//Также можно выводить тип переменной через:
//fmt.Println(reflect.TypeOf(X))