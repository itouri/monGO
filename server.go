package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"fmt"

	"./handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type dbinfo struct {
	Name       string
	Collection string
	Session    mgo.Session
	collection mgo.Collection
}

// そうとう危険な気がする．これはどこでやるのが正しい？
var (
	db         = "test"
	collection = "stamp"
	sess, _    = mgo.Dial("127.0.0.1:27017")
	conn       = sess.DB(db).C(collection)
)

//----------
// Handlers
//----------
func getListPrefecture(c echo.Context) error {
	// 構造体の配列
	spots := []spot{}

	// TODO goのパラメータでJSON返してる
	err := conn.Find(bson.M{}).Select(bson.M{"prefecture": 1}).All(&spots)
	if err != nil {
		return c.JSON(http.StatusOK, "Find:"+err.Error())
	}

	return c.JSON(http.StatusOK, spots)
}

func getList(c echo.Context) error {
	retJSON := &spot{}

	// idをObjectID型に変換
	idStr := c.Param("id")
	if !bson.IsObjectIdHex(idStr) {
		return c.JSON(http.StatusOK, "id can not convert to ObjectID")
	}
	id := bson.ObjectIdHex(idStr)

	// TODO goのパラメータでJSON返してる
	err := conn.FindId(id).One(&retJSON)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, retJSON)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// これを入れないと OPTION のメソッドがさばけずエラーになる
	// TODO これがない時のエラー原因を理解する
	e.Use(middleware.CORS())

	e.GET("/api/list/prefecture", getListPrefecture)
	e.GET("/api/list", getList)

	e.GET("/api/spot/:id", handlers.GetSpot)
	e.POST("/api/spot", handlers.PostSpot)

	e.File("/", "public/index.html")

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":1323")
}
