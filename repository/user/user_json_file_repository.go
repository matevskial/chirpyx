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

func (r *UserJsonFileRepository) Create(email string) (userDomain.User, error) {
	return r.db.CreateUser(email)
}
