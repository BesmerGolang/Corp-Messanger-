package auth

import "errors"

var usersDB = make(map[string]User)
var idCounter = 1

type Repository struct{}

func NewRepisitory() *Repository {
	return &Repository{}
}
func (r Repository) CreateUser(username string, passwordHash string) (User, error) {
	if _, exist := usersDB[username]; exist {
		return User{}, errors.New("User already exist")
	}
	user := User{
		ID:           idCounter,
		Username:     username,
		PasswordHash: passwordHash,
	}
	usersDB[username] = user
	idCounter++
	return user, nil
}
func (r Repository) GetUserByUsername(username string) (User, error) {
	user, exist := usersDB[username]
	if !exist {
		return User{}, errors.New("User doesn't exist")
	}
	return user, nil
}
