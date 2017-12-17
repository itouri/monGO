package confing

type (
	Config struct {
		App *App
		Db  *Db
	}

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

func (c *Config) Init() {
	c.App = &App{
		Name:  "Blog Manage",
		Port:  8006,
		Debug: true,
	}

	c.Db = &Db{
		Host:     "192.168.59.103",
		Port:     27017,
		Database: "blog",
	}
}
