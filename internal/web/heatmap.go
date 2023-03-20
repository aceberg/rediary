package web

import (
	"strconv"
	"time"

	"github.com/aceberg/rediary/internal/models"
)

func generateHeatmap(records []models.Record) []models.HeatMap {
	var heatMap []models.HeatMap
	var heat models.HeatMap
	totalByDay := make(map[string]int)

	max := time.Now()
	min := max.AddDate(0, 0, -35)

	monthRecords := fromDateToDate(records, min.Format("2006-01-02"), max.Format("2006-01-02"))

	for _, rec := range monthRecords {
		val, exists := totalByDay[rec.Date]
		if exists {
			totalByDay[rec.Date] = val + rec.Total
		} else {
			totalByDay[rec.Date] = rec.Total
		}
	}

	i := 5
	for d := max; !d.Before(min); d = d.AddDate(0, 0, -1) {
		heat.D = d.Format("2006-01-02")
		val, exists := totalByDay[heat.D]
		if exists {
			heat.V = val
		} else {
			heat.V = 0
		}
		heat.X = d.Weekday().String()[0:2]
		heat.Y = strconv.Itoa(i)

		heatMap = append(heatMap, heat)

		if heat.X == "Mo" {
			i--
		}
		if i < 1 {
			break
		}
	}

	return heatMap
}
