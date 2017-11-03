package ops

import "snaptron_api/src/data"


func Import_Dataframe(data_frame data.Dataframe) *[]Frame {
	var frames []Frame

	for _, data := range data_frame.Frames() {
		var frame Frame
		frame.junction_id = data.Junction_id
		for sample_id, count := range data.Samples {
			frame.sample_id = sample_id
			frame.count = count
			frames = append(frames, frame)
		}
	}

	return &frames
}


func convert_map_to_slice(m map[int]Frame) *[]Frame {
	var frames []Frame

	for _, frame := range m {
		frames = append(frames, frame)
	}

	return &frames
}


/* Part of intermediate_old.go */

type Pair struct {
	Key int
	Value int
}

type PairSlice []Pair

func (p PairSlice) Len() int {
	return len(p)
}

func (p PairSlice) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairSlice) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}


func ConvertMapToSlice(old map[interface{}][]interface{}) PairSlice {
	pl := make(PairSlice, len(old))
	i := 0

	for k, v := range old {
		// Only takes the first element in []interface{} v
		pl[i] = Pair{k.(int), v[0].(int)}
		i++
	}

	// Descending order, by value
	// sort.Sort(sort.Reverse(pl))
	return pl
}
