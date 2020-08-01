package book

import (
	"github.com/campbellD84/gobookapp/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Book model for DB
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks retrieves all books from DB
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

// GetBook retrieves single book from DB
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

// NewBook creates new book in DB
func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

// DeleteBook deletes book from DB
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("Book with given id not found")
		return
	}
	db.Delete(&book)
	c.Send("Book was deleted successfully")
}
