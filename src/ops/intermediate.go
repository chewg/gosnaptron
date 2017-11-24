package ops


var group_frames_by interface{}

func init() {
	/* Default Group By Sample ID */
	group_frames_by = By_Sample_ID{}
}


func Bind(frames_group ...*[]Frame) *[]Frame {
	var group []Frame

	// create a group by appending multiples frames together
	for _, frames := range frames_group {
		group = append(group, *frames...)
	}

	return &group
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


// Ensures distinctness
func Group_By(group_by_func func(), frames *[]Frame) *[]Frame {
	group_by_func()

	m := convert_slice_to_map(frames)
	s := convert_map_to_slice(m)

	return s
}


func Intersect(frames_1 *[]Frame, more_frames ...*[]Frame) *[]Frame {
	m := *convert_slice_to_map(frames_1)

	// Intersection procedure
	for _, frames := range more_frames {
		common_m := map[int]Frame{}

		for _, frame := range *frames {
			frame_id := get_frame_id(&frame)

			// Check if have intersection
			if m_frame, in_m := m[frame_id]; in_m {
				if common_frame, in_common := common_m[frame_id]; in_common {
					switch group_frames_by.(type) {
					case By_Sample_ID:
						common_m[frame_id] = update_by_sample_id(common_frame, frame)
					case By_Junction_ID:
						common_m[frame_id] = update_by_junction_id(common_frame, frame)
					}
				} else {
					// Not in common_m, so update count and add to common_m
					switch group_frames_by.(type) {
					case By_Sample_ID:
						common_m[frame_id] = update_by_sample_id(frame, m_frame)
					case By_Junction_ID:
						common_m[frame_id] = update_by_junction_id(frame, m_frame)
					}
				}
			}
		}

		m = common_m
	}

	return convert_map_to_slice(&m)
}


func Order(frames *[]Frame, order ...order_func) *[]Frame {
	order_by(order...).Sort(*frames)
	return frames
}


func Summarize(frames *[]Frame, stats ...stat_func) *[]Frame {
	for _, stat := range stats {
		frames = stat(frames)
	}

	return frames
}


func Union(frames_1 *[]Frame, more_frames ...*[]Frame) *[]Frame {
	m := *convert_slice_to_map(frames_1)

	// Union procedure
	for _, frames := range more_frames {
		for _, frame := range *frames {
			frame_id := get_frame_id(&frame)

			if saved_frame, exist := m[frame_id]; exist {
				switch group_frames_by.(type) {
				case By_Sample_ID:
					m[frame_id] = update_by_sample_id(saved_frame, frame)
				case By_Junction_ID:
					m[frame_id] = update_by_junction_id(saved_frame, frame)
				}
			} else {
				m[frame_id] = frame
			}
		}
	}

	return convert_map_to_slice(&m)
}

