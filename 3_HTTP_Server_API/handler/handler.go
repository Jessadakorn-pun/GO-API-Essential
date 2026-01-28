package handler

import (
	"os"
	"strconv"

	"github.com/Jessadakorn-pun/go-example/domain"
	"github.com/gofiber/fiber/v2"
)

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} domain.Book
// @Router /books [get]
func GetBooks(c *fiber.Ctx) error {
	return c.JSON(domain.Books)
}

func GetBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range domain.Books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book ID Not Found")
}

func CreateBook(c *fiber.Ctx) error {
	book := new(domain.Book) // create struct address
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	domain.Books = append(domain.Books, *book) // pass value to append method
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookupdate := new(domain.Book)
	if err := c.BodyParser(bookupdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range domain.Books {
		if book.ID == bookID {
			book.Title = bookupdate.Title
			book.Author = bookupdate.Author
			// Update Value in Array
			domain.Books[index] = book
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book ID Not Found")

}

func DeleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range domain.Books {
		if book.ID == bookID {
			domain.Books = append(domain.Books[:index], domain.Books[index+1:]...) // ... กระจาย element จาก slice
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book ID Not Found")
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File Upload Complete !")
}

func TestHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Hello World",
		"Greeting": "Hello PunPun",
	})
}

func GetEnv(c *fiber.Ctx) error {

	// if value, exits := os.LookupEnv("SECRET"); exits {
	// 	return c.JSON(fiber.Map{
	// 		"SECRET": value,
	// 	})
	// }

	// return c.JSON(fiber.Map{
	// 	"SECRET": "default secret",
	// })

	secret := os.Getenv("SECRET")
	if secret != "" {
		return c.JSON(fiber.Map{
			"SECRET": secret,
		})
	}

	return c.JSON(fiber.Map{
		"SECRET": "Default Secret",
	})

}
