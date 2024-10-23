package legacy

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/config"
	"intbackend/bdhgateway/errormsg"
	"intbackend/bdhgateway/locallogging"
)

func GetCreditLimitWithPerformanceIndicators(request apimodel.GetCreditLimitWithPerformanceIndicatorsRequest, methodName string, uuid string,
	backendResponseList *apimodel.BackendResponseList) apimodel.GetCreditLimitWithPerformanceIndicatorsListInfo {

	appConfig, _ := config.LoadConfig()
	serviceConfig := config.GetService("getCreditLimitWithPerformanceIndicators", appConfig)
	var responseObject apimodel.GetCreditLimitWithPerformanceIndicatorsListInfo
	var accessToken = serviceConfig.TokenAccess
	bearer := "Bearer " + accessToken

	// Write log Legacy
	log := locallogging.LocalLoggingLegacy{}
	log.SetLegacyInputLoggerStart(uuid, appConfig.Application.Profile, methodName, time.Now())

	reqB, _ := json.Marshal(request)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, er := http.NewRequest(http.MethodPost, serviceConfig.Endpoint, bytes.NewBuffer(reqB))
	if er != nil {
		SetBackendError(backendResponseList, errormsg.ERR_CD_HTTP_500, errormsg.ERR_MSG_HTTP_500, serviceConfig, BACKEND_RESP_URL)
		return responseObject
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	resp, er := client.Do(req)
	if er != nil {
		SetBackendError(backendResponseList, errormsg.ERR_CD_HTTP_500, errormsg.ERR_MSG_HTTP_500, serviceConfig, BACKEND_RESP_URL)
		return responseObject
	}

	bodyBytes, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		var httpCode string
		var httpMsg string
		if resp.StatusCode == 503 {
			httpCode = errormsg.ERR_CD_HTTP_503
			httpMsg = errormsg.ERR_MSG_HTTP_503
		} else if resp.StatusCode == 404 {
			httpCode = errormsg.ERR_CD_HTTP_404
			httpMsg = errormsg.ERR_MSG_HTTP_404
		} else {
			httpCode = errormsg.ERR_CD_HTTP_500
			httpMsg = errormsg.ERR_MSG_HTTP_500
		}
		SetBackendError(backendResponseList, httpCode, httpMsg, serviceConfig, BACKEND_RESP_URL)
		return responseObject
	}

	er = json.Unmarshal(bodyBytes, &responseObject)
	if er != nil {
		var httpCode string
		var httpMsg string
		if resp.StatusCode == 503 {
			httpCode = errormsg.ERR_CD_HTTP_503
			httpMsg = errormsg.ERR_MSG_HTTP_503
		} else if resp.StatusCode == 404 {
			httpCode = errormsg.ERR_CD_HTTP_404
			httpMsg = errormsg.ERR_MSG_HTTP_404
		} else {
			httpCode = errormsg.ERR_CD_HTTP_500
			httpMsg = errormsg.ERR_MSG_HTTP_500
		}
		SetBackendError(backendResponseList, httpCode, httpMsg, serviceConfig, BACKEND_RESP_URL)
		return responseObject
	}

	defer resp.Body.Close()

	errorCode := strconv.Itoa(resp.StatusCode)
	message := ""

	if resp.StatusCode == 200 {
		fmt.Println("Successfully !!!")

		if len(responseObject.Results) > 0 {
			for _, results := range responseObject.Results {
				errorCode = strconv.Itoa(results.Code)
				message = results.Message
			}
		} else {
			message = errormsg.ERR_MSG_SUCCESS
		}
	} else {
		fmt.Println("Error !!!")
		message = resp.Status
	}

	log.SetLegacyInputLoggerEnd(string(reqB), string(bodyBytes), errorCode, message, serviceConfig.Endpoint, time.Now())
	log.WriteLogLegacy()

	SetBackendError(backendResponseList, errorCode, message, serviceConfig, BACKEND_RESP_URL)

	return responseObject
}
