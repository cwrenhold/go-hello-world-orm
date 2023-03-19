package utils

func Filter[T interface{}](slice []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, elem := range slice {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}
