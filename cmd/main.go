package main

import (
	"github.com/labstack/echo/v4"
	"submission_promotion_api/internal/app/repositories"
	"submission_promotion_api/internal/app/services"
	"submission_promotion_api/internal/configs"
	"submission_promotion_api/internal/delivery"
)

func main() {

	configs.LoadViperEnv()

	db := configs.InitDatabase()

	e := echo.New()

	PromotionRepo := repositories.NewPromotionRepository(db)

	PromoService := services.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}
