package identity

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-auth/config"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/controllers"
	"github.com/shinhagunn/shop-auth/models"
	"github.com/shinhagunn/shop-auth/services"
	"github.com/shinhagunn/shop-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	payload := new(RegisterPayload)

	if err := c.BodyParser(payload); err != nil {
		return controllers.ErrServerInvalidBody
	}

	user := &models.User{
		UID:      utils.RandomUID(),
		Email:    payload.Email,
		Password: services.HashPassword(payload.Password),
		State:    "Pending",
		Role:     "Member",
	}

	if err := collection.User.Create(user); err != nil {
		panic(err)
	}

	services.EmailProducer("new-user", user)

	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	session.Set("uid", user.UID)
	session.Save()

	return c.SendStatus(200)
}

func ReSendEmailCode(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)

	if err != nil {
		return controllers.ErrAuthzInvalidSession
	}

	uid := session.Get("uid")

	if uid == nil {
		return controllers.ErrAuthzInvalidSession
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"uid": uid}).Decode(user); err != nil {
		return controllers.ErrAuthzUserNotExist
	}

	if user.State != "Pending" {
		return controllers.ErrAuthzUserNotPending
	}

	// Check old code and change it's status
	codes := []models.Code{}
	if err := collection.Code.SimpleFind(&codes, bson.M{"user_id": user.ID}); err != nil {
		panic(err)
	}

	for _, code := range codes {
		code.State = "Delete"
		if err := collection.Code.Update(&code); err != nil {
			panic(err)
		}
	}

	// Create new code
	services.EmailProducer("new-user", user)

	return c.JSON(200)
}

type VerifyPayload struct {
	Code string `json:"code"`
}

func VerificationCode(c *fiber.Ctx) error {
	payload := new(VerifyPayload)

	if err := c.BodyParser(payload); err != nil {
		return controllers.ErrServerInvalidBody
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

	if err := collection.User.FindOne(context.Background(), bson.M{"uid": uid}).Decode(user); err != nil {
		return controllers.ErrAuthzUserNotExist
	}

	if user.State != "Pending" {
		return controllers.ErrAuthzUserNotPending
	}

	code := new(models.Code)
	if err := collection.Code.FindOne(context.Background(), bson.M{"user_id": user.ID, "state": "Active"}).Decode(&code); err != nil {
		panic(err)
	}

	if code.Code == "" {
		return controllers.ErrUnprocessableEntity
	}

	if code.Code != payload.Code {
		return controllers.ErrUnprocessableEntity
	}

	code.State = "Delete"
	if err := collection.Code.Update(code); err != nil {
		panic(err)
	}

	user.State = "Active"
	if err := collection.User.Update(user); err != nil {
		panic(err)
	}

	session.Set("uid", user.UID)
	session.Save()

	return c.JSON(200)
}
