package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func DataFromCh(c chan int){
	for val := range c{
		fmt.Println(val)
	}
}


func main() {
	fmt.Println("Введите количество секунд работы программы:")
	var N int
	fmt.Scan(&N)
	c := make(chan int)

	go DataFromCh(c)
	for j := 0; j < 15; j++ {
		c <- j
	}
	close(c)

	time.Sleep(time.Second*time.Duration(N))
}