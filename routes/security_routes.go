package routes

import (
	configs "fiber-mongo-api/config"
	"fiber-mongo-api/controllers" //add this

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SecuirtyRoute(app *fiber.App) {

	private := app.Group("/secure")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.EnvJWTSecret()),
	}))
	private.Post("/verify", controllers.VerifyJWT)

	app.Get("/jwt", controllers.GetJWT)
}
