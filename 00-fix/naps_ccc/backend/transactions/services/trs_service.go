package services

/*
	Proceso de de datos para interacción.
*/
import (
	trss "github.com/devpablocristo/transactions/models/trss"
	trsRepository "github.com/devpablocristo/transactions/repositories"
	"github.com/devpablocristo/transactions/utils/errors"
)

func CreateTrs(t trss.Trs) (*trss.Trs, *errors.RestErr) {
	newT, err := trsRepository.CreateTrs(t)
	if err != nil {
		return nil, err
	}

	return newT, nil
}

func GetTrss() (*trss.Trss, *errors.RestErr) {
	urs, err := trsRepository.GetTrss()
	if err != nil {
		return nil, err
	}

	return urs, nil
}

func GetTrs(tId string) (*trss.Trs, *errors.RestErr) {
	u, err := trsRepository.GetTrs(tId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateTrs(t trss.Trs, tId string) (*trss.Trs, *errors.RestErr) {
	ur, err := trsRepository.UpdateTrs(t, tId)
	if err != nil {
		return nil, err
	}

	return ur, nil
}

func DeleteTrs(tId string) (*int64, *errors.RestErr) {
	del, err := trsRepository.DeleteTrs(tId)
	if err != nil {
		return nil, err
	}

	return del, nil
}
