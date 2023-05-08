package storage

type Item struct {
	Service  string
	Password string
}

type Storage interface {
	Set(username string, item *Item) error
	Get(username string, service string) (*Item, error)
	Delete(username string, service string) (int64, error)
	TearDown() error
}
