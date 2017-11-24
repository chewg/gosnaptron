package ops


type Filter_tuple struct {
	filter 		filter_func
	value 		int
}


type filter_func func(f *Frame, value int) bool


func Sample_Count_Gt(f *Frame, value int) bool {
	return f.First_Count() > value
}


func Sample_Count_Geq(f *Frame, value int) bool {
	return f.First_Count() >= value
}
