package ops

import (
	"snaptron_api/src/data"
	"math"
	"snaptron_api/src/web"
)


func order_by(order ...order_func) *sorter {
	return &sorter{
		order: order,
	}
}


func append_ints(s1 []int, s2 []int) []int {
	return append(s1, s2...)
}


func convert_map_to_slice(m *map[int]Frame) *[]Frame {
	var frames []Frame

	// no guarantee with map's order
	for _, frame := range *m {
		frames = append(frames, frame)
	}

	return &frames
}


func convert_slice_to_map(f *[]Frame) *map[int]Frame {
	m := map[int]Frame{}

	switch group_frames_by.(type) {
	case By_Sample_ID:
		for _, frame := range *f {
			for _, frame_id := range frame.sample_id {
				if saved_frame, exist := m[frame_id]; exist {
					m[frame_id] = update_by_sample_id(frame, saved_frame)
				} else {
					new_frame := frame
					new_frame.sample_id = []int{frame_id}
					m[frame_id] = new_frame
				}
			}
		}
	case By_Junction_ID:
		for _, frame := range *f {
			for _, frame_id := range frame.junction_id {
				if saved_frame, exist := m[frame_id]; exist {
					m[frame_id] = update_by_junction_id(frame, saved_frame)
				} else {
					new_frame := frame
					new_frame.junction_id = []int{frame_id}
					m[frame_id] = new_frame
				}
			}
		}
	}

	return &m
}


func get_frame_id(frame *Frame) int {
	var frame_id int

	switch group_frames_by.(type) {
	case By_Junction_ID:
		frame_id = frame.First_Junction_ID()
	case By_Sample_ID:
		frame_id = frame.First_Sample_ID()
	default:
		frame_id = frame.First_Sample_ID()
	}

	return frame_id
}


func slice_less_than(s1, s2 []int) bool {
	// single element comparison
	if len(s1) == 1 && len(s2) == 1 {
		return s1[0] < s2[0]
	}

	// multi-element comparison
	max_index := int(math.Min(float64(len(s1)), float64(len(s2))))

	for i := 0; i < max_index; i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}

	return true
}


func Dataframe_To_Frames(data_frame *data.Dataframe) *[]Frame {
	var frames []Frame

	for _, data := range data_frame.Frames() {
		frame := New_Frame()

		frame.junction_id = []int{data.Junction_id}

		for sample_id, count := range data.Samples {
			frame.sample_id = []int{sample_id}
			frame.count = []int{count}

			frames = append(frames, *frame)
		}
	}

	return &frames
}

/* Load Metadata should only be done right before printing out frame. */
func Load_Metadata_Into_Frames(f *[]Frame, url string) *[]Frame {
	metadata_offset := 60

	var frames []Frame

	metadata_map := *web.Import_Metadata(url)

	for _, frame := range *f {
		data_slice := metadata_map[frame.First_Sample_ID()]

		if len(data_slice) > 0 {
			metadata := data_slice[metadata_offset]
			frame.Set_Metadata(metadata)
		}

		frames = append(frames, frame)
	}

	return &frames
}


func Group_Frames_By_Junction_ID() {
	group_frames_by = By_Junction_ID{}
}

func Group_Frames_By_Sample_ID() {
	group_frames_by = By_Sample_ID{}
}
