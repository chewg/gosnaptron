package ops

import (
	"sort"
	"fmt"
)

/* This slice implementation keeps track of the original indices
Can be useful. Will keep here for now.
*/

type Slice struct {
	sort.Interface
	index []int
}

func (s Slice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.index[i], s.index[j] = s.index[j], s.index[i]
}


func NewSlice(n sort.Interface) *Slice {
	s := &Slice{Interface: n, index: make([]int, n.Len())}

	for i := range s.index {
		s.index[i] = i
	}
	return s
}

func NewIntSlice(n ...int) *Slice {
	return NewSlice(sort.IntSlice(n))
}


// todo remove entire file
func test_slice() {
	s := NewIntSlice(5, 8, 3, 10, 1)
	sort.Sort(s)
	fmt.Println(s.Interface, s.index)
}