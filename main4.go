package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func Data(i int, c <- chan int,wg *sync.WaitGroup) {
	for v := range c {
		fmt.Printf("Worker № %v, значение: %v\n", i, v)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Введите количество воркеров")
	var N int
	fmt.Scan(&N)
	c := make(chan int)

	for i:=1; i<N+1; i++{
		wg.Add(1)
		go Data(i,c,&wg)
	}
func(){
	for j:=0; j< 15; j++{
		c <-j
	}
	close(c)
	wg.Wait()
	time.Sleep(time.Hour)
}()



}