package dummy

type DummyRepository interface {
	GetRizz(dummy string) (int, error)
}
