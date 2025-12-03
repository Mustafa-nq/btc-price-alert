package handlers

import (
	"btc-price-alert/internal/models"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlertHandler struct {
	DB *mongo.Client
}

func (h *AlertHandler) Create(c echo.Context) error {
	var alert models.Alert
	if err := c.Bind(&alert); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := h.DB.Database("crypto").Collection("alerts")

	res, err := collection.InsertOne(context.Background(), alert)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Assign the inserted ID to alert.ID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		alert.ID = oid
	}

	return c.JSON(http.StatusOK, alert)
}

func (h *AlertHandler) GetAll(c echo.Context) error {

	cur, err := h.DB.Database("crypto").Collection("alerts").Find(context.Background(), bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var alerts []models.Alert
	cur.All(context.Background(), &alerts)

	return c.JSON(http.StatusOK, alerts)
}
