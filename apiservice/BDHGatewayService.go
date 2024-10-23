package apiservice

import (
	"fmt"
	"strconv"
	"strings"

	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/errormsg"
	"intbackend/bdhgateway/legacy"
)

func UsageConsumption(req apimodel.UsageConsumptionRequest, methodName string, uuid string) apimodel.UsageConsumptionResponse {

	var backendResponseList apimodel.BackendResponseList
	var resp apimodel.UsageConsumptionResponse
	var respUsageConsumption legacy.UsageConsumptionResp
	var mapResp []apimodel.GetRatedPerformanceIndicatorsList

	respUsageConsumption = legacy.UsageConsumption(req, methodName, uuid, &backendResponseList)
	resp.BackendResponseList = &backendResponseList
	if strings.Contains(resp.BackendResponseList.BackendResponseInfoArray[0].ErrorCode, "200") && len(respUsageConsumption.Bucket) > 0 {
		mapResp = mapUsageConsumptionResponse(respUsageConsumption)

		resp.GetRatedPerformanceIndicatorsList = mapResp

		resp.ErrorCode = errormsg.ERR_CD_SUCCESS
		resp.Message = errormsg.ERR_MSG_SUCCESS
	} else if strings.Contains(resp.BackendResponseList.BackendResponseInfoArray[0].ErrorCode, "200") {
		resp.ErrorCode = errormsg.ERR_CD_DNF
		resp.Message = errormsg.ERR_MSG_DNF
	} else {
		resp.ErrorCode = errormsg.ERR_CD_BACKEND_ERROR
		resp.Message = errormsg.GenBackendErrorMsg(resp.BackendResponseList.BackendResponseInfoArray[0].APIName, resp.BackendResponseList.BackendResponseInfoArray[0].Message)
	}

	return resp

}

func mapUsageConsumptionResponse(response legacy.UsageConsumptionResp) []apimodel.GetRatedPerformanceIndicatorsList {
	var respRatedPerformance []apimodel.GetRatedPerformanceIndicatorsList

	size := len(response.Bucket)

	for i := 0; i < size; i++ {
		var attr []apimodel.AttributeList
		var resp apimodel.GetRatedPerformanceIndicatorsList

		fmt.Printf("i : "+"%d", i)
		fmt.Println(response.Bucket[i].BucketSpecification.Name)
		var extractName apimodel.AttributeList
		extractName.AttributeName = "Extract name"
		extractName.AttributeValue = response.Bucket[i].BucketSpecification.Name
		attr = append(attr, extractName)

		var customerId apimodel.AttributeList
		customerId.AttributeName = "Customer ID"
		customerId.AttributeValue = response.RelatedParty.ID
		attr = append(attr, customerId)

		if len(response.Bucket[i].BucketCounter) > 0 {
			var accumulatedCharge apimodel.AttributeList
			accumulatedCharge.AttributeName = "Accumulated charge"
			accumulatedCharge.AttributeValue = fmt.Sprintf("%.f", response.Bucket[i].BucketCounter[0].Value)
			attr = append(attr, accumulatedCharge)
		}

		if len(response.Bucket[i].BucketBalance) > 0 {
			var uom apimodel.AttributeList
			uom.AttributeName = "UOM"
			uom.AttributeValue = response.Bucket[i].BucketBalance[0].Unit
			attr = append(attr, uom)
		}

		for j := 0; j < len(response.Bucket[i].BucketCharacteristic); j++ {

			if "lastCallDate" == response.Bucket[i].BucketCharacteristic[j].Name {

				var lastCallDate apimodel.AttributeList
				lastCallDate.AttributeName = "Last call date"
				lastCallDate.AttributeValue = response.Bucket[i].BucketCharacteristic[j].Value
				attr = append(attr, lastCallDate)

			} else if "rerateInd" == response.Bucket[i].BucketCharacteristic[j].Name {

				var rerateInd apimodel.AttributeList
				rerateInd.AttributeName = "Rerate Ind"
				rerateInd.AttributeValue = response.Bucket[i].BucketCharacteristic[j].Value
				attr = append(attr, rerateInd)
			} else if "AccumulatedQuantity" == response.Bucket[i].BucketCharacteristic[j].Name {

				var accumulatedQuantity apimodel.AttributeList
				accumulatedQuantity.AttributeName = "Accumulated quantity"
				accumulatedQuantity.AttributeValue = response.Bucket[i].BucketCharacteristic[j].Value
				attr = append(attr, accumulatedQuantity)
			} else if "serviceType" == response.Bucket[i].BucketCharacteristic[j].Name {

				var serviceType apimodel.AttributeList
				serviceType.AttributeName = "Service Type"
				serviceType.AttributeValue = response.Bucket[i].BucketCharacteristic[j].Value
				attr = append(attr, serviceType)
			}

		}
		resp.AttributeList = attr
		respRatedPerformance = append(respRatedPerformance, resp)
	}

	return respRatedPerformance

}

