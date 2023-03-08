package main

import (
	"flag"

	_ "time/tzdata"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/models"
	"github.com/aceberg/rediary/internal/web"
)

const confPath = "/data/rediary/config.yaml"

func main() {
	var conf models.Conf

	confPtr := flag.String("c", confPath, "Path to config yaml file")
	flag.Parse()

	conf.ConfPath = *confPtr

	check.Path(conf.ConfPath)

	web.Gui(conf)
}
