package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User datos del usuario
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Fullname  string             `json:"fullname"`
	Phone     string             `json:"phone"`
	Role      string             `json:"role"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

// Users lista de usuarios
type Users []*User

/*func (u *User) Validate() *error.RestErr {

}*/

/******************************************************************/
/******************************************************************/
/******************************************************************/
/******************************************************************/
/******************************************************************/

//este para despues
/*
type User2 struct {
	Id        int       `json:"id"`
	User      string    `json:"user"`
	Password  string    `json:"password"`
	Employee  Employee  `json:"employee,omitempty"`
	Rol       string    `json:"rol"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

*/

/*
{
"id":1,
"user":"ccc1",
"password":"1234",
"rol":"agente",
}
*/
