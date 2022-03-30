package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/controllers"
	"github.com/shinhagunn/shop-auth/models"
)

type MessagePayload struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Message  string `json:"message"`
}

func CustommerMessage(c *fiber.Ctx) error {
	payload := new(MessagePayload)

	if err := c.BodyParser(payload); err != nil {
		return controllers.ErrServerInvalidBody
	}

	custommer := &models.Custommer{
		Fullname: payload.Fullname,
		Email:    payload.Email,
		Phone:    payload.Phone,
		Address:  payload.Address,
		Message:  payload.Message,
	}

	if err := collection.Custommer.Create(custommer); err != nil {
		panic(err)
	}

	return c.JSON(200)
}
