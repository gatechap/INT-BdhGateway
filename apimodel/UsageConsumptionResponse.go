package apimodel

type UsageConsumptionResponse struct {
	Uuid                              string                              `json:"uuid,omitempty"`
	ErrorCode                         string                              `json:"errorCode,omitempty"`
	Message                           string                              `json:"message,omitempty"`
	BackendResponseList               *BackendResponseList                `json:"backendResponseList,omitempty"`
	GetRatedPerformanceIndicatorsList []GetRatedPerformanceIndicatorsList `json:"getRatedPerformanceIndicatorsList,omitempty"`
}

type GetRatedPerformanceIndicatorsList struct {
	PIType                 string                 `json:"pIType,omitempty"`
	AttributeList          []AttributeList        `json:"attributeList,omitempty"`
	PerformanceIndicatorID PerformanceIndicatorID `json:"performanceIndicatorID,omitempty"`
}

type AttributeList struct {
	AttributeValue string `json:"attributeValue,omitempty"`
	AttributeName  string `json:"attributeName,omitempty"`
}

type PerformanceIndicatorID struct {
	CycleInfo              CycleInfo         `json:"cycleInfo,omitempty"`
	AgreementID            string            `json:"agreementId,omitempty"`
	OfferInstance          string            `json:"offerInstance,omitempty"`
	PartitionID            string            `json:"partitionId,omitempty"`
	PerformanceIndicatorID string            `json:"performanceIndicatorId,omitempty"`
	FlexibleCycleInfo      FlexibleCycleInfo `json:"flexibleCycleInfo,omitempty"`
}

type CycleInfo struct {
	Month string `json:"month,omitempty"`
	Year  string `json:"year,omitempty"`
	Code  string `json:"code,omitempty"`
}

type FlexibleCycleInfo struct {
	Month string `json:"month,omitempty"`
	Year  string `json:"year,omitempty"`
	Code  string `json:"code,omitempty"`
}
