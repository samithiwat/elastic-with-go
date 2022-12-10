package common

type Management interface {
	Register(Subscriber)
	Serve() error
	GracefulShutdown() error
	GetErrCh() chan error
}
