package apiservice

import "intbackend/bdhgateway/apimodel"

func CheckDNFByKeyValues(keyValues []apimodel.KeyValueInfo) (bool, error) {
	if len(keyValues) > 0 {
		for i := 0; i < len(keyValues); i++ {
			keyValue := keyValues[i]
			if keyValue.Value != "" {
				return false, nil
			}
		}
	}
	return true, nil
}
