package handlers

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetAdminLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": "admin login loaded",
	})
}

func AdminLogin(c *fiber.Ctx) error {
	Body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	AdminDetails := models.Admin{}

	c.BodyParser(&Body)

	fmt.Println("From Request", Body)

	if Body.Email == "" {
		fmt.Println("Email shouldn't be empty")
		return c.JSON(fiber.Map{"failed": "Email field shouldn't be empty"})
	}

	result := initializers.DB.Raw(`SELECT * FROM admins WHERE email = ?`, Body.Email).Scan(&AdminDetails)

	if result.Error != nil {
		fmt.Println("Error Occured while fetching Admin", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "DB Error"})
	}

	if result.RowsAffected < 1 {
		fmt.Println("Admin with provided email don't exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "no admin exist with email entered"})
	}

	fmt.Println("From DB", AdminDetails)

	if Body.Email != AdminDetails.Email || Body.Password != AdminDetails.Password {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"failed": "Invalid Email or Password"})
	}

	token, err := helpers.CreateToken(c, "Admin", time.Hour*24, AdminDetails)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"failed": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Admin", Value: token})

	c.Locals("adminDetails", AdminDetails)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  token,
		// "admin":  c.Locals("userModel"),
		"admin": AdminDetails,
	})
}

func AdminDashboard(c *fiber.Ctx) error {
	fmt.Println("FibDboordLocaaaalsss", c.Locals("adminDetails"))

	return c.JSON(fiber.Map{
		"status":    "success",
		"dashboard": "dashboard data will be generated here",
		"admin":     c.Locals("userModel"),
	})

}
