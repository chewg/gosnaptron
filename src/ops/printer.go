package ops

import "fmt"


func Print_jir(group *[]Frame) {
	fmt.Printf("%-10v|%-10v|%-10v|%-10v|%-25v\n", "JIR", "Count 1", "Count 2", "Sample ID", "Metadata")

	for _, frame := range *group {
		frame_id := get_frame_id(&frame)
		fmt.Printf("%-10.5v %-10v %-10v %-10v %-25v\n",
			frame.stat[0], frame.stat[1], frame.stat[2], frame_id, frame.metadata)
	}
}


func Print_ssc(group *[]Frame) {
	fmt.Printf("%-10v|%-25v\n", "Sample ID", "Count")

	for _, frame := range *group {
		frame_id := get_frame_id(&frame)
		fmt.Printf("%-10v %-25v\n", frame_id, frame.First_Count())

	}
}


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
