package main

import "fmt"

type AndroidAdapter struct {
	AndroidTelephone *android
}

func (a *AndroidAdapter) insertIntoLightningPort(){
	fmt.Println("Adapter converts Lightning signal to Type-C.")
	a.AndroidTelephone.insertIntoTypeCPort()
}