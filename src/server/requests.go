package server

import (
	"net/http"
	"io/ioutil"
	"time"
	"fmt"
)

var timeout_wait = time.Second * 10
// TODO implement retry (and backoff) in Get


/*****
Get

Generic function that does a HTTP Get

Parameters: string that contains url address
Output: string that returned from HTTP Get
*****/
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


/*****
Get_mRNA_From_Server

Get RNA data from snaptron server.

Parameters: parameter string of url, datasource (srav1, gtex, etc.)
Output: string that is returned from snaptron server
*****/
func Get_mRNA_From_Server(params, datasource string) string {
	var server_address = fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/snaptron?%v", datasource, params)
	return Get(server_address)
}


/*****
Get_Metadata_From_Server

Get metadata from snaptron server.

Parameters: datasource (srav1, gtex, etc.)
Output: string that is returned from snaptron server
*****/
func Get_Metadata_From_Server(datasource string) string {
	var server_address = fmt.Sprintf("http://snaptron.cs.jhu.edu/%v/samples?all=1", datasource)
	return Get(server_address)
}