func GetUsage(req apimodel.GetUsageRequest, methodName string, uuid string) apimodel.GetUsageResponse {

	var backendResponseList apimodel.BackendResponseList
	var resp apimodel.GetUsageResponse
	var mapResp apimodel.GetRatedEventsListInfoResponse

	respGetUsageist := legacy.GetUsage(req, methodName, uuid, &backendResponseList)
	resp.BackendResponseList = &backendResponseList

	var ratedEventsInfo apimodel.RatedEventsInfoList
	ratedEventsInfo.ErrorMessageInfo.ErrorCode = backendResponseList.BackendResponseInfoArray[0].ErrorCode
	ratedEventsInfo.ErrorMessageInfo.ErrorMessage = backendResponseList.BackendResponseInfoArray[0].Message

	if len(respGetUsageist) > 0 {

		for _, respGetUsage := range respGetUsageist {

			if len(respGetUsage.UsageCharacteristic) > 0 {

				usageCharMap := convertToMapUsageChar(respGetUsage.UsageCharacteristic)

				var ratedInfo apimodel.RatedEventsInfo
				ratedInfo.DisplayOnBill = getValMapUsageChar("displayOnBill", usageCharMap)
				ratedInfo.DiscountAmt = getValMapUsageChar("discountAmount", usageCharMap)
				ratedInfo.CallingNumber = getValMapUsageChar("callingNumber", usageCharMap)
				ratedInfo.CallFwdInd = getValMapUsageChar("callForwardingInd", usageCharMap)
				ratedInfo.Imsi = getValMapUsageChar("imsiResource", usageCharMap)
				ratedInfo.RoundedUnits = getValMapUsageChar("roundedUnits", usageCharMap)
				ratedInfo.OperatorName = getValMapUsageChar("operatorName", usageCharMap)
				ratedInfo.RoamInd = getValMapUsageChar("roamInd", usageCharMap)
				ratedInfo.Duration = getValMapUsageChar("duration", usageCharMap)
				ratedInfo.EventState = respGetUsage.EventState
				ratedInfo.CustomerId = getValMapUsageChar("targetCustomerId", usageCharMap)
				ratedInfo.PabxOrigSubscriber = getValMapUsageChar("pabxOrigSubscriber", usageCharMap)
				ratedInfo.StartTime = respGetUsage.Date
				ratedInfo.CalledNumber = getValMapUsageChar("calledNumber", usageCharMap)
				ratedInfo.Apn = getValMapUsageChar("apn", usageCharMap)
				ratedInfo.VventId = respGetUsage.ID
				ratedInfo.Period = getValMapUsageChar("period", usageCharMap)
				ratedInfo.ChargeAmt = getValMapUsageChar("taxExcludedRatingAmount", usageCharMap)
				ratedInfo.PaymentCategory = getValMapUsageChar("paymentCategory", usageCharMap)
				ratedInfo.RoundedUnitsUom = getValMapUsageChar("roundedUnitsUom", usageCharMap)
				ratedInfo.EventType = respGetUsage.EventTypeName
				ratedInfo.DestinationCountry = getValMapUsageChar("destinationCountry", usageCharMap)
				ratedInfo.AdditionalCharge = getValMapUsageChar("additionalCharge", usageCharMap)
				ratedInfo.EventStateReasonCode = getValMapUsageChar("eventStateReasonCode", usageCharMap)
				ratedInfo.BillingArrangement = respGetUsage.RelatedParty[0].ID
				ratedInfo.DirectionCode = getValMapUsageChar("directionCode", usageCharMap)
				ratedInfo.DestinationCode = getValMapUsageChar("destinationCode", usageCharMap)
				ratedInfo.EventTypeId = respGetUsage.EventTypeID
				ratedInfo.GuidingResourceId = respGetUsage.Resource[0].Value
				ratedInfo.OfferCode = getValMapUsageChar("offerCode", usageCharMap)
				ratedInfo.EventTypeName = getValMapUsageChar("eventTypeName", usageCharMap)
				ratedInfo.PayChannel = respGetUsage.PartyAccount.ID
				ratedInfo.EndTime = getValMapUsageChar("endTime", usageCharMap)
				ratedInfo.StartDate = respGetUsage.Date
				ratedInfo.FreeAmt = fmt.Sprintf("%f", respGetUsage.RatedProductUsage[0].FreeChargeAmount)

				ratedEventsInfo.RatedEventsInfo = append(ratedEventsInfo.RatedEventsInfo, ratedInfo)
				ratedEventsInfo.TotalSize = strconv.Itoa(len(ratedEventsInfo.RatedEventsInfo))
				mapResp.RatedEventsInfoList = &ratedEventsInfo
			}
		}

		resp.GetRatedEventsListInfoResponse = &mapResp
		resp.ErrorCode = errormsg.ERR_CD_SUCCESS
		resp.Message = errormsg.ERR_MSG_SUCCESS
	} else {
		resp.ErrorCode = errormsg.ERR_CD_DNF
		resp.Message = errormsg.ERR_MSG_DNF
	}

	return resp

}

