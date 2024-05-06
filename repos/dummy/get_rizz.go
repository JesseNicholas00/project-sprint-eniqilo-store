package dummy

func (repo *dummyRepositoryImpl) GetRizz(dummy string) (int, error) {
	res := Rizz{}
	err := repo.db.Get(
		&res,
		"SELECT rizz FROM rizz WHERE dummy=$1",
		dummy,
	)
	if err != nil {
		return 0, err
	}
	return res.Rizz, nil
}
