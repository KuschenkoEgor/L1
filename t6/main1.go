
package main

import "fmt"

func worker(die chan int) {
	for val :=range die{
		fmt.Println(val)
	}
}

func main() {
	die := make(chan int)

	for i := 0; i < 3; i++ {
		go worker(die)
	}
	die <- 1
	die <- 2
	die <- 3

	// Остановка всех worker'ов.
	close(die)
}
