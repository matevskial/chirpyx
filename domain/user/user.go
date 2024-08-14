package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id    int
	Email string
}

type UserWithPassword struct {
	Id             int
	Email          string
	HashedPassword string
}

type UserRepository interface {
	Create(email string, hashedPassword string) (User, error)
	ExistsByEmail(email string) (bool, error)
	GetUserWithPasswordByEmail(email string) (UserWithPassword, error)
}
