package web

import (
	"github.com/aceberg/rediary/internal/models"
)

func countTotal(records []models.Record) models.Record {
	var total models.Record

	for _, rec := range records {
		total.Minus = total.Minus + rec.Minus
		total.Plus = total.Plus + rec.Plus
		total.Total = total.Total + rec.Total
	}

	return total
}

func countTags(records []models.Record) models.ChartJS {
	var chart models.ChartJS

	countTagMap := make(map[string]int)

	for _, rec := range records {
		_, exists := countTagMap[rec.Tag]
		if exists {
			countTagMap[rec.Tag] = countTagMap[rec.Tag] + rec.Total
		} else {
			countTagMap[rec.Tag] = rec.Total
		}
	}

	for tag, count := range countTagMap {
		chart.Tag = append(chart.Tag, tag)
		chart.Color = append(chart.Color, AppConfig.TagMap[tag])
		chart.Count = append(chart.Count, count)
	}

	return chart
}

func countNames(records []models.Record) models.ChartJS {
	var chart models.ChartJS

	countTagMap := make(map[string]int)

	for _, rec := range records {
		_, exists := countTagMap[rec.Name]
		if exists {
			countTagMap[rec.Name] = countTagMap[rec.Name] + rec.Total
		} else {
			countTagMap[rec.Name] = rec.Total
		}
	}

	for tag, count := range countTagMap {
		chart.Tag = append(chart.Tag, tag)
		chart.Count = append(chart.Count, count)
	}

	return chart
}

func countCloud(records []models.Record) models.ChartJS {
	var chart models.ChartJS

	countTagMap := make(map[string]int)

	for _, rec := range records {
		_, exists := countTagMap[rec.Name]
		if exists {
			countTagMap[rec.Name] = countTagMap[rec.Name] + 1
		} else {
			countTagMap[rec.Name] = 1
		}
	}

	for tag, count := range countTagMap {
		chart.Tag = append(chart.Tag, tag)
		chart.Count = append(chart.Count, count*10)
	}

	return chart
}
