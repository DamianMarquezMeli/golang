package domain

import "time"

type Person struct {
	UUID       string    `json:"uuid" form:"idPerson" gorm:"primary_key"`
	Name       string    `json:"name" form:"nombresPerson" binding:"required"`
	Lastname   string    `json:"lastname" form:"lastname" binding:"required"`
	Age        int       `json:"age" form:"gender" binding:"required"`
	Gender     string    `json:"gender" form:"gender" binding:"required"`
	CreatedAt  time.Time `gorm:"-" bson:"created_at" json:"created_at"`
	UpdartedAt time.Time `gorm:"-" bson:"updated_at" json:"updated_at,omitempty"`
	DeletedAt  time.Time `gorm:"-" bson:"deleted_at" json:"deleted_at,omitempty"`
}
