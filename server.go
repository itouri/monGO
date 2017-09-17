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
		ID          bson.ObjectId `bson:"_id"`
		Name        string        `bson:"spot_name"`
		UserID      int           `bson:"user_id"`
		ImageIDs    []int         `bson:"image_ids"`
		Latitude    float64       `bson:"latitude"`
		Longtitude  float64       `bson:"longtitude"`
		Prefecture  string        `bson:"prefecture"`
		City        string        `bson:"city"`
		Description string        `bson:"description"`
		Hint        string        `bson:"hint"`
		Favorites   int           `bson:"favorites"`
		Created     time.Time     `bson:"created"`
		Modified    time.Time     `bson:"modified"`
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

// func postSpot(c echo.Context) error {

// }

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/spot/:id", getSpot)
	//e.POST("/api/spot/:id", postSpot)

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":1323")
}
