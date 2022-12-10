package common

type Subscriber interface {
	Listen()
	Close() error
	RegisterHandler(...MessageHandler)
}
