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

func CreateSlide(c *fiber.Ctx) error {
	slide := new(models.Slide)

	if err := c.BodyParser(slide); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	if err := collection.Slide.Create(slide); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(slide)
}

func UpdateSlide(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	payload := new(models.Slide)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	SlideID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	slide := new(models.Slide)
	collection.Slide.FindOne(context.Background(), bson.M{"_id": SlideID}).Decode(&slide)

	if slide.Title == "" {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	if payload.Description != "" {
		slide.Description = payload.Description
	}

	if payload.Title != "" {
		slide.Title = payload.Title
	}

	if payload.Image != "" {
		slide.Image = payload.Image
	}

	if err := collection.Slide.Update(slide); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(slide)
}

func DeleteSlide(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	SlideID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	slide := new(models.Slide)

	collection.Slide.FindOne(context.Background(), bson.M{"_id": SlideID}).Decode(slide)

	if slide.Title == "" {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	if err := collection.Slide.Delete(slide); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(200)
}
