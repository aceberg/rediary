package web

import (
	"log"
	"net/http"
	"sort"

	"github.com/aceberg/red/internal/conf"
	"github.com/aceberg/red/internal/db"
	"github.com/aceberg/red/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.DB = r.FormValue("db")
	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func addActionHandler(w http.ResponseWriter, r *http.Request) {

	action := r.FormValue("name")

	if action != "" {
		AppConfig.Actions = append(AppConfig.Actions, action)
		sort.Strings(AppConfig.Actions)
		conf.Write(AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delActionHandler(w http.ResponseWriter, r *http.Request) {
	var newActions []string

	action := r.FormValue("name")

	for _, a := range AppConfig.Actions {
		if a != action {
			newActions = append(newActions, a)
		}
	}
	AppConfig.Actions = newActions

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: deleting all records from DB")

	db.Clear(AppConfig.DB)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
