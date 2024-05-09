package auth

type AuthRepository interface {
	CreateStaff(staff Staff) (Staff, error)
	FindStaffByPhone(phone string) (Staff, error)
}
