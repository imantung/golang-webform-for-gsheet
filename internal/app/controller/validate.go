package controller

import (
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	translator, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, translator)
}

func valdnErrMsg(err error) string {
	errs := err.(validator.ValidationErrors)
	var msgs []string
	for _, val := range errs.Translate(translator) {
		msgs = append(msgs, val)
	}
	return strings.Join(msgs, ",")
}

// NOTE: put custom validator here
