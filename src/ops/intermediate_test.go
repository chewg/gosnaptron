package ops

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	type args struct {
		frames_1    *[]Frame
		more_frames []*[]Frame
	}

	frame_slice_1 := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{2, 4}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 3}, []int{21}, []int{5}, []float32{0.0}, "",
		},
	}

	frame_slice_2 := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{1, 3, 5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
	}

	test_args_1 := args{
		&frame_slice_1,
		[]*[]Frame{&frame_slice_2},
	}

	frame_slice_result_1 := []Frame{
		Frame{
			[]int{1, 1}, []int{20}, []int{1, 2, 3, 4, 5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 3}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
	}

	frame_slice_3 := []Frame{
		Frame{
			[]int{10, 20}, []int{20}, []int{10, 15, 20}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{311}, []int{3}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{411}, []int{4}, []float32{0.0}, "",
		},
	}

	test_args_2 := args{
		&frame_slice_1,
		[]*[]Frame{&frame_slice_2, &frame_slice_3},
	}

	frame_slice_result_2 := []Frame{
		Frame{
			[]int{1, 1, 10, 20}, []int{20}, []int{1, 2, 3, 4, 5, 10, 15, 20}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 3}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{311}, []int{3}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{411}, []int{4}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"Union of 2", test_args_1, &frame_slice_result_1},
		{"Union of 3", test_args_2, &frame_slice_result_2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.frames_1, tt.args.more_frames...); !reflect.DeepEqual(Order(got, Incr_Sample_ID), Order(tt.want, Incr_Sample_ID)) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		frames_1    *[]Frame
		more_frames []*[]Frame
	}

	frame_slice_1 := []Frame{
		Frame{
			[]int{1, 2, 5}, []int{20}, []int{2, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
	}

	frame_slice_2 := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{1, 3}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
	}

	test_args_1 := args{
		&frame_slice_1,
		[]*[]Frame{&frame_slice_2},
	}

	frame_slice_result_1 := []Frame{
		Frame{
			[]int{1, 1, 2, 5}, []int{20}, []int{1, 2, 3, 6, 8}, []float32{0.0}, "",
		},
	}

	frame_slice_3 := []Frame{
		Frame{
			[]int{10, 20}, []int{20}, []int{10, 15, 20}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{311}, []int{3}, []float32{0.0}, "",
		},
		Frame{
			[]int{10}, []int{411}, []int{4}, []float32{0.0}, "",
		},
	}

	test_args_2 := args{
		&frame_slice_1,
		[]*[]Frame{&frame_slice_2, &frame_slice_3},
	}

	frame_slice_result_2 := []Frame{
		Frame{
			[]int{1, 1, 2, 5, 10, 20}, []int{20}, []int{1, 2, 3, 6, 8, 10, 15, 20}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"Intersection of 2", test_args_1, &frame_slice_result_1},
		{"Intersection of 3", test_args_2, &frame_slice_result_2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.frames_1, tt.args.more_frames...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_By(t *testing.T) {

	frame_slice_1 := &[]Frame{
		Frame{
			[]int{1, 2, 3}, []int{20, 21}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
	}

	result_by_junction_id := &[]Frame{
		Frame{
			[]int{1}, []int{20, 21, 21}, []int{2, 4, 5, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{2}, []int{20, 21}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{3}, []int{20, 21}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
	}

	result_by_sample_id := &[]Frame{
		Frame{
			[]int{1, 1, 2, 3}, []int{21}, []int{2, 4, 5, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 2, 3}, []int{20}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		group_by_func func()
		frames *[]Frame
		want *[]Frame
	}{
		{"By_Junction_ID", Group_Frames_By_Junction_ID, frame_slice_1, result_by_junction_id},
		{"By_Sample_ID", Group_Frames_By_Sample_ID, frame_slice_1, result_by_sample_id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Group_By(tt.group_by_func, tt.frames); !reflect.DeepEqual(Order(got, Incr_Junction_ID), tt.want) {
				t.Errorf("Group_By() = %v, want %v with global %v", got, tt.want)
			}
		})
	}
}


func TestOrder(t *testing.T) {
	type args struct {
		frames *[]Frame
		order  []order_func
	}

	frame_slice_1 := &[]Frame{
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{7}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{1}, []float32{0.0}, "",
		},
	}

	result_frame := &[]Frame{
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{7}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{1}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"Order", args{frame_slice_1, []order_func{Incr_Sample_ID, Incr_Count}}, result_frame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Order(tt.args.frames, tt.args.order...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBind(t *testing.T) {
	type args struct {
		frames_group []*[]Frame
	}

	frame_slice_1 := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{2}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
	}

	frame_slice_2 := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{1}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
	}

	test_args := args{
		[]*[]Frame{&frame_slice_1, &frame_slice_2},
	}

	frame_slice_result := []Frame{
		Frame{
			[]int{1}, []int{20}, []int{2}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{20}, []int{1}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{31}, []int{3}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"2 Frames", test_args, &frame_slice_result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bind(tt.args.frames_group...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		frames        *[]Frame
		filter_tuples []Filter_tuple
	}

	frame_slice_1 := &[]Frame{
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{2, 4}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{10}, []float32{0.0}, "",
		},
	}

	result_frame := &[]Frame{
		Frame{
			[]int{1}, []int{21}, []int{10}, []float32{0.0}, "",
		},
	}

	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"Filter > 5", args{frame_slice_1, []Filter_tuple{{Sample_Count_Gt, 5}}}, result_frame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.frames, tt.args.filter_tuples...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummarize(t *testing.T) {
	type args struct {
		frames *[]Frame
		stats  []stat_func
	}

	frame_slice_1 := &[]Frame{
		Frame{
			[]int{1, 2, 3}, []int{22}, []int{2, 4, 6, 8}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{5}, []float32{0.0}, "",
		},
		Frame{
			[]int{1}, []int{21}, []int{10}, []float32{0.0}, "",
		},
		Frame{
			[]int{3}, []int{22}, []int{10}, []float32{0.0}, "",
		},
	}

	result_frame_slice := &[]Frame{
		Frame{
			[]int{1, 1}, []int{21}, []int{15}, []float32{0.0}, "",
		},
		Frame{
			[]int{1, 2, 3, 3}, []int{22}, []int{30}, []float32{0.0}, "",
		},
	}


	tests := []struct {
		name string
		args args
		want *[]Frame
	}{
		{"Sum Count", args{frame_slice_1, []stat_func{Sum_Count}}, result_frame_slice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summarize(tt.args.frames, tt.args.stats...); !reflect.DeepEqual(Order(got, Incr_Junction_ID), tt.want) {
				t.Errorf("Summarize() = %v, want %v", got, tt.want)
			}
		})
	}
}
