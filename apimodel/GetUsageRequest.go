package apimodel

type GetUsageRequest struct {
	CorrelatedId  string `json:"correlatedId" binding:"required"`
	CallerUuid    string `json:"callerUuid"`
	CustomerId    string `json:"customerId" binding:"required"`
	SubscriberId  string `json:"SubscriberId"`
	Ban           string `json:"ban" binding:"required"`
	CycleCode     string `json:"cycleCode" binding:"required"`
	CycleYear     string `json:"cycleYear" binding:"required"`
	CycleInstance string `json:"cycleInstance" binding:"required"`
	PageSize      int32  `json:"pageSize"`
	PageNumber    int32  `json:"pageNumber"`
}
