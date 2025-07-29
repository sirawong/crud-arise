package utils

func GetValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func SetPtr[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}
