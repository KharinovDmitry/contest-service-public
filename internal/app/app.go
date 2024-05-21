package app

import (
	"contest/cmd/migrator"
	"contest/config"
	_ "contest/docs"
	"contest/internal/server/router"
	service "contest/internal/service"
	"contest/internal/storage"
	"time"
)

func Run(cfg *config.Config) error {
	store := storage.NewStorage()
	if err := store.Init(cfg.СonnStr, 1*time.Minute); err != nil {
		return err
	}

	migrator.Run(cfg.СonnStr, cfg.MigrDir)

	services := service.NewServiceManager()
	if err := services.Init(store, cfg.Env); err != nil {
		return err
	}

	services.Logger.Info("app started")

	if err := router.Run(store, services, cfg.JwtSecret, cfg.Port); err != nil {
		return err
	}

	return nil
}
