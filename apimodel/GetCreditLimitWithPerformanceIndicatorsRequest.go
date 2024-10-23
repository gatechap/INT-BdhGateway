package apimodel

type GetCreditLimitWithPerformanceIndicatorsRequest struct {
	RequestList []RequestList `json:"requestList" validate:"required"`
}

type RequestList struct {
	AccountNo    int `json:"accountNo"`
	SubscriberNo int `json:"subscriberNo"`
}
