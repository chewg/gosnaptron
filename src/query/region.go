package query

import (
	"errors"
	"fmt"
)


type region struct {
	chromosome_base string
	start_pos int
	end_pos int
	gene_symbol string
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

func (r *region) Reset_Chromosome() *region {
	r.chromosome_base = ""
	r.start_pos = -1
	r.end_pos = -1
	return r
}

func (r *region) Reset_Gene() *region {
	r.gene_symbol = ""
	return r
}

func (r *region) Initialized() bool {
	if r.chromosome_base != "" || r.gene_symbol != "" {
		return true
	}
	return false
}

func (r *region) Export() (string, error) {
	if r.chromosome_base != "" {
		if r.start_pos == -1 || r.end_pos == -1 {
			return "", errors.New("Both start and end positions required for chromosome")
		}

		// Region, start_pos, end_pos specified
		return fmt.Sprintf("regions=%s:%d-%d", r.chromosome_base, r.start_pos, r.end_pos), nil
	}


	if r.gene_symbol != "" {
		return fmt.Sprintf("regions=%s", r.gene_symbol), nil
	}

	return "", errors.New("A valid chromosome base or gene symbol required")
}



