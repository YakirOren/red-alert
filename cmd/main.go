package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"polling-to-ws/broadcast"
	"polling-to-ws/config"
	"polling-to-ws/redAlert"
)

func main() {
	configProvider := config.EnvConfigProvider{}
	conf, err := configProvider.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(conf.LogLevel)

	r := gin.Default()
	hub := broadcast.NewHub()
	go hub.Run()

	r.GET("/ws", func(ctx *gin.Context) {
		broadcast.ServeWs(hub, ctx.Writer, ctx.Request)

	})

	ra := redAlert.NewRedAlertClient(hub)
	go ra.Run()

	addr := fmt.Sprintf(":%s", conf.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Could not start http server: %s", err)
	}

}
