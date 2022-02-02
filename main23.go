package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	sl := []int{22,12,45,1,77}
	var n int
	fmt.Println("Введите номер элемента, который хотите удалить:")
	fmt.Scan(&n)
	//Операции именно для номера элемента, а не для индекса
	sl = append(sl[:n-1],sl[n:]...)
	fmt.Println("Обновленный массив:",sl)

}