package port

import "github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"

type Service interface {
	GetBook(string) *domain.Book
	SaveBook(*domain.Book) error
	GetInventory() ([]domain.BookStock, error)
}
