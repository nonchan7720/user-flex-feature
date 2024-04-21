package collection

func Filter[T any](elms []T, fn func(T) bool) []T {
	outputs := make([]T, 0)
	for _, elm := range elms {
		if fn(elm) {
			outputs = append(outputs, elm)
		}
	}
	return outputs
}
