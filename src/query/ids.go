package query

import (
	"fmt"
	"errors"
	"strings"
)


type junction_ids struct {
	ids []int
}

func Junction_IDs() junction_ids {
	var j junction_ids
	j.ids = make([]int, 0)
	return j
}

func (j *junction_ids) Add(ids ...int) *junction_ids {
	for id := range ids {
		j.ids = append(j.ids, id)
	}

	return j
}

func (j *junction_ids) Initialized() bool {
	if len(j.ids) > 0 {
		return true
	}

	return false
}

func (j *junction_ids) Export() (string, error) {
	if j.Initialized() {
		ids_string := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(j.ids)), ","), "[]")
		return fmt.Sprintf("sids=%s", ids_string), nil
	}

	return "", errors.New("No valid junction ids")
}



type sample_ids struct {
	ids []int
}

func Sample_IDs() sample_ids {
	var s sample_ids
	s.ids = make([]int, 0)
	return s
}

func (s *sample_ids) Add(ids ...int) *sample_ids {
	for id := range ids {
		s.ids = append(s.ids, id)
	}
	return s
}

func (s *sample_ids) Initialized() bool {
	if len(s.ids) > 0 {
		return true
	}

	return false
}

func (s *sample_ids) Export() (string, error) {
	if s.Initialized() {
		ids_string := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s.ids)), ","), "[]")
		return fmt.Sprintf("ids=%s", ids_string), nil
	}

	return "", errors.New("No valid sample ids")
}


