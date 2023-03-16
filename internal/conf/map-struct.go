package conf

import (
	"github.com/aceberg/rediary/internal/models"
)

func structToMap(tagStruct []models.TagType) map[string]string {

	tagMap := make(map[string]string)

	for _, tag := range tagStruct {

		tagMap[tag.Name] = tag.Color
	}

	return tagMap
}

func mapToStruct(tagMap map[string]string) []models.TagType {

	var tagStruct []models.TagType
	var tag models.TagType

	for key, value := range tagMap {
		tag.Name = key
		tag.Color = value
		tagStruct = append(tagStruct, tag)
	}

	return tagStruct
}
