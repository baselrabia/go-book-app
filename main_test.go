package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/baselrabia/go-book-app/handlers"
	"github.com/baselrabia/go-book-app/models"
	"github.com/baselrabia/go-book-app/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	bookJSON = `{"title":"Sample Book","author":"John Doe","published":2020}`
)

func TestCreateBook(t *testing.T) {
	// Database setup
	db := models.InitDB("test.db")
	models.Migrate(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//create handler
	repo := repository.NewGormBookRepository(models.DB)
	h := handlers.NewBookHandler(repo)

	// Assertions
	if assert.NoError(t, h.CreateBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Verify the created book in the database
		var createdBook models.Book
		if err := db.First(&createdBook, 1).Error; err != nil {
			t.Errorf("Error retrieving created book from the database: %v", err)
		}
		assert.Equal(t, "Sample Book", createdBook.Title)
	}
}

func TestGetBook(t *testing.T) {
	// Database setup
	db := models.InitDB("test.db")
	models.Migrate(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()

	// Initialize the BookHandler with the GORM database connection
	repo := repository.NewGormBookRepository(db)
	h := handlers.NewBookHandler(repo)

	// Create a request with valid book data
	bJSON := `{"title": "created Sample Book", "author": "John Doe", "published": 2020}`
	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bJSON))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()

	// Perform the request
	createContext := e.NewContext(createReq, createRec)

	// Create a book first
	if assert.NoError(t, h.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/books/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, h.GetBook(c)) {

		// Verify the retrieved book's title
		var retrievedBook models.Book
		if err := repo.DB.First(&retrievedBook, 1).Error; err != nil {
			t.Errorf("Error retrieving book from the database: %v", err)
		}
		assert.Equal(t, "created Sample Book", retrievedBook.Title)
	}

}

func TestUpdateBook(t *testing.T) {
	// Database setup
	db := models.InitDB("test.db")
	models.Migrate(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()
	// Initialize the BookHandler with the GORM database connection
	repo := repository.NewGormBookRepository(db)
	bookHandler := handlers.NewBookHandler(repo)

	// Create a request with valid book data
	bJSON := `{"title": "created Sample Book", "author": "John Doe", "published": 2020}`
	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bJSON))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()

	// Perform the request
	createContext := e.NewContext(createReq, createRec)

	// Create a book first
	if assert.NoError(t, bookHandler.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	// Now, test updating the created book
	updateReq := httptest.NewRequest(http.MethodPut, "/api/books/1", strings.NewReader(`{"title":"Updated Book"}`))
	updateReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	updateRec := httptest.NewRecorder()
	updateContext := e.NewContext(updateReq, updateRec)
	updateContext.SetParamNames("id")
	updateContext.SetParamValues("1")

	// Assertions
	if assert.NoError(t, bookHandler.UpdateBook(updateContext)) {
		assert.Equal(t, http.StatusOK, updateRec.Code)

		// Verify the updated book's title
		var updatedBook models.Book
		if err := db.First(&updatedBook, 1).Error; err != nil {
			t.Errorf("Error retrieving updated book from the database: %v", err)
		}
		assert.Equal(t, "Updated Book", updatedBook.Title)
	}
}

func TestDeleteBook(t *testing.T) {
	// Database setup
	db := models.InitDB("test.db")
	models.Migrate(db)
	defer db.Migrator().DropTable(&models.Book{})

	e := echo.New()
	// Initialize the BookHandler with the GORM database connection
	repo := repository.NewGormBookRepository(db)
	bookHandler := handlers.NewBookHandler(repo)

	// Create a request with valid book data
	bJSON := `{"title": "created Sample Book", "author": "John Doe", "published": 2020}`
	createReq := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(bJSON))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()

	// Perform the request
	createContext := e.NewContext(createReq, createRec)

	// Create a book first
	if assert.NoError(t, bookHandler.CreateBook(createContext)) {
		assert.Equal(t, http.StatusCreated, createRec.Code)
	}

	// Now, test deleting the created book
	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/books/1", nil) // Assuming 1 is the ID of the created book
	deleteRec := httptest.NewRecorder()
	deleteContext := e.NewContext(deleteReq, deleteRec)
	deleteContext.SetParamNames("id")
	deleteContext.SetParamValues("1")

	// Assertions
	if assert.NoError(t, bookHandler.DeleteBook(deleteContext)) {
		assert.Equal(t, http.StatusNoContent, deleteRec.Code)

		// Verify that the book has been deleted
		var deletedBook models.Book
		if err := db.First(&deletedBook, 1).Error; err != gorm.ErrRecordNotFound {
			t.Errorf("Expected the book to be deleted, but it still exists in the database.")
		}
	}
}
