package main

import "fmt"

func FindCoincidences(A,B []int)  {
	M := make(map[int]int)
	answer := make([]int,0)
	for _,val:=range A{
		M[val] = val
	}
	for _,vb := range B{
		if _,ok:= M[vb]; ok {
			answer = append(answer,vb)
		}
	}
	fmt.Println("Пересечения массивов:",answer)

}

func main() {
	A:= []int{11,22,5,1,-10}
	B:= []int{21,5,1,-1}
	FindCoincidences(A,B)
}