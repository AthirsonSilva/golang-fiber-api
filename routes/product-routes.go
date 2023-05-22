package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/models"
	"github.com/sixfwa/fiber-api/utils"
)

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body Product true "Product"
// @Success 201 {object} ProductResponse
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": product,
		})
	}

	product.ID = uuid.New().String()

	database.Database.Db.Create(&product)

	return c.Status(201).JSON(utils.CreateResponseProduct(product))
}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} ProductResponse
// @Router /api/v1/products [get]
func GetAllProduct(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)

	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "There is no product in the database",
		})
	}

	return c.Status(200).JSON(products)
}

// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} ProductResponse
// @Router /api/v1/products/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&product)

	if product.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return c.Status(200).JSON(utils.CreateResponseProduct(product))
}

// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body Product true "Product"
// @Success 200 {object} ProductResponse
// @Router /api/v1/products/{id} [put]
func FindProduct(id string, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == "" {
		return errors.New("product does not exist")
	}

	return nil
}

// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body Product true "Product"
// @Success 200 {object} ProductResponse
// @Router /api/v1/products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&product)

	if product.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": product,
		})
	}

	database.Database.Db.Save(&product)

	return c.Status(200).JSON(utils.CreateResponseProduct(product))
}

// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Router /api/v1/products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&product)

	if product.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	database.Database.Db.Delete(&product)

	return c.SendStatus(204)
}
