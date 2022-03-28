package admin

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryPayload struct {
	Name string `json:"name"`
}

func CreateCategory(c *fiber.Ctx) error {
	payload := new(CategoryPayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	category := &models.Category{
		Name: payload.Name,
	}

	if err := collection.Category.Create(category); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	payload := c.Params("id")

	if payload == "" {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	CategoryID, err := primitive.ObjectIDFromHex(payload)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	category := new(models.Category)
	collection.Category.FindOne(context.Background(), bson.M{"_id": CategoryID}).Decode(&category)

	if category.Name == "" {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	if err := collection.Category.Delete(category); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(200)
}
