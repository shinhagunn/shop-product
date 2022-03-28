package resource

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentPayload struct {
	Content string `json:"content"`
}

func CreateComment(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	ProductID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	product := new(models.Product)

	if err := collection.Product.FindOne(context.Background(), bson.M{"_id": ProductID}).Decode(&product); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	user := c.Locals("CurrentUser").(*models.User)

	payload := new(CommentPayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	comment := models.Comment{
		UserID:       user.ID,
		ProductID:    product.ID,
		Content:      payload.Content,
		LikeCount:    0,
		DislikeCount: 0,
	}

	if err := collection.Comment.Create(&comment); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(200)
}

func LikeComment(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	CommentID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	comment := new(models.Comment)

	if err := collection.Comment.FindOne(context.Background(), bson.M{"_id": CommentID}).Decode(&comment); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	comment.LikeCount++

	if err := collection.Comment.Update(comment); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(200)
}

func DislikeComment(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	CommentID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	comment := new(models.Comment)

	if err := collection.Comment.FindOne(context.Background(), bson.M{"_id": CommentID}).Decode(&comment); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	comment.DislikeCount++

	if err := collection.Comment.Update(comment); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(200)
}

func GetCommentsInProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	ProductID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	comments := []models.Comment{}

	if err := collection.Comment.SimpleFind(&comments, bson.M{"product_id": ProductID}); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	return c.JSON(comments)
}
