package cartesian

// Product returns cartesian product of arguments. Each argument must be a slice
// of the same kind.
func Product[T any](dimensions ...[]T) [][]T {
	if len(dimensions) == 0 {
		return nil
	}

	head, tail := dimensions[0], dimensions[1:]
	tailProduct := Product[T](tail...)

	var result [][]T
	for _, v := range head {
		if len(tailProduct) == 0 {
			result = append(result, []T{v})
		} else {
			for _, ttail := range tailProduct {
				element := []T{
					v,
				}
				element = append(element, ttail...)
				result = append(result, element)
			}
		}

	}
	return result
}
