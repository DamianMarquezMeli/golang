package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `laralapassword`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	//contraseña := `laralapassword`
	contraseña := `qwerty`

	err = bcrypt.CompareHashAndPassword(bs, []byte(contraseña))
	if err != nil {
		fmt.Println("No podes loguearte")
	} else {
		fmt.Println("Estas logueado")
	}

}
