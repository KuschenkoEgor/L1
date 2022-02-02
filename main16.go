package main

import (
	"fmt"
	"sort"
)

//Реализовать быструю сортировку
//массива (quicksort) встроенными методами языка.

func main() {
	A := []int{3,1,4,33,21}
	fmt.Println("Массив до сортировки:",A)
	sort.Ints(A)
	fmt.Println("Массив после сортировки:",A)

}