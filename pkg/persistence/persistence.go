package persistence

const (
	LC = "local"
)

type Storage interface {
	Store(docs string) error
}
