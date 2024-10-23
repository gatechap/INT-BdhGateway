package apimodel

import "time"

type GetCreditLimitWithPerformanceIndicatorsResponse struct {
	Uuid                                        string                                       `json:"uuid,omitempty"`
	ErrorCode                                   string                                       `json:"errorCode,omitempty"`
	Message                                     string                                       `json:"message,omitempty"`
	BackendResponseList                         *BackendResponseList                         `json:"backendResponseList,omitempty"`
	GetCreditLimitWithPerformanceIndicatorsList *GetCreditLimitWithPerformanceIndicatorsList `json:"getCreditLimitWithPerformanceIndicatorsList"`
}

type GetCreditLimitWithPerformanceIndicatorsList struct {
	Results []Results `json:"results"`
}

type GetCreditLimitWithPerformanceIndicatorsListInfo struct {
	Results []Results `json:"results"`
}

type Results struct {
	Code         int    `json:"code"`
	Data         Data   `json:"data,omitempty"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	AccountNo    int    `json:"accountNo"`
	SubscriberNo int    `json:"subscriberNo"`
}

type Data struct {
	AccountInfo               AccountInfo                 `json:"accountInfo"`
	NameAddressInfoList       []NameAddressInfoList       `json:"nameAddressInfoList"`
	CustomerCyclesInfos       []CustomerCyclesInfos       `json:"customerCyclesInfos"`
	PerformanceIndicatorInfos []PerformanceIndicatorInfos `json:"performanceIndicatorInfos"`
}

type AccountInfo struct {
	AccountNo                   int         `json:"accountNo"`
	CustomerNo                  int         `json:"customerNo"`
	ExternalID                  interface{} `json:"externalId"`
	AccountSubType              string      `json:"accountSubType"`
	AccountPriority             string      `json:"accountPriority"`
	AccountingBalancePolicy     string      `json:"accountingBalancePolicy"`
	AgreementID                 int         `json:"agreementId"`
	AtbCharityCode              interface{} `json:"atbCharityCode"`
	AutoBlacklistInd            string      `json:"autoBlacklistInd"`
	AutoBlacklistRsnCd          interface{} `json:"autoBlacklistRsnCd"`
	AutoBlacklistUpDate         interface{} `json:"autoBlacklistUpDate"`
	ManualBlacklistInd          string      `json:"manualBlacklistInd"`
	ManualBlacklistRsnCd        interface{} `json:"manualBlacklistRsnCd"`
	ManualBlacklistUpDate       interface{} `json:"manualBlacklistUpDate"`
	OblgCalcFormUpdDate         string      `json:"oblgCalcFormUpdDate"`
	ObligationCalcFormula       string      `json:"obligationCalcFormula"`
	ClWaiverUpdDate             string      `json:"clWaiverUpdDate"`
	ColLastActDate              interface{} `json:"colLastActDate"`
	ColRsnCd                    string      `json:"colRsnCd"`
	ColStatus                   string      `json:"colStatus"`
	CompanyCode                 string      `json:"companyCode"`
	ConvEffectiveDate           interface{} `json:"convEffectiveDate"`
	ConvergenceCode             interface{} `json:"convergenceCode"`
	CrdLmtFixPolicyCd           interface{} `json:"crdLmtFixPolicyCd"`
	CreditClass                 string      `json:"creditClass"`
	CreditClassUpdDate          string      `json:"creditClassUpdDate"`
	CreditLastActDate           interface{} `json:"creditLastActDate"`
	CreditLimitExpDate          interface{} `json:"creditLimitExpDate"`
	CreditLimitRsnCode          string      `json:"creditLimitRsnCode"`
	CreditLimitWaiverExpDate    interface{} `json:"creditLimitWaiverExpDate"`
	CreditLimitWaiverInd        string      `json:"creditLimitWaiverInd"`
	CreditRsnCd                 interface{} `json:"creditRsnCd"`
	CreditStatus                string      `json:"creditStatus"`
	CustBranchNo                string      `json:"custBranchNo"`
	CustTaxID                   string      `json:"custTaxId"`
	DocumentType                string      `json:"documentType"`
	IddIndicator                string      `json:"iddIndicator"`
	InitiationReason            interface{} `json:"initiationReason"`
	IrIndicator                 string      `json:"irIndicator"`
	HistoryInd                  interface{} `json:"historyInd"`
	AccountOpenDate             string      `json:"accountOpenDate"`
	BillingCurrency             string      `json:"billingCurrency"`
	TaxExemptionDate            interface{} `json:"taxExemptionDate"`
	BusinessEntityID            int         `json:"businessEntityId"`
	AccStsBan                   int         `json:"accStsBan"`
	LegacyBan                   int         `json:"legacyBan"`
	CollectionInd               string      `json:"collectionInd"`
	CollectionStartDate         interface{} `json:"collectionStartDate"`
	CollectionFixCsrCd          interface{} `json:"collectionFixCsrCd"`
	CollectionFixPolicy         string      `json:"collectionFixPolicy"`
	CollectionPermanentWaiveInd string      `json:"collectionPermanentWaiveInd"`
	CollWaiverExpDate           interface{} `json:"collWaiverExpDate"`
	CollectionStatus            string      `json:"collectionStatus"`
	FullSuspensionIndicator     string      `json:"fullSuspensionIndicator"`
	SuspensionRsn               string      `json:"suspensionRsn"`
	SuspensionType              interface{} `json:"suspensionType"`
	PunishmentLevel             interface{} `json:"punishmentLevel"`
	OperatorID                  int         `json:"operatorId"`
	OperatorName                string      `json:"operatorName"`
	ParentAccount               int         `json:"parentAccount"`
	PartnerCode                 interface{} `json:"partnerCode"`
	PersonalClUpdDate           string      `json:"personalClUpdDate"`
	PersonalCreditLimit         int         `json:"personalCreditLimit"`
	RelatedOu                   int         `json:"relatedOu"`
	SpecialInstructions         interface{} `json:"specialInstructions"`
	TempClUpdDate               interface{} `json:"tempClUpdDate"`
	TempCreditLimit             int         `json:"tempCreditLimit"`
	WhtCertiNo                  interface{} `json:"whtCertiNo"`
	WhtInd                      string      `json:"whtInd"`
	WhtTaxUpDate                interface{} `json:"whtTaxUpDate"`
	AddressInfo                 interface{} `json:"addressInfo"`
	NameInfo                    interface{} `json:"nameInfo"`
	CustomerType                string      `json:"customerType"`
	CustomerSubType             string      `json:"customerSubType"`
	DfltCreditLimit             int         `json:"dfltCreditLimit"`
	InitThreshold               int         `json:"initThreshold"`
	IncrThreshold               int         `json:"incrThreshold"`
	FinObligationFrm            string      `json:"finObligationFrm"`
	AltFinObligationFrm         string      `json:"altFinObligationFrm"`
	MinimumPay                  string      `json:"minimumPay"`
}

type NameAddressInfoList struct {
	NameInfo    NameInfo    `json:"nameInfo,omitempty"`
	AddressInfo AddressInfo `json:"addressInfo,omitempty"`
}

type NameInfo struct {
	NameType           string    `json:"nameType"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Identification     string    `json:"identification"`
	Gender             string    `json:"gender"`
	NameLine1          string    `json:"nameLine1"`
	LinkType           string    `json:"linkType"`
	IdentificationType string    `json:"identificationType"`
	Title              string    `json:"title"`
	EffectiveDate      time.Time `json:"effectiveDate"`
	MaritalStatus      string    `json:"maritalStatus"`
}

