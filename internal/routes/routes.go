package routes

import (
	"btc-price-alert/internal/handlers"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(e *echo.Echo, Client *mongo.Client) {
	h := &handlers.AlertHandler{DB: Client}

	e.POST("/alerts", h.Create)
	e.GET("/alerts", h.GetAll)
}
