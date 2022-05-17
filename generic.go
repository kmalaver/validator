package validator

type GenericValidator[T any] struct {
	value       T
	name        string
	validations []func() error
}

func NewGeneric[T any](value T, name string) *GenericValidator[T] {
	return &GenericValidator[T]{value: value, name: name}
}

func (g *GenericValidator[T]) Validate() error {
	for _, f := range g.validations {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
