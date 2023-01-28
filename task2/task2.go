package task2

func SliceShift(data []int, shift int) []int {
	res := make([]int, len(data))

	for shift > len(data) {
		shift -= len(data)
	}

	if shift == len(data) {
		copy(res, data)
		return res
	}

	for i, val := range data {
		nextPosition := i + shift
		if nextPosition > len(data)-1 {
			nextPosition -= len(data)
		}

		res[nextPosition] = val
	}

	return res
}
