package main

import (
	"fmt"
	"time"
)

func Work() string {
	return "Hello"
}

func main() {

	ch := make(chan string)


	go func() {
		timer := time.After(time.Second*1)
		for {
			select {
			case ch <- Work():
			case <-timer:
				fmt.Println("Ожидание окончено")
				close(ch)
				return
			}
			time.Sleep(time.Millisecond * 200)
		}
	}()


	for data := range ch {
		fmt.Printf("%s\n", data)
	}
}
