package main

//Реализовать бинарный поиск встроенными методами языка.
import (
	"fmt"
	"sort"
)

func main() {
	B := []int{3, 1, 4, 33, 21, 24}
	sort.Ints(B) //Бинарный поиск работает ТОЛЬКО с отсортированным массивом
	x := 21
	i := sort.Search(len(B), func(i int) bool { return B[i] >= x })
	if i < len(B) && B[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, B)
	} else {
		fmt.Printf("%d not found in %v\n", x, B)
	}

}
