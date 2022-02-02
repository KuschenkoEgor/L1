package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную
//запись данных в map.
func main() {

	c := make(map[int]int)
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	n := 200
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c[i] = i
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(c))

}