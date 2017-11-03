package ops


type stat_func func(f *[]Frame) *[]Frame


func Sum_Count_By_Sample_ID(frames *[]Frame) *[]Frame {
	m := map[int]Frame{}

	for _, frame := range *frames {
		if value, exist := m[frame.sample_id]; exist {
			value.count += frame.count
			m[frame.sample_id] = value
		} else {
			// not exist
			m[frame.sample_id] = frame
		}
	}

	return convert_map_to_slice(m)
}





/* For Intermediate Old*/

type junction_id string
type sample_id string
type sample_count string


type mean string
type max string
type min string
type count string
type sum string


type eq string
type neq string
type gt string
type lt string
type gteq string
type lteq string
