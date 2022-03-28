package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"github.com/shinhagunn/shop-product/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Input []CartProduct

type Cart struct {
	Product  models.Product `json:"product"`
	Quantity int64          `json:"quantity"`
}

func HandleOrder(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)

	result := []models.Product{}

	ids := []primitive.ObjectID{}

	for _, product := range user.Cart {
		ids = append(ids, product.ProductID)
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}

	if err := collection.Product.SimpleFind(&result, filter); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	order := new(models.Order)

	order.UserID = user.ID
	order.Status = "Pending"

	for i, v := range result {
		order.OrderProduct = append(order.OrderProduct, models.OrderProduct{
			Name:  v.Name,
			Image: v.Image,
			Price: (v.Price * (1 - v.Discount)) * float64(user.Cart[i].Quantity),
		})
	}

	if err := collection.Order.Create(order); err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	services.DeliverProducer("deliver", order)

	return c.JSON(200)
}
