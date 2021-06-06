package main

import "encoding/json"

type JsonParser struct{}

func (j *JsonParser) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *JsonParser) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *JsonParser) MimeType() string {
	return "application/json"
}
