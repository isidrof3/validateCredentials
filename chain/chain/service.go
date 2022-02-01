package chain

type Service interface {
	ValidatePassword(user, pass string) error
}

type ServiceImplementation struct {
	validateChain Chain
}

func NewServiceImplementation(
	validateChain Chain) *ServiceImplementation {
	return &ServiceImplementation{
		validateChain: validateChain,
	}
}

func (s *ServiceImplementation) ValidatePassword(user, pass string) error {

	//inicia cadena de validaciones
	err := s.validateChain.Next(user, pass)

	return err
}
