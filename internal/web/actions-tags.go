package web

import (
	// "log"
	"net/http"
	"sort"
	"strings"

	"github.com/aceberg/rediary/internal/conf"
	"github.com/aceberg/rediary/internal/models"
)

func addActionHandler(w http.ResponseWriter, r *http.Request) {
	var action models.Action

	action.Name = r.FormValue("name")
	action.Tag = r.FormValue("tag")

	if action.Name != "" {
		AppConfig.Actions = append(AppConfig.Actions, action)

		sort.Slice(AppConfig.Actions, func(i, j int) bool {
			return AppConfig.Actions[i].Name < AppConfig.Actions[j].Name
		})

		sort.Slice(AppConfig.Actions, func(i, j int) bool {
			return AppConfig.Actions[i].Tag < AppConfig.Actions[j].Tag
		})

		conf.Write(AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delActionHandler(w http.ResponseWriter, r *http.Request) {
	var newActions []models.Action

	tagname := r.FormValue("name")
	tnslice := strings.Split(tagname, ":")
	tag := tnslice[0]
	name := tnslice[1]

	for _, a := range AppConfig.Actions {
		if a.Tag != tag || a.Name != name {
			newActions = append(newActions, a)
		}
	}
	AppConfig.Actions = newActions

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func addTagHandler(w http.ResponseWriter, r *http.Request) {

	tag := r.FormValue("tag")
	color := r.FormValue("color")

	if tag != "" {
		AppConfig.TagMap[tag] = color
		conf.Write(AppConfig)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delTagHandler(w http.ResponseWriter, r *http.Request) {

	tag := r.FormValue("tag")

	delete(AppConfig.TagMap, tag)

	conf.Write(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
