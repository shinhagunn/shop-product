package middlewares

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-auth/config"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/controllers"
	"github.com/shinhagunn/shop-auth/models"
	"github.com/shinhagunn/shop-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func MustAuth(c *fiber.Ctx) error {
	path := strings.Replace(c.Path(), "/api/v2/auth", "", 1)
	if strings.Contains(path, "/api/v2/myauth/identity") || strings.Contains(path, "/api/v2/myauth/public") || strings.Contains(path, "api/v2/product/public") {
		return c.SendStatus(200)
	}

	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	uid := session.Get("uid")

	if uid == nil {
		return controllers.ErrAuthzInvalidSession
	}

	user := new(models.User)
	if err := collection.User.FindOne(context.Background(), bson.M{"uid": uid}).Decode(&user); err != nil {
		return controllers.ErrAuthzUserNotExist
	}

	return c.Next()
}

func MustPending(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)

	if user.State != "Pending" {
		return controllers.ErrAuthzUserNotPending
	}

	return c.Next()
}

func MustGuest(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	uid := session.Get("uid")

	if uid == nil {
		return c.Next()
	}

	return controllers.ErrAuthzUserNotGuest
}

func CheckRequest(c *fiber.Ctx) error {
	jwt_auth, err := utils.CheckJWT(strings.Replace(c.Get("Authorization"), "Bearer ", "", -1))

	if err != nil {
		return controllers.ErrJWTDecodeAndVerify
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"uid": jwt_auth.UID}).Decode(&user); err != nil {
		return controllers.ErrAuthzUserNotExist
	}

	c.Locals("CurrentUser", user)

	return c.Next()
}
