package main

import (
	"net/http"

	"github.com/Gaurav-malwe/login-service/config"
	"github.com/Gaurav-malwe/login-service/internal/controller"
	"github.com/Gaurav-malwe/login-service/internal/service"
	"github.com/Gaurav-malwe/login-service/mongodb"
	mconfig "github.com/Gaurav-malwe/login-service/mongodb/config"
	"github.com/Gaurav-malwe/login-service/router"
	log "github.com/Gaurav-malwe/login-service/utils/logging"
)

func main() {
	cfg := config.GetConfig()

	// Intialize logging

	// Intialize Root App

	// Mongo db intialization
	mongodbConf := mconfig.NewConfigFromEnv()
	mongoProvider := mongodb.New(mongodbConf)
	err := mongoProvider.Init()
	if err != nil {
		log.DebugWithFields("MongoDB initialization failed", log.Fields{"error": err.Error()})
	}
	log.Info("%s initialized", mongodbConf.AppName)

	// Router, ctl, repo, svc and Api server intialization
	loginServiceAPIServer := router.NewLoginServiceAPIServer(mongoProvider)
	repo := loginServiceAPIServer.Repository
	svc := service.New(repo, cfg)
	ctl := controller.New(svc, cfg)

	// Register handlers
	loginServiceAPIServer.RegisterHandlers(ctl)

	r := loginServiceAPIServer.Router

	// Listen and serve
	if err := http.ListenAndServe(":"+cfg.GetString("server_port"), r); err != nil {
		log.Fatal("http.ListenAndServe Failed")
	}
}
