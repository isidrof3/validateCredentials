package chain

type Chain interface {
	Next(user, pass string) error
}
