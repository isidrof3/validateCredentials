package chain

import (
	"errors"
	"fmt"
	"unicode"
)

type nivelDosStep struct {
	next Chain
}

func NewNivelDosStepStep(
	next Chain) *nivelDosStep {
	return &nivelDosStep{
		next: next,
	}
}

func (nds *nivelDosStep) Next(user, pass string) error {
	tieneMayus := false
	tieneMin := false

	//evaluando nivel 2 de validaciones

	//evaluando mayusculas
	for _, r := range pass {
		if unicode.IsUpper(r) {
			tieneMayus = true
			break
		}
	}

	//evaluando minusculas
	for _, r := range pass {
		if unicode.IsLower(r) {
			tieneMin = true
			break
		}
	}

	//evaluando que el pass tenga min y mayus
	if !tieneMayus || !tieneMin {
		fmt.Println("Contraseña debe de contener al menos una letra mayuscula y una letra minuscula")
		return errors.New("pass dont has min or mayus")
	}

	if nds.next != nil && user == "GERENTE" {
		errNext := nds.next.Next(user, pass)
		if errNext != nil {
			return errNext
		}
	}
	return nil
}
