package ops

import (
	"reflect"
	"testing"
)

type fields struct {
	junction_id []int
	sample_id   []int
	count       []int
	stat        []float32
	metadata    string
}

func setup_default_fields() fields {
	f := fields{}
	f.junction_id = []int{1001, 1002, 1003, 1004, 1005}
	f.sample_id = []int{1, 2, 3, 4, 5}
	f.count = []int{3, 6, 9, 12, 15}
	f.stat = []float32{100.100}
	f.metadata = "go-snap"
	return f
}

func TestFrame_First_Junction_ID(t *testing.T) {
	default_fields := setup_default_fields()

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"First Junction ID", default_fields, 1001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.First_Junction_ID(); got != tt.want {
				t.Errorf("Frame.First_Junction_ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_First_Sample_ID(t *testing.T) {
	default_fields := setup_default_fields()

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"First Sample ID", default_fields, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.First_Sample_ID(); got != tt.want {
				t.Errorf("Frame.First_Sample_ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_First_Count(t *testing.T) {
	default_fields := setup_default_fields()

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"First Count", default_fields, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.First_Count(); got != tt.want {
				t.Errorf("Frame.First_Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_Add_Junction_ID(t *testing.T) {
	default_fields := setup_default_fields()

	result_frame := New_Frame()
	result_frame.junction_id = []int{123, 1001, 1002, 1003, 1004, 1005}
	result_frame.sample_id = []int{1, 2, 3, 4, 5}
	result_frame.count = []int{3, 6, 9, 12, 15}
	result_frame.stat = []float32{100.100}
	result_frame.metadata = "go-snap"

	type args struct {
		new_id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Frame
	}{
		{"Add 123", default_fields, args{123}, result_frame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.Add_Junction_ID(tt.args.new_id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Frame.Add_Junction_ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_Add_Sample_ID(t *testing.T) {
	default_fields := setup_default_fields()

	result_frame := New_Frame()
	result_frame.junction_id = []int{1001, 1002, 1003, 1004, 1005}
	result_frame.sample_id = []int{1, 2, 3, 4, 5, 123}
	result_frame.count = []int{3, 6, 9, 12, 15}
	result_frame.stat = []float32{100.100}
	result_frame.metadata = "go-snap"

	type args struct {
		new_id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Frame
	}{
		{"Add 123", default_fields, args{123}, result_frame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.Add_Sample_ID(tt.args.new_id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Frame.Add_Sample_ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrame_Aggregate_Count(t *testing.T) {
	default_fields := setup_default_fields()

	result_frame := New_Frame()
	result_frame.junction_id = []int{1001, 1002, 1003, 1004, 1005}
	result_frame.sample_id = []int{1, 2, 3, 4, 5}
	result_frame.count = []int{45}
	result_frame.stat = []float32{100.100}
	result_frame.metadata = "go-snap"

	type args struct {
		fn aggreg_func
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Frame
	}{
		{"Sum Count", default_fields, args{aggreg_sum}, result_frame},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Frame{
				junction_id: tt.fields.junction_id,
				sample_id:   tt.fields.sample_id,
				count:       tt.fields.count,
				stat:        tt.fields.stat,
				metadata:    tt.fields.metadata,
			}
			if got := f.Aggregate_Count(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Frame.Aggregate_Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
