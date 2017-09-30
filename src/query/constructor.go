package query

import (
	"bytes"
	"snaptron_api/src/web"
)


var server_address string = "http://snaptron.cs.jhu.edu/srav2/snaptron?"

func Execute(params ...interface{}) string {
	query_string := build_query(params...)
	return web.Get(query_string)
}


func build_query(params ...interface{}) string {
	var b bytes.Buffer

	b.WriteString(server_address)

	appending := false
	region_exist := false

	for _, param := range params {

		switch param.(type) {
		case region:
			r := param.(region)
			if r.Initialized() {
				region_exist = true
				if appending {
					b.WriteString("&")
				} else {
					appending = true
				}

				params_str, _ := r.Export()
				b.WriteString(params_str)
			}
		case filter:
			f := param.(filter)
			if f.Initialized() {
				if appending {
					b.WriteString("&")
				} else {
					appending = true
				}

				params_str, _ := f.Export()
				b.WriteString(params_str)
			}
		case metadata:
			m := param.(metadata)
			if m.Initialized() {
				if appending {
					b.WriteString("&")
				} else {
					appending = true
				}

				params_str, _ := m.Export()
				b.WriteString(params_str)
			}
		case junction_ids:
			j := param.(junction_ids)
			if j.Initialized() {
				if appending {
					b.WriteString("&")
				} else {
					appending = true
				}

				params_str, _ := j.Export()
				b.WriteString(params_str)
			}
		case sample_ids:
			s := param.(sample_ids)
			if s.Initialized() {
				if appending {
					b.WriteString("&")
				} else {
					appending = true
				}

				params_str, _ := s.Export()
				b.WriteString(params_str)
			}
		}
	}

	if !region_exist {
		// TODO should do something with an error
	}

	return b.String()
}


