package handlers

import (
	"net/http"

	"github.com/baselrabia/go-book-app/dto"
	"github.com/baselrabia/go-book-app/models"
	"github.com/baselrabia/go-book-app/repository"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type BookHandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(dto.Book)
	if err := c.Bind(book); err != nil {
		log.Error(err.Error())
		return err
	}
	// Validate the book data
	if err := validateBook(book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.Repo.CreateBook(book); err != nil {
		log.Error("Failed to create the book. ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create the book.")
	}

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error("Invalid book ID: ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}

	book, err := h.Repo.GetBookByID(id)
	if err != nil {
		log.Error("Book not found ", err.Error())
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.Repo.GetAllBooks()
	if err != nil {
		log.Error("Failed to retrieve books. ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve books.")
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error("Invalid book ID: ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}
	// Create a new dto.Book instance and bind the request data to it
	var updatedBook dto.Book
	if err := c.Bind(&updatedBook); err != nil {
		log.Error("Invalid request data ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid request data")
	}

	book := new(models.Book)
	if book, err = h.Repo.UpdateBook(id, &updatedBook); err != nil {
		log.Error("Failed to update the book ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update the book.")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error("Invalid book ID: ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}

	if err := h.Repo.DeleteBook(id); err != nil {
		log.Error("Book not found ", err.Error())
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.NoContent(http.StatusNoContent)
}
