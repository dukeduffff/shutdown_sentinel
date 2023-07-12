package main

import (
	"context"
	"fmt"
	"github.com/shutdown_sentinel/check"
	"github.com/shutdown_sentinel/command"
	"github.com/shutdown_sentinel/config"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

func LoopDo(ctx context.Context, config *config.Config) {
	for {
		select {
		case <-ctx.Done():
			log.Info("ping功能退出")
		default:
			for i := 0; i < config.FailRetry; i++ {
				ping, cost := check.Ping(config)
				log.WithFields(
					log.Fields{
						"是否可达": ping,
						"耗时":   fmt.Sprintf("%v\tms", math.Round(cost*1000)/1000),
						"ip":   config.SentinelIp,
					}).Info("当前ping信息")
				if ping {
					break
				}
				// 一直失败则执行命令
				if i == config.FailRetry-1 {
					command.ExecuteCommand(config)
				}
			}
			time.Sleep(time.Second * config.Interval)
		}

	}
}
