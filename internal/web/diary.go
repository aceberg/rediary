package web

import (
	"net/http"

	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

func diaryHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	show := r.URL.Query().Get("show")
	tag := r.URL.Query().Get("tag")

	if tag == "" {
		if show == "" {
			AllRecords = db.Select(AppConfig.DB)
			show = "week"
		}
		guiData.Records = filterRecords(AllRecords, show)
		guiData.OneRec = countTotal(guiData.Records)
		guiData.Chart = countTags(guiData.Records)
	} else {
		AllRecords = filterByTag(AllRecords, tag)
		guiData.Records = filterRecords(AllRecords, "week")
		guiData.OneRec = countTotal(guiData.Records)
		guiData.Chart = countNames(guiData.Records)
	}

	execTemplate(w, "diary", guiData)
}

func diaryShowHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	guiData.Records = fromDateToDate(AllRecords, from, to)

	guiData.OneRec = countTotal(guiData.Records)
	guiData.Chart = countTags(guiData.Records)

	execTemplate(w, "diary", guiData)
}
