package web

import (
	"net/http"
	"time"

	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	AllRecords = db.Select(AppConfig.DB)
	guiData.Heat = generateHeatmap(AllRecords)

	today := time.Now().Format("2006-01-02")

	guiData.Records = fromDateToDate(AllRecords, today, today)
	guiData.OneRec = countTotal(guiData.Records)
	guiData.OneRec.Date = today

	execTemplate(w, "index", guiData)
}
