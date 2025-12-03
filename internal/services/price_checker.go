package services

import (
	"btc-price-alert/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PriceResponse struct {
	Bitcoin  float64 `json:"bitcoin"`
	Ethereum float64 `json:"ethereum"`
}

func getPrices() (*PriceResponse, error) {

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]map[string]float64
	json.NewDecoder(resp.Body).Decode(&data)

	return &PriceResponse{
		Bitcoin:  data["bitcoin"]["usd"],
		Ethereum: data["ethereum"]["usd"],
	}, nil

}

func CheckAlerts(ctx context.Context, client *mongo.Client) {
	prices, err := getPrices()
	if err != nil {
		return
	}

	col := client.Database("crypto").Collection("Alerts")
	cur, _ := col.Find(ctx, bson.M{})

	var alerts []models.Alert
	cur.All(ctx, &alerts)

	for _, a := range alerts {
		current := prices.Bitcoin
		if a.Coin == "ETH" {
			current = prices.Ethereum
		}

		trigger := (a.Direction == "above" && current > a.Price) || (a.Direction == "below" && current < a.Price)

		if trigger {
			//Send email logic here
			fmt.Printf("Alert! %s is %s %.2f USD (current: %.2f USD)  \n ", a.Coin, a.Direction, a.Price, current)
		}

	}

}
