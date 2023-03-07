package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/red/internal/check"
	"github.com/aceberg/red/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("DB", "/data/red/sqlite.db")
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
	config.Actions = viper.GetStringSlice("ACTIONS")

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

	err := viper.WriteConfig()
	check.IfError(err)
}
