package chain

import "errors"

//MOCK para service
type nivelUnoStepMock struct{}

func NewNivelUnoStepMock() *nivelUnoStepMock { return &nivelUnoStepMock{} }

func (rr *nivelUnoStepMock) Next(user, pass string) error {

	return nil
}

//MOCK error para nivel uno
type nivelUnoStepErrMock struct{}

func NewNivelUnoStepErrMock() *nivelUnoStepErrMock { return &nivelUnoStepErrMock{} }

func (rr *nivelUnoStepErrMock) Next(user, pass string) error {

	return errors.New("pass less than 8")
}

//MOCK para nivel 2
type nivelDosStepMock struct{}

func NewNivelDosStepMock() *nivelDosStepMock { return &nivelDosStepMock{} }

func (rr *nivelDosStepMock) Next(user, pass string) error {

	return errors.New("pass dont has min or mayus")
}
