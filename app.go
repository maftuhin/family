package main

import (
	"database/sql"
	"family/database"
	"family/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Run Application
func main() {
	app := fiber.New(fiber.Config{})
	database.DBConn = gormDb()
	app.Static("/public", "./public")
	routes.SetUpRoute(app)

	app.Post("", func(c *fiber.Ctx) error {
		return c.SendString("")
	})

	log.Fatal(app.Listen(":8000"))
}

func gormDb() *gorm.DB {
	sqlDB, _ := sql.Open("mysql", "maftuhin:8Belas0694@@tcp(103.150.196.232:3306)/family")
	// sqlDB, _ := sql.Open("mysql", "root:8Belas0694@@tcp(localhost:3306)/family")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalln("s", err)
	}
	return gormDB
}
