package user

type User struct {
	Id    int
	Email string
}

type UserRepository interface {
	Create(email string) (User, error)
}
