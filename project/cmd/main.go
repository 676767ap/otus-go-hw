package main

import (
	"github.com/676767ap/project/internal/app"
	"github.com/676767ap/project/internal/config"
	"github.com/676767ap/project/util/log"
)

// @title        banner-rotator
// @version      0.1
// @description  Проект "Ротация баннеров"
func main() {
	cfg := config.LoadConfig()

	log.Init(cfg)

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
