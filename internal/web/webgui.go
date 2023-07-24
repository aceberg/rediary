package web

import (
	"log"
	"net/http"

	"github.com/aceberg/rediary/internal/auth"
	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/conf"
	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

// Gui - start web server
func Gui(config models.Conf) {

	AppConfig, authConf = conf.Get(config.ConfPath)
	AppConfig.ConfPath = config.ConfPath
	AppConfig.Icon = Icon

	db.Create(AppConfig.DB)
	db.Migrate(AppConfig.DB)
	AllRecords = db.Select(AppConfig.DB)

	log.Println("INFO: starting web gui with", AppConfig.ConfPath)

	address := AppConfig.Host + ":" + AppConfig.Port
	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/login/", loginHandler)

	http.HandleFunc("/", auth.Auth(indexHandler, &authConf))
	http.HandleFunc("/add_action/", auth.Auth(addActionHandler, &authConf))
	http.HandleFunc("/add_record/", auth.Auth(addRecordHandler, &authConf))
	http.HandleFunc("/add_tag/", auth.Auth(addTagHandler, &authConf))
	http.HandleFunc("/auth_conf/", auth.Auth(authConfHandler, &authConf))
	http.HandleFunc("/auth_save/", auth.Auth(saveAuthHandler, &authConf))
	http.HandleFunc("/clear/", auth.Auth(clearHandler, &authConf))
	http.HandleFunc("/config/", auth.Auth(configHandler, &authConf))
	http.HandleFunc("/del_action/", auth.Auth(delActionHandler, &authConf))
	http.HandleFunc("/del_record/", auth.Auth(delRecordHandler, &authConf))
	http.HandleFunc("/del_tag/", auth.Auth(delTagHandler, &authConf))
	http.HandleFunc("/diary/", auth.Auth(diaryHandler, &authConf))
	http.HandleFunc("/diary_show/", auth.Auth(diaryShowHandler, &authConf))
	http.HandleFunc("/save_colors/", auth.Auth(saveColorsHandler, &authConf))
	http.HandleFunc("/save_config/", auth.Auth(saveConfigHandler, &authConf))
	http.HandleFunc("/stat/", auth.Auth(statHandler, &authConf))
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
