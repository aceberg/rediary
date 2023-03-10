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
	guiData.Records = AllRecords

	for _, rec := range AllRecords {
		guiData.OneRec.Minus = guiData.OneRec.Minus + rec.Minus
		guiData.OneRec.Plus = guiData.OneRec.Plus + rec.Plus
		guiData.OneRec.Total = guiData.OneRec.Total + rec.Total
	}

	execTemplate(w, "diary", guiData)
}
