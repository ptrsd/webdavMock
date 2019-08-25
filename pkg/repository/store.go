package repository

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	Find(string) ([]byte, error)
}

type WriteRepository interface {
	Put(string, []byte) error
}
