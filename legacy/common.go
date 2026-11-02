package legacy

type WrappedResponse[T any] struct {
	Data T `json:"data"`
}

func Pointer[T any](v T) *T {
	return &v
}
