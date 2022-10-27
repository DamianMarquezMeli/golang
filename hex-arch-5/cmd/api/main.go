package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/devpablocristo/golang/hex-arch-5/internal/core/domain"
	patientservice "github.com/devpablocristo/golang/hex-arch-5/internal/core/service"
	ginhandler "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/handlers/gin"
	memkvsrepo "github.com/devpablocristo/golang/hex-arch-5/internal/infrastructure/repositories/kvs"
)

const webServerPort string = ":8088"

func main() {

	patPerson := domain.Person{
		UUID:     "1",
		Name:     "Homero",
		Lastname: "Simpson",
		DNI:      12345,
		Gender:   "m",
	}

	docPerson := domain.Person{
		UUID:     "2",
		Name:     "Nick",
		Lastname: "Riviera",
		DNI:      63435,
		Gender:   "m",
	}

	doctor := domain.Doctor{
		Doctor:     docPerson,
		Speciality: "Surgery",
	}

	patient := &domain.Patient{
		Patient:   patPerson,
		Doctor:    doctor,
		Hospital:  "General",
		Diagnosis: "Cancer",
	}

	b, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	patientRepository := memkvsrepo.NewMemKVS()
	patientService := patientservice.NewPatientService(patientRepository)
	patientHandler := ginhandler.NewGinHandler(patientService)
	// patientHandler.SetupRoutes()
	// patientHandler.Run(webServerPort)

	router := gin.New()

	router.GET("/patient/:id", patientHandler.GetPatient)
	router.POST("/patient", patientHandler.CreatePatient)

	router.Run(webServerPort)
}
