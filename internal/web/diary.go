package web

import (
	"net/http"

	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

func diaryHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	AllRecords = db.Select(AppConfig.DB)

	show := r.URL.Query().Get("show")
	tag := r.URL.Query().Get("tag")

	if tag == "" {
		guiData.Records = filterRecords(AllRecords, show)
		guiData.OneRec = countTotal(guiData.Records)
		guiData.Chart = countTags(guiData.Records)
	} else {
		guiData.Records = filterByTag(AllRecords, tag)
		guiData.OneRec = countTotal(guiData.Records)
		guiData.Chart = countNames(guiData.Records)
	}

	execTemplate(w, "diary", guiData)
}

func diaryShowHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	AllRecords = db.Select(AppConfig.DB)

	from := r.FormValue("from")
	to := r.FormValue("to")

	guiData.Records = fromDateToDate(AllRecords, from, to)

	guiData.OneRec = countTotal(guiData.Records)
	guiData.Chart = countTags(guiData.Records)

	execTemplate(w, "diary", guiData)
}
