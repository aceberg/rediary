package web

import (
	// "log"
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

	// log.Println("MAP = ", totalByDay)

	i := 0
	for d := min; !d.After(max); d = d.AddDate(0, 0, 1) {
		heat.D = d.Format("2006-01-02")
		val, exists := totalByDay[heat.D]
		if exists {
			heat.V = val
		} else {
			heat.V = 0
		}
		heat.X = d.Weekday().String()[0:2]

		if heat.X == "Mo" && i == 0 {
			i = 1
		}
		if i != 0 {
			heat.Y = strconv.Itoa(i)
			heatMap = append(heatMap, heat)

			// log.Println("HEAT = ", heat)

			if heat.X == "Su" {
				i++
			}
		}
	}

	return heatMap
}
