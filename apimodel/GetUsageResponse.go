package apimodel

type GetUsageResponse struct {
	Uuid                           string                          `json:"uuid,omitempty"`
	ErrorCode                      string                          `json:"errorCode,omitempty"`
	Message                        string                          `json:"message,omitempty"`
	BackendResponseList            *BackendResponseList            `json:"backendResponseList,omitempty"`
	GetRatedEventsListInfoResponse *GetRatedEventsListInfoResponse `json:"getRatedEventsListInfoResponse,omitempty"`
}

type GetRatedEventsListInfoResponse struct {
	RatedEventsInfoList *RatedEventsInfoList `json:"ratedEventsInfoMList,omitempty"`
}

type RatedEventsInfoList struct {
	TransactionId    string            `json:"transactionId,omitempty"`
	ErrorMessageInfo ErrorMessageInfo  `json:"errorMessage,omitempty"`
	TotalSize        string            `json:"totalSize,omitempty"`
	RatedEventsInfo  []RatedEventsInfo `json:"ratedEventsInfoM,omitempty"`
}

type RatedEventsInfo struct {
	Code                 string `json:"code,omitempty"`
	DisplayOnBill        string `json:"displayOnBill,omitempty"`
	DiscountAmt          string `json:"discountAmt,omitempty"`
	CallingNumber        string `json:"callingNumber,omitempty"`
	CallFwdInd           string `json:"callFwdInd,omitempty"`
	Imsi                 string `json:"imsi,omitempty"`
	RoundedUnits         string `json:"roundedUnits,omitempty"`
	OperatorName         string `json:"operatorName,omitempty"`
	RoamInd              string `json:"roamInd,omitempty"`
	Duration             string `json:"duration,omitempty"`
	EventState           string `json:"eventState,omitempty"`
	CustomerId           string `json:"customerId,omitempty"`
	PabxOrigSubscriber   string `json:"pabxOrigSubscriber,omitempty"`
	StartTime            string `json:"startTime,omitempty"`
	CalledNumber         string `json:"calledNumber,omitempty"`
	Apn                  string `json:"apn,omitempty"`
	VventId              string `json:"eventId,omitempty"`
	Period               string `json:"period,omitempty"`
	ChargeAmt            string `json:"chargeAmt,omitempty"`
	PaymentCategory      string `json:"paymentCategory,omitempty"`
	RoundedUnitsUom      string `json:"roundedUnitsUom,omitempty"`
	EventType            string `json:"eventType,omitempty"`
	DestinationCountry   string `json:"destinationCountry,omitempty"`
	AdditionalCharge     string `json:"additionalCharge,omitempty"`
	EventStateReasonCode string `json:"eventStateReasonCode,omitempty"`
	BillingArrangement   string `json:"billingArrangement,omitempty"`
	DirectionCode        string `json:"directionCode,omitempty"`
	DestinationCode      string `json:"destinationCode,omitempty"`
	EventTypeId          string `json:"eventTypeId,omitempty"`
	GuidingResourceId    string `json:"guidingResourceId,omitempty"`
	OfferCode            string `json:"offerCode,omitempty"`
	EventTypeName        string `json:"eventTypeName,omitempty"`
	PayChannel           string `json:"payChannel,omitempty"`
	EndTime              string `json:"endTime,omitempty"`
	StartDate            string `json:"startDate,omitempty"`
	FreeAmt              string `json:"freeAmt,omitempty"`
}

type ErrorMessageInfo struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
}
