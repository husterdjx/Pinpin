package util

import "github.com/go-playground/validator/v10"

func ErrorHandler(err error, m map[string]string) (msg string) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return "输入参数错误"
	}
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErrs {
			switch val.Tag() {
			case "required":
				return m[val.Field()] + "不能为空哦"
			case "len":
				switch val.Field() {
				case "Tel":
					return m[val.Field()] + "必须为11位"
				}
			}
		}
	}
	return ""
}
