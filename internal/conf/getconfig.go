package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf
	var actions []models.Action
	var tags []models.TagType

	viper.SetDefault("DB", "/data/rediary/sqlite.db")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8847")
	viper.SetDefault("THEME", "minty")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DB, _ = viper.Get("DB").(string)
	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)

	err = viper.UnmarshalKey("actions", &actions)
	check.IfError(err)
	config.Actions = actions

	err = viper.UnmarshalKey("tags", &tags)
	check.IfError(err)
	config.TagMap = structToMap(tags)

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("db", config.DB)
	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("actions", config.Actions)
	viper.Set("tags", mapToStruct(config.TagMap))

	err := viper.WriteConfig()
	check.IfError(err)
}
