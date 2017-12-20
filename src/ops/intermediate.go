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

	m := Convert_Slice_To_Map(frames)
	s := Convert_Map_To_Slice(m)

	return s
}


func Intersect(frames_1 *[]Frame, more_frames ...*[]Frame) *[]Frame {
	m := *Convert_Slice_To_Map(frames_1)

	// Intersection procedure
	for _, frames := range more_frames {
		intersected_m := map[int]Frame{}

		for _, frame := range *frames {
			frame_id := get_frame_id(&frame)

			// Check if have intersection
			if stored_frame, intersection_exists := m[frame_id]; intersection_exists {
				if intersected_frame, stored_in_intersected_m := intersected_m[frame_id]; stored_in_intersected_m {
					// In intersected_m, so update intersected_frame
					switch group_frames_by.(type) {
					case By_Sample_ID:
						intersected_m[frame_id] = update_by_sample_id(intersected_frame, frame)
					case By_Junction_ID:
						intersected_m[frame_id] = update_by_junction_id(intersected_frame, frame)
					}
				} else {
					// Not in intersected_m, so add stored_frame to intersected_m
					switch group_frames_by.(type) {
					case By_Sample_ID:
						intersected_m[frame_id] = update_by_sample_id(frame, stored_frame)
					case By_Junction_ID:
						intersected_m[frame_id] = update_by_junction_id(frame, stored_frame)
					}
				}
			}
		}

		m = intersected_m
	}

	return Convert_Map_To_Slice(&m)
}


func Order(frames *[]Frame, order ...order_func) *[]Frame {
	order_by(order...).Sort(*frames)
	return frames
}


/*****
Summarize

Takes a slice of frames and performs summary statistic on each of the frames in the slice.
For example, one can pass in ops.Sum_Count as a stat_func which does a summation of each frame's count

Parameters: pointer to 1 slice of frames, 1 or more stat functions found in stat.go
Output: address of 1 slice of frames, which has had stat_func performed on it
*****/
func Summarize(frames *[]Frame, stats ...stat_func) *[]Frame {
	for _, stat := range stats {
		frames = stat(frames)
	}

	return frames
}


/*****
Union

Takes slices of frames, and based on the frame id the frames are grouped by and is globally set,
the different slices are unioned together.

Parameters: pointer to 1 slice of frames, pointer(s) to 1 or more slice(s) of frames
Output: address of 1 slice of frames, which is the union of all frames grouped by the frame id
*****/
func Union(frames_1 *[]Frame, more_frames ...*[]Frame) *[]Frame {
	m := *Convert_Slice_To_Map(frames_1)

	// iterate through more_frames and union each one with frames_1
	for _, frames := range more_frames {
		for _, frame := range *frames {
			frame_id := get_frame_id(&frame)

			// union by frame_id
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

	return Convert_Map_To_Slice(&m)
}
