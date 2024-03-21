package app

import (
	"Snai.CMS.Api/common/logging"
	"Snai.CMS.Api/common/message"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var validate *validator.Validate
var trans ut.Translator

func InitValidator() {
	validate = validator.New()
	// 中文翻译器
	uniTrans := ut.New(zh.New())
	trans, _ = uniTrans.GetTranslator("zh")
	// 注册翻译器到验证器
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logging.Error("registerDefaultTranslations fail: %s\n", err.Error())
	}
}

func BindAndValid(c *gin.Context, params interface{}, bindType string) *message.Message {
	var err error

	//绑定
	if bindType == "json" {
		err = c.BindJSON(params)
	} else {
		err = c.Bind(params)
	}
	if err != nil {
		return &message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)}
	}

	//校验
	errs := validate.Struct(params)
	if errs != nil {
		verrs, ok := errs.(validator.ValidationErrors)
		if !ok {
			logging.Error("InvalidValidationError")
			return &message.Message{Code: message.ValidParamsError, Msg: message.GetMsg(message.ValidParamsError)}
		}
		var errStr []string
		for errv := range verrs.Translate(trans) {
			errStr = append(errStr, errv)
		}
		return &message.Message{Code: message.ValidParamsError, Msg: message.GetMsg(message.ValidParamsError), Result: errStr}
	}

	return &message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
}
