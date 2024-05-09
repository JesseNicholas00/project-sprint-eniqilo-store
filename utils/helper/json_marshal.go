package helper

import "encoding/json"

// panics on marshal failure
func MustMarshalJson(obj interface{}) []byte {
	res, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return res
}
