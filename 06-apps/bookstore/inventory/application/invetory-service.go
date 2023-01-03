package application

import (
	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/application/port"
	"github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
)

type InventoryService struct {
	storage port.Service
}

func NewInventoryService(st port.Service) *InventoryService {
	return &InventoryService{
		storage: st,
	}
}

func (i *InventoryService) GetBook(ISBN string) *domain.Book {
	//i.Book.Title = strings.ToLower(i.Book.Title)
	return &domain.Book{}
}

func (i *InventoryService) SaveBook(book *domain.Book) error {
	//return r.handler.SaveBook(book)
	return nil
}

func (i *InventoryService) GetInventory() ([]domain.BookStock, error) {
	//results, _ := r.handler.ListInventory()
	return domain.Inventory, nil
}
