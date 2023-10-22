package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
)

func UserSignUp(c *fiber.Ctx) error {
	user := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
	}{}

	c.BodyParser(&user)

	fmt.Println(user)

	if user.Email == "" || user.Password == "" || user.FirstName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "The fields shouldn't be empty"})
	}

	res := initializers.DB.Exec(`SELECT email FROM users WHERE email = ?`, user.Email)
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": res.Error})
	}
	if res.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "user with email already exist"})
	}

	err := helpers.SendOtp(user.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "OTP Error", "error": err})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "success", "message": "verify otp at /signup/verifyOtp"})
}

func VerifyOtp(c *fiber.Ctx) error {
	user := struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
		OTP       string `json:"otp"`
	}{}
	c.BodyParser(&user)

	status, err := helpers.VerifyOtp(user.Phone, user.OTP)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "OTP Error", "error": err})
	}
	if status != "approved" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "OTP is invalid"})
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}
	result.Row().Scan(&newUser)

	token, err := helpers.CreateToken(c, "User", time.Hour*24, newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "user": c.Locals("UserModel"), "token": token})
}

func UserLogin(c *fiber.Ctx) error {
	user := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	c.BodyParser(&user)

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Fields shouldn't be empty"})
	}

	dbUser := models.User{}
	result := initializers.DB.Raw(`SELECT * FROM users WHERE email = ?`, user.Email).Scan(&dbUser)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "No user registered with this email"})
	}

	if user.Password != dbUser.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "Wrong Password"})
	}

	token, err := helpers.CreateToken(c, "User", time.Hour*24, dbUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}

	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User Logged In successfully", "user": c.Locals("UserModel"), "token": token})
}

func GetDishPagewise(c *fiber.Ctx) error {
	dishList := []models.Dish{}
	page, err := strconv.ParseInt(c.Params("page"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Error occured while parsing URL params", "error": err.Error})
	}
	offset := (page - 1) * 10

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE deleted_at IS NULL LIMIT 10 OFFSET ?`, offset).Scan(&dishList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "dishList": dishList, "user": c.Locals("UserModel")})
}
