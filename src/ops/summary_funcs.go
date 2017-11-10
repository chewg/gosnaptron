package ops


type stat_func func(f *[]Frame) *[]Frame


// Only guarantee with junction_id is that it is first that occurs
func Sum_Count_By_Sample_ID(frames *[]Frame) *[]Frame {
	m := map[int]Frame{}

	for _, frame := range *frames {
		if sample_frame, exist := m[frame.sample_id]; exist {
			sample_frame.count += frame.count
			m[frame.sample_id] = sample_frame
		} else {
			// sample id not exist
			m[frame.sample_id] = frame
		}
	}

	return convert_map_to_slice(m)
}


// todo adhere to stat_func, or generalize stat_func to func(f ...*[]Frame)
func Junction_Inclusion_Ratio(f1, f2 *[]Frame) *[]Frame {
	f2_map := map[int]int{}

	// assume no duplicate sample ids. Else need to do something akin to Sum_Count_By_Sample_ID
	// todo extrapolate similarities between Sum_Count_By_Sample_ID and Junction_Inclusion_Ratio
	// create map for f2
	for _, frame := range *f2 {
		f2_map[frame.sample_id] = frame.count
	}

	// is intersection op
	i := 0
	for _, frame := range *f1 {
		if f2_count, exist := f2_map[frame.sample_id]; exist {
			f1_count := frame.count
			frame.stat = jir_formula(f1_count, f2_count)

			(*f1)[i] = frame
			i++
		}
	}

	jir_frames := (*f1)[:i]

	return &jir_frames
}


func jir_formula(c1, c2 int) float32 {
	c1_float := float32(c1)
	c2_float := float32(c2)

	return ((c1_float - c2_float) / (c1_float + c2_float + 1))
}