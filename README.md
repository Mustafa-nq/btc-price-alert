A lightweight API service built with Go, Echo, MongoDB, and Cron Jobs that allows users to register price alerts for BTC/ETH and triggers a notification when the price crosses their defined threshold.

This project includes:

 -REST API (Echo Framework)
 -MongoDB storage
 -Cron job price checker
 -Console alerts when price conditions are met
 -Modular folder structure
 -Ready to extend with Email/Telegram notifications

Features

API Features
-Create crypto alert rules
-Fetch all existing alert rules
-Supports BTC & ETH

Alerts based on:
above a price
below a price

Installation
1. Clone the repository
git clone https://github.com/Mustafa-nq/btc-price-alert.git
cd btc-price-alert

2. Install dependencies
go mod tidy

3. Set up MongoDB
Run Mongo locally or use Docker:
docker run -d -p 27017:27017 --name mongo mongo

4. Run the server
go run cmd/server/main.go

Server starts on:
http://localhost:8080

API Endpoints

Create Alert
POST /alerts

Request Body:
{
  "coin": "BTC",
  "price": 90000,
  "direction": "above",
  "email": "user@example.com"
}

Get All Alerts
GET /alerts


