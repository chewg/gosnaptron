package query

import (
	"fmt"
	"errors"
)

type metadata struct {
	key string
	op string
	value string
}

func Metadata() metadata {
	var m metadata
	return m
}

func (m *metadata) Key(key string) *metadata {
	m.key = key
	return m
}

func (m *metadata) Value(op string, value string) *metadata {
	m.op = op
	m.value = value
	return m
}

func (m *metadata) Initialized() bool {
	if m.key != "" && m.value != "" {
		return true
	}

	return false
}

func (m *metadata) Export() (string, error) {
	if m.Initialized() {
		return fmt.Sprintf("sfilter=%s%s:%s", m.key, m.op, m.value), nil
	}

	return "", errors.New("No valid metadata")
}
