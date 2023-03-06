package web

import (
	"log"
	"net/http"

	"github.com/aceberg/red/internal/check"
	"github.com/aceberg/red/internal/conf"
	"github.com/aceberg/red/internal/models"
)

// Gui - start web server
func Gui(config models.Conf) {

	AppConfig = conf.Get(config.ConfPath)
	AppConfig.Icon = Icon

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port
	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
