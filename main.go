package main

import (
	"github.com/pascallimeux/ocms/api"
	"github.com/pascallimeux/ocms/common"
	"github.com/pascallimeux/ocms/hyperledger"
	"github.com/pascallimeux/ocms/hyperledger/content"
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
	var configuration common.Configuration
	err := utils.Read_Conf(config_file, &configuration)
	if err != nil {
		panic(err.Error())
	}

	// Init logger
	f := log.Init_log(configuration.LogFileName, configuration.Logger)
	defer f.Close()

	// Get local IP address if possible
	ipAddress, err := utils.GetOutboundIP()
	if err != nil {
		log.Error(log.Here(), " Impossible to retrieve the IP address of this machine")
	} else {
		configuration.HttpHostUrl = ipAddress + ":8030"
	}

	// Write configuration in log
	log.Info(log.Here(), utils.Get_fields(configuration))

	// Init Hyperledger helpers
	HP_helper := hyperledger.HP_Helper{HttpHyperledger: configuration.HttpHyperledger}
	Consent_Helper := consent.Consent_Helper{HP_helper: HP_helper, ChainCodePath: configuration.ChainCodePath, ChainCodeName: configuration.ChainCodeName, EnrollID: configuration.EnrollID, EnrollSecret: configuration.EnrollSecret}

	// Init application context
	appContext := api.AppContext{Consent_helper: Consent_Helper, Configuration: configuration}

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
