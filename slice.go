package validator

import "fmt"

type SliceType interface {
	[]any
}

type SliceV[T []V, V any] struct {
	*GenericValidator[T]
}

func Slice[T []V, V any](value T, name string) *SliceV[T, V] {
	return &SliceV[T, V]{NewGeneric(value, name)}
}

func (v *SliceV[T, V]) Min(min int, msg ...string) *SliceV[T, V] {
	v.validations = append(v.validations, func() error {
		if len(v.value) < min {
			return v.NewError(fmt.Sprintf(msgMin, min), msg)
		}
		return nil
	})
	return v
}

func (v *SliceV[T, V]) Max(max int, msg ...string) *SliceV[T, V] {
	v.validations = append(v.validations, func() error {
		if len(v.value) > max {
			return v.NewError(fmt.Sprintf(msgMax, max), msg)
		}
		return nil
	})
	return v
}

func (v *SliceV[T, V]) Length(length int, msg ...string) *SliceV[T, V] {
	v.validations = append(v.validations, func() error {
		if len(v.value) != length {
			return v.NewError(fmt.Sprintf(msgLength, length), msg)
		}
		return nil
	})
	return v
}

func (v *SliceV[T, V]) Check(checkFun func(value V, name string) Validator) *SliceV[T, V] {
	v.validations = append(v.validations, func() error {
		for i, value := range v.value {
			if err := checkFun(value, fmt.Sprintf("%s: at position %d", v.name, i)).Validate(); err != nil {
				return err
			}
		}
		return nil
	})
	return v
}
