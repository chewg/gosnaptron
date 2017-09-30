package data

/* what is returned from snaptron server */
type data struct {
	datasource string
	id string
	chromosome string
	start int
	end int
	strand byte
	annotated bool
	left_motif string
	right_motif	string
	left_annotated string
	right_annotated	string
	samples	[]int
	samples_count int
	coverage_sum float32
	coverage_avg float32
	coverage_median float32
	source_dataset_id int
}
