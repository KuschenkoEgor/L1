package main

import "fmt"

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

var Arr = [5]int{2,4,6,8,10}

func Square1(c chan int) {
	for _,val := range Arr{
		c <- val*val
	}
	close(c)
}

func main() {
	var sum int
	c := make(chan int)
	go Square1(c)
	for val := range c{
		sum +=val
	}
	fmt.Println("Сумма квадратов массива:",sum)
}