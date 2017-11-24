package web

import (
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"strconv"
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


func Import_Metadata(url string) *map[int][]string {
	metadata_string := Get(url)
	metadata_slice := strings.Split(metadata_string, "\n")

	var row_slice [][]string

	for _, metadata := range metadata_slice {
		fields := strings.Split(metadata, "\t")
		row_slice = append(row_slice, fields)
	}

	row_slice = row_slice[1:]

	metadata_map := map[int][]string{}

	for _, row := range row_slice {
		i32, _ := strconv.ParseInt(row[0], 10, 32)
		sample_id := int(i32)

		metadata_map[sample_id] = row
	}

	return &metadata_map
}