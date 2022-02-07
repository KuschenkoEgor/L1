package main

import (
	"fmt"
	"math"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.


func ConvertToBigInt(x, y int64) (*big.Int, *big.Int) {
	Newx := big.NewInt(x)
	Newy := big.NewInt(y)
	return Newx,Newy
}
//Add a and b (a+b)
func Add(x,y *big.Int) *big.Int{
	summ := new(big.Int).Add(x,y)
	return summ
}
//Sub a and b (a-b)
func Sub(x,y *big.Int) *big.Int{
	sub := new(big.Int).Sub(x,y)
	return sub
}
//Divide a on b (a/b)
func Div(x,y *big.Int) *big.Int{
	div := new(big.Int).Div(x,y)
	return div
}
//Multiply a and b (a*b)
func Mul(x,y *big.Int) *big.Int{
	mul := new(big.Int).Mul(x,y)
	return mul
}

func main() {
	a := math.Pow(2,40)
	b := math.Pow(2,35)

	ConvA,ConvB := ConvertToBigInt(int64(a),int64(b))

	S := Add(ConvA,ConvB)
	D := Div(ConvA,ConvB)
	M := Mul(ConvA,ConvB)
	Sb:=Sub(ConvA,ConvB)
	fmt.Printf("A:%v\nB:%v\n",ConvA,ConvB)
	fmt.Printf("Add: %v\nDivide: %v\nMultiply: %v\nSub: %v\n",S,D,M,Sb)
}