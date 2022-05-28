package main

import (
	configs "fiber-mongo-api/config"
	"fiber-mongo-api/routes" //add this

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//add the routes by passing the app to the UserRoute function
	routes.UserRoute(app) //add this
	routes.SecuirtyRoute(app)

	app.Listen(":6000")
}
