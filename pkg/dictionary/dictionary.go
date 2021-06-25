package dictionary

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

func wordMeaning(word string) {
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

	response, err := httpClient.Get("https://api.dictionaryapi.dev/api/v2/entries/en_US/" + word)
	if err != nil {
		log.Error(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(response.Body)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var responseModel []ResponseModel

	if response.Body != nil && response.StatusCode == 200 {
		err := json.Unmarshal(bodyBytes, &responseModel)
		if err != nil {
			return
		}

		mut.Lock()
		defer mut.Unlock()
		WordMeaningDecorator(responseModel)
	}
}

func Dictionary(c *cli.Context) error {
	for _, word := range c.Args().Slice() {
		go wordMeaning(word)
		wg.Add(1)
	}
	wg.Wait()
	return nil
}
