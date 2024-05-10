package auth

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var findStaffByPhoneLogger = logging.GetLogger(
	"authRepository",
	"findStaffByPhone",
)

func (repo *authRepostioryImpl) FindStaffByPhone(
	phone string,
) (res Staff, err error) {
	query := `
		SELECT
			*
		FROM
			staffs
		WHERE
			staff_phone_number = :phone_number
	`
	rows, err := repo.db.NamedQuery(query, map[string]interface{}{
		"phone_number": phone,
	})
	if err != nil {
		findStaffByPhoneLogger.Printf(
			"error executing query: %s",
			err,
		)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			findStaffByPhoneLogger.Printf(
				"error reading result: %s",
				err,
			)
		}
	}

	if res.Id == "" {
		err = ErrPhoneNumberNotFound
		return
	}

	return
}
