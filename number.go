package validator

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type NumberV[T Numeric] struct {
	*GenericValidator[T]
}

func Number[T Numeric](value T, name string) *NumberV[T] {
	return &NumberV[T]{NewGeneric(value, name)}
}

func (v *NumberV[T]) Required(msg ...string) *NumberV[T] {
	v.validations = append(v.validations, func() error {
		var zero T
		if v.value == zero {
			return v.NewError(msgRequired, msg)
		}
		return nil
	})
	return v
}

func (v *NumberV[T]) Min(min T, msg ...string) *NumberV[T] {
	v.validations = append(v.validations, func() error {
		if v.value < min {
			return v.NewError(fmt.Sprintf(msgMin, min), msg)
		}
		return nil
	})
	return v
}

func (v *NumberV[T]) Max(max T, msg ...string) *NumberV[T] {
	v.validations = append(v.validations, func() error {
		if v.value > max {
			return v.NewError(fmt.Sprintf(msgMax, max), msg)
		}
		return nil
	})
	return v
}
