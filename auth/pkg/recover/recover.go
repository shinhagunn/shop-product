package recover

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shinhagunn/shop-auth/controllers"
)

type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// EnableStackTrace enables handling stack trace
	//
	// Optional. Default: false
	EnableStackTrace bool

	// StackTraceHandler defines a function to handle stack trace
	//
	// Optional. Default: defaultStackTraceHandler
	StackTraceHandler func(c *fiber.Ctx, e interface{}) string
}

func New(config Config) fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				msg := config.StackTraceHandler(c, r)

				log.Println(msg)

				var ok bool
				if err, ok = r.(error); !ok {
					// Set error that will call the global error handler
					err = controllers.ErrServerInternal
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
