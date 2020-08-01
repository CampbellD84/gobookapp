package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add a new Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete a Book")
}
