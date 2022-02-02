package main

import "fmt"

//Поменять местами два числа
//без создания временной переменной.

func main() {
	 A:= []int{3,10}

	 for i:=0;i<len(A)-1;i++{
	 	A[i],A[i+1] = A[i+1],A[i]
	 }
	 fmt.Println(A)
}