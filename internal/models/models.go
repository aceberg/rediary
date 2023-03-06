package models

// Conf - web gui config
type Conf struct {
	DB       string
	Host     string
	Port     string
	Theme    string
	Icon     string
	ConfPath string
}

// Record - write to DB
type Record struct {
	ID   int    `db:"ID"`
	Date string `db:"DATE"`
	Name string `db:"NAME"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Themes []string
}
