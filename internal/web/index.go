package web

import (
	"net/http"

	"github.com/aceberg/red/internal/db"
	"github.com/aceberg/red/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Themes = []string{"cerulean", "cosmo", "cyborg"}

	AllRecords = db.Select(AppConfig.DB)
	guiData.Records = AllRecords

	execTemplate(w, "index", guiData)
}
