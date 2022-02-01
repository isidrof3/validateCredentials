package chain

import (
	"errors"
	"fmt"
)

type nivelUnoStep struct {
	next Chain
}

func NewNivelUnoStepStep(
	next Chain) *nivelUnoStep {
	return &nivelUnoStep{
		next: next,
	}
}

func (nus *nivelUnoStep) Next(user, pass string) error {

	//evaluando nivel 1 de validaciones
	if len(pass) <= 8 {
		fmt.Println("Contraseña debe de poseer una longitud mínima de 8 caracteres")
		return errors.New("pass less than 8")
	}

	if pass == "InfoBlindajes" {
		fmt.Println("Contraseña no debe de contener el nombre de la empresa(“InfoBlindajes”)")
		return errors.New("no validate pass")
	}

	//si es coordinador o gerente se procede con validaciones de nivel 2
	if nus.next != nil && (user == "COORDINADOR" || user == "GERENTE") {
		errNext := nus.next.Next(user, pass)
		if errNext != nil {
			return errNext
		}
	}
	return nil
}
