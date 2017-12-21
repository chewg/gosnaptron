package ops


type Filter_tuple struct {
	filter 		filter_func
	value 		int
}


/*****
filter_func

Takes as parameters a frame and a value, then does a comparison between a Frame field and the value

Parameters: 1 pointer to Frame, a int value
Output: bool of whether frame field comparison was true or false
*****/
type filter_func func(f *Frame, value int) bool


func Sample_Count_Gt(f *Frame, value int) bool {
	return f.First_Count() > value
}


func Sample_Count_Geq(f *Frame, value int) bool {
	return f.First_Count() >= value
}
