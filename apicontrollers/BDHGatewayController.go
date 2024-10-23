package apicontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/apiservice"
	"intbackend/bdhgateway/errormsg"
	"intbackend/bdhgateway/httphandler"
	"intbackend/bdhgateway/intutilities"
	"intbackend/bdhgateway/locallogging"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// bdhgateway godoc
// @Summary bdhgateway
// @Tags bdh-gateway-controller
// @Accept  json
// @Produce  json
// @param X-Channel header string false "X-Channel"
// @param X-GatewayType header string false "X-GatewayType"
// @param X-LegacyUsername header string false "X-LegacyUsername"
// @param X-Username header string true "X-Username"
// @Param Request body apimodel.UsageConsumptionRequest true "requestInfo"
// @Success 200 {array} apimodel.UsageConsumptionResponse
// @Failure 400,404 {object} apimodel.HttpErrorResponse
// @Failure 500 {object} apimodel.HttpErrorResponse
// @Failure default {object} apimodel.HttpErrorResponse
// @Router /bdhgateway/usageconsumption [post]
func UsageConsumption(c *gin.Context) {

	const CURRENT_FUNCTION = "usageconsumption"
	const METHOD_NAME = "/usageconsumption"

	startTime := time.Now()

	var isError bool
	var errCode string
	var errMsg string
	var resultStatus string
	var response apimodel.UsageConsumptionResponse
	var httpResponse int

	// initial error hanler info
	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: AppConfig.Application.Name,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        PATH_CONTROLLER_BDH_GATEWAY,
		ErrorFunction:    CURRENT_FUNCTION,
	}

	// generate uuid
	struuid := uuid.New().String()

	// Validate request (bad request)
	var request apimodel.UsageConsumptionRequest
	var bReq []byte
	var err error
	err = c.ShouldBindJSON(&request)

	// check bad request
	if err != nil {
		fmt.Println(strings.ToUpper(err.Error()))
		isError = true
		isNullParam := errormsg.IsNotNullParam(err.Error(), USAGECONSUMPTION_REQ_FIELD)
		if isNullParam != "" {
			errCode = errormsg.ERR_CD_REQUIRED_FIELD
			errMsg = errormsg.GenRequiredFieldMsg(isNullParam)
		} else {
			httpResponse = http.StatusBadRequest
			httpErrResponse := apimodel.HttpErrorResponse{
				Timestamp: intutilities.GetCurrentISO8601(),
				Status:    httpResponse,
				Error:     http.StatusText(httpResponse),
				Message:   "",
				Path:      c.Request.URL.String(),
			}

			c.JSON(httpResponse, httpErrResponse)
			return

		}
		errHanlerInfo.Error = err

	} else {
		bReq, err = json.Marshal(&request)
		if err != nil {
			isError = true
			errCode = errormsg.ERR_CD_INTERNAL_FAILURE
			errMsg = errormsg.GenInternalFailureMsg(err.Error())
			errHanlerInfo.Error = err
		}
	}

	// httpheader
	httpHeader, requestparam := httphandler.GetHttpHeaderInfo(c, &struuid)
	username := intutilities.CheckUsername(httpHeader.XChannel, requestparam.Username, httpHeader.XUsername)
	if !isError && username == "" {
		isError = true
		errCode = errormsg.ERR_CD_REQUIRED_FIELD
		errMsg = errormsg.GenRequiredFieldMsg("Username")
		errHanlerInfo.Error = nil
	}

	// Write log request
	log := locallogging.LocalLogging{}
	log.SetRequestInputLogger(httpHeader, requestparam, AppConfig, bReq, startTime)
	log.WriteLogRequest()

	if !isError {
		response = apiservice.UsageConsumption(request, METHOD_NAME, struuid)
	}

	if isError {
		response = apimodel.UsageConsumptionResponse{
			Uuid:      struuid,
			ErrorCode: errCode,
			Message:   errMsg,
		}
		httpResponse = http.StatusOK
		resultStatus = "F"

		log.SetErrorInputLogger(errCode, errMsg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.WriteLogError()
	} else {

		if response.ErrorCode == "INT10001" { // DNF
			errCode = errormsg.ERR_CD_DNF
			errMsg = errormsg.ERR_MSG_DNF
			resultStatus = "NA"

		} else {
			errCode = errormsg.ERR_CD_SUCCESS
			errMsg = errormsg.ERR_MSG_SUCCESS
			resultStatus = "S"

		}

		response.Uuid = struuid
		httpResponse = http.StatusOK
	}

	log.SetResponseInputLogger(errCode, errMsg, resultStatus, "", "", time.Now())
	log.WriteLogResponse()
	c.JSON(httpResponse, response)
}

