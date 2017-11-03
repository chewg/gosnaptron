package ops


func Arrange(frames *[]Frame, less ...less_func) *[]Frame {
	Order_By(less...).Sort(*frames)
	return frames
}


// Create a group by appending multiples frames together
func Group(frames_group ...*[]Frame) *[]Frame {
	var group []Frame

	for _, frames := range frames_group {
		group = append(group, *frames...)
	}

	return &group
}


//func Select(frames []Frame, less ...less_func) []Frame {
//
//
//}

func Summarize(frames *[]Frame, stats ...stat_func) *[]Frame {
	for _, stat := range stats {
		frames = stat(frames)
	}

	return frames
}



// todo
func Filter(frames []Frame) []Frame {
	return frames
}