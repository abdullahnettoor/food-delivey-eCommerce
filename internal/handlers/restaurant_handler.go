package handlers

import (
	"fmt"
	"time"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/helpers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
)

func RestuarantSignUp(c *fiber.Ctx) error {
	Body := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}{}

	c.BodyParser(&Body)

	if Body.Email == "" || Body.Password == "" || Body.Name == "" || Body.Description == "" {
		fmt.Println("All fields should be filled")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed! All fields should be filled"})
	}
	fmt.Println("Finding email", Body.Email)
	result := initializers.DB.Exec(`SELECT email FROM restaurants WHERE email = ?`, Body.Email)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}
	if result.RowsAffected != 0 {
		fmt.Println("Restaurant with provided email already exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed! Restaurant with email entered already exist"})
	}

	restaurant := models.Restaurant{
		Email:       Body.Email,
		Password:    Body.Password,
		Name:        Body.Name,
		Description: Body.Description,
	}

	result = initializers.DB.Create(&restaurant)
	if result.Error != nil {
		fmt.Println("Error occured while creating new Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}

	result = initializers.DB.Raw(`SELECT * FROM restaurants WHERE email = ?`, Body.Email).Scan(&restaurant)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}

	token, err := helpers.CreateToken(c, "Restaurant", time.Hour*24, restaurant)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! JWT Error", "error": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Restaurant", Value: token})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"message":    "You can add items to sell, only after the verification made by admin",
		"restaurant": c.Locals("RestaurantModel"),
	})
}

func RestaurantLogin(c *fiber.Ctx) error {
	Body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	RestaurantDetails := models.Restaurant{}
	c.BodyParser(&Body)

	fmt.Println("From Request", Body)

	if Body.Email == "" {
		fmt.Println("Email shouldn't be empty")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Failed! Email field shouldn't be empty"})
	}

	result := initializers.DB.Raw(`SELECT * FROM restaurants WHERE email = ?`, Body.Email).Scan(&RestaurantDetails)
	if result.Error != nil {
		fmt.Println("Error Occured while fetching Restaurant", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}

	if result.RowsAffected < 1 {
		fmt.Println("Restaurant with provided email don't exist")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed! No Restaurant exist with email entered"})
	}

	fmt.Println("From DB", RestaurantDetails)

	if Body.Email != RestaurantDetails.Email || Body.Password != RestaurantDetails.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed! Invalid Email or Password"})
	}

	token, err := helpers.CreateToken(c, "Restaurant", time.Hour*24, RestaurantDetails)
	if err != nil {
		fmt.Println("Error Creating token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! JWT Error", "error": err})
	}
	fmt.Println("Token created")
	c.Cookie(&fiber.Cookie{Name: "Authorize Restaurant", Value: token})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"token":      token,
		"restaurant": c.Locals("RestaurantModel"),
	})
}

func RestaurantDashboard(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":     "success",
		"restaurant": c.Locals("RestaurantModel"),
		"dashboard":  "Dasboard data will be passed here",
	})
}

func AddDish(c *fiber.Ctx) error {

	dish := models.Dish{}
	c.BodyParser(&dish)

	if 0 > dish.Price || dish.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed! Given datas are invalid"})
	}

	// Add new dish to DB
	dishId := initializers.DB.Create(&dish)
	if dishId.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": dishId.Error})
	}

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE id = ?`, dish.ID).Scan(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": dishId.Error})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Dish Added", "dish": dish, "restaurant": c.Locals("RestaurantModel")})
}

func GetDishes(c *fiber.Ctx) error {
	restaurant := c.Locals("RestaurantModel").(map[string]interface{})
	dishesList := []models.Dish{}

	fmt.Println("Restaurant ID is", restaurant["restaurantId"])

	result := initializers.DB.Raw(`SELECT * FROM dishes WHERE restaurant_id = ? AND deleted_at IS NULL`, restaurant["restaurantId"]).Scan(&dishesList)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed! DB Error", "error": result.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "dishes": dishesList, "restaurant": c.Locals("RestaurantModel")})
}

func DeleteDish(c *fiber.Ctx) error {
	dishId := c.Params("id")

	fmt.Println("Dish ID to Delete", dishId)

	result := initializers.DB.Exec(`UPDATE dishes SET deleted_at = NOW() WHERE id = ?`, dishId)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed! DB Error",
			"error":  result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Dish deleted"})
}
