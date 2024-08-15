package database

import userDomain "github.com/matevskial/chirpyx/domain/user"

type User struct {
	Id             int
	Email          string
	HashedPassword string
	IsChirpyRed    bool
}

func (db *JsonFileDB) CreateUser(email string, hashedPassword string) (userDomain.User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return userDomain.User{}, err
	}
	userToAdd := User{Id: dbStructure.UserIdSeq, Email: email, HashedPassword: hashedPassword}
	addedUser := dbStructure.addUser(userToAdd)
	err = db.writeDB(dbStructure)
	if err != nil {
		return userDomain.User{}, err
	}
	return addedUser, nil
}

func (db *JsonFileDB) ExistsUserByEmail(email string) (bool, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return false, err
	}

	for _, value := range dbStructure.Users {
		if value.Email == email {
			return true, nil
		}
	}

	return false, nil
}

func (db *JsonFileDB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, value := range dbStructure.Users {
		if value.Email == email {
			return value, nil
		}
	}

	return User{}, userDomain.ErrUserNotFound
}

func (db *JsonFileDB) ExistsUserByEmailAndIdIsNot(email string, id int) (bool, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return false, err
	}

	for _, value := range dbStructure.Users {
		if value.Email == email && value.Id != id {
			return true, nil
		}
	}

	return false, nil
}

func (db *JsonFileDB) UpdateUser(id int, email string, hashedPassword string) (userDomain.User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return userDomain.User{}, err
	}

	user, err := dbStructure.updateUser(id, email, hashedPassword)
	if err != nil {
		return userDomain.User{}, err
	}
	err = db.writeDB(dbStructure)
	if err != nil {
		return userDomain.User{}, err
	}
	return user, nil
}
