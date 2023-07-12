package command

import (
	"github.com/shutdown_sentinel/config"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func ExecuteCommand(config *config.Config) {
	split := strings.Split(config.TodoCommand, " ")
	name := split[0]
	arg := split[1:]
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		log.WithField("err", err).Info("execute command failed")
	}
	log.WithField("output", string(output)).Info("command execute success")
}
