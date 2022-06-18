package utils

import "encoding/json"

func JSONEncode(object interface{}) string {
	json, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(json)
}