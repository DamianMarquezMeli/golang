package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	domain "github.com/devpablocristo/interviews/bookstore/src/inventory/domain"
	"github.com/gorilla/mux"
)

func init() {
	book1 := Book{
		Author: Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	book2 := Book{
		Author: Person{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
	}

	book3 := Book{
		Author: Person{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := Book{
		Author: Person{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El 'Martín Fierro'",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	Inventory = []InventoryInfo{
		{
			Book:  book1,
			Stock: 41,
		},
		{
			Book:  book2,
			Stock: 32,
		},
		{
			Book:  book3,
			Stock: 12,
		},
		{
			Book:  book4,
			Stock: 93,
		},
	}
}

var inventory []domain.Book
var id uint

func books() {
	id = 4

	book1 := domain.Book{
		ID: 1,
		Author: domain.Author{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
		Stock: 9,
	}

	book2 := domain.Book{
		ID: 2,
		Author: domain.Author{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
		Stock: 15,
	}

	book3 := domain.Book{
		ID: 3,
		Author: domain.Author{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := domain.Book{
		ID: 4,
		Author: domain.Author{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El Aleph",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	inventory = []domain.Book{
		book1,
		book2,
		book3,
		book4,
	}

}

func GetBook(w http.ResponseWriter, r *http.Request) domain.Book {

	book1 := domain.Book{
		Author: domain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	return book1

}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myRest.Inventory)
}

func AddBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	received_JSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	newBooksSlice := make([]domain.InventoryInfo, 0)
	err = json.Unmarshal(received_JSON, &newBooksSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
		return
	}

	inventoryMap := make(map[string]int64)
	for i := 0; i < len(domain.Inventory); i++ {
		inventoryMap[domain.Inventory[i].Book.ISBN] = domain.Inventory[i].Stock
	}

	for _, v := range newBooksSlice {
		if _, found := inventoryMap[v.Book.ISBN]; !found {
			domain.Inventory = append(domain.Inventory, v)
		} else {
			for i := range domain.Inventory {
				if domain.Inventory[i].Book.ISBN == v.Book.ISBN {
					domain.Inventory[i].Stock += v.Stock
					break
				}
			}
		}
	}

	// fmt.Fprintf(w, "inventoryMap: %+v", inventoryMap)
	// fmt.Fprintf(w, "Inventory: %+v", Inventory)

	fmt.Fprintln(w, "Inventory:")
	for _, v := range domain.Inventory {
		fmt.Fprintln(w, "Book:", v.Book.Title)
		fmt.Fprintln(w, "Stock:", v.Stock)
		fmt.Fprintln(w, "------------------")
	}
}

type HTTPInteractor struct {
	handler usecases.UseCasesInteractor
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func NewHTTPInteractor(handler usecases.UseCasesInteractor) *HTTPInteractor {
	return &HTTPInteractor{handler}
}

func MakeHTTPInteractor(handler usecases.UseCasesInteractor) HTTPInteractor {
	return HTTPInteractor{handler}
}

func (h HTTPInteractor) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book inventory.Book

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	err2 := h.handler.SaveBook(book)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h HTTPInteractor) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	results, err := h.handler.ListInventory()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(results)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   "pong",
	})
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	isbn := r.URL.Query().Get("isbn")
	if isbn != "" {
		var sliceBooks []domain.Book
		for _, v := range inventory {
			if v.ISBN == isbn {
				sliceBooks = append(sliceBooks, v)
			}
		}

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: 200,
			Data:   sliceBooks,
		})
		return
	}

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   inventory,
	})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook domain.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	id = id + 1
	newBook.ID = id

	inventory = append(inventory, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusCreated,
		Data:   newBook,
	})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	var newBook domain.Book
	err = json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	for i, book := range inventory {
		if book.ID == uint(idBook) {
			newBook.ID = book.ID
			inventory[i] = newBook
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
}

func updateBookByPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	index := search_book_index_byID(idBook, inventory)

	//book := inventory[index]

	parametros_patcheados := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&parametros_patcheados)

	if err != nil {
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error decoding request body",
		})
		return
	}

	var newBook domain.Book

	bookString, _ := json.Marshal(parametros_patcheados)
	json.Unmarshal(bookString, &newBook)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	if _, ok := parametros_patcheados["id"]; ok {
		inventory[index].ID = newBook.ID
	}

	if _, ok := parametros_patcheados["price"]; ok {
		inventory[index].Price = newBook.Price
	}

	if _, ok := parametros_patcheados["author"]; ok {
		inventory[index].Author.Firstname = newBook.Author.Firstname
		inventory[index].Author.Lastname = newBook.Author.Lastname
	}

	if _, ok := parametros_patcheados["title"]; ok {
		inventory[index].Title = newBook.Title
	}

	if _, ok := parametros_patcheados["isbn"]; ok {
		inventory[index].ISBN = newBook.ISBN
	}

	if _, ok := parametros_patcheados["stock"]; ok {
		inventory[index].Stock = newBook.Stock
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   inventory[index],
	})
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	var newBook domain.Book
	for _, book := range inventory {
		if book.ID == uint(idBook) {
			newBook = book
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
}

func isbn_containes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	isbnParam := param["isbn"]

	var newBook domain.Book
	for _, book := range inventory {
		if strings.Contains(book.ISBN, isbnParam) {
			newBook = book
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newBook,
	})
}

func deleteBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	idParam := param["id"]

	idBook, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: " + idParam,
		})
		return
	}

	book_index := search_book_index_byID(idBook, inventory)

	switch book_index {
	case 0:
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusBadRequest,
			Data:   "error: ID inexistente",
		})
		return
	default:
		inventory = append(inventory[:book_index], inventory[(book_index+1):]...)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusOK,
			Data:   inventory,
		})
	}
}

func search_book_index_byID(neccesaryID uint64, inVentory []domain.Book) int {
	for i, book := range inVentory {
		if book.ID == uint(neccesaryID) {
			return i
		}
	}
	return 0
}

// List all books on the inventory
func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Inventory)
}

// Add books to inventory
func addBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	received_JSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	newBooksSlice := make([]InventoryInfo, 0)
	err = json.Unmarshal(received_JSON, &newBooksSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
		return
	}

	inventoryMap := make(map[string]int64)
	for i := 0; i < len(Inventory); i++ {
		inventoryMap[Inventory[i].Book.ISBN] = Inventory[i].Stock
	}

	for _, v := range newBooksSlice {
		if _, found := inventoryMap[v.Book.ISBN]; !found {
			Inventory = append(Inventory, v)
		} else {
			for i := range Inventory {
				if Inventory[i].Book.ISBN == v.Book.ISBN {
					Inventory[i].Stock += v.Stock
					break
				}
			}
		}
	}

	// fmt.Fprintf(w, "inventoryMap: %+v", inventoryMap)
	// fmt.Fprintf(w, "Inventory: %+v", Inventory)

	fmt.Fprintln(w, "Inventory:")
	for _, v := range Inventory {
		fmt.Fprintln(w, "Book:", v.Book.Title)
		fmt.Fprintln(w, "Stork:", v.Stock)
		fmt.Fprintln(w, "------------------")
	}

}
