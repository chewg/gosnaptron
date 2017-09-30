package query


import (
	"bytes"
	"fmt"
)

type filter struct {
	length int
	length_op string

	no_annotation     bool
	set_no_annotation bool

	strand_plus  bool
	strand_minus bool
	set_strand   bool

	samples_count int
	samples_count_op string
	coverage_sum int
	coverage_sum_op string
	coverage_avg float32
	coverage_avg_op string
	coverage_median float32
	coverage_median_op string
}

func Filter() filter {
	var f filter
	return f
}

func (f *filter) Length(op string, length int) *filter {
	f.length = length
	f.length_op = op
	return f
}

func (f *filter) Annotated(none bool) *filter {
	f.no_annotation = none
	f.set_no_annotation = true
	return f
}

func (f *filter) Strand_Plus(require bool) *filter {
	f.strand_plus = require
	f.set_strand = true
	return f
}

func (f *filter) Strand_Minus(require bool) *filter {
	f.strand_minus = require
	f.set_strand = true
	return f
}

func (f *filter) Samples_Count(op string, num int) *filter {
	f.samples_count = num
	f.samples_count_op = op
	return f
}

func (f *filter) Coverage_Sum(op string, num int) *filter {
	f.coverage_sum = num
	f.coverage_sum_op = op
	return f
}

func (f *filter) Coverage_Average(op string, num float32) *filter {
	f.coverage_avg = num
	f.coverage_avg_op = op
	return f
}

func (f *filter) Coverage_Median(op string, num float32) *filter {
	f.coverage_median = num
	f.coverage_median_op = op
	return f
}

func (f *filter) Initialized() bool {
	if f.length == 0 && f.samples_count == 0 && f.coverage_sum == 0 && f.coverage_avg == 0 && f.coverage_median == 0 {
		return false
	}

	return true
}

func (f *filter) Export() (string, error) {
	var b bytes.Buffer
	appending := false
	
	if f.length != 0 {
		appending = true
		b.WriteString(fmt.Sprintf("rfilter=intron_length%s:%d", f.length_op, f.length))
	}

	if f.samples_count != 0 {
		if appending {
			b.WriteString("&")
		}
		appending = true
		
		b.WriteString(fmt.Sprintf("rfilter=samples_count%s:%d", f.samples_count_op, f.samples_count))
	}


	if f.coverage_sum != 0 {
		if appending {
			b.WriteString("&")
		}
		appending = true
		
		b.WriteString(fmt.Sprintf("rfilter=coverage_sum%s:%d", f.coverage_sum_op, f.coverage_sum))
	}

	if f.coverage_median != 0 {
		if appending {
			b.WriteString("&")
		}
		appending = true
		
		b.WriteString(fmt.Sprintf("rfilter=coverage_median%s:%g", f.coverage_median_op, f.coverage_median))
	}

	if f.coverage_avg != 0 {
		if appending {
			b.WriteString("&")
		}
		appending = true
		
		b.WriteString(fmt.Sprintf("rfilter=coverage_avg%s:%g", f.coverage_avg_op, f.coverage_avg))
	}


	if f.set_no_annotation {
		if appending {
			b.WriteString("&")
		}
		appending = true

		 if f.no_annotation {
		 	b.WriteString(fmt.Sprintf("rfilter=annotated:0"))
		 } else {
		 	b.WriteString(fmt.Sprintf("rfilter=annotated:1"))
		 }
	}

	if f.set_strand {
		if f.strand_plus {
			if appending {
				b.WriteString("&")
			}
			appending = true

			b.WriteString(fmt.Sprintf("rfilter=strand:+"))
		} else if f.strand_minus {
			if appending {
				b.WriteString("&")
			}
			appending = true

			b.WriteString(fmt.Sprintf("rfilter=strand:-"))
		}
	}


	return b.String(), nil
}
