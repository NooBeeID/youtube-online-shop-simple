package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRotue := router.Group("products")
	{
		productRotue.Get("", handler.GetListProducts)
		productRotue.Post("", handler.CreateProduct)
		productRotue.Get("/sku/:sku", handler.GetProductDetail)
	}
}
