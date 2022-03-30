package routes

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shinhagunn/shop-auth/config"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/controllers"
	"github.com/shinhagunn/shop-auth/controllers/identity"
	"github.com/shinhagunn/shop-auth/controllers/public"
	"github.com/shinhagunn/shop-auth/controllers/resource"
	"github.com/shinhagunn/shop-auth/models"
	"github.com/shinhagunn/shop-auth/pkg/recover"
	"github.com/shinhagunn/shop-auth/routes/middlewares"
	"github.com/shinhagunn/shop-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func InitRouter() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "server.internal_error"

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			return c.Status(code).JSON(message)
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) string {
			buf := make([]byte, 2048)
			buf = buf[:runtime.Stack(buf, false)]

			return fmt.Sprintf("panic: %v\n%s\n", e, buf)
		},
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// /api/v2/auth/*
	app.All("/api/v2/auth/*", middlewares.MustAuth, func(c *fiber.Ctx) error {
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

		if user.State != "Active" {
			return controllers.ErrAuthzUserNotActive
		}

		jwt_token, err := utils.GenerateJWT(user)

		if err != nil {
			return controllers.ErrJWTDecodeAndVerify
		}

		jwt_token = "Bearer " + jwt_token
		c.Set("Authorization", jwt_token)

		return c.SendStatus(200)
	})

	api_identity := app.Group("/api/v2/identity")
	{
		// Login
		api_identity.Post("/login", middlewares.MustGuest, identity.Login)
		// Logout
		api_identity.Get("/logout", middlewares.MustAuth, identity.Logout)
		// Register
		api_identity.Post("/register", middlewares.MustGuest, identity.Register)

		// Resend email code
		api_identity.Get("/resendemail", identity.ReSendEmailCode)
		// Verify code
		api_identity.Post("/verifycode", identity.VerificationCode)
	}

	api_public := app.Group("/api/v2/public")
	{
		// Post message custommer
		api_public.Post("/sendmessage", public.CustommerMessage)
	}

	api_resource := app.Group("/api/v2/resource", middlewares.CheckRequest)
	{
		// Update Password
		api_resource.Post("/user/update/password", resource.UpdatePassword)
	}

	app.Listen(":3001")
}
