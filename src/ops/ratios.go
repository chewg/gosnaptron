package ops

type ratio_func func(...int) float32


func JIR_Ratio(nums ...int) float32 {
	if len(nums) != 2 {
		panic("jir_ratio needs exactly 2 ints passed in.")
	}

	num1 := float32(nums[0])
	num2 := float32(nums[1])

	return ((num2 - num1) / (num1 + num2 + 1))
}


func PSI_Ratio(nums ...int) float32 {
	if len(nums) != 3 {
		panic("psi_ratio needs exactly 3 ints passed in.")
	}

	num1 := float32(nums[0])
	num2 := float32(nums[1])
	num3 := float32(nums[2])

	mean := (num1 + num2) / 2.0

	return (mean / (mean + num3))
}


// assume that there is only 1 count value within each slice of counts for each *[]Frame
func Calculate_Ratio(ratio ratio_func, f1 *[]Frame, fs ...*[]Frame) *[]Frame {
	union := Union(f1, fs...)
	NUM_OPERANDS := 1 + len(fs)

	frames_with_ratio := make([]Frame, 0)

	for _, frame := range *union {
		operands := generate_operands(frame.count, NUM_OPERANDS)

		stat := ratio(operands...)

		frame.stat = append(frame.stat, stat)
		for _, op := range operands {
			frame.stat = append(frame.stat, float32(op))
		}

		frames_with_ratio = append(frames_with_ratio, frame)
	}

	return &frames_with_ratio
}


func generate_operands(operands []int, total int) []int {
	return_operands := make([]int, 0)
	return_operands = append_ints(return_operands, operands)

	remaining := total - len(operands)

	for (remaining > 0) {
		return_operands = append(return_operands, 0)
		remaining--
	}

	return return_operands
}