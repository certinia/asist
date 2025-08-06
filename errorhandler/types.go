package errorhandler

type UserError struct {
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}
