// todo, this file will be superceded by intermediate.go

package ops

import (
	"sync"
	"sort"
)


type sample struct {
	junction_id int
	id int
	count int
}

// need to be able to rearrange according to the struct field one wants
func arrange(samples []sample, field string, op string) []sample {
	if field == "id" {
		if op == "<" {
			sort.Slice(samples, func(i, j int) bool {
				return samples[i].id < samples[j].id
			})
		} else if op == ">" {
			sort.Slice(samples, func(i, j int) bool {
				return samples[i].id > samples[j].id
			})
		}
	} else if field == "count" {
		if op == "<" {
			sort.Slice(samples, func(i, j int) bool {
				return samples[i].count < samples[j].count
			})
		} else if op == ">" {
			sort.Slice(samples, func(i, j int) bool {
				return samples[i].count > samples[j].count
			})
		}
	}

	return samples
}


//func summarize(samples []sample, op string) (stat int) {
//	if op == "min" {
//		if len(samples) > 0 {
//			stat = samples[0].count
//		}
//
//		for _, e := range samples {
//			if e.count < stat {
//				stat = e.count
//			}
//		}
//	} else if op == "max" {
//		if len(samples) > 0 {
//			stat = samples[0].count
//		}
//
//		for _, e := range samples {
//			if e.count > stat {
//				stat = e.count
//			}
//		}
//	}
//
//	return stat
//}


//func filter_by(op string, samples []sample, sample_ids ...int) (result []sample) {
//	// Build map from ...int for O(1) lookup
//	sample_id_map := make(map[int]bool)
//
//	for _, id := range sample_ids {
//		sample_id_map[id] = true
//	}
//
//	result = []sample{}
//
//	if op == "==" {
//		for _, s  := range samples {
//			if sample_id_map[s.id] {
//				result = append(result, s)
//			}
//		}
//	} else if op == "!=" {
//		for _, s  := range samples {
//			if !sample_id_map[s.id] {
//				result = append(result, s)
//			}
//		}
//	}
//
//	return result
//}

/*
TODO Streamline and cleanup the chain

What the chain is currently constructed like,
in accordance to the data structures

The chain has an input of []sample

---> 1. group_by
---> 2. summarize_by or filter_by
---> 3. arrange_by

Then the chain has an output of []interface{}

 */


/*
Sort op
gt is in descending order
lt is in ascending order
 */
func arrange_by(op interface{}, field interface{}, data map[interface{}][]interface{}) []interface{} {
	arrangement := map[interface{}][]interface{}{}

	// TODO sort by key, or sort by value. for sort by value refer to conversion.go

	switch field.(type) {
	case junction_id:
	case sample_id:
		switch op.(type) {
		case gt:
		case lt:
		}
	case sample_count:
		switch op.(type) {
		case gt:
		case lt:
		}
	}

	return arrangement["lala"]
}




func Test_Group_By() {
	var j junction_id
	var s []sample
	group_by(j, s)
}

func group_by(id_type interface{}, samples []sample) map[interface{}][]interface{} {

	groups := map[interface{}][]interface{}{}

	switch id_type.(type) {
	case junction_id:
		for _, s := range samples {
			groups[s.junction_id] = append(groups[s.junction_id], s.count)
		}
	case sample_id:
		for _, s := range samples {
			groups[s.id] = append(groups[s.id], s.count)
		}
	}

	return groups
}


func summarize_by(op interface{}, data map[interface{}][]interface{}) map[interface{}][]interface{} {

	summary_stats := map[interface{}][]interface{}{}

	switch op.(type) {
	case count:
		for id, count_slice := range data {
			var sum_slice []interface{}
			summary_stats[id] = append(sum_slice, len(count_slice))
		}
	case max:
		for id, count_slice := range data {
			max := 0
			for _,e := range count_slice {
				if max < e.(int) {
					max = e.(int)
				}
			}

			var max_slice []interface{}
			summary_stats[id] = append(max_slice, max)
		}
	case mean:
		for id, count_slice := range data {
			average := 0
			for _,e := range count_slice {
				average += e.(int)
			}

			average = average / len(count_slice)

			var avg_slice []interface{}
			summary_stats[id] = append(avg_slice, average)
		}
	case min:
		for id, count_slice := range data {
			min := 0
			for _,e := range count_slice {
				if min > e.(int) {
					min = e.(int)
				}
			}

			var min_slice []interface{}
			summary_stats[id] = append(min_slice, min)
		}

	case sum:
		for id, count_slice := range data {
			sum := 0
			for _, e := range count_slice {
				sum += e.(int)
			}

			var sum_slice []interface{}
			summary_stats[id] = append(sum_slice, sum)
		}
	}

	return summary_stats
}


func filter_by(op interface{}, data map[interface{}][]interface{}, ids ...int) map[interface{}][]interface{} {

	filtered := map[interface{}][]interface{}{}

	switch op.(type) {
	case eq:
		for _, id := range ids {
			if value, exist := data[id]; exist {
				filtered[id] = value
			}
		}
	case neq:
		for _, id := range ids {
			if _, exist := data[id]; exist {
				delete(data, id)
			}
		}

		filtered = data
	}

	return filtered
}


/* TODO replace old joins */

func outer_join(first map[int]int, second map[int]int, ch chan<- interface{}, wg *sync.WaitGroup) {
	var smaller map[int]int
	var larger map[int]int

	if len(first) < len(second) {
		smaller = first
		larger = second
	} else {
		smaller = second
		larger = first
	}

	for key, smaller_val := range smaller {
		if larger_val, exist := larger[key]; exist {
			larger[key] = smaller_val + larger_val
		} else {
			larger[key] = smaller_val
		}
	}

	ch <- larger
	wg.Done()
}


func inner_join(first map[int]int, second map[int]int, ch chan<- interface{}, wg *sync.WaitGroup) {
	var smaller map[int]int
	var larger map[int]int

	if len(first) < len(second) {
		smaller = first
		larger = second
	} else {
		smaller = second
		larger = first
	}

	for key, smaller_val := range smaller {
		if larger_val, exist := larger[key]; exist {
			smaller[key] = smaller_val + larger_val
		} else {
			// fmt.Println("Deleting %d", key)
			delete(smaller, key)
		}
	}

	ch <- smaller
	wg.Done()
}
