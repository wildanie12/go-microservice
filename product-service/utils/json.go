package utils

import "encoding/json"

// JSONEncode encodes any object to json string,
// useful for debugging an error
func JSONEncode(object interface{}) string {
	json, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(json)
}