// bdhgateway godoc
// @Summary bdhgateway
// @Tags bdh-gateway-controller
// @Accept  json
// @Produce  json
// @param X-Channel header string false "X-Channel"
// @param X-GatewayType header string false "X-GatewayType"
// @param X-LegacyUsername header string false "X-LegacyUsername"
// @param X-Username header string true "X-Username"
// @Param Request body apimodel.GetUsageRequest true "requestInfo"
// @Success 200 {array} apimodel.GetUsageResponse
// @Failure 400,404 {object} apimodel.HttpErrorResponse
// @Failure 500 {object} apimodel.HttpErrorResponse
// @Failure default {object} apimodel.HttpErrorResponse
// @Router /bdhgateway/getusage [post]
func GetUsage(c *gin.Context) {

	const CURRENT_FUNCTION = "getusage"
	const METHOD_NAME = "/getusage"

	startTime := time.Now()

	var isError bool
	var errCode string
	var errMsg string
	var resultStatus string
	var response apimodel.GetUsageResponse
	var httpResponse int

	// initial error hanler info
	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: AppConfig.Application.Name,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        PATH_CONTROLLER_BDH_GATEWAY,
		ErrorFunction:    CURRENT_FUNCTION,
	}

	// generate uuid
	struuid := uuid.New().String()

	// Validate request (bad request)
	var request apimodel.GetUsageRequest
	var bReq []byte
	var err error
	err = c.ShouldBindJSON(&request)

	// check bad request
	if err != nil {
		fmt.Println(strings.ToUpper(err.Error()))
		isError = true
		isNullParam := errormsg.IsNotNullParam(err.Error(), GetUsage_REQ_FIELD)

		if isNullParam != "" {
			errCode = errormsg.ERR_CD_REQUIRED_FIELD
			errMsg = errormsg.GenRequiredFieldMsg(isNullParam)
		} else {
			httpResponse = http.StatusBadRequest
			httpErrResponse := apimodel.HttpErrorResponse{
				Timestamp: intutilities.GetCurrentISO8601(),
				Status:    httpResponse,
				Error:     http.StatusText(httpResponse),
				Message:   "",
				Path:      c.Request.URL.String(),
			}

			c.JSON(httpResponse, httpErrResponse)
			return

		}
		errHanlerInfo.Error = err

	} else {
		bReq, err = json.Marshal(&request)
		if err != nil {
			isError = true
			errCode = errormsg.ERR_CD_INTERNAL_FAILURE
			errMsg = errormsg.GenInternalFailureMsg(err.Error())
			errHanlerInfo.Error = err
		}
	}

	// httpheader
	httpHeader, requestparam := httphandler.GetHttpHeaderInfo(c, &struuid)
	username := intutilities.CheckUsername(httpHeader.XChannel, requestparam.Username, httpHeader.XUsername)

	if !isError && username == "" {
		isError = true
		errCode = errormsg.ERR_CD_REQUIRED_FIELD
		errMsg = errormsg.GenRequiredFieldMsg("Username")
		errHanlerInfo.Error = nil
	}

	// Write log request
	log := locallogging.LocalLogging{}
	log.SetRequestInputLogger(httpHeader, requestparam, AppConfig, bReq, startTime)
	log.WriteLogRequest()

	if !isError {
		response = apiservice.GetUsage(request, METHOD_NAME, struuid)
	}

	if isError {
		response = apimodel.GetUsageResponse{
			Uuid:      struuid,
			ErrorCode: errCode,
			Message:   errMsg,
		}

		httpResponse = http.StatusOK
		resultStatus = "F"

		log.SetErrorInputLogger(errCode, errMsg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.WriteLogError()

	} else {
		if response.ErrorCode == "INT10001" { // DNF
			errCode = errormsg.ERR_CD_DNF
			errMsg = errormsg.ERR_MSG_DNF
			resultStatus = "NA"

		} else {
			errCode = errormsg.ERR_CD_SUCCESS
			errMsg = errormsg.ERR_MSG_SUCCESS
			resultStatus = "S"
		}

		response.Uuid = struuid
		httpResponse = http.StatusOK
	}

	log.SetResponseInputLogger(errCode, errMsg, resultStatus, "", "", time.Now())
	log.WriteLogResponse()
	c.JSON(httpResponse, response)
}

