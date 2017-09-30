package query

import (
	"errors"
	"fmt"
	"bytes"
)


type region struct {
	chromosome_base string
	start_pos int
	end_pos int

	gene_symbol string

	/* Modifiers */
	contains bool
	exact bool
	either_start bool
	either_end bool

	set_contains bool
	set_exact bool
	set_either bool
}

func Region() region {
	var r region
	r.start_pos = -1
	r.end_pos = -1
	return r
}

func (r *region) Chromosome(base string) *region {
	r.chromosome_base = base
	return r
}

func (r *region) Start_Pos(pos int) *region {
	r.start_pos = pos
	return r
}

func (r *region) End_Pos(pos int) *region {
	r.end_pos = pos
	return r
}

func (r *region) Gene(symbol string) *region {
	r.gene_symbol = symbol
	return r
}

func (r *region) Contains(contains_some bool) *region {
	r.contains = contains_some
	r.set_contains = true
	return r
}

func (r *region) Exact(contains_all bool) *region {
	r.exact = contains_all
	r.set_exact = true
	return r
}

func (r *region) Either_Start(contains_start bool) *region {
	r.either_start = contains_start
	r.set_either = true
	return r
}

func (r *region) Either_End(contains_end bool) *region {
	r.either_end = contains_end
	r.set_either = true
	return r
}

func (r *region) Initialized() bool {
	if r.chromosome_base != "" || r.gene_symbol != "" {
		return true
	}
	return false
}

func (r *region) Export() (string, error) {
	var b bytes.Buffer

	if r.chromosome_base != "" {
		if r.start_pos == -1 || r.end_pos == -1 {
			return "", errors.New("Both start and end positions required for chromosome")
		}

		// Region, start_pos, end_pos specified
		b.WriteString(fmt.Sprintf("regions=%s:%d-%d", r.chromosome_base, r.start_pos, r.end_pos))

	} else if r.gene_symbol != "" {
		b.WriteString(fmt.Sprintf("regions=%s", r.gene_symbol))

	} else {
		return "", errors.New("A valid chromosome base or gene symbol required")
	}

	if r.set_contains {
		b.WriteString("&")

		if r.contains {
			b.WriteString("contains=1")
		} else {
			b.WriteString("contains=0")
		}

	} else if r.set_either {
		b.WriteString("&")

		if r.either_start && r.either_end {
			b.WriteString("either=0")
		} else if r.either_start {
			b.WriteString("either=1")
		} else if r.either_end {
			b.WriteString("either=2")
		}

	} else if r.set_exact {
		b.WriteString("&")

		if r.exact {
			b.WriteString("exact=1")
		} else {
			b.WriteString("exact=0")
		}
	}

	return b.String(), nil
}




