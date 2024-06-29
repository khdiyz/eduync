package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Custom validation function for Uzbek phone numbers starting with +998
func uzbPhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	var uzbPhonePattern = `^\+998\d{9}$`
	re := regexp.MustCompile(uzbPhonePattern)
	return re.MatchString(phone)
}
