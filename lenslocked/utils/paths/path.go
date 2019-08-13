package paths

import (
	"os"
)

func BuildPath(config string) string {
	if os.Getenv("PROFILE") == "PROD" {
		return (config + "_prod")
	}
	if os.Getenv("PROFILE") == "STG" {
		return (config + "_stg")
	}
	if os.Getenv("PROFILE") == "DEV" {
		return (config + "_dev")
	}

	// default "_dev"
	return config + "_dev"
}
