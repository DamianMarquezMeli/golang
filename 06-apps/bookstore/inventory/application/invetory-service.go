package application

import (
	"strings"

	inventory "github.com/devpablocristo/interviews/bookstore/src/inventory/domain"
)

func GetBook(i inventory.InventoryInfo) inventory.InventoryInfo {
	i.Book.Title = strings.ToLower(i.Book.Title)
	return i
}

type RepositoryInteractor struct {
	handler RepositoryInteractorRespository
}

func NewRepositoryInteractor(handler RepositoryInteractorRespository) *RepositoryInteractor {
	return &RepositoryInteractor{handler}
}

func (r RepositoryInteractor) SaveBook(book inventory.Book) error {
	return r.handler.SaveBook(book)
}

func (r RepositoryInteractor) ListInventory() ([]inventory.Book, error) {
	results, _ := r.handler.ListInventory()
	return results, nil
}
