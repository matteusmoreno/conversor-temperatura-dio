package main

import "fmt"

const boilingPointK int = 373

func main() {

	boilingPointC := (boilingPointK - 273)

	fmt.Println("O ponto de ebulição em Kelvin é de:", boilingPointK)
	fmt.Println("O ponto de ebulição em Celsius é de:", boilingPointC)
}
