package ops

type ratio_func func(...int) float32


func JIR_Ratio(cs ...int) float32 {
	if len(cs) != 2 {
		panic("jir_ratio needs exactly 2 ints passed in.")
	}

	c1_float := float32(cs[0])
	c2_float := float32(cs[1])

	return ((c1_float - c2_float) / (c1_float + c2_float + 1))
}


// todo fully implement Percent Spliced In
func psi_ratio(cs ...int) float32 {
	if len(cs) != 3 {
		panic("psi_ratio needs exactly 3 ints passed in.")
	}

	c1_float := float32(cs[0])
	c2_float := float32(cs[1])
	c3_float := float32(cs[2])

	mean := (c1_float + c2_float) / 2.0

	return (mean / (mean + c3_float))
}



func Calculate_Ratio(ratio ratio_func, f1, f2 *[]Frame) *[]Frame {
	f2_map := *convert_slice_to_map(f2)

	// is intersection op
	i := 0
	for _, f1_frame := range *f1 {
		if f2_frame, exist := f2_map[f1_frame.First_Sample_ID()]; exist {
			f1_count := f1_frame.Aggregate_Count(aggreg_sum).First_Count()
			f2_count := f2_frame.Aggregate_Count(aggreg_sum).First_Count()

			f1_frame.stat = []float32{ratio(f1_count, f2_count), float32(f1_count), float32(f2_count)}

			(*f1)[i] = f1_frame
			i++
		}
	}

	jir_frames := (*f1)[:i]

	return &jir_frames
}