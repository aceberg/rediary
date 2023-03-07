package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/red/internal/check"
	"github.com/aceberg/red/internal/models"
	"github.com/aceberg/red/internal/web"
)

const confPath = "/data/red/config.yaml"

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr

	check.Path(conf.ConfPath)

	web.Gui(conf)
}
