package domain

type Person struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}
