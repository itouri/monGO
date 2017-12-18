package handlers

import (
	"net/http"

	"../models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func GetSpot(c echo.Context) error {
	retJSON := &models.Spot{}

	// idをObjectID型に変換
	idStr := c.Param("id")
	if !bson.IsObjectIdHex(idStr) {
		return c.JSON(http.StatusOK, "id can not convert to ObjectID")
	}
	id := bson.ObjectIdHex(idStr)

	// err := conn.FindId(id).One(&retJSON)
	// TODO goのパラメータでJSON返してる
	spot := new(models.Spot)
	err := spot.FindId("unicorns", id).One(&retJSON)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, retJSON)
}

func PostSpot(c echo.Context) error {
	spot := new(models.Spot)
	if err := c.Bind(spot); err != nil {
		return c.JSON(http.StatusBadRequest, "Bind: "+err.Error())
	}

	// err := conn.Insert(models.spots)
	err := spot.Insert()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Insert: "+err.Error())
	}
	return c.NoContent(http.StatusOK)
}
