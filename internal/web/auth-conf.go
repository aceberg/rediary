package web

import (
	// "log"
	"net/http"

	// "github.com/aceberg/rediary/internal/auth"
	"github.com/aceberg/rediary/internal/conf"
	"github.com/aceberg/rediary/internal/models"
)

func authConfHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Auth.Auth = authConf.Auth
	guiData.Auth.User = authConf.User
	guiData.Auth.ExpStr = authConf.ExpStr

	execTemplate(w, "auth", guiData)
}

func saveAuthHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.DB = r.FormValue("db")
	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.BgColor = r.FormValue("bgcolor")

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
