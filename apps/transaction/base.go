package transaction

import (
	infrafiber "nbid-online-shop/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	trxRoute := router.Group("transactions")
	{
		// menggunakan middleware
		trxRoute.Use(infrafiber.CheckAuth())

		// route dibawahnya akan menggunakan middleware tersebut
		trxRoute.Post("/checkout", handler.CreateTransaction)
		trxRoute.Get("/user/histories", handler.GetTransactionByUser)
	}
}
