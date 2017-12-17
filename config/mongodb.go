package confing

type (
	App struct {
		Name  string
		Port  uint
		Debug bool
	}

	Db struct {
		Host     string
		Port     uint
		User     string
		Pwd      string
		Database string
	}
)

var (
	APPCONFIG *App
	DBCONFIG  *Db
)

// init()は勝手に呼ばれる
func init() {
	APPCONFIG = &App{
		Name:  "Blog Manage",
		Port:  8006,
		Debug: true,
	}

	DBCONFIG = &Db{
		Host:     "192.168.59.103",
		Port:     27017,
		Database: "blog",
	}
}
