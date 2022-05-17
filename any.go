package validator

type AnyV struct {
	*GenericValidator[any]
}

func Any(value any, name string) *AnyV {
	return &AnyV{NewGeneric(value, name)}
}

func (v *AnyV) Required(msg ...string) *AnyV {
	v.validations = append(v.validations, func() error {
		if v.value != nil {
			return nil

		}
		return v.NewError(msgRequired, msg)
	})
	return v
}

func (v *AnyV) Validate() error {
	for _, validation := range v.validations {
		if err := validation(); err != nil {
			return err
		}
	}

	if a, ok := v.value.(Validator); ok && a != nil {
		return a.Validate()
	}

	return nil
}
