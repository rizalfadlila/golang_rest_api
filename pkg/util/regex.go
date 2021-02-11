package util

import "regexp"

// IsValidNumber :nodoc:
func IsValidNumber(param interface{}) bool {
	isValid, err := regexp.MatchString(`^[0-9]`, param.(string))
	if err != nil {
		return false
	}
	return isValid
}
