package controller

import (
	"family/database"
	"family/models"

	"github.com/gofiber/fiber/v2"
)

func FamilyTree(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var person models.Person
	db.First(&person, "uid=?", id)
	// father
	var father models.Person
	db.First(&father, "uid=?", person.Father)
	// mother
	var mother models.Person
	db.First(&mother, "uid=?", person.Mother)
	// spouse
	var spouse models.Person
	db.First(&spouse, "uid=?", person.Spouse)
	// brother
	var brothers []models.Person
	db.Where("father=? AND mother=? AND uid!=?", person.Father, person.Mother, person.UID).Limit(10).Find(&brothers)
	// child
	var child []models.Person
	db.Where("father=? OR mother=?", person.UID, person.UID).Limit(10).Find(&child)

	return c.JSON(fiber.Map{
		"id":       person.ID,
		"uid":      person.UID,
		"name":     person.Name,
		"image":    person.Image,
		"gender":   person.Gender,
		"address":  person.Address,
		"dad":      father,
		"mom":      mother,
		"partner":  spouse,
		"brothers": brothers,
		"children": child,
	})
}
