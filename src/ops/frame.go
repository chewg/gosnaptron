package ops

import (
	"sort"
)


type Frame struct {
	junction_id int
	sample_id int
	count int
	stat float32
	data string
}


type sorter struct {
	frames 	[]Frame
	less    []less_func
}


func (s *sorter) Sort(frames []Frame) {
	s.frames = frames
	sort.Sort(s)
}

// part of sort.Interface.
func (s *sorter) Swap(i, j int) {
	s.frames[i], s.frames[j] = s.frames[j], s.frames[i]
}

// part of sort.Interface
func (s *sorter) Less(i, j int) bool {
	frame_i, frame_j := &s.frames[i], &s.frames[j]

	// All but last less func comparison
	var l int
	for l = 0; l < len(s.less) - 1; l++ {
		less := s.less[l]

		switch {
		case less(frame_i, frame_j):	// frame_i < frame_j
			return true
		case less(frame_j, frame_i):	// frame_i > frame_j
			return false
		}
		// frame_i == frame_j; go to next less func comparison
	}

	// All prev less func comparisons equal, so return last less func
	return s.less[l](frame_i, frame_j)
}

// part of sort.Interface
func (s *sorter) Len() int {
	return len(s.frames)
}
