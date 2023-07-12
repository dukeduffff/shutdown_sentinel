package config

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Config struct {
	Interval    time.Duration `json:"interval"`
	FailRetry   int           `json:"fail_retry"`
	SentinelIp  string        `json:"sentinel_ip"`
	TodoCommand string        `json:"todo_command"`
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
	return &config, nil
}

func checkCommand(command string) error {
	if len(command) == 0 {
		return errors.New("config.TodoCommand can not be empty")
	}
	return nil
}
