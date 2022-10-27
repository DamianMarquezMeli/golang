package chihandler

import "net/http"

func GetPatient(w http.ResponseWriter, r *http.Request) string {
	return "patient"
}

func CreatePatient(w http.ResponseWriter, r *http.Request) string {
	return "created"
}
