package validator

type Validator interface {
	Validate() error
}

type ErrorList []error

func (e ErrorList) Error() string {
	msg := ""
	for _, err := range e {
		msg += err.Error() + "\n"
	}
	return msg
}

func ValidateAll(validators ...Validator) error {
	var errors ErrorList
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}

func Validate(validators ...Validator) error {
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func Ptr[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
