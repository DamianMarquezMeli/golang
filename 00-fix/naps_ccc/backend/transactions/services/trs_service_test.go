package services

import (
	"fmt"
	"testing"
	"time"

	trss "github.com/devpablocristo/transactions/models/trss"
	trsRepository "github.com/devpablocristo/transactions/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	En VS no sale t.log por la salida estandar
	Correr en en la consola:

	go test -run TestCreateTrs

	go test -v trs_service_test.go

	Para ver los diferentes tests.
*/

func TestCreateTrs(t *testing.T) {

	oid := primitive.NewObjectID()
	l, _ := time.LoadLocation("America/Buenos_Aires") // or the name of your time zone
	tm := time.Now()

	// si no se cambia algo del los datos, no los guarda,
	// no se pq., supungo que tiene que ver con algo sobre seguridad
	trs := trss.Trs{
		Id:          oid,
		Description: "poiupoiu",
		Username:    "r.Garcia",
		Password:    "12345",
		Fullname:    "Ronaldo Garcia",
		Phone:       "777777",
		Date:        tm.In(l),
	}

	u, rErr := CreateTrs(trs)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Error(u)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario creado correctamente:\n", u, "\n\n")
	}
}

func TestGetTrss(t *testing.T) {
	urs, rErr := GetTrss()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	if len(*urs) < 1 {
		t.Error("No hay documentos en la base de datos.")
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Se obtuvieron todos los documentos correctamente:")
		for _, v := range *urs {
			fmt.Println(v)
		}
		fmt.Println()
	}
}

func TestGetTrs(t *testing.T) {
	lastDocument, rErr := trsRepository.GetIdLastInseted()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	tId := lastDocument["_id"].(primitive.ObjectID).Hex()

	u, rErr := GetTrs(tId)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario obternido correctamente:\n", u, "\n\n")
	}
}

func TestUpdateTrs(t *testing.T) {

	lastDocument, rErr := trsRepository.GetIdLastInseted()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	tId := lastDocument["_id"].(primitive.ObjectID).Hex()

	l, _ := time.LoadLocation("America/Buenos_Aires") // or the name of your time zone
	tm := time.Now()

	/*u := trs.Trs{
		Trsname:  "Abuelo888.simpson",
		Password:  "5432124234234242",
		Email:     "b.simpson@fox.com",
		UpdatedAt: tm.In(l),
	}*/

	trs := trss.Trs{
		Description: "qwerty qwerty",
		Username:    "m.lopez",
		Password:    "098765",
		Fullname:    "Mario Lopez",
		Phone:       "0001010",
		Date:        tm.In(l),
	}

	ur, rErr := UpdateTrs(trs, tId)

	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario actualizado correctamente:\n", ur, "\n\n")
	}
}

func TestDeleteTrs(t *testing.T) {

	lastDocument, rErr := trsRepository.GetIdLastInseted()
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	}

	uId := lastDocument["_id"].(primitive.ObjectID).Hex()

	_, rErr = DeleteTrs(uId)
	if rErr != nil {
		t.Error(rErr.Message)
		t.Fail()
	} else {
		t.Log("\n\nEXITO! Usuario eliminado:\n", lastDocument, "\n\n")
	}
}
