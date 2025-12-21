package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Init() (*validator.Validate, ut.Translator) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()

	// Register JSON tag name for validation errors
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		fmt.Println("Validator Translation Registration Failed:", err)
	}

	return validate, trans
}
