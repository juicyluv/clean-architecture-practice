package apperror

type AppError struct {
	Message string    `json:"message"`
	Type    ErrorType `json:"-"`
}

func (ae AppError) Error() string {
	return ae.Message
}
