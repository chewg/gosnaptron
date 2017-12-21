package ops


/*****
stat_func

Takes as parameters 1 pointer to a slice of frames. The function body contains the summary statistic operation that is
performed on each frame's field.

Parameters: 1 pointer to a slice of frames
Output: 1 address to a slice of frames, with the summary statistic been performed
*****/
type stat_func func(f *[]Frame) *[]Frame


func Sum_Count(frames *[]Frame) *[]Frame {
	saved := map[int]Frame{}

	for _, frame := range *frames {
		frame_id := get_frame_id(&frame)

		if saved_frame, exist := saved[frame_id]; exist {
			switch group_frames_by.(type) {
			case By_Sample_ID:
				saved[frame_id] = sum_count_by_sample_id(saved_frame, frame)
			case By_Junction_ID:
				saved[frame_id] = sum_count_by_junction_id(saved_frame, frame)
			}
		} else {
			// no frame saved for id
			frame.Aggregate_Count(aggreg_sum)
			saved[frame_id] = frame
		}
	}

	return Convert_Map_To_Slice(&saved)
}


func sum_count_by_junction_id(f_keep, f_additional Frame) Frame {
	f_keep = update_by_junction_id(f_keep, f_additional)
	f_keep.Aggregate_Count(aggreg_sum)
	return f_keep
}


func sum_count_by_sample_id(f_keep, f_additional Frame) Frame {
	f_keep = update_by_sample_id(f_keep, f_additional)
	f_keep.Aggregate_Count(aggreg_sum)
	return f_keep
}
