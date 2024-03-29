package main

import (
	"github.com/676767ap/otus-go-hw/project/internal/app"
	"github.com/676767ap/otus-go-hw/project/internal/config"
	"github.com/676767ap/otus-go-hw/project/util/log"
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
