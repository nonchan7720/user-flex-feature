package collection

func ToInterface[T any](values []T) []interface{} {
	results := make([]interface{}, len(values))
	for idx, value := range values {
		results[idx] = value
	}
	return results
}

func ToValues[T any](values []*T) []T {
	result := make([]T, len(values))
	for idx, v := range values {
		result[idx] = *v
	}
	return result
}
