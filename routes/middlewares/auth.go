package middlewares

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"github.com/shinhagunn/shop-product/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func MustAdmin(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)

	if user.Role != "Admin" {
		return c.Status(401).JSON(controllers.AuthzInvalidPermissionErr)
	}

	return c.Next()
}

func CheckRequest(c *fiber.Ctx) error {
	jwt_auth, err := utils.CheckJWT(strings.Replace(c.Get("Authorization"), "Bearer ", "", -1))

	if err != nil {
		return c.Status(500).JSON(controllers.FailedToParseJWT)
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"uid": jwt_auth.UID}).Decode(&user); err != nil {
		user = &models.User{
			UID:   jwt_auth.UID,
			Email: jwt_auth.Email,
			Role:  jwt_auth.Role,
			Cart:  []models.CartProduct{},
		}

		if err := collection.User.Create(user); err != nil {
			return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
		}
	} else {
		if err := collection.User.Update(user); err != nil {
			return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
		}
	}

	c.Locals("CurrentUser", user)

	return c.Next()
}
