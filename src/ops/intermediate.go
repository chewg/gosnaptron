package ops


func Arrange(frames *[]Frame, order ...order_func) *[]Frame {
	order_by(order...).Sort(*frames)
	return frames
}


func order_by(order ...order_func) *sorter {
	return &sorter{
		order: order,
	}
}


func Filter(frames *[]Frame, filter_tuples ...Filter_tuple) *[]Frame {
	i := 0

	for _, frame := range *frames {
		keep := true

		for _, tuple := range filter_tuples {
			if !tuple.filter(&frame, tuple.value) {
				keep = false
				break
			}
		}

		if keep {
			(*frames)[i] = frame
			i++
		}
	}

	filtered_frames := (*frames)[:i]

	return &filtered_frames
}


func Group(frames_group ...*[]Frame) *[]Frame {
	var group []Frame

	// create a group by appending multiples frames together
	for _, frames := range frames_group {
		group = append(group, *frames...)
	}

	return &group
}


func Summarize(frames *[]Frame, stats ...stat_func) *[]Frame {
	for _, stat := range stats {
		frames = stat(frames)
	}

	return frames
}
