package dummy

import dummyRepo "github.com/KerakTelor86/GoBoiler/repos/dummy"

type dummyServiceImpl struct {
	repo        dummyRepo.DummyRepository
	dummyCfgVal int
}

func NewDummyService(
	repo dummyRepo.DummyRepository,
	dummyCfgVal int,
) DummyService {
	return &dummyServiceImpl{repo: repo, dummyCfgVal: dummyCfgVal}
}
