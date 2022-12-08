package task2

func SliceShift(data []int, shift int) []int {
	length := len(data)
	res := make([]int, length)

	if shift == length {
		copy(res, data)
		return res
	}

	for i, val := range data {
		nextPosition := i + shift
		for nextPosition > length-1 {
			nextPosition -= length
		}

		res[nextPosition] = val
	}

	return res
}
