package validator

import (
	"reflect"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	val "gopkg.in/go-playground/validator.v9"
	jaTranslations "gopkg.in/go-playground/validator.v9/translations/ja"
)

type ValidationTranslator struct {
	validator  *val.Validate
	translator ut.Translator
}

func NewValidationTranslator() *ValidationTranslator {
	ja := ja.New()
	uni := ut.New(ja, ja)
	trans, _ := uni.GetTranslator("ja")
	validate := val.New()
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
	err := this.validator.Struct(obj)
	errs, ok := err.(val.ValidationErrors)
	if ok {
		return errs.Translate(this.translator)
	}

	return map[string]string{}
}
