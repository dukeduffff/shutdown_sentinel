package check

import (
	"fmt"
	"github.com/shutdown_sentinel/config"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"syscall"
	"time"
)

const (
	timeout = time.Second * 1
	port    = 8000
)

var errnoMap = map[int]interface{}{
	int(syscall.ETIME): nil,
	113:                nil, // go未定义的异常码
}

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

	if opError, ok := err.(*net.OpError); ok {
		if opError.Timeout() {
			return false, 0
		}
		if syscallError, ok := opError.Err.(*os.SyscallError); ok {
			if errno_ptr, ok := syscallError.Unwrap().(syscall.Errno); ok {
				log.WithFields(log.Fields{
					"errno": fmt.Sprintf("0x%x", int(errno_ptr)),
					"error": syscallError.Error(),
				}).Debug("ping error")
				if errno_ptr == syscall.EHOSTUNREACH {
					return false, 0
				}
			}
		}
	}

	return true, float64(time.Now().UnixMicro()-t1) / 1000
}

func checkError(err error) {
	if err != nil {
		log.WithField("error", err).Error("ping error")
	}
}
