package dummy

func (svc *dummyServiceImpl) GetDummy(req GetDummyReq, res *GetDummyRes) error {
	rizz, err := svc.repo.GetRizz(req.Id)

	if err != nil {
		return err
	}

	res.Rizz = rizz
	return nil
}
