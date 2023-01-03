package port

import "github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"

type Repository interface {
	SaveBook(domain.Book) error
	ListInventory() ([]domain.Book, error)
}
