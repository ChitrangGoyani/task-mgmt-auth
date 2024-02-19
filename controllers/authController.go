package controllers

import (
	"strconv"
	"time"

	"github.com/ChitrangGoyani/task-mgmt-auth/database"
	"github.com/ChitrangGoyani/task-mgmt-auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = "secret"

func GetUser(c *fiber.Ctx) error {
	// Query to postgres from params and get the user row based on the id taken from the cookies
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(&fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	return c.Status(200).JSON(user)
}

func Login(c *fiber.Ctx) error {
	// login a user into the account by comparing the password hash from the database
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(&fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadGateway)
		return c.JSON(&fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // alive for a day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(&fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true, // front-end should not be able to access the cookie
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"message": "Successful Login",
	})
}

func Register(c *fiber.Ctx) error {
	// Create a user and put it into the database, return 200 or 500
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(404).SendString("Could not hash password")
	}
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)
	return c.Status(200).JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Create a dummy cookie with the same name, expire it and return
	cookie := fiber.Cookie{
		Name:     "task-mgmt-jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
	})
}
