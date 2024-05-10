package auth

import "github.com/JesseNicholas00/EniqiloStore/utils/logging"

var createStaffLogger = logging.GetLogger("authRepository", "createStaff")

func (repo *authRepostioryImpl) CreateStaff(
	staff Staff,
) (res Staff, err error) {
	query := `
		INSERT INTO staffs(
			staff_id,
			staff_name,
			staff_phone_number,
			staff_password
		) VALUES (
			:staff_id,
			:staff_name,
			:staff_phone_number,
			:staff_password
		) RETURNING
			staff_id,
			staff_name,
			staff_phone_number,
			staff_password,
			created_at,
			updated_at
	`
	rows, err := repo.db.NamedQuery(query, staff)
	if err != nil {
		createStaffLogger.Printf("error executing query: %s", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			createStaffLogger.Printf("error reading result: %s", err)
			return
		}
	}

	return
}
