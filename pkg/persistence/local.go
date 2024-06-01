package persistence

type Local struct {
}

func (l *Local) Store() {

}

func NewLocal() Storage {
	return &Local{}
}
