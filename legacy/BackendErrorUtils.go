package legacy

import (
	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/config"
)

func SetBackendError(backendResponseList *apimodel.BackendResponseList, errorCode string, errorMsg string, serviceConfig config.Service, backendUrl string) {
	var backendResponseInfoArray apimodel.BackendResponseInfoArray
	backendResponseInfoArray.APIName = serviceConfig.Name
	backendResponseInfoArray.System = serviceConfig.System
	backendResponseInfoArray.URL = backendUrl
	backendResponseInfoArray.ErrorCode = errorCode
	backendResponseInfoArray.Message = errorMsg
	backendResponseList.BackendResponseInfoArray = append(backendResponseList.BackendResponseInfoArray, backendResponseInfoArray)
	backendResponseList.Size = len(backendResponseList.BackendResponseInfoArray)
}
