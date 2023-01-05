package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	cdomain "github.com/devpablocristo/golang/06-apps/bookstore/commons/domain"
	port "github.com/devpablocristo/golang/06-apps/bookstore/inventory/application/port"
	domain "github.com/devpablocristo/golang/06-apps/bookstore/inventory/domain"
	mapdb "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/repository/inmemory/mapdb"
	slicedb "github.com/devpablocristo/golang/06-apps/bookstore/inventory/infrastructure/repository/inmemory/slicedb"
	pdomain "github.com/devpablocristo/golang/06-apps/bookstore/person/domain"
)

var (
	mapDB   = mapdb.NewMapDB()
	sliceDB = slicedb.NewSliceDB()
)

func init() {
	book1 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	book2 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Stanislaw",
			Lastname:  "Lem",
		},
		Title: "Solaris",
		Price: 65.20,
		ISBN:  "0156027607",
	}

	book3 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Arthur C.",
			Lastname:  "Clarck",
		},
		Title: "Rendezvous with Rama",
		Price: 53.50,
		ISBN:  "0-575-01587-X",
	}

	book4 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Jorge Luis",
			Lastname:  "Borges",
		},
		Title: "El Aleph",
		Price: 42.75,
		ISBN:  "84-206-1933-7",
	}

	domain.Inventory = []domain.BookStock{
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

	mapDB.SaveBook(b1)
	mapDB.SaveBook(b2)

	sliceDB.SaveBook(b1)
	sliceDB.SaveBook(b2)

	fmt.Println(*mapDB)
	fmt.Println(sliceDB)
}

type ErrorResponse struct {
	Message string `json:"error"`
}

type Handler struct {
	inventory port.Service
}

func NewHandler(in port.Service) *Handler {
	return &Handler{
		inventory: in,
	}
}

func (h Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	book1 := domain.Book{
		Author: pdomain.Person{
			Firstname: "Isaac",
			Lastname:  "Asimov",
		},
		Title: "Fundation",
		Price: 28.50,
		ISBN:  "0-553-29335-4",
	}

	fmt.Println(book1)
}

func (h Handler) GetInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	inventory, err := h.inventory.GetInventory()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(cdomain.ResponseInfo{
		Status: http.StatusOK,
		Data:   inventory,
	})
}

func (h Handler) AddBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	received_JSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	newBookStock := make([]domain.BookStock, 0)
	err = json.Unmarshal(received_JSON, &newBookStock)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"Error unmarshaling the request: %v"}`, err)
		return
	}

	inventoryMap := make(map[string]int64)
	for i := 0; i < len(domain.Inventory); i++ {
		inventoryMap[domain.Inventory[i].Book.ISBN] = domain.Inventory[i].Stock
	}

	for _, v := range newBookStock {
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

func (h Handler) AddBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book domain.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	err2 := h.handler.SaveBook(book)

	// id = id + 1
	// newBook.ID = id
	// inventory = append(inventory, newBook)

	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) UpdateBookByPatch(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) ISBNContaines(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
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
