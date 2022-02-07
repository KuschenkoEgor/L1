package main
//Реализовать паттерн «адаптер» на любом примере.

func main() {
	client := &client{}
	iphone := &iphone{}

	client.insertLightningConnectorIntoTelephone(iphone)

	android := &android{}
	androidAdapter := &AndroidAdapter{
		android,
	}
	client.insertLightningConnectorIntoTelephone(androidAdapter)

}