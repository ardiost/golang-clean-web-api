package validation

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)
var logger = logging.NewLogger(config.GetConfig())
// /^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$/gm
func IranianMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	res, err := regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		logger.Error(logging.Validation,logging.MobileValidation,err.Error(),nil)
		log.Print(err.Error())
	}
	return res
}
