package fims

func Apply[T, R any](
	collection []T,
	iteratee func(T) R,
) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item)
	}

	return result
}

func Reduce[T any, R any](
	collection []T,
	iteratee func(acc R, item T, index int) R,
) R {
	var accumulator R

	for i, item := range collection {
		accumulator = iteratee(accumulator, item, i)
	}

	return accumulator
}

func FilterSlice[T any](
	collection []T,
	predicate func(item T, index int) bool,
) []T {
	result := make([]T, 0)

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}
