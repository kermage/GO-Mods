package fims

func Columns[T, R comparable](
	collection []T,
	key string,
	iteratee func(T) R,
) map[R]T {
	result := make(map[R]T, len(collection))

	for _, item := range collection {
		result[iteratee(item)] = item
	}

	return result
}

func Flip[K comparable, V comparable](
	m map[K]V,
) map[V]K {
	result := make(map[V]K, len(m))

	for k, v := range m {
		result[v] = k
	}

	return result
}
