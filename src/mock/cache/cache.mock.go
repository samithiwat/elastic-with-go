package cache

import (
	"github.com/samithiwat/elastic-with-go/src/pb"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
	V map[string]interface{}
}

func (t *RepositoryMock) SaveCache(key string, v interface{}, ttl int) error {
	args := t.Called(key, v, ttl)

	t.V[key] = v

	return args.Error(0)
}

func (t *RepositoryMock) GetCache(key string, v interface{}) error {
	args := t.Called(key, v)

	if args.Get(0) != nil {
		switch v.(type) {
		case *[]*pb.Course:
			*v.(*[]*pb.Course) = *args.Get(0).(*[]*pb.Course)
		}
	}

	return args.Error(1)
}

func (t *RepositoryMock) RemoveCache(key string) (err error) {
	args := t.Called(key)

	return args.Error(0)
}
