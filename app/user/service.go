package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(userID int) (User, error)
	GetAllUser() ([]User, error)
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

func (s *service) Login(input LoginUserInput) (User, error){
	// ambil email dan passeeord dari user
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return user, err
	}

	// cek email
	if user.ID == 0 {
		return user, errors.New("Email not register")
	}

	// cek password
	result := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if result != nil {
		return user, errors.New("Password no match")
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error){
	// ambil email dari user
	email := input.Email

	// cari email user
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error){
	// dapatkan user berdasarkan ID
	user, err := s.repository.FindByID(ID)
	if err != nil{
		return user, err
	}

	// update attribute avatar filename
	user.AvatarFileName = fileLocation

	// simpan perubahan avatar filename
	updateUser, err := s.repository.Update(user)
	if err != nil{
		return updateUser, err
	}

	return updateUser, nil
}

func (s *service) GetUserByID(userID int) (User, error){
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	// cek usernya
	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

// Service for get all data user
func (s *service) GetAllUser() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}