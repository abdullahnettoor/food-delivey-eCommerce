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

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error Occured while fetching Restaurant", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  hashedPassword,
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

	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "verify otp @ /verifyOtp", "user": c.Locals("UserModel"), "token": token})
}

func VerifyOtp(c *fiber.Ctx) error {
	body := struct {
		OTP string `json:"otp"`
	}{}
	c.BodyParser(&body)

	u := c.Locals("UserModel").(map[string]interface{})

	status, err := helpers.VerifyOtp(fmt.Sprintf("%v", u["phone"]), body.OTP)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "OTP Error", "error": err})
	}
	if status != "approved" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "message": "OTP is invalid"})
	}

	result := initializers.DB.Exec(`UPDATE users SET status = 'Active' WHERE id = ?`, u["userId"])
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	var user models.User
	result = initializers.DB.Raw(`SELECT * FROM users WHERE id = ?`, u["userId"]).Scan(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "DB Error", "error": result.Error})
	}

	token, err := helpers.CreateToken(c, "User", time.Hour*24, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed!", "message": "JWT Error", "error": err})
	}
	c.Locals("UserModel", user)
	c.Cookie(&fiber.Cookie{Name: "Authorize User", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "user": c.Locals("UserModel"), "message": "User verified successfully"})
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

	if ok, err := helpers.CompareHashedPassword(dbUser.Password, user.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Bcrypt Error", "error": err})
	} else if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed!", "message": "Password is wrong"})
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
