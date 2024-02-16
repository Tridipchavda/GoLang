package bookapi

import (
	"crud_postgres/dbconn"
	"crud_postgres/models"

	"gorm.io/gorm"
)

// Book struct with DB and BookService Interface
type Book struct {
	db          *gorm.DB
	BookService BookService
}

// Function to get New Book
func NewService(app *dbconn.App) Book {
	book := &Book{
		db: app.DB,
	}
	return *book
}

// Interface with 5 Services Compulsory to Execute
type BookService interface {
	CreateBook(book models.Book) error
	GetAllBook() ([]models.Book, error)
	GetOneBook() (models.Book, error)
	UpdateOneBook(book models.Book, ISBN string) error
	DeleteOneBook(ISBN string) error
}

// Get the Book data using GORM
func (b *Book) GetAllBook() ([]models.Book, error) {
	var books []models.Book
	err := b.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Get one Book data using GORM
func (b *Book) GetOneBook(ISBN string) (*models.Book, error) {
	var book models.Book
	err := b.db.First(&book, "ISBN = ?", ISBN).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Create Book data using GORM
func (b *Book) CreateBook(book models.Book) (models.Book, error) {
	err := b.db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

// Update Book Data using GORM by ISBN
func (b *Book) UpdateOneBook(book models.Book, ISBN string) error {

	err := b.db.Exec("UPDATE BOOKS SET title=?,author=?,year=?,publisher=?,genre=? WHERE ISBN = ?", book.Title, book.Author, book.Year, book.Publisher, book.Genre, book.ISBN).Error
	if err != nil {
		return err
	}
	return nil
}

// Hard Delete Book Data from Table using GORM
func (b *Book) DeleteOneBook(ISBN string) error {

	err := b.db.Exec("DELETE FROM BOOKS WHERE ISBN = ?", ISBN).Error
	if err != nil {
		return err
	}
	return nil
}
