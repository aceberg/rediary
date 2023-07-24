package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/rediary/internal/auth"
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

	authConf.User = r.FormValue("user")
	authConf.ExpStr = r.FormValue("expire")

	authStr := r.FormValue("auth")
	pw := r.FormValue("password")

	if authStr == "on" {
		authConf.Auth = true
	} else {
		authConf.Auth = false
	}
	AppConfig.Auth = authConf.Auth

	if pw != "" {
		authConf.Password = auth.HashPassword(pw)
	}

	authConf.Expire = auth.ToTime(authConf.ExpStr)

	conf.Write(AppConfig, authConf)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
