package apperror

type ErrorType int8

const (
	_ ErrorType = iota
	ErrorTypeInternal
	ErrorTypeNotFound
	ErrorTypeInvalidRequest
)
