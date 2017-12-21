package ops

import "fmt"


/*****
Print_jir

Takes as parameters 1 pointer to a slice of frames. The function body prints to stdout the results of the JIR High
Level example.

Parameters: 1 pointer to a slice of frames
Output: none is returned
*****/
func Print_jir(group *[]Frame) {
	fmt.Printf("%-10v|%-10v|%-10v|%-10v|%-25v\n", "JIR", "Count 1", "Count 2", "Sample ID", "Metadata")

	for _, frame := range *group {
		frame_id := get_frame_id(&frame)
		fmt.Printf("%-10.5v %-10v %-10v %-10v %-25v\n",
			frame.stat[0], frame.stat[1], frame.stat[2], frame_id, frame.metadata)
	}
}


/*****
Print_ssc

Takes as parameters 1 pointer to a slice of frames. The function body prints to stdout the results of the SSC High
Level example.

Parameters: 1 pointer to a slice of frames
Output: none is returned
*****/
func Print_ssc(group *[]Frame) {
	fmt.Printf("%-10v|%-25v\n", "Sample ID", "Count")

	for _, frame := range *group {
		frame_id := get_frame_id(&frame)
		fmt.Printf("%-10v %-25v\n", frame_id, frame.First_Count())

	}
}


/*****
Print_tsv

Takes as parameters 1 pointer to a slice of frames that is the intersect group and another pointer that is the union
group. The check of whether there is presence is performed in the function body. The function then prints to stdout
the results of the TSV High Level example.

Parameters: 1 pointer to a slice of frames, 1 pointer to a lice of frames
Output: none is returned
*****/
func Print_tsv(intersect_group, union_group *[]Frame) {
	intersect_map := *Convert_Slice_To_Map(intersect_group)

	fmt.Printf("%-10v|%-3v|%-25v\n", "Sample ID", "0/1", "Metadata")

	for _, frame := range *union_group {
		frame_id := get_frame_id(&frame)

		if _, exist := intersect_map[frame_id]; exist {
			fmt.Printf("%-10v %-3v %-25v\n", frame_id, "1", frame.metadata)
		}
	}
}
