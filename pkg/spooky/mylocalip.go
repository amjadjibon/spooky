package spooky

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

// GetOutboundIP Get preferred outbound ip of this machine
func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer func() error {
		if err := conn.Close(); err != nil {
			return err
		}
		return nil
	}()

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
