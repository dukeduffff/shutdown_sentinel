package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shutdown_sentinel/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
		return
	}
	if len(cmd.configFile) == 0 {
		log.Error("please support config file, -c xxx.json")
		return
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableQuote:  true,
	})
	conf, err := config.LoadConfig(cmd.configFile)
	if err != nil {
		log.WithField("error", err).Error("load config error")
		return
	}
	jsonStr, err := json.Marshal(conf)
	log.WithField("config_json", string(jsonStr)).Info("load config success!")

	go LoopDo(ctx, conf)

	// 等待用户kill
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	cancelFunc()
}
