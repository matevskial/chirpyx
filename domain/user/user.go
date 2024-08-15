package user

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	Id          int
	Email       string
	IsChirpyRed bool
}

type UserWithPassword struct {
	Id             int
	Email          string
	HashedPassword string
	IsChirpyRed    bool
}

type UserRepository interface {
	Create(email string, hashedPassword string) (User, error)
	ExistsByEmail(email string) (bool, error)
	GetUserWithPasswordByEmail(email string) (UserWithPassword, error)
	ExistsByEmailAndIdIsNot(email string, id int) (bool, error)
	Update(id int, email string, hashedPassword string) (User, error)
}
