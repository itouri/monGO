package main

import (
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	dbinfo struct {
		Name       string
		Collection string
		Session    mgo.Session
		collection mgo.Collection
	}

	spot struct {
		// ID          bson.ObjectId `bson:"_id"`
		// TODO `bson:"spot_name"`のときのメンバの変数名はSpotNameではない
		Spot_name   string    `bson:"spot_name"`
		User_id     int       `bson:"user_id"`
		Image_ids   []int     `bson:"image_ids"`
		Latitude    float64   `bson:"latitude"`
		Longitude   float64   `bson:"longitude"`
		Prefecture  string    `bson:"prefecture"`
		City        string    `bson:"city"`
		Description string    `bson:"description"`
		Hint        string    `bson:"hint"`
		Favorites   int       `bson:"favorites"`
		Created     time.Time `bson:"created"`
		Modified    time.Time `bson:"modified"`
	}
)

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

func getSpot(c echo.Context) error {
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

func postSpot(c echo.Context) error {
	spots := new(spot)
	if err := c.Bind(spots); err != nil {
		return c.JSON(http.StatusBadRequest, "Bind: "+err.Error())
	}

	err := conn.Insert(spots)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Insert: "+err.Error())
	}
	return c.NoContent(http.StatusOK)
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

	e.GET("/api/spot/:id", getSpot)
	e.POST("/api/spot", postSpot)

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":1323")
}
