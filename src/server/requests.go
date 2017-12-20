package server

import (
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
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


func Get_mRNA_From_Server(params, datasource string) string {
	var server_address = fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/snaptron?%v", datasource, params)
	return Get(server_address)
}


func Get_Metadata_From_Server(datasource string) string {
	var server_address = fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/samples?all=1", datasource)
	return Get(server_address)
}
