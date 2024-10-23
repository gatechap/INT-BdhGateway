package legacy

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"intbackend/bdhgateway/apimodel"
	"intbackend/bdhgateway/config"
	"intbackend/bdhgateway/errormsg"
	"intbackend/bdhgateway/intutilities"
	"intbackend/bdhgateway/locallogging"
)

func GetUsage(request apimodel.GetUsageRequest, methodName string, uuid string, backendResponseList *apimodel.BackendResponseList) []GetUsageResp {

	appConfig, _ := config.LoadConfig()
	serviceConfig := config.GetService("usage", appConfig)
	var responseObject []GetUsageResp
	var accessToken = serviceConfig.TokenAccess
	bearer := "Bearer " + accessToken

	// Write log Legacy
	log := locallogging.LocalLoggingLegacy{}
	log.SetLegacyInputLoggerStart(uuid, appConfig.Application.Profile, methodName, time.Now())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	reqB, _ := json.Marshal(request)
	req, er := http.NewRequest(http.MethodGet, serviceConfig.Endpoint, nil)
	pageSize := strconv.Itoa(int(request.PageSize))
	q := req.URL.Query()
	q.Add("billingCycleSpecificationId", request.CycleCode)
	q.Add("billingCycleInstance", request.CycleInstance)
	q.Add("customerId", request.CustomerId)
	q.Add("billingCycleYear", request.CycleYear)
	q.Add("filters", "Usage.eventTypeId=="+request.Ban)
	q.Add("limit", intutilities.GetLimit(pageSize))
	q.Add("offset", intutilities.GetOffset(int(request.PageNumber), int(request.PageSize)))

	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

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
	contentLength := int(resp.ContentLength)

	if strings.Contains(resp.Status, "200") && contentLength == 0 {
		SetBackendError(backendResponseList, resp.Status, errormsg.ERR_MSG_DNF, serviceConfig, BACKEND_RESP_URL)
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
		message = errormsg.ERR_MSG_SUCCESS
	} else {
		message = resp.Status
	}

	log.SetLegacyInputLoggerEnd(string(reqB), string(bodyBytes), errorCode, message, serviceConfig.Endpoint, time.Now())
	log.WriteLogLegacy()

	SetBackendError(backendResponseList, errorCode, message, serviceConfig, BACKEND_RESP_URL)

	return responseObject
}

type GetUsageResp struct {
	ID                  string                `json:"id"`
	Href                string                `json:"href"`
	Date                string                `json:"date"`
	EventTypeID         string                `json:"eventTypeId"`
	EventTypeName       string                `json:"eventTypeName"`
	EventState          string                `json:"eventState"`
	UsageCharacteristic []UsageCharacteristic `json:"usageCharacteristic"`
	RelatedParty        []UMRelatedParty      `json:"relatedParty"`
	Resource            []Resource            `json:"resource"`
	PartyAccount        UMPartyAccount        `json:"partyAccount"`
	RatedProductUsage   []RatedProductUsage   `json:"ratedProductUsage"`
}

type UsageCharacteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UMRelatedParty struct {
	ID string `json:"id"`
}

type Resource struct {
	Value string `json:"value"`
}

type UMPartyAccount struct {
	ID string `json:"id"`
}

type RatedProductUsage struct {
	TaxExcludedRatingAmount float64 `json:"taxExcludedRatingAmount"`
	FreeChargeAmount        float64 `json:"freeChargeAmount"`
}
