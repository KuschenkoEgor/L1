//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования)

package main

import "fmt"

type Human struct {
	Name string
	Age int
	Growth float64
}

type Action struct {
	WalkinSpeed int
	Human
}

func (h Human) ConvertGrowth() float64 {
	return h.Growth/100
}
//Вызываем метод, описанный для структуры Human
//Т.к. данная структура является частью структуры Action,
//к любым методам, объявленным в тиме Human, можно получить доступ
//через поле Action
func main() {
	h := Human{Name: "Egor",Age: 25,Growth: 178}
	action := Action{WalkinSpeed: 6,Human:h}
	fmt.Println("Рост в метрах:",action.ConvertGrowth())
}