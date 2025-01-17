package config

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("port", "8080")
	viper.SetDefault("admin-username", "vpakh")
	viper.SetDefault("group-name", "golang_ua_official")
	viper.SetDefault("bot-name", "GolangUA Butler Local")
	viper.SetDefault("bot-description", "Бот для адміністрування української спільноти Go.")

	viper.MustBindEnv("port", "PORT")
	viper.MustBindEnv("bot-token", "BOT_TOKEN")
	viper.MustBindEnv("webhook-url", "WEBHOOK_URL")
	viper.MustBindEnv("project-id", "PROJECT_ID")
	viper.MustBindEnv("project-region", "PROJECT_REGION")
	viper.MustBindEnv("k-service", "K_SERVICE")
	viper.MustBindEnv("log-level", "LOG_LEVEL")
	viper.MustBindEnv("log-format", "LOG_FORMAT")
	viper.MustBindEnv("log-source", "LOG_SOURCE")
}
