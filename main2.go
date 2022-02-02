package main

import "fmt"

//Написать программу, которая конкурентно рассчитает
//значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.
var A = [5]int{2,4,6,8,10}

func Square(c chan int) {
	for _,val := range A{
		c <- val*val
	}
	close(c)
}

func main() {
	c := make(chan int)
	go Square(c)
	for val := range c{
		fmt.Println(val)
	}

}