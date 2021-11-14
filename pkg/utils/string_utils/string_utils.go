package string_utils

import (
	"encoding/json"
)

// IsJSONString checks if the JSON is string
// by attempting to dump it on an interface
func IsJSONString(s string) bool {
	var js interface{}
	err := json.Unmarshal([]byte(s), &js)
	if err != nil {
		return false
	}
	return true
}

// ToJSON returns a variable as type JSON... Or "error" (string) if not
func ToJSON(i interface{}) string {
	_json, err := json.Marshal(i)
	if err != nil {
		return "error marshaling"
	} else {
		return string(_json)
	}
}