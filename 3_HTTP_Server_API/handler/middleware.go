package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/Jessadakorn-pun/go-example/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CheckMiddleware(c *fiber.Ctx) error {

	// log middleware
	start := time.Now()
	fmt.Printf("URL = %s, MEHOD = %s, TIME = %s \n", c.OriginalURL(), c.Method(), start)

	// read from jwtware local context and convert to jwt token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// check role from jwt token
	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	user := new(domain.User)
	// check if request and parse
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != domain.MemberUser.Email || user.Password != domain.MemberUser.Password {
		return fiber.ErrUnauthorized // Unauthorized to use api when user + password wrong
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   t,
	})
}
