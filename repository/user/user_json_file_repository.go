package user

import (
	"github.com/matevskial/chirpyx/database"
	userDomain "github.com/matevskial/chirpyx/domain/user"
)

type UserJsonFileRepository struct {
	db *database.JsonFileDB
}

func NewUserJsonFileRepository(db *database.JsonFileDB) *UserJsonFileRepository {
	return &UserJsonFileRepository{db: db}
}

func (r *UserJsonFileRepository) Create(email string, hashedPassword string) (userDomain.User, error) {
	return r.db.CreateUser(email, hashedPassword)
}

func (r *UserJsonFileRepository) ExistsByEmail(email string) (bool, error) {
	return r.db.ExistsUserByEmail(email)
}

func (r *UserJsonFileRepository) GetUserWithPasswordByEmail(email string) (userDomain.UserWithPassword, error) {
	user, err := r.db.GetUserByEmail(email)
	if err != nil {
		return userDomain.UserWithPassword{}, err
	}
	return userDomain.UserWithPassword{Id: user.Id, Email: user.Email, HashedPassword: user.HashedPassword}, nil
}
