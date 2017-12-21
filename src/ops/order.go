package ops


/*****
order_func

Takes as parameters 2 pointers to Frames. Within the function body, a comparison is made between the respective field
of each Frame. A bool is returned of whether the comparison was true or false

Parameters: 2 pointers to slices of frames
Output: bool of whether frame field comparison was true or false
*****/
type order_func func(f1, f2 *Frame) bool


func Incr_Junction_ID(f1, f2 *Frame) bool {
	 return slice_less_than(f1.junction_id, f2.junction_id)
}


func Decr_Junction_ID(f1, f2 *Frame) bool {
	return slice_less_than(f2.junction_id, f1.junction_id)
}


func Incr_Sample_ID(f1, f2 *Frame) bool {
	return slice_less_than(f1.sample_id, f2.sample_id)
}


func Decr_Sample_ID(f1, f2 *Frame) bool {
	return slice_less_than(f2.sample_id, f1.sample_id)
}


func Incr_Count(f1, f2 *Frame) bool {
	return slice_less_than(f1.count, f2.count)
}


func Decr_Count(f1, f2 *Frame) bool {
	return slice_less_than(f2.count, f1.count)
}


func Incr_Stat(f1, f2 *Frame) bool {
	return f1.First_Stat() < f2.First_Stat()
}


func Decr_Stat(f1, f2 *Frame) bool {
	return f1.First_Stat() > f2.First_Stat()
}


func Incr_Metadata(f1, f2 *Frame) bool {
	return f1.metadata < f2.metadata
}


func Decr_Metadata(f1, f2 *Frame) bool {
	return f1.metadata > f2.metadata
}
