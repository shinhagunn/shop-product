package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
)

func GetUserCurrent(c *fiber.Ctx) error {

	user := c.Locals("CurrentUser").(*models.User)

	return c.JSON(user)
}

type ProfilePayload struct {
	Fullname string `json:"fullname"`
	Age      string `json:"age"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
}

func UpdateUserProfile(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(*models.User)
	payload := new(ProfilePayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	if payload.Fullname == "" {
		user.UserProfile.Fullname = "none"
	} else {
		user.UserProfile.Fullname = payload.Fullname
	}

	if payload.Address == "" {
		user.UserProfile.Address = "none"
	} else {
		user.UserProfile.Address = payload.Address
	}

	if payload.Age == "" {
		user.UserProfile.Age = "none"
	} else {
		user.UserProfile.Age = payload.Age
	}

	if payload.Gender == "" {
		user.UserProfile.Gender = "none"
	} else {
		user.UserProfile.Gender = payload.Gender
	}

	if payload.Phone == "" {
		user.UserProfile.Phone = "none"
	} else {
		user.UserProfile.Phone = payload.Phone
	}

	if err := collection.User.Update(user); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	return c.JSON(200)
}
