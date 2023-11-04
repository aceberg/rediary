package models

import (
	"github.com/aceberg/rediary/internal/auth"
)

// Conf - web gui config
type Conf struct {
	DB         string
	Host       string
	Port       string
	Theme      string
	BgColor    string
	Icon       string
	ConfPath   string
	ColorPlus  string
	ColorMinus string
	MoodMax    string
	Actions    []Action
	TagMap     map[string]string
	Auth       bool
}

// Action - one action
type Action struct {
	Name string `mapstructure:"name"`
	Tag  string `mapstructure:"tag"`
}

// TagType - one tag
type TagType struct {
	Name  string `mapstructure:"name"`
	Color string `mapstructure:"color"`
}

// ChartJS - data for charts
type ChartJS struct {
	Tag   []string
	Color []string
	Count []int
}

// HeatMap - data for heatmap
type HeatMap struct {
	X string
	Y string
	D string
	V int
}

// Record - write to DB
type Record struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Tag   string `db:"TAG"`
	Name  string `db:"NAME"`
	Color string `db:"COLOR"`
	Note  string `db:"NOTE"`
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
	Heat    []HeatMap
	Auth    auth.Conf
}
