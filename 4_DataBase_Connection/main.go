package main

import (
	"fmt"
	database "sql-connection/Database"
	handler "sql-connection/Handler"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	// Conncect to Database
	database.PostgresConnect(psqlInfo)

	// Connection Database Sucess
	fmt.Println(database.DB)

	// Close Connection after complete all process
	defer database.DB.Close()

	// Create Product using handling function
	// var productTest = domain.Product{
	// 	Name:  "Go product",
	// 	Price: 222,
	// }

	// err := handler.CreateProduct(&productTest)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create New Product Successful !")

	// Get Product
	// product, err := handler.GetProduct(5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Get Successful : ", product)

	// Update Product
	// product, err = handler.UpdateProduct(5, &domain.Product{Name: "YYY", Price: 88888})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Update Successful! New Value: Name = %s, Price = %d \n", product.Name, product.Price)

	// Delete Prouct
	// err := handler.DeleteProduct(5)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Delete Product Successful!")

	// Get Mutiple Prodcuts
	// products, err := handler.GetProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, item := range products {
	// 	fmt.Printf("Id: %d	Name: %s	Price: %d \n", item.ID, item.Name, item.Price)
	// }
	// fmt.Println("Get All Products Successful!")

	// Initiate Fiber
	app := fiber.New()

	// Get API
	app.Get("/product", handler.GetProductsHandler)
	app.Get("/product/:id", handler.GetProductHandler)

	// Post API
	app.Post("/product", handler.CreateProductHandler)

	// Put API
	app.Put("/product", handler.UpdateProductHandler)

	// Delete API
	app.Delete("/product/:id", handler.DeleteProductHandler)

	// Defind Fiber Port
	app.Listen(":8080")

}
