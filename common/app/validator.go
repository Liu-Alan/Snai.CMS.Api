package app

import (
	"reflect"
	"strings"

	"Snai.CMS.Api/common/logging"
	"Snai.CMS.Api/common/message"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var G_Validate *validateCtx

type validateCtx struct {
	*validator.Validate
	trans ut.Translator
}

func InitValidator() {
	validate := validator.New()
	// 中文翻译器
	uniTrans := ut.New(zh.New())
	trans, _ := uniTrans.GetTranslator("zh")

	//通过自定义标签label来替换字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 注册翻译器到验证器
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logging.Error("registerDefaultTranslations fail: %s\n", err.Error())
	}

	G_Validate = &validateCtx{validate, trans}
}

func (vc *validateCtx) GetError(errs error) []string {
	var errStr []string
	for _, err := range errs.(validator.ValidationErrors) {
		if vc.trans != nil {
			errStr = append(errStr, strings.Replace(err.Translate(vc.trans), "Password", "密码", -1))
		} else {
			errStr = append(errStr, err.Field()+"验证不符合"+err.Tag())
		}
	}
	return errStr
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
	errs := G_Validate.Struct(params)
	if errs != nil {
		errStr := G_Validate.GetError(errs)
		return &message.Message{Code: message.ValidParamsError, Msg: message.GetMsg(message.ValidParamsError), Result: errStr}
	}

	return &message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
}
