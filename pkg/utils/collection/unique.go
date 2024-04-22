package collection

func Uniq[T comparable](elms []T) []T {
	outputs := make([]T, 0, len(elms))
	m := make(map[T]bool)
	for _, elm := range elms {
		if _, ok := m[elm]; !ok {
			m[elm] = true
			outputs = append(outputs, elm)
		}
	}
	return outputs
}
