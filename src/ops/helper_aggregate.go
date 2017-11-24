package ops


type aggreg_func func(v1, v2 int) int


func aggreg_diff(v1, v2 int) int {
	return v1 - v2
}


func aggreg_sum(v1, v2 int) int {
	return v1 + v2
}
