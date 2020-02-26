package validator

import (
	"reflect"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	jaTranslations "gopkg.in/go-playground/validator.v9/translations/ja"
)

type ValidationTranslator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidationTranslator(locale string) *ValidationTranslator {
	ja := ja.New()
	uni := ut.New(ja, ja)
	trans, _ := uni.GetTranslator("ja")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("trans")
	})
	jaTranslations.RegisterDefaultTranslations(validate, trans)

	return &ValidationTranslator{
		validate,
		trans,
	}
}

func (this *ValidationTranslator) Validate(obj interface{}) map[string]string {
	err := this.Validator.Struct(obj)
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		return errs.Translate(this.Translator)
	}

	return map[string]string{}
}
