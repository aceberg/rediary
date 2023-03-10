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

	countTagMap := make(map[string]int)
	for _, rec := range guiData.Records {
		_, exists := countTagMap[rec.Tag]
		if exists {
			countTagMap[rec.Tag] = countTagMap[rec.Tag] + rec.Total
		} else {
			countTagMap[rec.Tag] = rec.Total
		}
	}

	for tag, count := range countTagMap {
		guiData.Chart.Tag = append(guiData.Chart.Tag, tag)
		guiData.Chart.Color = append(guiData.Chart.Color, AppConfig.TagMap[tag])
		guiData.Chart.Count = append(guiData.Chart.Count, count)
	}

	execTemplate(w, "diary", guiData)
}
