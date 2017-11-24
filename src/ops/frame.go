package ops

import (
	"sort"
)


type Frame struct {
	junction_id 	[]int
	sample_id 		[]int
	count 			[]int
	stat 			[]float32
	metadata 		string
}


func New_Frame() *Frame {
	f := Frame{}
	f.junction_id = make([]int, 0)
	f.sample_id = make([]int, 0)
	f.count = make([]int, 0)
	f.stat = make([]float32, 0)
	return &f
}


func (f *Frame) First_Count() int {
	return f.count[0]
}


func (f *Frame) First_Junction_ID() int {
	return f.junction_id[0]
}


func (f *Frame) First_Sample_ID() int {
	return f.sample_id[0]
}


func (f *Frame) First_Stat() float32 {
	return f.stat[0]
}


func (f *Frame) Add_Count(new_count int) *Frame {
	f.count = append(f.count, new_count)
	sort.Ints(f.count)
	return f
}


func (f *Frame) Add_Junction_ID(new_id int) *Frame {
	f.junction_id = append(f.junction_id, new_id)
	sort.Ints(f.junction_id)
	return f
}


func (f *Frame) Add_Sample_ID(new_id int) *Frame {
	f.sample_id = append(f.sample_id, new_id)
	sort.Ints(f.sample_id)
	return f
}


func (f *Frame) Add_Stat(new_stat float32) *Frame {
	f.stat = append(f.stat, new_stat)
	return f
}


func (f *Frame) Set_Metadata(s string) *Frame {
	f.metadata = s
	return f
}


func (f *Frame) Aggregate_Count(fn aggreg_func) *Frame {
	new_count := 0

	for _, value := range f.count {
		new_count = fn(new_count, value)
	}

	f.count = []int{new_count}

	return f
}


// then have a function that takes in these function and performs them
func update_count(f_keep, f_additional *Frame) *Frame {
	f_keep.count = append_ints(f_keep.count, f_additional.count)
	sort.Ints(f_keep.count)
	return f_keep
}


func update_junction_id(f_keep, f_additional *Frame) *Frame {
	f_keep.junction_id = append_ints(f_keep.junction_id, f_additional.junction_id)
	sort.Ints(f_keep.junction_id)
	return f_keep
}


func update_sample_id(f_keep, f_additional *Frame) *Frame {
	f_keep.sample_id = append_ints(f_keep.sample_id, f_additional.sample_id)
	sort.Ints(f_keep.sample_id)
	return f_keep
}


func update_by_junction_id(f_keep, f_additional Frame) Frame {
	f_keep = *update_count(&f_keep, &f_additional)
	f_keep = *update_sample_id(&f_keep, &f_additional)
	return f_keep
}


func update_by_sample_id(f_keep, f_additional Frame) Frame {
	f_keep = *update_count(&f_keep, &f_additional)
	f_keep = *update_junction_id(&f_keep, &f_additional)
	return f_keep
}


type By_Junction_ID struct{}
type By_Sample_ID struct {}


type sorter struct {
	frames 		[]Frame
	order    	[]order_func
}


func (s *sorter) Sort(frames []Frame) {
	s.frames = frames
	sort.Sort(s)
}

// part of sort.Interface
func (s *sorter) Len() int {
	return len(s.frames)
}

// part of sort.Interface
func (s *sorter) Less(i, j int) bool {
	frame_i, frame_j := &s.frames[i], &s.frames[j]

	// All but last less func comparison
	var l int
	for l = 0; l < len(s.order) - 1; l++ {
		less := s.order[l]

		switch {
		case less(frame_i, frame_j):	// frame_i < frame_j
			return true
		case less(frame_j, frame_i):	// frame_i > frame_j
			return false
		}
		// frame_i == frame_j; go to next less func comparison
	}

	// All prev less func comparisons equal, so return last less func
	return s.order[l](frame_i, frame_j)
}

// part of sort.Interface.
func (s *sorter) Swap(i, j int) {
	s.frames[i], s.frames[j] = s.frames[j], s.frames[i]
}
