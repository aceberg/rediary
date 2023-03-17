package web

import (
	"log"
	"net/http"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/conf"
	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

// Gui - start web server
func Gui(config models.Conf) {

	AppConfig = conf.Get(config.ConfPath)
	AppConfig.ConfPath = config.ConfPath
	AppConfig.Icon = Icon

	db.Create(AppConfig.DB)
	AllRecords = db.Select(AppConfig.DB)

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port
	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add_action/", addActionHandler)
	http.HandleFunc("/add_record/", addRecordHandler)
	http.HandleFunc("/add_tag/", addTagHandler)
	http.HandleFunc("/clear/", clearHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/del_action/", delActionHandler)
	http.HandleFunc("/del_record/", delRecordHandler)
	http.HandleFunc("/del_tag/", delTagHandler)
	http.HandleFunc("/diary/", diaryHandler)
	http.HandleFunc("/diary_show/", diaryShowHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
