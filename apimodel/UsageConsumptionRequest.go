package apimodel

type UsageConsumptionRequest struct {
	CorrelatedId      string `json:"correlatedId" binding:"required"`
	CallerUuid        string `json:"callerUuid"`
	CustomerId        string `json:"customerId" binding:"required"`
	AgreementId       string `json:"agreementId" binding:"required"`
	CycleMonth        string `json:"cycleMonth" binding:"required"`
	CycleCode         string `json:"cycleCode" binding:"required"`
	CycleYear         string `json:"cycleYear" binding:"required"`
	PageSize          string `json:"pageSize"`
	PageNumber        string `json:"pageNumber"`
	ExternalStructure string `json:"externalStructure" binding:"required"`
}
