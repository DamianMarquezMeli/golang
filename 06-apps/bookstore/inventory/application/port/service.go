package port

type InventoryApp interface {
	GetBook(i inventory.InventoryInfo) inventory.InventoryInfo
}
