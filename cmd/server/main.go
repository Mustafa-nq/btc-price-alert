package main

import (
	"context"
	"log"
	"os"

	"btc-price-alert/internal/db"
	"btc-price-alert/internal/routes"
	"btc-price-alert/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := db.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("Mongo connection error: ", err)
	}

	e := echo.New()
	routes.RegisterRoutes(e, client)

	c := cron.New()
	c.AddFunc("*/2 * * * *", func() { // every 2 minutes
		services.CheckAlerts(context.Background(), client)
	})
	c.Start()

	e.Logger.Fatal(e.Start(":8080"))

}
