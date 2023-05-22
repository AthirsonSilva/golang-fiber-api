package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/models"
	"github.com/sixfwa/fiber-api/utils"
)

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User"
// @Success 201 {object} UserResponse
// @Router /api/v1/users [post]
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": user,
		})
	}

	user.ID = uuid.New().String()
	database.Database.Db.Create(&user)

	return c.Status(201).JSON(utils.CreateResponseUser(user))
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse
// @Router /api/v1/users [get]
func GetAllUser(c *fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "There is no user in the database",
		})
	}

	return c.Status(200).JSON(users)
}

// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Router /api/v1/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&user)

	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(200).JSON(utils.CreateResponseUser(user))
}

// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body User true "User"
// @Success 200 {object} UserResponse
// @Router /api/v1/users/{id} [put]
func FindUser(id string, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == "" {
		return errors.New("user does not exist")
	}

	return nil
}

// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body User true "User"
// @Success 201 {object} UserResponse
// @Router /api/v1/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	var user models.User

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&user)

	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": user,
		})
	}

	database.Database.Db.Save(&user)

	return c.Status(201).JSON(utils.CreateResponseUser(user))
}

// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Router /api/v1/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	var user models.User

	database.Database.Db.Where("id = ?", id).First(&user)

	if user.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	database.Database.Db.Delete(&user)

	return c.SendStatus(204)
}
