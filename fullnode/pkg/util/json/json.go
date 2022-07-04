package json

import (
	jsoniter "github.com/json-iterator/go"
)

// Defining JSON operations
var (
	json          = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)

func MarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

func UnmarshalFromString(str string, v interface{}) error {
	err := jsoniter.UnmarshalFromString(str, v)
	if err != nil {
		return err
	}
	return nil
}
