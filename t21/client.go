package main

import "fmt"

type client struct {

}
func (c *client) insertLightningConnectorIntoTelephone (t telephone){
	fmt.Println("Client inserts Lightning connector into telephone.")
	t.insertIntoLightningPort()
}