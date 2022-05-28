package controllers

import (
	"fiber-mongo-api/security"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetJWT(c *fiber.Ctx) error {
	token, exp, err := security.CreateJWTToken()
	if err != nil {
		log.Panic("Failed creating jwt", err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": token, "exp": exp})
}

func VerifyJWT(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": true})

}
