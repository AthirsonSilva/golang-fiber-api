package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/models"
	"github.com/sixfwa/fiber-api/utils"
)

// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body Order true "Order"
// @Success 201 {object} OrderResponse
// @Router /api/v1/orders [post]
func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	var user models.User
	var product models.Product

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": order,
		})
	}

	if err := FindUser(order.UserID, &user); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"user":    order.UserID,
		})
	} else if err := FindProduct(order.ProductID, &product); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"product": order.ProductID,
		})
	}

	order.ID = uuid.New().String()
	database.Database.Db.Create(&order)

	responseUser := utils.CreateResponseUser(user)
	responseProduct := utils.CreateResponseProduct(product)
	responseOrder := utils.CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(201).JSON(responseOrder)
}

// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} OrderResponse
// @Router /api/v1/orders [get]
func GetAllOrder(c *fiber.Ctx) error {
	var orders []models.Order

	database.Database.Db.Find(&orders)

	if len(orders) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "There is no order in the database",
		})
	}

	return c.Status(200).JSON(orders)
}

// @Summary Get a order
// @Description Get a order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} OrderResponse
// @Router /api/v1/orders/{id} [get]
func GetOrder(c *fiber.Ctx) error {
	var order models.Order

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&order)

	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	return c.Status(200).JSON(order)
}

// @Summary Update a order
// @Description Update a order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body Order true "Order"
// @Success 200 {object} OrderResponse
// @Router /api/v1/orders/{id} [put]
func UpdateOrder(c *fiber.Ctx) error {
	var order models.Order
	var user models.User
	var product models.Product

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&order)

	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"object": order,
		})
	}

	if err := FindUser(order.UserID, &user); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"user":    order.UserID,
		})
	} else if err := FindProduct(order.ProductID, &product); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"product": order.ProductID,
		})
	}

	database.Database.Db.Save(&order)

	responseUser := utils.CreateResponseUser(user)
	responseProduct := utils.CreateResponseProduct(product)
	responseOrder := utils.CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

// @Summary Delete a order
// @Description Delete a order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 204
// @Router /api/v1/orders/{id} [delete]
func DeleteOrder(c *fiber.Ctx) error {
	var order models.Order

	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID is required",
		})
	}

	database.Database.Db.Where("id = ?", id).First(&order)

	if order.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	database.Database.Db.Delete(&order)

	return c.SendStatus(204)
}
