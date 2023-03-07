package web

import (
	"net/http"
	"time"

	"github.com/aceberg/red/internal/db"
	"github.com/aceberg/red/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	AllRecords = db.Select(AppConfig.DB)

	currentTime := time.Now()

	guiData.OneRec.Date = currentTime.Format("2006-01-02")

	for _, rec := range AllRecords {
		if rec.Date == guiData.OneRec.Date {
			guiData.Records = append(guiData.Records, rec)
			guiData.OneRec.Minus = guiData.OneRec.Minus + rec.Minus
			guiData.OneRec.Plus = guiData.OneRec.Plus + rec.Plus
			guiData.OneRec.Total = guiData.OneRec.Total + rec.Total
		}
	}

	execTemplate(w, "index", guiData)
}
