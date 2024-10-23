package main

import (
	"intbackend/bdhgateway/apirouter"
	"intbackend/bdhgateway/config"
	"intbackend/bdhgateway/errormsg"
	"intbackend/bdhgateway/locallogging"
)

// @title [BDH] Swagger APIs
// @version 00.01
// @description [BDH] Implement for support project Billing Upgrade.
// @BasePath /
// @schemes http
// @securityDefinitions.basic BasicAuth
// @in header
// @Username Authorization
// @Password Authorization
func main() {

	appConfig, errHanlerInfo := config.LoadConfig()

	if errHanlerInfo != nil && errHanlerInfo.Error != nil {
		// Set locallogging
		log := locallogging.LocalLogging{}
		errDesc := errHanlerInfo.Error.Error()
		errCode := errormsg.ERR_CD_INTERNAL_FAILURE
		msg := errormsg.GenInternalFailureMsg(errDesc)

		errLog := log.BuildErrorInputLoggerBeforeController(appConfig.Application.Profile, appConfig.Application.Name, errCode, msg, errHanlerInfo.Error, errHanlerInfo.ErrorApplication, errHanlerInfo.ErrorModule, errHanlerInfo.ErrorFile, errHanlerInfo.ErrorFunction)
		log.SetErrorInputLoggerBeforeController(errLog)
		log.WriteLogErrorBeforeController(errLog)
	}

	router := apirouter.SetupAPIRouter(appConfig)
	router.Run(appConfig.Server.Port)
}
