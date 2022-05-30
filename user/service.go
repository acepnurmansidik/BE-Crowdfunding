package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error){
	// struct dari user
	user := User{}
	// melakukan hash password
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil{
		return user, err
	}

	// mapping input dari user(strunct RegisterUser) ke dalam struct User
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.PasswordHash = string(PasswordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil{
		return newUser, err
	}

	return newUser, nil
}