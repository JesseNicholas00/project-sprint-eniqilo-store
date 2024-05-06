package dummy

type DummyService interface {
	GetDummy(req GetDummyReq, res *GetDummyRes) error
}
