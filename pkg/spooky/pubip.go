package spooky

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func GetPubIP(c *cli.Context) error {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	httpClient := http.Client{
		Transport: netTransport,
		Timeout:   time.Second * 10,
	}

	response, err := httpClient.Get("https://api.ipify.org/?format=text")
	if err != nil {
		return err
	}

	defer func() error {
		if err := response.Body.Close(); err != nil {
			return err
		}
		return nil
	}()

	if response.StatusCode != 200 {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(string(bodyBytes))

	return nil
}
