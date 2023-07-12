package config

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type Config struct {
	Interval    time.Duration `json:"interval"`
	FailRetry   int           `json:"fail_retry"`
	SentinelIp  string        `json:"sentinel_ip"`
	TodoCommand string        `json:"todo_command"`
	LogLevel    string        `json:"log_level"`
}

var levelMap = map[string]log.Level{
	"trace": log.TraceLevel,
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
	"fatal": log.FatalLevel,
	"panic": log.PanicLevel,
}

func LoadConfig(configDir string) (*Config, error) {
	content, err := os.ReadFile(configDir)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err := json.Unmarshal(content, &config); err != nil {
		return nil, err
	}
	if err := checkCommand(config.TodoCommand); err != nil {
		return nil, err
	}
	log.SetLevel(log.InfoLevel)
	if level, ok := levelMap[strings.ToLower(config.LogLevel)]; ok {
		log.SetLevel(level)
	}
	return &config, nil
}

func checkCommand(command string) error {
	if len(command) == 0 {
		return errors.New("config.TodoCommand can not be empty")
	}
	return nil
}
