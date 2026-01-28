package main

import (
	"log"
	"os"

	_ "github.com/Jessadakorn-pun/go-example/docs"
	"github.com/Jessadakorn-pun/go-example/domain"
	"github.com/Jessadakorn-pun/go-example/handler"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load .env error")
	}

	// define render engin
	engine := html.New("./views", ".html")

	// Initate Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// swagger aip documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	domain.Books = append(domain.Books, domain.Book{ID: 1, Title: "Mikelopster", Author: "Mike"})
	domain.Books = append(domain.Books, domain.Book{ID: 2, Title: "MM", Author: "Mike"})

	// Login Process
	app.Post("/login", handler.Login)

	// Middleware
	// app.Use(handler.CheckMiddleware)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(handler.CheckMiddleware)

	// GET Method
	app.Get("/books", handler.GetBooks)
	app.Get("/books/:id", handler.GetBook)
	app.Get("/test-HTML", handler.TestHTML)
	app.Get("/api/config", handler.GetEnv)

	// POST Method
	app.Post("/books", handler.CreateBook)

	// PUT Method
	app.Put("/books/:id", handler.UpdateBook)

	// DELETE Method
	app.Delete("/books/:id", handler.DeleteBook)

	// POST : Upload file
	app.Post("/upload", handler.UploadFile)

	// Defind Fiber Port
	app.Listen(":8080")
}

// func main() {
// 	app := fiber.New()
// 	app.Get("/hello", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello World")
// 	})

// 	app.Listen(":8080")
// }

// "fmt"
// "log"
// "net/http"

// func main() {
// 	http.HandleFunc("/hello", helloHandler)

// 	fmt.Printf("Starting Server at Port 8080 \n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not Found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "404 not Found", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello World")
// }
