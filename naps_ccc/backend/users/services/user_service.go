package services

/*
	Proceso de de datos para interacción.
*/
import (
	"time"

	users "github.com/devpablocristo/users/models/users"
	userRepository "github.com/devpablocristo/users/repositories"
	"github.com/devpablocristo/users/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	l, _ := time.LoadLocation("America/Buenos_Aires")
	t := time.Now()

	u.CreatedAt = t.In(l)
	u.UpdatedAt = t.In(l)

	newU, err := userRepository.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return newU, nil
}

func GetUsers() (*users.Users, *errors.RestErr) {
	urs, err := userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return urs, nil
}

func GetUser(uId string) (*users.User, *errors.RestErr) {
	u, err := userRepository.GetUser(uId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(u users.User, uId string) (*users.User, *errors.RestErr) {
	ur, err := userRepository.UpdateUser(u, uId)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func DeleteUser(userId string) (*int64, *errors.RestErr) {
	del, err := userRepository.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	return del, nil
}
