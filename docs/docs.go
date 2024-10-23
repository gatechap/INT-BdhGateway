// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bdhgateway/getcreditlimitwithperformanceindicators": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bdh-gateway-controller"
                ],
                "summary": "bdhgateway",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Channel",
                        "name": "X-Channel",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-GatewayType",
                        "name": "X-GatewayType",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-LegacyUsername",
                        "name": "X-LegacyUsername",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-Username",
                        "name": "X-Username",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "requestInfo",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.GetCreditLimitWithPerformanceIndicatorsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apimodel.GetCreditLimitWithPerformanceIndicatorsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    }
                }
            }
        },
        "/bdhgateway/getusage": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bdh-gateway-controller"
                ],
                "summary": "bdhgateway",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Channel",
                        "name": "X-Channel",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-GatewayType",
                        "name": "X-GatewayType",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-LegacyUsername",
                        "name": "X-LegacyUsername",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-Username",
                        "name": "X-Username",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "requestInfo",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.GetUsageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apimodel.GetUsageResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    }
                }
            }
        },
        "/bdhgateway/usageconsumption": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bdh-gateway-controller"
                ],
                "summary": "bdhgateway",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Channel",
                        "name": "X-Channel",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-GatewayType",
                        "name": "X-GatewayType",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-LegacyUsername",
                        "name": "X-LegacyUsername",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "X-Username",
                        "name": "X-Username",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "requestInfo",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.UsageConsumptionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/apimodel.UsageConsumptionResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/apimodel.HttpErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimodel.AccountInfo": {
            "type": "object",
            "properties": {
                "accStsBan": {
                    "type": "integer"
                },
                "accountNo": {
                    "type": "integer"
                },
                "accountOpenDate": {
                    "type": "string"
                },
                "accountPriority": {
                    "type": "string"
                },
                "accountSubType": {
                    "type": "string"
                },
                "accountingBalancePolicy": {
                    "type": "string"
                },
                "addressInfo": {
                    "type": "object"
                },
                "agreementId": {
                    "type": "integer"
                },
                "altFinObligationFrm": {
                    "type": "string"
                },
                "atbCharityCode": {
                    "type": "object"
                },
                "autoBlacklistInd": {
                    "type": "string"
                },
                "autoBlacklistRsnCd": {
                    "type": "object"
                },
                "autoBlacklistUpDate": {
                    "type": "object"
                },
                "billingCurrency": {
                    "type": "string"
                },
                "businessEntityId": {
                    "type": "integer"
                },
                "clWaiverUpdDate": {
                    "type": "string"
                },
                "colLastActDate": {
                    "type": "object"
                },
                "colRsnCd": {
                    "type": "string"
                },
                "colStatus": {
                    "type": "string"
                },
                "collWaiverExpDate": {
                    "type": "object"
                },
                "collectionFixCsrCd": {
                    "type": "object"
                },
                "collectionFixPolicy": {
                    "type": "string"
                },
                "collectionInd": {
                    "type": "string"
                },
                "collectionPermanentWaiveInd": {
                    "type": "string"
                },
                "collectionStartDate": {
                    "type": "object"
                },
                "collectionStatus": {
                    "type": "string"
                },
                "companyCode": {
                    "type": "string"
                },
                "convEffectiveDate": {
                    "type": "object"
                },
                "convergenceCode": {
                    "type": "object"
                },
                "crdLmtFixPolicyCd": {
                    "type": "object"
                },
                "creditClass": {
                    "type": "string"
                },
                "creditClassUpdDate": {
                    "type": "string"
                },
                "creditLastActDate": {
                    "type": "object"
                },
                "creditLimitExpDate": {
                    "type": "object"
                },
                "creditLimitRsnCode": {
                    "type": "string"
                },
                "creditLimitWaiverExpDate": {
                    "type": "object"
                },
                "creditLimitWaiverInd": {
                    "type": "string"
                },
                "creditRsnCd": {
                    "type": "object"
                },
                "creditStatus": {
                    "type": "string"
                },
                "custBranchNo": {
                    "type": "string"
                },
                "custTaxId": {
                    "type": "string"
                },
                "customerNo": {
                    "type": "integer"
                },
                "customerSubType": {
                    "type": "string"
                },
                "customerType": {
                    "type": "string"
                },
                "dfltCreditLimit": {
                    "type": "integer"
                },
                "documentType": {
                    "type": "string"
                },
                "externalId": {
                    "type": "object"
                },
                "finObligationFrm": {
                    "type": "string"
                },
                "fullSuspensionIndicator": {
                    "type": "string"
                },
                "historyInd": {
                    "type": "object"
                },
                "iddIndicator": {
                    "type": "string"
                },
                "incrThreshold": {
                    "type": "integer"
                },
                "initThreshold": {
                    "type": "integer"
                },
                "initiationReason": {
                    "type": "object"
                },
                "irIndicator": {
                    "type": "string"
                },
                "legacyBan": {
                    "type": "integer"
                },
                "manualBlacklistInd": {
                    "type": "string"
                },
                "manualBlacklistRsnCd": {
                    "type": "object"
                },
                "manualBlacklistUpDate": {
                    "type": "object"
                },
                "minimumPay": {
                    "type": "string"
                },
                "nameInfo": {
                    "type": "object"
                },
                "oblgCalcFormUpdDate": {
                    "type": "string"
                },
                "obligationCalcFormula": {
                    "type": "string"
                },
                "operatorId": {
                    "type": "integer"
                },
                "operatorName": {
                    "type": "string"
                },
                "parentAccount": {
                    "type": "integer"
                },
                "partnerCode": {
                    "type": "object"
                },
                "personalClUpdDate": {
                    "type": "string"
                },
                "personalCreditLimit": {
                    "type": "integer"
                },
                "punishmentLevel": {
                    "type": "object"
                },
                "relatedOu": {
                    "type": "integer"
                },
                "specialInstructions": {
                    "type": "object"
                },
                "suspensionRsn": {
                    "type": "string"
                },
                "suspensionType": {
                    "type": "object"
                },
                "taxExemptionDate": {
                    "type": "object"
                },
                "tempClUpdDate": {
                    "type": "object"
                },
                "tempCreditLimit": {
                    "type": "integer"
                },
                "whtCertiNo": {
                    "type": "object"
                },
                "whtInd": {
                    "type": "string"
                },
                "whtTaxUpDate": {
                    "type": "object"
                }
            }
        },
        "apimodel.AddressInfo": {
            "type": "object",
            "properties": {
                "addressLine1": {
                    "type": "string"
                },
                "addressLine2": {
                    "type": "string"
                },
                "addressLine3": {
                    "type": "string"
                },
                "addressLine4": {
                    "type": "string"
                },
                "addressType": {
                    "type": "string"
                },
                "amphur": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "effectiveDate": {
                    "type": "string"
                },
                "houseNo": {
                    "type": "string"
                },
                "linkType": {
                    "type": "string"
                },
                "moo": {
                    "type": "string"
                },
                "streetName": {
                    "type": "string"
                },
                "tumbon": {
                    "type": "string"
                },
                "typeOfAccomodation": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "apimodel.AttributeInfoList": {
            "type": "object",
            "properties": {
                "attributeName": {
                    "type": "string"
                },
                "attributeValue": {
                    "type": "string"
                }
            }
        },
        "apimodel.AttributeList": {
            "type": "object",
            "properties": {
                "attributeName": {
                    "type": "string"
                },
                "attributeValue": {
                    "type": "string"
                }
            }
        },
        "apimodel.BackendResponseInfoArray": {
            "type": "object",
            "properties": {
                "apiName": {
                    "type": "string"
                },
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "system": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "apimodel.BackendResponseList": {
            "type": "object",
            "properties": {
                "backendResponseInfoArray": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.BackendResponseInfoArray"
                    }
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "apimodel.CustomerCyclesInfos": {
            "type": "object",
            "properties": {
                "customerNo": {
                    "type": "integer"
                },
                "customerPartitionId": {
                    "type": "integer"
                },
                "cycleCode": {
                    "type": "integer"
                },
                "cycleInstance": {
                    "type": "integer"
                },
                "cycleSeqNo": {
                    "type": "integer"
                },
                "cycleYear": {
                    "type": "integer"
                },
                "endDate": {
                    "type": "string"
                },
                "runDate": {
                    "type": "object"
                },
                "startDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "apimodel.CycleInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "month": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "apimodel.Data": {
            "type": "object",
            "properties": {
                "accountInfo": {
                    "$ref": "#/definitions/apimodel.AccountInfo"
                },
                "customerCyclesInfos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.CustomerCyclesInfos"
                    }
                },
                "nameAddressInfoList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.NameAddressInfoList"
                    }
                },
                "performanceIndicatorInfos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.PerformanceIndicatorInfos"
                    }
                }
            }
        },
        "apimodel.ErrorMessageInfo": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                }
            }
        },
        "apimodel.FlexibleCycleInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "month": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "apimodel.GetCreditLimitWithPerformanceIndicatorsList": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.Results"
                    }
                }
            }
        },
        "apimodel.GetCreditLimitWithPerformanceIndicatorsRequest": {
            "type": "object",
            "required": [
                "requestList"
            ],
            "properties": {
                "requestList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.RequestList"
                    }
                }
            }
        },
        "apimodel.GetCreditLimitWithPerformanceIndicatorsResponse": {
            "type": "object",
            "properties": {
                "backendResponseList": {
                    "$ref": "#/definitions/apimodel.BackendResponseList"
                },
                "errorCode": {
                    "type": "string"
                },
                "getCreditLimitWithPerformanceIndicatorsList": {
                    "$ref": "#/definitions/apimodel.GetCreditLimitWithPerformanceIndicatorsList"
                },
                "message": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "apimodel.GetRatedEventsListInfoResponse": {
            "type": "object",
            "properties": {
                "ratedEventsInfoMList": {
                    "$ref": "#/definitions/apimodel.RatedEventsInfoList"
                }
            }
        },
        "apimodel.GetRatedPerformanceIndicatorsList": {
            "type": "object",
            "properties": {
                "attributeList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.AttributeList"
                    }
                },
                "pIType": {
                    "type": "string"
                },
                "performanceIndicatorID": {
                    "$ref": "#/definitions/apimodel.PerformanceIndicatorID"
                }
            }
        },
        "apimodel.GetUsageRequest": {
            "type": "object",
            "required": [
                "ban",
                "correlatedId",
                "customerId",
                "cycleCode",
                "cycleInstance",
                "cycleYear"
            ],
            "properties": {
                "SubscriberId": {
                    "type": "string"
                },
                "ban": {
                    "type": "string"
                },
                "callerUuid": {
                    "type": "string"
                },
                "correlatedId": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "cycleCode": {
                    "type": "string"
                },
                "cycleInstance": {
                    "type": "string"
                },
                "cycleYear": {
                    "type": "string"
                },
                "pageNumber": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "apimodel.GetUsageResponse": {
            "type": "object",
            "properties": {
                "backendResponseList": {
                    "$ref": "#/definitions/apimodel.BackendResponseList"
                },
                "errorCode": {
                    "type": "string"
                },
                "getRatedEventsListInfoResponse": {
                    "$ref": "#/definitions/apimodel.GetRatedEventsListInfoResponse"
                },
                "message": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "apimodel.HttpErrorResponse": {
            "type": "object",
            "properties": {
                "@timestamp": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "apimodel.NameAddressInfoList": {
            "type": "object",
            "properties": {
                "addressInfo": {
                    "$ref": "#/definitions/apimodel.AddressInfo"
                },
                "nameInfo": {
                    "$ref": "#/definitions/apimodel.NameInfo"
                }
            }
        },
        "apimodel.NameInfo": {
            "type": "object",
            "properties": {
                "effectiveDate": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "identification": {
                    "type": "string"
                },
                "identificationType": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "linkType": {
                    "type": "string"
                },
                "maritalStatus": {
                    "type": "string"
                },
                "nameLine1": {
                    "type": "string"
                },
                "nameType": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "apimodel.PerformanceIndicatorID": {
            "type": "object",
            "properties": {
                "agreementId": {
                    "type": "string"
                },
                "cycleInfo": {
                    "$ref": "#/definitions/apimodel.CycleInfo"
                },
                "flexibleCycleInfo": {
                    "$ref": "#/definitions/apimodel.FlexibleCycleInfo"
                },
                "offerInstance": {
                    "type": "string"
                },
                "partitionId": {
                    "type": "string"
                },
                "performanceIndicatorId": {
                    "type": "string"
                }
            }
        },
        "apimodel.PerformanceIndicatorInfos": {
            "type": "object",
            "properties": {
                "agreementNo": {
                    "type": "integer"
                },
                "attributeInfoList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.AttributeInfoList"
                    }
                },
                "cycleCode": {
                    "type": "integer"
                },
                "cycleInstance": {
                    "type": "integer"
                },
                "cycleYear": {
                    "type": "integer"
                },
                "offerInstance": {
                    "type": "integer"
                },
                "partitionId": {
                    "type": "integer"
                },
                "performanceIndicatorId": {
                    "type": "integer"
                }
            }
        },
        "apimodel.RatedEventsInfo": {
            "type": "object",
            "properties": {
                "additionalCharge": {
                    "type": "string"
                },
                "apn": {
                    "type": "string"
                },
                "billingArrangement": {
                    "type": "string"
                },
                "callFwdInd": {
                    "type": "string"
                },
                "calledNumber": {
                    "type": "string"
                },
                "callingNumber": {
                    "type": "string"
                },
                "chargeAmt": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "destinationCode": {
                    "type": "string"
                },
                "destinationCountry": {
                    "type": "string"
                },
                "directionCode": {
                    "type": "string"
                },
                "discountAmt": {
                    "type": "string"
                },
                "displayOnBill": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "endTime": {
                    "type": "string"
                },
                "eventId": {
                    "type": "string"
                },
                "eventState": {
                    "type": "string"
                },
                "eventStateReasonCode": {
                    "type": "string"
                },
                "eventType": {
                    "type": "string"
                },
                "eventTypeId": {
                    "type": "string"
                },
                "eventTypeName": {
                    "type": "string"
                },
                "freeAmt": {
                    "type": "string"
                },
                "guidingResourceId": {
                    "type": "string"
                },
                "imsi": {
                    "type": "string"
                },
                "offerCode": {
                    "type": "string"
                },
                "operatorName": {
                    "type": "string"
                },
                "pabxOrigSubscriber": {
                    "type": "string"
                },
                "payChannel": {
                    "type": "string"
                },
                "paymentCategory": {
                    "type": "string"
                },
                "period": {
                    "type": "string"
                },
                "roamInd": {
                    "type": "string"
                },
                "roundedUnits": {
                    "type": "string"
                },
                "roundedUnitsUom": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                }
            }
        },
        "apimodel.RatedEventsInfoList": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "$ref": "#/definitions/apimodel.ErrorMessageInfo"
                },
                "ratedEventsInfoM": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.RatedEventsInfo"
                    }
                },
                "totalSize": {
                    "type": "string"
                },
                "transactionId": {
                    "type": "string"
                }
            }
        },
        "apimodel.RequestList": {
            "type": "object",
            "properties": {
                "accountNo": {
                    "type": "integer"
                },
                "subscriberNo": {
                    "type": "integer"
                }
            }
        },
        "apimodel.Results": {
            "type": "object",
            "properties": {
                "accountNo": {
                    "type": "integer"
                },
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/apimodel.Data"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "subscriberNo": {
                    "type": "integer"
                }
            }
        },
        "apimodel.UsageConsumptionRequest": {
            "type": "object",
            "required": [
                "agreementId",
                "correlatedId",
                "customerId",
                "cycleCode",
                "cycleMonth",
                "cycleYear",
                "externalStructure"
            ],
            "properties": {
                "agreementId": {
                    "type": "string"
                },
                "callerUuid": {
                    "type": "string"
                },
                "correlatedId": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "cycleCode": {
                    "type": "string"
                },
                "cycleMonth": {
                    "type": "string"
                },
                "cycleYear": {
                    "type": "string"
                },
                "externalStructure": {
                    "type": "string"
                },
                "pageNumber": {
                    "type": "string"
                },
                "pageSize": {
                    "type": "string"
                }
            }
        },
        "apimodel.UsageConsumptionResponse": {
            "type": "object",
            "properties": {
                "backendResponseList": {
                    "$ref": "#/definitions/apimodel.BackendResponseList"
                },
                "errorCode": {
                    "type": "string"
                },
                "getRatedPerformanceIndicatorsList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimodel.GetRatedPerformanceIndicatorsList"
                    }
                },
                "message": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "00.01",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "[BDH] Swagger APIs",
	Description: "[BDH] Implement for support project Billing Upgrade.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
