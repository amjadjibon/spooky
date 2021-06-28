package spooky

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net"
)

// getOutboundIP Get preferred outbound ip of this machine
func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func MyLocalIP(c *cli.Context) error {
	ip, err := getOutboundIP()
	if err != nil {
		return err
	}
	fmt.Println(ip.String())
	return nil
}
