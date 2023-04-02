package web

import (
	"time"

	"github.com/aceberg/rediary/internal/models"
)

func filterRecords(records []models.Record, show string) []models.Record {
	var filtered []models.Record

	currentTime := time.Now()
	max := currentTime.Format("2006-01-02")

	switch show {
	case "week":
		min := currentTime.AddDate(0, 0, -7).Format("2006-01-02")
		filtered = fromDateToDate(records, min, max)
	case "month":
		min := currentTime.AddDate(0, -1, 0).Format("2006-01-02")
		filtered = fromDateToDate(records, min, max)
	default:
		filtered = records
	}

	return filtered
}

func fromDateToDate(records []models.Record, min, max string) []models.Record {
	var filtered []models.Record

	if max == "" {
		currentTime := time.Now()
		max = currentTime.Format("2006-01-02")
	}

	for _, rec := range records {
		if rec.Date >= min && rec.Date <= max {
			filtered = append(filtered, rec)
		}
	}

	return filtered
}

func filterByTag(records []models.Record, tag string) []models.Record {
	var filtered []models.Record

	for _, rec := range records {
		if rec.Tag == tag {
			filtered = append(filtered, rec)
		}
	}

	return filtered
}
