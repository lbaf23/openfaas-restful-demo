package util

import "encoding/json"

func JsonToStruct(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func StructToJson(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}
