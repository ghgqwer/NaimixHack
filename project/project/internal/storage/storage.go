package storage

type Storage struct {
	Inner map[string]string
}

func NewStorage() Storage {
	return Storage{
		Inner: make(map[string]string),
	}
}
