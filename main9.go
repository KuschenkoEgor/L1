package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй —
//результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func TakeData(one,two chan int){
	for val := range one{
		//fmt.Println("Значение из 1-го канала:",val)
		two <- val*val
	}
}

func Square2(two chan int, wg *sync.WaitGroup){

	for v:=range two{
		fmt.Println("Значение из 2-го канала:",v)
		wg.Done()
	}
}

func main() {
	wg := sync.WaitGroup{}
	one := make(chan int)
	two := make(chan int)
	N := 5
	wg.Add(N)
	go TakeData(one,two)
	go Square2(two,&wg)
	for i:=1; i<N+1; i++{
		one <- i
	}
	wg.Wait()
	defer close(one)
	defer close(two)
}