package tools

func Merge[T any](cs ...chan T) chan T {
	var res = make(chan T)

	for _, ch := range cs {
		go func(from, to chan T) {
			for v := range from {
				to <- v
			}
		}(ch, res)
	}
	return res
}

func Append[T any](slices ...[]T) []T {
	var res []T
	for _, slice := range slices {
		res = append(res, slice...)
	}
	return res
}
