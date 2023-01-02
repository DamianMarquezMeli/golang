package trss

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Trs datos del usuario
type Trs struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Description string             `json:"description"`
	Username    string             `json:"trsname"`
	Password    string             `json:"password"`
	Fullname    string             `json:"fullname"`
	Phone       string             `json:"phone"`
	Date        time.Time          `bson:"date" json:"date"`
}

// Trss lista de transacciones
type Trss []*Trs
