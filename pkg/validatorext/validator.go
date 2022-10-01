package validatorext

import (
	"context"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	mobileRegexString = `^(((\+|00)(90|963|964|971|98))|0)?([1-9]\d*)$`
)

var (
	validate          = validator.New()
	mobileRegex       = regexp.MustCompile(mobileRegexString)
	refEnumValidation = reflect.TypeOf((*EnumValidation)(nil)).Elem()
	_                 = validate.RegisterValidation("enum", ValidateEnum)
)

func init() {
	_ = validate.RegisterValidation("mobile", ValidateMobile)
}

type EnumValidation interface {
	InRange(v interface{}) bool
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if !fl.Field().Type().Implements(refEnumValidation) {
		return false
	}
	v := fl.Field().Interface()
	return v.(EnumValidation).InRange(v)
}

func ValidateMobile(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	return mobileRegex.MatchString(v)
}

func Struct(ctx context.Context, s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}
