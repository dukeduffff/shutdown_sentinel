package check

import (
	"fmt"
	"github.com/shutdown_sentinel/config"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

const (
	timeout = time.Second * 1
	port    = 8000
)

func Ping(config *config.Config) (bool, float64) {
	t1 := time.Now().UnixMicro()
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.SentinelIp, port))
	checkError(err)
	tcpConn, err := net.DialTimeout("tcp", tcpAddr.String(), timeout)
	defer func(tcpConn net.Conn) {
		if tcpConn == nil {
			return
		}
		err := tcpConn.Close()
		checkError(err)
	}(tcpConn)
	if err.(*net.OpError).Timeout() {
		return false, 0
	}
	return true, float64(time.Now().UnixMicro()-t1) / 1000
}

func checkError(err error) {
	if err != nil {
		log.WithField("error", err).Error("ping error")
	}
}
