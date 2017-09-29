package query



type junction_ids struct {
	id []int
}

func Junction_IDs() junction_ids {
	var j junction_ids
	j.id = make([]int, 0)
	return j
}

func (j *junction_ids) Add(id int) *junction_ids {
	j.id = append(j.id, id)
	return j
}



type sample_ids struct {
	id []int
}

func Sample_IDs() sample_ids {
	var s sample_ids
	s.id = make([]int, 0)
	return s
}

func (s *sample_ids) Add(id int) *sample_ids {
	s.id = append(s.id, id)
	return s
}



