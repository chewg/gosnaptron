package ops


type less_func func(f1, f2 *Frame) bool


// returns sorter that sorts using less functions, done in param order listed
func Order_By(less ...less_func) *sorter {
	return &sorter{
		less: less,
	}
}


func Incr_Junction_ID(f1, f2 *Frame) bool {
	return f1.junction_id < f2.junction_id
}


func Decr_Junction_ID(f1, f2 *Frame) bool {
	return f1.junction_id > f2.junction_id
}


func Incr_Sample_ID(f1, f2 *Frame) bool {
	return f1.sample_id < f2.sample_id
}


func Decr_Sample_ID(f1, f2 *Frame) bool {
	return f1.sample_id > f2.sample_id
}


func Incr_Count(f1, f2 *Frame) bool {
	return f1.count < f2.count
}


func Decr_Count(f1, f2 *Frame) bool {
	return f1.count > f2.count
}


func Incr_Data(f1, f2 *Frame) bool {
	return f1.data < f2.data
}


func Decr_Data(f1, f2 *Frame) bool {
	return f1.data > f2.data
}
