package models

// Conf - web gui config
type Conf struct {
	DB       string
	Host     string
	Port     string
	Theme    string
	Icon     string
	ConfPath string
	Actions  []Action
	TagMap   map[string]string
}

// Action - one action
type Action struct {
	Name string `mapstructure:"name"`
	Tag  string `mapstructure:"tag"`
}

// ChartJS - data for charts
type ChartJS struct {
	Tag   []string
	Color []string
	Count []int
}

// Record - write to DB
type Record struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Tag   string `db:"TAG"`
	Name  string `db:"NAME"`
	Color string `db:"COLOR"`
	Minus int    `db:"MINUS"`
	Plus  int    `db:"PLUS"`
	Total int    `db:"TOTAL"`
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Records []Record
	OneRec  Record
	Themes  []string
	Chart   ChartJS
}
