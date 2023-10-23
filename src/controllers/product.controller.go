package controllers

import (
	"strconv"
	"strings"
	"time"

	"thex-products/initializers"
	"thex-products/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandlerHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Golang, Fiber, SQLite, and GORM",
	})
}

func ListProducts(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var products []models.Product
	results := initializers.DB.Limit(intLimit).Offset(offset).Find(&products)
	
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(products), "products": products})
}

func GetProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")

	var product models.Product
	result := initializers.DB.First(&product, "id = ?", productId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No product with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"product": product}})
}

func CreateProduct(c *fiber.Ctx) error {
	var payload *models.CreateProductSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()

	newProduct := models.Product{
		Name:     		payload.Name,
		Description:   	payload.Description,
		Price:  		payload.Price,
		Image: 			payload.Image,
		Category: 		payload.Category,
		Published: 		payload.Published,
		CreatedAt: 		now,
		UpdatedAt: 		now,
	}

	result := initializers.DB.Create(&newProduct)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(
			fiber.Map{
				"status": "fail", 
				"message": "Title already exist, please use another title",
			})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"status": "error", 
				"message": result.Error.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"status": "success", 
			"data": fiber.Map{
				"product": newProduct,
			},
		})
}

func UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")

	var payload *models.UpdateProductSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var product models.Product
	result := initializers.DB.First(&product, "id = ?", productId)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No product with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Name != "" {
		updates["name"] = payload.Name
	}
	if payload.Description != "" {
		updates["description"] = payload.Description
	}
	if payload.Category != "" {
		updates["category"] = payload.Category
	}

	updates["updated_at"] = time.Now()

	initializers.DB.Model(&product).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"product": product}})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")

	result := initializers.DB.Delete(&models.Product{}, "id = ?", productId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No product with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
