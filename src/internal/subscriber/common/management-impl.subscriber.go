package common

type management struct {
	errCh          chan error
	subscriberList []Subscriber
}

func NewSubscriberManagement() (Management, error) {
	return &management{
		make(chan error),
		[]Subscriber{},
	}, nil
}

func (s *management) Register(subscriber Subscriber) {
	s.subscriberList = append(s.subscriberList, subscriber)
}

func (s *management) Serve() error {
	for _, subscriber := range s.subscriberList {
		go subscriber.Listen()
	}

	return <-s.errCh
}

func (s *management) GracefulShutdown() error {
	for _, subscriber := range s.subscriberList {
		if err := subscriber.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (s *management) GetErrCh() chan error {
	return s.errCh
}
