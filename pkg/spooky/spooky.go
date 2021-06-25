package spooky

import (
	"fmt"
	"github.com/amjadjibon/spooky/pkg/constant"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"net"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

func sendRequest(website string) {
	defer wg.Done()

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

	response, err := httpClient.Get("https://" + website)
	if err != nil {
		log.Errorf("%s: %s", constant.ApplicationName, err)
		return
	}
	if err = response.Body.Close(); err != nil {
		log.Errorf("%s: %s", constant.ApplicationName, err)
		return
	}

	mut.Lock()
	defer mut.Unlock()

	fmt.Printf("[%d] %s\n", response.StatusCode, website)
}

func GetStatusCode(c *cli.Context) error {
	for _, website := range c.Args().Slice() {
		go sendRequest(website)
		wg.Add(1)
	}
	wg.Wait()
	return nil
}
