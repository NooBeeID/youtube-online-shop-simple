package main

import (
	"log"
	"nbid-online-shop/apps/auth"
	"nbid-online-shop/apps/product"
	"nbid-online-shop/external/database"
	"nbid-online-shop/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	product.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
