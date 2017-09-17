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
	spot struct {
		// ID          bson.ObjectId `bson:"_id"`
		// TODO `bson:"spot_name"`のときのメンバの変数名はSpotNameではない
		Spotname string `bson:"spotname"`
		UserID   int    `bson:"user_id"`
		// ImageIDs    []int     `bson:"image_ids"`
		// Latitude    float64   `bson:"latitude"`
		// Longtitude  float64   `bson:"longtitude"`
		Prefecture  string    `bson:"prefecture"`
		City        string    `bson:"city"`
		Description string    `bson:"description"`
		Hint        string    `bson:"hint"`
		Favorites   int       `bson:"favorites"`
		Created     time.Time `bson:"created"`
		Modified    time.Time `bson:"modified"`
	}
)

// var (
// 	tablename = "userinfo"
// 	seq       = 1
// 	conn, _   = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/test", nil)
// 	sess      = conn.NewSession(nil)
// )

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
func getSpot(c echo.Context) error {
	retJSON := &spot{}

	// idをObjectID型に変換
	idStr := c.Param("id")
	if !bson.IsObjectIdHex(idStr) {
		return c.JSON(http.StatusOK, "id can not convert to ObjectID")
	}
	id := bson.ObjectIdHex(idStr)

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
	// user_id, err := strconv.Atoi(c.Param("user_id"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// latitude, err := strconv.ParseFloat(c.Param("latitude"), 32)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// longitude, err := strconv.ParseFloat(c.Param("longitude"), 32)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	// insJSON := &spot{
	// 	Name: c.Param("spot_name"),
	// 	// UserID:      user_id,
	// 	// Latitude:    latitude,
	// 	// Longtitude:  longitude,
	// 	Prefecture:  c.Param("prefecture"),
	// 	City:        c.Param("city"),
	// 	Description: c.Param("description"),
	// 	Hint:        c.Param("hint"),
	// }

	// err := conn.Insert(insJSON)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
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

	e.GET("/api/spot/:id", getSpot)
	e.POST("/api/spot", postSpot)

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":1323")
}
