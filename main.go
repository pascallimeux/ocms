package main

import (
	"github.com/pascallimeux/ocms/api"
	"github.com/pascallimeux/ocms/utils"
	"github.com/pascallimeux/ocms/utils/log"
	"net/http"
	"os"
	"time"
)

func main() {

	// get arguments
	config_file := "config.json"
	args := os.Args[1:]
	if len(args) == 1 {
		config_file = args[0]
	}

	// Init configuration
	configuration, err := utils.Readconf(config_file)
	if err != nil {
		panic(err.Error())
	}

	// Init logger
	f := log.Init_log(configuration.LogFileName, configuration.Logger)
	defer f.Close()
	log.Info(log.Here(), configuration.To_string())

	// Init application context
	appContext := api.AppContext{}

	// Start http server
	router := appContext.CreateRoutes()
	log.Info(log.Here(), "Listening on: ", configuration.HttpHostUrl)
	s := &http.Server{
		Addr:         configuration.HttpHostUrl,
		Handler:      router,
		ReadTimeout:  configuration.ReadTimeout * time.Second,
		WriteTimeout: configuration.WriteTimeout * time.Second,
	}
	log.Fatal(log.Here(), s.ListenAndServe().Error())
}
