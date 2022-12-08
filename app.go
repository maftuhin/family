package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Run Application
func main() {
	app := fiber.New(fiber.Config{})
	// database.DBConn = gormDb()
	app.Static("/public", "./public")
	// routes.SetUpRoute(app)

	app.Post("", func(c *fiber.Ctx) error {
		return c.SendString("")
	})

	log.Fatal(app.Listen(":8000"))
}

func gormDb() *gorm.DB {
	// sqlDB, err := sql.Open("mysql", "remote:8Belas0694@tcp(limaefdua.com:3306)/family")
	sqlDB, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/family")
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
