package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/rediary/internal/auth"
	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/models"
)

// Get - read config from file or env
func Get(path string) (models.Conf, auth.Conf) {
	var config models.Conf
	var actions []models.Action
	var tags []models.TagType

	var authConf auth.Conf

	viper.SetDefault("DB", "/data/rediary/sqlite.db")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8847")
	viper.SetDefault("THEME", "minty")
	viper.SetDefault("BGCOLOR", "light")
	viper.SetDefault("COLORPLUS", "#ff3300")
	viper.SetDefault("COLORMINUS", "#00aeff")
	viper.SetDefault("MOODMAX", "5")
	viper.SetDefault("AUTH_USER", "")
	viper.SetDefault("AUTH_PASSWORD", "")
	viper.SetDefault("AUTH_EXPIRE", "7d")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DB, _ = viper.Get("DB").(string)
	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.BgColor, _ = viper.Get("BGCOLOR").(string)
	config.ColorPlus, _ = viper.Get("COLORPLUS").(string)
	config.ColorMinus, _ = viper.Get("COLORMINUS").(string)
	config.MoodMax, _ = viper.Get("MOODMAX").(string)
	authConf.Auth = viper.GetBool("AUTH")
	authConf.User, _ = viper.Get("AUTH_USER").(string)
	authConf.Password, _ = viper.Get("AUTH_PASSWORD").(string)
	authConf.ExpStr, _ = viper.Get("AUTH_EXPIRE").(string)

	authConf.Expire = auth.ToTime(authConf.ExpStr)
	config.Auth = authConf.Auth

	err = viper.UnmarshalKey("actions", &actions)
	check.IfError(err)
	config.Actions = actions

	err = viper.UnmarshalKey("tags", &tags)
	check.IfError(err)
	config.TagMap = structToMap(tags)

	return config, authConf
}

// Write - write config to file
func Write(config models.Conf, authConf auth.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("db", config.DB)
	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("BGCOLOR", config.BgColor)
	viper.Set("colorminus", config.ColorMinus)
	viper.Set("colorplus", config.ColorPlus)
	viper.Set("moodmax", config.MoodMax)
	viper.Set("actions", config.Actions)
	viper.Set("tags", mapToStruct(config.TagMap))

	viper.Set("auth", authConf.Auth)
	viper.Set("auth_user", authConf.User)
	viper.Set("auth_password", authConf.Password)
	viper.Set("auth_expire", authConf.ExpStr)

	err := viper.WriteConfig()
	check.IfError(err)
}
