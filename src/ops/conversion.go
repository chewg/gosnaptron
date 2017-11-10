package ops

import "snaptron_api/src/data"


func Import_Dataframe(data_frame *data.Dataframe) *[]Frame {
	var frames []Frame

	for _, data := range data_frame.Frames() {
		var frame Frame
		frame.junction_id = data.Junction_id
		for sample_id, count := range data.Samples {
			frame.sample_id = sample_id
			frame.count = count
			frames = append(frames, frame)
		}
	}

	return &frames
}


func convert_map_to_slice(m map[int]Frame) *[]Frame {
	var frames []Frame

	// no guarantee with map's order
	for _, frame := range m {
		frames = append(frames, frame)
	}

	return &frames
}

