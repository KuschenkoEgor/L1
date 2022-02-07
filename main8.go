package main

import (
	"encoding/binary"
	"fmt"
)

//Дана переменная int64. Разработать программу,
//которая устанавливает i-й бит в 1 или 0.

func main() {
	var index uint
	var number int64
	var NewNum uint
	fmt.Println("Введите число:")
	fmt.Scan(&number)
	fmt.Println("Введите индекс бита,который хотите изменить (0-7):")
	fmt.Scan(&index)
	fmt.Println("Введите значение, на которое хотите изменить выбранный ранее бит (0 или 1):")
	fmt.Scan(&NewNum)
	if index >= 0 && index < 8 && (NewNum ==0 || NewNum == 1) {
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(number))
		fmt.Printf("Введенное число: %v\n", number)
		fmt.Printf("Введенное чисто в байтах: %v\n", bytes)
		bytes[index] = byte(NewNum)
		New := binary.LittleEndian.Uint64(bytes)

		fmt.Printf("Измененное число в байтах: %v\n", bytes)
		fmt.Printf("Измененное число: %v\n", New)
	} else {
		fmt.Println("Введенное значение превышает заданный диапазон")
	}
}