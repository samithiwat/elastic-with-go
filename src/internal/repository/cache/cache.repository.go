package cache

type Repository interface {
	SaveCache(string, interface{}, int) error
	GetCache(string, interface{}) error
	RemoveCache(string) error
}
