package handler

import (
	domain "sql-connection/Domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// integrated with fiber
func GetProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi((c.Params("id")))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product, err := GetProduct(productId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(product)
}

func GetProductsHandler(c *fiber.Ctx) error {
	products, err := GetProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(products)
}

func CreateProductHandler(c *fiber.Ctx) error {
	newProduct := new(domain.Product)

	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// insert into database
	err := CreateProduct(newProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(newProduct)
}

func UpdateProductHandler(c *fiber.Ctx) error {
	updateProduct := new(domain.Product)
	if err := c.BodyParser(updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// update process
	product, err := UpdateProduct(updateProduct.ID, updateProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(product)
}

func DeleteProductHandler(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = DeleteProduct(productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)

}
