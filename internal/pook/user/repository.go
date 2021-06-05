package user

// Repository is the contract between DB to the application
type Repository interface {
	FindAll() ([]User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) error
}
