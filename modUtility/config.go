package modUtility

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// x api key: for operator only;
const CG_Key_xApiKey = "XAPIKEY"
const CG_Key_HttpPort = "HTTPPORT"
const CG_Key_DBUrl = "DBURL"
const CG_Key_DBToken = "DBTOKEN"

const GD_Default_XApiKey = "aaXAPIKEYaa"
const GD_Default_HttpPort = 10000

var G_XApiKey = ""
var G_HttpPort = 0
var G_DBUrl = ""
var G_DBToken = ""

func config_Initialize() error {
	G_XApiKey = getEnvKey(CG_Key_xApiKey)
	if G_XApiKey == "" {
		G_XApiKey = GD_Default_XApiKey
	}
	portStr := getEnvKey(CG_Key_HttpPort)
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			G_HttpPort = port
		}
	}
	if G_HttpPort < 10 {
		G_HttpPort = GD_Default_HttpPort
	}
	G_DBUrl = getEnvKey(CG_Key_DBUrl)
	if G_DBUrl == "" {
		return errors.New("database url empty")
	}

	G_DBToken = getEnvKey(CG_Key_DBToken)
	if G_DBToken == "" {
		return errors.New("database token empty")
	}

	return nil
}

func getEnvKey(key string) string {
	portStr := os.Getenv(key)
	portStr = strings.TrimSpace(portStr)
	portStr = strings.Trim(portStr, "\r")
	portStr = strings.Trim(portStr, "\n")
	portStr = strings.Trim(portStr, "\r")

	return portStr
}
