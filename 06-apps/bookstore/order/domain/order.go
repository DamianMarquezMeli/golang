package domain

import "time"

type Order struct {
	Customer Person       `json:"client"`
	Date     time.Time    `json:"date"`
	Details  []OrderItems `json:"details"`
}

type OrderItems struct {
	Book     Book  `json:"books"`
	Quantity int64 `json:"quantity"`
}

var Orders []Order
