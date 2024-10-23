package httphandler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHttpHeaderInfo(c *gin.Context, uuid *string) (*HttpHeaderInfo, *HttpRequestParamInfo) {
	var clientIP string
	var xChannel string
	var restpath string
	var xUsername string
	var xGateway string
	var xLegacyUsername string

	// Find client ip
	if len(c.Request.Header[http.CanonicalHeaderKey("Real-Client-IP")]) > 0 {
		clientIP = c.Request.Header[http.CanonicalHeaderKey("Real-Client-IP")][0]
	} else {
		clientIP = c.Request.RemoteAddr
	}

	// Find xChannel
	if len(c.Request.Header["X-Channel"]) > 0 {
		xChannel = c.Request.Header["X-Channel"][0]
	}

	// Find Gateway type
	if len(c.Request.Header[http.CanonicalHeaderKey("X-GatewayType")]) > 0 {
		xGateway = c.Request.Header[http.CanonicalHeaderKey("X-GatewayType")][0]
	}

	// Find legacy username
	if len(c.Request.Header[http.CanonicalHeaderKey("X-LegacyUsername")]) > 0 {
		xLegacyUsername = c.Request.Header[http.CanonicalHeaderKey("X-LegacyUsername")][0]
	}

	// Find rest path
	arrPath := strings.Split(c.Request.URL.Path, "/")
	if len(arrPath) > 0 {
		restpath = "/" + arrPath[len(arrPath)-1]
	}

	// Find username
	if len(c.Request.Header["X-Username"]) > 0 {
		xUsername = c.Request.Header["X-Username"][0]
	}
	username := findGatewayUser(c, xUsername)

	return &HttpHeaderInfo{
			ClientIP:        clientIP,
			XChannel:        xChannel,
			XGatewayType:    xGateway,
			XUsername:       xUsername,
			XLegacyUsername: xLegacyUsername,
		},
		&HttpRequestParamInfo{
			RestPath:     restpath,
			Username:     username,
			Uuid:         *uuid,
		}
}

func findGatewayUser(c *gin.Context, xuser string) string {
	// Get the Basic Authentication credentials
	user, _, hasAuth := c.Request.BasicAuth()
	if hasAuth {
		return user
	} else {
		return xuser
	}
}
