package functional

func Map[T any, S any](source []T, f func(T) S) []S {
	var result []S
	for _, value := range source {
		result = append(result, f(value))
	}
	return result
}

func MapIndexed[T any, S any](source []T, f func(T, int) S) []S {
	var result []S
	for index, value := range source {
		result = append(result, f(value, index))
	}
	return result
}

func Filter[T any](source []T, f func(T) bool) []T {
	var result []T
	for _, value := range source {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
