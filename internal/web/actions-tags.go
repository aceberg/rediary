package web

import (
	"net/http"
	"sort"

	"github.com/aceberg/rediary/internal/conf"
	// "github.com/aceberg/rediary/internal/models"
)

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