// bdhgateway godoc
// @Summary bdhgateway
// @Tags bdh-gateway-controller
// @Accept  json
// @Produce  json
// @param X-Channel header string false "X-Channel"
// @param X-GatewayType header string false "X-GatewayType"
// @param X-LegacyUsername header string false "X-LegacyUsername"
// @param X-Username header string true "X-Username"
// @Param Request body apimodel.GetCreditLimitWithPerformanceIndicatorsRequest true "requestInfo"
// @Success 200 {array} apimodel.GetCreditLimitWithPerformanceIndicatorsResponse
// @Failure 400,404 {object} apimodel.HttpErrorResponse
// @Failure 500 {object} apimodel.HttpErrorResponse
// @Failure default {object} apimodel.HttpErrorResponse
// @Router /bdhgateway/getcreditlimitwithperformanceindicators [post]
func GetCreditLimitWithPerformanceIndicators(c *gin.Context) {

	const CURRENT_FUNCTION = "getcreditlimitwithperformanceindicators"
	const METHOD_NAME = "/getcreditlimitwithperformanceindicators"

	startTime := time.Now()

	var isError bool
	var errCode string
	var errMsg string
	var resultStatus string
	var response apimodel.GetCreditLimitWithPerformanceIndicatorsResponse
	var httpResponse int

	// initial error hanler info
	errHanlerInfo := errormsg.ErrorHandlerInfo{
		ErrorApplication: AppConfig.Application.Name,
		ErrorModule:      CURRENT_MODULE,
		ErrorFile:        PATH_CONTROLLER_BDH_GATEWAY,
		ErrorFunction:    CURRENT_FUNCTION,
	}

	// generate uuid
	struuid := uuid.New().String()

	// Validate request (bad request)
	var request apimodel.GetCreditLimitWithPerformanceIndicatorsRequest
	var bReq []byte
	var err error
	err = c.ShouldBindJSON(&request)

	// check bad request
	if err = validator.New().Struct(&request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			//fmt.Println("==========StructField : ", err.StructField())

			field, _ := reflect.TypeOf(apimodel.GetCreditLimitWithPerformanceIndicatorsRequest{}).FieldByName(err.StructField())

			result, _ := field.Tag.Lookup("json")

			if err.Tag() == "required" {
				isError = true
				if err.Tag() == "required" {
					errCode = errormsg.ERR_CD_REQUIRED_FIELD
					errMsg = errormsg.GenRequiredFieldMsg(result)
				}
				errHanlerInfo.Error = err
				break
			}
		}
	}

	// check RequestList
	if !isError {
		if len(request.RequestList) <= 0 {
			isError = true
			errCode = errormsg.ERR_CD_REQUIRED_FIELD
			errMsg = errormsg.GenRequiredFieldMsg(GetCreditLimit_REQ_FIELD_RequestList)
			errHanlerInfo.Error = err
		} /*else {
			for i := 0; i < len(request.RequestList); i++ {
				if request.RequestList[i].AccountNo == "" {
					isError = true
					errMsg = errormsg.GenRequiredFieldMsg(GetCreditLimit_REQ_FIELD_RequestList)
				}

				if isError {
					errCode = errormsg.ERR_CD_REQUIRED_FIELD
					errHanlerInfo.Error = err
				}
			}
		} */
	}

	// httpheader
	httpHeader, requestparam := httphandler.GetHttpHeaderInfo(c, &struuid)
	username := intutilities.CheckUsername(httpHeader.XChannel, requestparam.Username, httpHeader.XUsername)
	if !isError && username == "" {
		isError = true
		errCode = errormsg.ERR_CD_REQUIRED_FIELD
		errMsg = errormsg.GenRequiredFieldMsg("Username")
		errHanlerInfo.Error = nil
	}

	// Write log request
	log := locallogging.LocalLogging{}
	log.SetRequestInputLogger(httpHeader, requestparam, AppConfig, bReq, startTime)
	log.WriteLogRequest()

	if !isError {
		response = apiservice.GetCreditLimitWithPerformanceIndicators(request, METHOD_NAME, struuid)
	}

	if isError {
		response = apimodel.GetCreditLimitWithPerformanceIndicatorsResponse{
			Uuid:      struuid,
			ErrorCode: errCode,
			Message:   errMsg,
		}
		httpResponse = http.StatusOK
		resultStatus = "F"

		log.SetErrorInputLogger(errCode, errMsg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.WriteLogError()

	} else {

		if response.ErrorCode == "INT10001" { // DNF
			errCode = errormsg.ERR_CD_DNF
			errMsg = errormsg.ERR_MSG_DNF
			resultStatus = "NA"

		} else {
			errCode = errormsg.ERR_CD_SUCCESS
			errMsg = errormsg.ERR_MSG_SUCCESS
			resultStatus = "S"

		}

		response.Uuid = struuid
		httpResponse = http.StatusOK
	}

	log.SetResponseInputLogger(errCode, errMsg, resultStatus, "", "", time.Now())
	log.WriteLogResponse()
	c.JSON(httpResponse, response)
}
