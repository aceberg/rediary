package web

import (
	"net/http"

	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

func statHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	AllRecords = db.Select(AppConfig.DB)

	execTemplate(w, "stat", guiData)
}
