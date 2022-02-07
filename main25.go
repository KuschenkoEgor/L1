package main

import (
	"fmt"
	"time"
)
//Реализовать собственную функцию sleep.
func MySleep1(sec int) {
	<-time.After(time.Second * time.Duration(sec))
	fmt.Printf("%v seconds has passed\n", sec)
}
func MySleep2(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	fmt.Printf("%v seconds has passed\n", sec)
}

func main() {
	s1 := 5
	s2 := 3
	fmt.Printf("Start 1st Function, time to sleep: %v sec\n",s1)
	MySleep1(s1)
	fmt.Printf("Start 2nd Function, time to sleep: %v sec\n",s2)
	MySleep2(s2)
}