type AddressInfo struct {
	Zip                string    `json:"zip"`
	TypeOfAccomodation string    `json:"typeOfAccomodation"`
	City               string    `json:"city"`
	AddressType        string    `json:"addressType"`
	Tumbon             string    `json:"tumbon"`
	Moo                string    `json:"moo"`
	Amphur             string    `json:"amphur"`
	StreetName         string    `json:"streetName"`
	AddressLine1       string    `json:"addressLine1"`
	HouseNo            string    `json:"houseNo"`
	LinkType           string    `json:"linkType"`
	AddressLine2       string    `json:"addressLine2"`
	AddressLine3       string    `json:"addressLine3"`
	AddressLine4       string    `json:"addressLine4"`
	EffectiveDate      time.Time `json:"effectiveDate"`
}

type CustomerCyclesInfos struct {
	CycleCode           int         `json:"cycleCode"`
	CycleYear           int         `json:"cycleYear"`
	CycleInstance       int         `json:"cycleInstance"`
	CycleSeqNo          int         `json:"cycleSeqNo"`
	StartDate           string      `json:"startDate"`
	EndDate             string      `json:"endDate"`
	Status              string      `json:"status"`
	CustomerNo          int         `json:"customerNo"`
	CustomerPartitionID int         `json:"customerPartitionId"`
	RunDate             interface{} `json:"runDate"`
}

type PerformanceIndicatorInfos struct {
	PerformanceIndicatorID int                 `json:"performanceIndicatorId"`
	OfferInstance          int                 `json:"offerInstance"`
	PartitionID            int                 `json:"partitionId"`
	AgreementNo            int                 `json:"agreementNo"`
	CycleCode              int                 `json:"cycleCode"`
	CycleYear              int                 `json:"cycleYear"`
	CycleInstance          int                 `json:"cycleInstance"`
	AttributeInfoList      []AttributeInfoList `json:"attributeInfoList"`
}

type AttributeInfoList struct {
	AttributeName  string `json:"attributeName"`
	AttributeValue string `json:"attributeValue"`
}
