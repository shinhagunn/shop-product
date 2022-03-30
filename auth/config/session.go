package config

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var SessionStore *session.Store

func InitSessionStore() {
	storage := redis.New(redis.Config{
		Host:  "127.0.0.1",
		Port:  6379,
		Reset: false,
	})

	SessionStore = session.New(session.Config{
		Storage:        storage,
		Expiration:     7 * time.Hour,
		CookiePath:     "/",
		CookieHTTPOnly: true,
	})

	log.Println("Initialize session success")
}
