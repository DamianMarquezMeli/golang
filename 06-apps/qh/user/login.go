package login

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func login() {
	s := `qwerty`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	password := `qwerty`

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println("incorrect password")
	} else {
		fmt.Println("Loggin successful")
	}

}
