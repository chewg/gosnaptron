package ops


/*****
ratio_func

Takes as parameters a variable number of integers. The function body contains the formula for the ratio and the ratio
is computed. The computed value, which is a float32 type, is returned.

Parameters: variable number of ints
Output: float32 which is the value from the ratio formula
*****/
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


/*****
Calculate_Ratio

Assume: In each *[]Frame, there is only 1 count value within each count field slice

Takes in a ratio_func and at least 1 slice of frames, then does a union of the frames and generates the results
of apply the ratio_func on the frames.

Parameters: 1 ratio_func, at least 1 pointer to a slice of frames
Output: address to a slice of frames, contains a populated stat field
*****/
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



/*****
generate_operands

The operands that the ratio_func takes in as parameters are created and returned from this function.

Parameters: 1 slice of int operands, 1 int value of how many operands ratio_func needs
Output: a slice of int containing all the operands the ratio_func needs
*****/
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
