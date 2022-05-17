package validator

const (
	msgRequired = "is required"
	msgEmail    = "is not a valid email"
	msgMin      = "must be at least %v"
	msgMax      = "must be less than %v"
	msgLength   = "must be %v characters"
	msgMatch    = "does not match the pattern"
)

type Error struct {
	Name    string
	Message string
}

func (e Error) Error() string {
	return e.Name + ": " + e.Message
}

func (g *GenericValidator[T]) NewError(defMessage string, customMessage []string) *Error {
	msg := defMessage
	if len(customMessage) > 0 {
		msg = customMessage[0]
	}
	return &Error{
		Name:    g.name,
		Message: msg,
	}
}
