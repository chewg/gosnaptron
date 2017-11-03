package examples

import (
	"sync"
	"snaptron_api/src/data"
	"fmt"
)

// multithreading example

func load_dataframes_to_channel(df data.Dataframe, ch chan interface{}) {
	frames := df.Frames()

	for i := 0; i < len(frames); i++ {
		select {
			case ch <- frames[i].Samples:
			default:
				fmt.Println("load_dataframes_to_channel: channel is full, discarding value.")
		}
	}
}


func Shared_Sample_Count(df data.Dataframe) map[int]int {

	queue := make(chan interface{}, len(df.Frames()))

	load_dataframes_to_channel(df, queue)

	//select {
	//case queue <- 5:
	//default:
	//	fmt.Println("Shared_Sample_Count: channel is full, discarding value.")
	//}
	//print(len(queue))

	wg := sync.WaitGroup{}

	for {
		for len(queue) >= 2 {
			first_value := <-queue
			second_value := <-queue

			wg.Add(1)
			go outer_join(first_value.(map[int]int), second_value.(map[int]int), queue, &wg)
		}

		wg.Wait()

		if len(queue) <= 1 {
			return_map := <-queue
			return return_map.(map[int]int)
		} else {
			// not finished, continue loop
			wg = sync.WaitGroup{}
		}
	}

}


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
		if larger_val, ok := larger[key]; ok {
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
		if larger_val, ok := larger[key]; ok {
			smaller[key] = smaller_val + larger_val
		} else {
			// fmt.Println("Deleting %d", key)
			delete(smaller, key)
		}
	}

	ch <- smaller
	wg.Done()
}



//func Unbounded_Channel() (chan<- interface{}, <-chan interface{}) {
//	input := make(chan interface{})
//	output := make(chan interface{})
//
//	go func() {
//		var in_queue []interface{}
//
//		out_chan := func() chan interface {} {
//			if len(in_queue) <= 0 {
//				return nil
//			}
//			return output
//		}
//
//		get_value := func() interface{} {
//			if len(in_queue) <= 0 {
//				return nil
//			}
//			return in_queue[0]
//		}
//
//		for len(in_queue) > 0 || input != nil {
//			select {
//			case value, ok := <-input:
//				if ok {
//					in_queue = append(in_queue, value)
//				} else {
//					input = nil
//				}
//			case out_chan() <- get_value():
//				in_queue = in_queue[1:]
//			}
//		}
//
//		close(output)
//	}()
//
//	return input, output
//}