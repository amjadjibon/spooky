package ipapi

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func requestIPAPI(ip string) {
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

	response, err := httpClient.Get("https://ipapi.co/" + ip + "/json/")
	if err != nil {
		log.Error(err)
		return
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Error(err)
			return
		}
	}()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(string(responseBytes))
}

func GetIPInformation(c *cli.Context) error {
	ip := c.Args().First()

	restyRequest(ip)

	return nil
}

func restyRequest(ip string)  {
	client := resty.New()
	response, err := client.R().Get("https://ipapi.co/" + ip + "/json/")
	if err != nil {
		return
	}

	fmt.Println(response.ReceivedAt().Sub(response.Request.Time).String())
	fmt.Println(string(response.Body()))
}
