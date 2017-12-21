package ops

import (
	"math"
	"snaptron_api/src/query"
	"strings"
	"strconv"
	"snaptron_api/src/server"
)


/*****
order_by

Loads the order_funcs into a sorter type struct to be used with sorting slices of Frames

Parameters: variable number of order_func
Output: address to a sorter type struct
*****/
func order_by(order ...order_func) *sorter {
	return &sorter{
		order: order,
	}
}


/*****
append_ints

For appending 2 slices of ints

Parameters: 2 slices of ints
Output: 1 slice of ints
*****/
func append_ints(s1 []int, s2 []int) []int {
	return append(s1, s2...)
}


/*****
Convert_Map_To_Slice

Converts a map of Frames to a slice of Frames. Commonly used to ensure intermediate.go functions have a uniform
return type

Parameters: pointer to a map
Output: address of a slice of Frames
*****/
func Convert_Map_To_Slice(m *map[int]Frame) *[]Frame {
	var frames []Frame

	// no guarantee with map's order
	for _, frame := range *m {
		frames = append(frames, frame)
	}

	return &frames
}


/*****
Convert_Slice_To_Map

Converts a slice of Frames to a map of Frames. Distinctness is an inherent property within the conversion

Parameters: pointer to a slice of Frames
Output: address of a map
*****/
func Convert_Slice_To_Map(f *[]Frame) *map[int]Frame {
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


/*****
get_frame_id

Commonly used when one needs to get the ID of the frame, whether by Sample ID or Junction ID depending on
what is globally set in the ops package.

Parameters: pointer to a Frame
Output: Frame ID, which is either the Frame's Sample ID or Junction ID
*****/
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


/*****
slice_less_than

For doing comparison of whether one slice of ints is less than another slice of ints

Parameters: 2 slices of ints
Output: bool of whether the first slice is less than the second slice
*****/
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


/*****
Dataframe_To_Frames

To convert the from the Dataframe type struct in src/query/ to the Frames type struct in src/ops/

Parameters: 1 pointer to slice of Dataframes
Output: address of a slice of frames
*****/
func Dataframe_To_Frames(data_frame *query.Dataframe) *[]Frame {
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


/*****
Load_Metadata_Into_Frames

Loads the metadata into each of the Frame's metadata field. Ideally, the function should only be run right before
printing so that the metadata stays even while other intermediate.go functions are run.

Parameters: 1 pointer to slice of Frames, string of datasource (srav1, gtex, etc.), and the offset of metadata
Output: address of a slice of frames with metadata loaded
*****/
func Load_Metadata_Into_Frames(f *[]Frame, datasource string, metadata_offset int) *[]Frame {
	var frames []Frame

	metadata_map := *Import_Metadata(datasource)

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


/*****
Import_Metadata

Fetches the metadata from snaptron server and formats it into a map for fast lookup when it will later
be use for filling in the metadata field in Frames.

Parameters: 1 string that specifies datasource (srav1, gtex, etc.)
Output: address of a map which contains the data formatted for internal use
*****/
func Import_Metadata(datasource string) *map[int][]string {
	metadata_string := server.Get_Metadata_From_Server(datasource)
	metadata_slice := strings.Split(metadata_string, "\n")

	var row_slice [][]string

	for _, metadata := range metadata_slice {
		fields := strings.Split(metadata, "\t")
		row_slice = append(row_slice, fields)
	}

	row_slice = row_slice[1:]

	metadata_map := map[int][]string{}

	for _, row := range row_slice {
		i32, _ := strconv.ParseInt(row[0], 10, 32)
		sample_id := int(i32)

		metadata_map[sample_id] = row
	}

	return &metadata_map
}


/*****
Group_Frames_By_Junction_ID

Helps set the group_frames_by variable (global variable in ops package) to group by Junction ID.

Parameters: none
Output: none
*****/
func Group_Frames_By_Junction_ID() {
	group_frames_by = By_Junction_ID{}
}


/*****
Group_Frames_By_Sample_ID

Helps set the group_frames_by variable (global variable in ops package) to group by Sample ID.

Parameters: none
Output: none
*****/
func Group_Frames_By_Sample_ID() {
	group_frames_by = By_Sample_ID{}
}
