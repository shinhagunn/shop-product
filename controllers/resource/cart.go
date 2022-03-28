package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func findIndexProductInCart(cart []models.CartProduct, productID string) (int, string) {
	for index, product := range cart {
		id, err := primitive.ObjectIDFromHex(productID)

		if err != nil {
			return 0, controllers.ServerInternalError
		}

		if product.ProductID == id {
			return index, ""
		}
	}

	return 0, controllers.FailedConnectDataInDatabase
}

type CartRespone struct {
	Product  models.Product `json:"product"`
	Quantity int64          `json:"quantity"`
}

func GetCartProducts(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)
	respone := []CartRespone{}
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

	for i, v := range result {
		respone = append(respone, CartRespone{
			Product:  v,
			Quantity: user.Cart[i].Quantity,
		})
	}

	return c.JSON(respone)
}

type ProductToCartPayload struct {
	ProductID string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

func AddProductToCart(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)

	payload := new(ProductToCartPayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	productID, err := primitive.ObjectIDFromHex(payload.ProductID)

	if err != nil {
		return c.Status(500).JSON(controllers.ServerInternalError)
	}

	index, error := findIndexProductInCart(user.Cart, payload.ProductID)

	if error == controllers.ServerInternalError {
		return c.Status(500).JSON(err)
	}

	if error != controllers.FailedConnectDataInDatabase {
		user.Cart[index].Quantity += payload.Quantity
	} else {
		cartProduct := new(models.CartProduct)

		cartProduct.ProductID = productID
		cartProduct.Quantity = payload.Quantity

		user.Cart = append(user.Cart, *cartProduct)
	}

	if err := collection.User.Update(user); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	return c.JSON(200)
}

func RemoveIndexCart(s []models.CartProduct, index int) []models.CartProduct {
	return append(s[:index], s[index+1:]...)
}

func RemoveProductInCart(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)

	payload := c.Params("id")

	if payload == "" {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	index, err := findIndexProductInCart(user.Cart, payload)

	if err != "" {
		return c.Status(500).JSON(err)
	}

	user.Cart = RemoveIndexCart(user.Cart, index)

	if err := collection.User.Update(user); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	return c.JSON(200)
}
