package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"gopkg.in/mgo.v2/bson"

	"../models"
)

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
			// return c.String(http.StatusOK, err.Error())
			return c.String(http.StatusOK, "ERROR")
		}
	}

	pricePram := c.QueryParam("price")

	f := c.QueryParam("function")
	var str string
	switch f {
	case "addstock":
		err = addStock(name, amount)
	case "checkstock":
		str, err = checkStock(name)
	case "sell":
		err = sell(name, amount, pricePram)
	case "checksales":
		str, err = checkSales()
	case "deleteall":
		deleteAll()
	default:
		// log.Println("ERROR:no function")
		err = fmt.Errorf("ERROR:no function")
	}

	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	if str != "" {
		// return c.String(http.StatusOK, "OK: "+str)
		return c.String(http.StatusOK, str)
	}
	// return c.String(http.StatusOK, "OK")
	return c.NoContent(http.StatusOK)

}

func addStock(name string, amount int) error {
	// validation
	if name == "" {
		// return c.String(http.StatusOK, "ERROR:addsotck")
		return fmt.Errorf("name must not nil")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must greater then 0")
	}

	stocker := new(models.Stocker)
	selector := bson.M{"name": name}
	upsert := bson.M{"$inc": bson.M{"amount": amount}}
	_, err := stocker.Upsert(selector, upsert)
	if err != nil {
		return err
	}
	return nil
}

func checkStock(name string) (string, error) {
	stocker := new(models.Stocker)
	if name != "" {
		query := bson.M{"name": name}
		result := new(models.Stocker)
		err := stocker.Find(stocker.ColName(), query).One(&result)
		if err != nil {
			return "", err
		}
		retStr := result.Name + ": " + strconv.Itoa(result.Amount) + "\n"
		return retStr, nil
	}
	results := []models.Stocker{}
	query := bson.M{"amount": bson.M{"$gt": 0}}
	err := stocker.Find(stocker.ColName(), query).All(&results)
	if err != nil {
		return "", err
	}
	var retStr string
	for _, result := range results {
		retStr += result.Name + ": " + strconv.Itoa(result.Amount) + "\n"
	}
	return retStr, nil
}

func sell(name string, amount int, pricePram string) error {
	// validation
	if name == "" {
		// return c.String(http.StatusOK, "ERROR:addsotck")
		return fmt.Errorf("name must not nil")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must greater then 0")
	}

	stocker := new(models.Stocker)

	// Check exsistence of name
	selector := bson.M{"name": name}
	err := stocker.Find(stocker.ColName(), selector).One(&stocker)
	if err != nil {
		return err
	}

	if stocker.Amount < amount {
		return fmt.Errorf("Amount was over commited")
	}

	upsert := bson.M{"$inc": bson.M{"amount": -amount}}
	_, err = stocker.Upsert(selector, upsert)
	if err != nil {
		return err
	}

	var price float64
	if pricePram != "" {
		price, err = strconv.ParseFloat(pricePram, 64)
		if err != nil {
			// return c.String(http.StatusOK, err.Error())
			return err
		}

		if price <= 0 {
			return fmt.Errorf("price must greater then 0")
		}

		inst := new(models.Sell)
		selector = bson.M{"sell": bson.M{"$gte": 0}}
		upsert = bson.M{"$inc": bson.M{"sell": float64(amount) * price}}
		_, err = inst.Upsert(selector, upsert)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkSales() (string, error) {
	inst := new(models.Sell)
	err := inst.Find(inst.ColName(), nil).One(&inst)
	if err != nil {
		return "", err
	}
	// 12.3456
	val := inst.Sell
	// 1234.56
	val *= 100
	// 1235.0
	val = math.Ceil(val)
	// 12.35
	val /= 100

	return fmt.Sprintf("sales: %v\n", val), nil
}

func deleteAll() {
	stocker := new(models.Stocker)
	stocker.DeleteAll()
	inst := new(models.Sell)
	inst.DeleteAll()
}
