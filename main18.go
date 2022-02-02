package main

//Реализовать структуру-счетчик, которая будет инкрементироваться
//в конкурентной среде. По завершению программа должна выводить
//итоговое значение счетчика.
import (
	"fmt"
	"sync"
)

func main() {
	var n int
	X:=200
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(X)
	for i:=0;i<X;i++{
		go func(){
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)

}