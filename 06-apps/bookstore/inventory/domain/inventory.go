package domain

import "time"

type Book struct {
	Author Person  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	ISBN   string  `json:"isbn"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type InventoryBook struct {
	Book      Book      `json:"book"`
	Stock     int64     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

var Inventory []InventoryBook
