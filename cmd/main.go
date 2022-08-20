package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
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
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	ra := redAlert.NewRedAlertClient(m)
	go ra.Run()

	addr := fmt.Sprintf(":%s", conf.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Could not start http server: %s", err)
	}

}
