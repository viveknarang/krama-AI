package main

import (
	"os"

	"github.com/romana/rlog"
)

//PROPERTYFILE property file
var PROPERTYFILE string

func loadEnvironmentVariables() {

	rlog.Debug("loadEnvironmentVariables() handle function invoked ...")

	PROPERTYFILE = os.Getenv("KRAMA_AI_PROPERTY_FILE")

	//If the environment variable is not set then loading from default location
	if PROPERTYFILE == "" {
		PROPERTYFILE = "/krama-AI.properties"
	}

}
