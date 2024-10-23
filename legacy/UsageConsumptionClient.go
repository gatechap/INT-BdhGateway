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

func UsageConsumption(request apimodel.UsageConsumptionRequest, methodName string, uuid string,
	backendResponseList *apimodel.BackendResponseList) UsageConsumptionResp {

	appConfig, _ := config.LoadConfig()
	serviceConfig := config.GetService("usageConsumption", appConfig)
	var responseObject UsageConsumptionResp

	accessToken := serviceConfig.TokenAccess

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

	pageNum := intutilities.StringToIngeter(request.PageNumber)
	pageSize := intutilities.StringToIngeter(request.PageSize)
	q := req.URL.Query()
	q.Add("billingCycleYear", request.CycleYear)
	q.Add("billingCycleInstance", request.CycleMonth)
	q.Add("billingCycleSpecificationId", request.CycleCode)
	q.Add("customerId", request.CustomerId)
	q.Add("mainSubscriptionProductId", request.AgreementId)
	q.Add("bucketRole", getBucketRole(request.ExternalStructure))
	q.Add("limit", intutilities.GetLimit(request.PageSize))
	q.Add("offset", intutilities.GetOffset(pageNum, pageSize))
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
		fmt.Println("Successfully !!!")
		message = errormsg.ERR_MSG_SUCCESS

	} else {
		fmt.Println("Error !!!")
		message = resp.Status
	}

	log.SetLegacyInputLoggerEnd(string(reqB), string(bodyBytes), errorCode, message, serviceConfig.Endpoint, time.Now())
	log.WriteLogLegacy()

	SetBackendError(backendResponseList, errorCode, message, serviceConfig, BACKEND_RESP_URL)

	return responseObject

}

func getBucketRole(external string) string {
	var role string

	if "True PI query accumulation list per ST and UOM" == external {
		role = "AdditionalRateAccumulation"

	} else if "TRUE PI query allowance list" == external {
		role = "ConsumptionAllowance"

	} else if "TRUE PI query rate list" == external {
		role = "MonetaryConsumptionAllowance"

	} else if "True PI query accumulation list per subscriber" == external {
		role = "RateAccumulation"

	} else if "True PI query credit limit" == external {
		role = "ConsumptionBudgetControl"
	}

	return role
}

type UsageConsumptionResp struct {
	Bucket        []Bucket     `json:"bucket"`
	EffectiveDate string       `json:"effectiveDate"`
	Name          string       `json:"name"`
	Role          string       `json:"role"`
	RelatedParty  RelatedParty `json:"relatedParty"`
}

type Bucket struct {
	BucketBalance             []BucketBalance        `json:"bucketBalance"`
	BucketCharacteristic      []BucketCharacteristic `json:"bucketCharacteristic"`
	BucketCounter             []BucketCounter        `json:"bucketCounter"`
	BucketSpecification       BucketSpecification    `json:"bucketSpecification"`
	ID                        string                 `json:"id"`
	MainSubscriptionProductID string                 `json:"mainSubscriptionProductId"`
	PartyAccount              PartyAccount           `json:"partyAccount"`
	Product                   Product                `json:"product"`
	Description               string                 `json:"description"`
}

type BucketBalance struct {
	RemainingValue int    `json:"remainingValue"`
	Unit           string `json:"unit"`
}

type BucketCharacteristic struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BucketCounter struct {
	Value float32 `json:"value"`
}

type BucketSpecification struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type PartyAccount struct {
	ID string `json:"id"`
}

type Product struct {
	PublicIdentifier string          `json:"publicIdentifier"`
	ProductOffering  ProductOffering `json:"productOffering"`
}

type ProductOffering struct {
	Name string `json:"name"`
}

type RelatedParty struct {
	ID string `json:"id"`
}
