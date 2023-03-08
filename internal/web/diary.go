package web

import (
	"net/http"

	"github.com/aceberg/red/internal/db"
	"github.com/aceberg/red/internal/models"
)

func diaryHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	AllRecords = db.Select(AppConfig.DB)
	guiData.Records = AllRecords

	execTemplate(w, "diary", guiData)
}
