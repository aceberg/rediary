package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/red/internal/check"
	"github.com/aceberg/red/internal/models"
	"github.com/aceberg/red/internal/web"
)

const confPath = "/data/red/config.yaml"
const dbPath = ""

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	dbPtr := flag.String("d", dbPath, "Path to DB file")
	flag.Parse()

	conf.ConfPath = *confPtr
	conf.DB = *dbPtr

	check.Path(conf.ConfPath)

	web.Gui(conf)
}
