package controller

import (
	"family/database"
	"family/models"
	"family/tools/paging"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePerson(c *fiber.Ctx) error {
	db := database.DBConn
	var body models.Person
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString(err.Error())
	}
	body.UID = uuid.New().String()
	create := db.Create(&body)
	if create.Error == nil {
		return c.JSON(body)
	} else {
		return c.JSON(fiber.Map{"message": create.Error})
	}
}

func SearchPerson(c *fiber.Ctx) error {
	db := database.DBConn
	// param
	page, _ := strconv.Atoi(c.Query("page", "1"))
	query := c.Query("q")
	// sql
	var member []models.Person
	sql := db.Where("name like ?", "%"+query+"%")
	// paging
	paginator := paging.Paging(&paging.Param{
		DB:    sql,
		Page:  page,
		Limit: 10,
	}, member)
	// result
	return c.JSON(paginator)
}

func PersonDetail(c *fiber.Ctx) error {
	db := database.DBConn
	uuid := c.Params("id")

	var person models.Person

	result := db.First(&person, "uuid=?", uuid)
	if result.RowsAffected > 0 {
		return c.JSON(person)
	}
	return c.Status(400).JSON(&fiber.Map{"message": "Tidak Ditemukan"})
}

func UpdatePerson(c *fiber.Ctx) error {
	db := database.DBConn
	var body models.Person
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString(err.Error())
	}
	result := db.Model(&body).Updates(body)
	if result.Error == nil {
		return c.JSON(body)
	} else {
		return c.JSON(result.Error)
	}
}
