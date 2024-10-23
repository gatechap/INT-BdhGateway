package apiservice

import (
	"strings"

	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/errormsg"
)

func ValidRequiredStringArray(stringArrayFieldName string, stringFieldName string, stringArray []string) (string, string) {
	if len(stringArray) <= 0 {
		return errormsg.ERR_CD_REQUIRED_FIELD, errormsg.GenRequiredFieldMsg(stringArrayFieldName)
	}

	for i := 0; i < len(stringArray); i++ {
		stringVal := stringArray[i]

		if stringVal == "" {
			return errormsg.ERR_CD_REQUIRED_FIELD, errormsg.GenRequiredFieldMsg(stringFieldName)
		}
	}

	return errormsg.ERR_CD_SUCCESS, errormsg.ERR_MSG_SUCCESS
}

func ValidRequiredKeyValue(keyValueInfoFieldName string, keyFieldName string, valueFieldName string, keyValues []apimodel.KeyValueInfo) (string, string) {
	if len(keyValues) <= 0 {
		return errormsg.ERR_CD_REQUIRED_FIELD, errormsg.GenRequiredFieldMsg(keyValueInfoFieldName)
	}

	for i := 0; i < len(keyValues); i++ {
		keyValue := keyValues[i]

		if keyValue.Key == "" {
			return errormsg.ERR_CD_REQUIRED_FIELD, errormsg.GenRequiredFieldMsg(keyFieldName)
		} else if keyValue.Value == "" {
			return errormsg.ERR_CD_REQUIRED_FIELD, errormsg.GenRequiredFieldMsg(valueFieldName)
		}
	}

	return errormsg.ERR_CD_SUCCESS, errormsg.ERR_MSG_SUCCESS
}

func ValidValueStringArray(value string, validFieldName string, validValue string) string {

	var isValid bool
	validArray := strings.Split(validValue, ",")

	for i := 0; i < len(validArray); i++ {
		stringVal := validArray[i]
		if stringVal == value {
			isValid = true
		}
	}

	if !isValid {
		return errormsg.GenParamInvalidMsg(validFieldName, validValue)
	} else {
		return "valid"
	}

}
