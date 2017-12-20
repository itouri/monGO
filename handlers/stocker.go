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

func checkStock(name string) string {
	stocker := new(models.Stocker)
	if name != "" {
		query := bson.M{"name": name}
		result := new(models.Stocker)
		err := stocker.Find(stocker.ColName(), query).One(&result)
		if err != nil {
			log.Fatalf("FINDONE: " + err.Error())
		}
		return result.Name + ": " + strconv.Itoa(result.Amount) + "\n"
	}
	results := []models.Stocker{}
	query := bson.M{"amount": bson.M{"$gt": 0}}
	err := stocker.Find(stocker.ColName(), query).All(&results)
	if err != nil {
		log.Fatalf("FINDALL: " + err.Error())
	}
	var retStr string
	for _, result := range results {
		retStr += result.Name + ": " + strconv.Itoa(result.Amount) + "\n"
	}
	return retStr
}

func sell(name string, amount int, price float64) {
	stocker := new(models.Stocker)
	selector := bson.M{"name": name}
	upsert := bson.M{"$inc": bson.M{"amount": -amount}}
	_, err := stocker.Upsert(selector, upsert)
	if err != nil {
		log.Fatalf("UPSERT: " + err.Error())
	}

	inst := new(models.Sell)
	selector = bson.M{"sell": bson.M{"$gte": 0}}
	upsert = bson.M{"$inc": bson.M{"sell": float64(amount) * price}}
	_, err = inst.Upsert(selector, upsert)
	if err != nil {
		log.Fatalf("UPSERT: " + err.Error())
	}
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
		if err != nil {
			return c.String(http.StatusOK, err.Error())
		}
	}

	pricePram := c.QueryParam("price")
	var price float64
	if pricePram != "" {
		price, err = strconv.ParseFloat(pricePram, 64)
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
		return c.String(http.StatusOK, checkStock(name))
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
