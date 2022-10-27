package domain

type Person struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	DNI      int64  `json:"dni"`
	Gender   string `json:"gender"`
}

type Doctor struct {
	Doctor     Person `json:"uuid"`
	Speciality string `json:"speciality"`
}

type Patient struct {
	Patient   Person `json:"patient"`
	Doctor    Doctor `json:"doctor"`
	Hospital  string `json:"hospital"`
	Diagnosis string `json:"diagnosis"`
}