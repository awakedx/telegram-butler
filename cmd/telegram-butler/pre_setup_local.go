//go:build local

package main

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/GolangUA/telegram-butler/internal/module/logger"
)

func preSetup(ctx context.Context, log logger.Logger) error {
	viper.SetConfigFile(".env.local")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	fw, err := ngrok.ListenAndForward(
		ctx,
		&url.URL{
			Scheme: "http",
			Host:   ":" + viper.GetString("port"),
		},
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return fmt.Errorf("start ngrok tunnel: %w", err)
	}
	log.Infof("Ngrok tunnel: %s", fw.URL())

	viper.Set("webhook-url", fw.URL()+"/webhook")

	return nil
}