package main

import (
	"contest/config"
	_ "contest/docs"
	"contest/internal/app"
	"flag"
)

// @title Contest Service API
// @version 1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "path", "", "path to config file")
	flag.Parse()
	if cfgPath == "" {
		panic("config file path is required")
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		panic("read config error: " + err.Error())
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err.Error())
	}
}
