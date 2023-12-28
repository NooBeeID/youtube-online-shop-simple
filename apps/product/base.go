package product

import (
	"nbid-online-shop/apps/auth"
	infrafiber "nbid-online-shop/infra/fiber"

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
		productRotue.Get("/sku/:sku", handler.GetProductDetail)

		// need authorization

		productRotue.Post("",
			infrafiber.CheckAuth(),
			infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}),
			handler.CreateProduct,
		)
	}
}
