package identity

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-auth/config"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/controllers"
	"github.com/shinhagunn/shop-auth/models"
	"github.com/shinhagunn/shop-auth/services"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	passwordWrongErr = "password is incorrect"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	payload := new(LoginPayload)

	if err := c.BodyParser(payload); err != nil {
		return controllers.ErrServerInvalidBody
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"email": payload.Email}).Decode(&user); err != nil {
		return controllers.ErrAuthzUserNotExist
	}

	if user.Role == "" {
		return controllers.ErrAuthzPermissionDenied
	}

	if user.State == "Delete" {
		return controllers.ErrAuthzPermissionDenied
	}

	if user.State == "Banned" {
		return controllers.ErrAuthzPermissionDenied
	}

	if result := services.CheckPasswordHash(payload.Password, user.Password); !result {
		return controllers.ErrUnprocessableEntity
	}

	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	session.Set("uid", user.UID)
	session.Save()

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	session.Destroy()
	session.Save()

	return c.JSON(200)
}
