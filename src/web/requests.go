package web

import (
	"net/http"
	"io/ioutil"
	"time"
)

var timeout_wait = time.Second * 10


// TODO implement retry (and backoff) in Get

func Get(url_addr string) string {
	var client = &http.Client{
		Timeout: timeout_wait,
	}

	resp, err := client.Get(url_addr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

