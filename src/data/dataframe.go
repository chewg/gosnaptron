package data

import (
	"strings"
	"strconv"
)

/* what is returned from snaptron server */

type Dataframe struct {
	frames []Data
}

func DataFrame() Dataframe {
	var df Dataframe
	return df
}

func (df * Dataframe) Frames() []Data {
	return df.frames
}

func (df *Dataframe) Load_DataFrames(frames ...[]string) *Dataframe {
	for _, frame := range frames {
		d := Data{}

		d.datasource = frame[0]

		i32, _ := strconv.ParseInt(frame[1], 10, 32)
		d.Junction_id = int(i32)

		d.chromosome = frame[2]

		i32, _ = strconv.ParseInt(frame[3], 10, 32)
		d.start = int(i32)

		i32, _ = strconv.ParseInt(frame[4], 10, 32)
		d.end = int(i32)

		i32, _ = strconv.ParseInt(frame[5], 10, 32)
		d.length = int(i32)

		d.strand = frame[6]

		// annotated is weird
		d.annotated, _ = strconv.ParseBool(frame[7])

		d.left_motif = frame[8]
		d.right_motif = frame[9]
		d.left_annotated = frame[10]
		d.right_annotated = frame[11]

		frame_12 := strings.Trim(frame[12], " ,")
		samples := strings.Split(frame_12,",")

		var samples_int = make(map[int]int)

		for _, pair := range samples {
			fields := strings.Split(pair, ":")

			key, _ := strconv.Atoi(fields[0])
			value, _ := strconv.Atoi(fields[1])

			samples_int[key] = value
		}

		d.Samples = samples_int

		i32, _ = strconv.ParseInt(frame[13], 10, 32)
		d.samples_count = int(i32)

		f32, _ := strconv.ParseFloat(frame[14], 32)
		d.coverage_sum = float32(f32)

		f32, _ = strconv.ParseFloat(frame[15], 32)
		d.coverage_avg = float32(f32)

		f32, _ = strconv.ParseFloat(frame[16], 32)
		d.coverage_median = float32(f32)

		i32, _ = strconv.ParseInt(frame[17], 10, 32)
		d.source_dataset_id = int(i32)

		df.frames = append(df.frames, d)
	}

	return df
}


type Data struct {
	datasource string
	Junction_id int
	chromosome string
	start int
	end int
	length int
	strand string
	annotated bool
	left_motif string
	right_motif	string
	left_annotated string
	right_annotated	string
	Samples map[int]int
	samples_count int
	coverage_sum float32
	coverage_avg float32
	coverage_median float32
	source_dataset_id int
}
