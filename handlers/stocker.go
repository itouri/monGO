package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"../models"

	"github.com/labstack/echo"
)

func addStock(name string, amount int) error {
	stocker := new(models.Stocker)
	selector := bson.M{"name": name}
	upsert := bson.M{"$inc": bson.M{"amount": amount}}
	_, err := stocker.Upsert(selector, upsert)
	if err != nil {
		log.Fatalf("UPSERT: " + err.Error())
	}
	return nil
}

func checkStock(name string) {

}

func sell(name string, amount int, price float64) {

}

func checkSales() {

}

func deleteAll() {
	stocker := new(models.Stocker)
	stocker.DeleteAll()
}

func GetStocker(c echo.Context) error {
	name := c.QueryParam("name")

	amountPram := c.QueryParam("amount")
	var amount int
	var err error
	if amountPram == "" {
		amount = 1
	} else {
		amount, err = strconv.Atoi(amountPram)
		fmt.Print(amount)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
	}

	pricePram := c.QueryParam("price")
	var price float64
	if pricePram != "" {
		price, err = strconv.ParseFloat(pricePram, 64)
		fmt.Print(price)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
	}
	f := c.QueryParam("function")
	switch f {
	case "addstock":
		if name == "" {
			return c.String(http.StatusOK, "ERROR:addsotck")
		}
		err := addStock(name, amount)
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
	case "checkstock":
		checkStock(name)
	case "sell":
		if name == "" {
			return c.String(http.StatusOK, "ERROR")
		}
		sell(name, amount, price)
	case "checksales":
		checkSales()
	case "deleteall":
		deleteAll()
	default:
		return c.String(http.StatusOK, "ERROR:no function")
	}
	return c.String(http.StatusOK, "OK")
}
