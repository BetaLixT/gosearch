package common

func Pointerify[T any](i T) *T {
	return &i
}
