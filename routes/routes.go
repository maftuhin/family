package routes

import (
	"family/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(app *fiber.App) {
	app.Post("person/create", controller.CreatePerson)
	app.Post("person/update/:id", controller.UpdatePerson)
	app.Get("person/search", controller.SearchPerson)
	app.Get("person/:id", controller.PersonDetail)

	app.Get("family/tree/:id", controller.FamilyTree)
	app.Get("family/fix", controller.FixNullValue)
}