func convertToMapUsageChar(data []legacy.UsageCharacteristic) map[string]string {
	result := make(map[string]string)
	for _, d := range data {
		result[d.Name] = d.Value
	}
	return result
}

func getValMapUsageChar(key string, mapData map[string]string) string {
	var result = ""
	if val, ok := mapData[key]; ok {
		result = val
	}
	return result
}

func GetCreditLimitWithPerformanceIndicators(req apimodel.GetCreditLimitWithPerformanceIndicatorsRequest, methodName string, uuid string) apimodel.GetCreditLimitWithPerformanceIndicatorsResponse {

	var backendResponseList apimodel.BackendResponseList
	var resp apimodel.GetCreditLimitWithPerformanceIndicatorsResponse
	var respGetCreditLimitInfo apimodel.GetCreditLimitWithPerformanceIndicatorsListInfo
	var getGetCreditLimitList apimodel.GetCreditLimitWithPerformanceIndicatorsList

	respGetCreditLimitInfo = legacy.GetCreditLimitWithPerformanceIndicators(req, methodName, uuid, &backendResponseList)
	resp.BackendResponseList = &backendResponseList

	if len(respGetCreditLimitInfo.Results) > 0 {

		for _, results := range respGetCreditLimitInfo.Results {

			if results.Message == "Success" {
				getGetCreditLimitList.Results = respGetCreditLimitInfo.Results

				resp.GetCreditLimitWithPerformanceIndicatorsList.Results = respGetCreditLimitInfo.Results
				resp.ErrorCode = errormsg.ERR_CD_SUCCESS
				resp.Message = errormsg.ERR_MSG_SUCCESS

			} else {
				resp.ErrorCode = errormsg.ERR_CD_BACKEND_ERROR
				resp.Message = errormsg.GenBackendErrorMsg(resp.BackendResponseList.BackendResponseInfoArray[0].APIName, resp.BackendResponseList.BackendResponseInfoArray[0].Message)
			}
		}
	} else {
		resp.ErrorCode = errormsg.ERR_CD_BACKEND_ERROR
		resp.Message = errormsg.GenBackendErrorMsg(resp.BackendResponseList.BackendResponseInfoArray[0].APIName, resp.BackendResponseList.BackendResponseInfoArray[0].Message)
	}

	// if strings.Contains(resp.BackendResponseList.BackendResponseInfoArray[0].ErrorCode, "200") && len(respGetCreditLimitInfo.Results) > 0 {
	// 	getGetCreditLimitList.Results = respGetCreditLimitInfo.Results

	// 	resp.ErrorCode = errormsg.ERR_CD_SUCCESS
	// 	resp.Message = errormsg.ERR_MSG_SUCCESS
	// } else if strings.Contains(resp.BackendResponseList.BackendResponseInfoArray[0].ErrorCode, "200") {
	// 	resp.ErrorCode = errormsg.ERR_CD_DNF
	// 	resp.Message = errormsg.ERR_MSG_DNF
	// } else {
	// 	resp.ErrorCode = errormsg.ERR_CD_BACKEND_ERROR
	// 	resp.Message = errormsg.GenBackendErrorMsg(resp.BackendResponseList.BackendResponseInfoArray[0].APIName, resp.BackendResponseList.BackendResponseInfoArray[0].Message)
	// }

	// resp.GetCreditLimitWithPerformanceIndicatorsList = &getGetCreditLimitList

	return resp

}
