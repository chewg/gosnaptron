package ops


/*****
aggreg_func

Provides a finer resolution of operation on a Frame field (slice of ints). Methods that use aggreg_func are
found in frame.go

Parameters: 2 ints
Output: 1 int which is the value from the formula
*****/
type aggreg_func func(v1, v2 int) int


func aggreg_diff(v1, v2 int) int {
	return v1 - v2
}


func aggreg_sum(v1, v2 int) int {
	return v1 + v2
}
