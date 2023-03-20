package web

import (
	"log"
	"net/http"

	"github.com/aceberg/rediary/internal/conf"
	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
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

func saveColorsHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.ColorMinus = r.FormValue("minus")
	AppConfig.ColorPlus = r.FormValue("plus")

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("INFO: deleting all records from DB")

	db.Clear(AppConfig.DB)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
