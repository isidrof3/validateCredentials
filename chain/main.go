package main

import (
	"chain/chain"
	"fmt"
)

func main() {
	//Inicializando los eslabones de la cadena
	twoStep := chain.NewNivelDosStepStep(nil)
	oneStep := chain.NewNivelUnoStepStep(twoStep)

	service := chain.NewServiceImplementation(oneStep)

	err := service.ValidatePassword("ANALISTA", "passworsssss")

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Credenciales validas :)")
	}
}
