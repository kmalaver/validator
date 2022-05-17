package validator

import (
	"fmt"
	"regexp"
)

type StringType interface {
}

type StringV struct {
	*GenericValidator[string]
}

func String(value string, name string) *StringV {
	return &StringV{NewGeneric(value, name)}
}

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	urlRegex   = regexp.MustCompile(`^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(([0-9]{1,5})?\/.*)?$`)
)

func (v *StringV) Required(msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if v.value == "" {
			return v.NewError(msgRequired, msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Email(msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if v.value != "" && !emailRegex.MatchString(v.value) {
			return v.NewError(msgEmail, msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Min(min int, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if len(v.value) < min {
			return v.NewError(fmt.Sprintf(msgMin, min), msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Max(max int, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if len(v.value) > max {
			return v.NewError(fmt.Sprintf(msgMax, max), msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Length(length int, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if len(v.value) != length {
			return v.NewError(fmt.Sprintf(msgLength, length), msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Url(msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		if v.value != "" && !urlRegex.MatchString(v.value) {
			return v.NewError("msgUrl", msg)
		}
		return nil
	})
	return v
}

func (v *StringV) Regex(pattern string, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		match, err := regexp.MatchString(pattern, v.value)
		if err != nil {
			return v.NewError(fmt.Sprintf("msgRegex: %s", err), msg)
		}
		if !match {
			return v.NewError("msgRegex", msg)
		}
		return nil
	})
	return v
}

func (v *StringV) OneOf(values []string, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		for _, value := range values {
			if v.value == value {
				return nil
			}
		}
		return v.NewError("msgOneOf", msg)
	})
	return v
}

func (v *StringV) NotOneOf(values []string, msg ...string) *StringV {
	v.validations = append(v.validations, func() error {
		for _, value := range values {
			if v.value == value {
				return v.NewError("msgNotOneOf", msg)
			}
		}
		return nil
	})
	return v
}
