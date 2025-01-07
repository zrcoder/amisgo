package comp

import js "encoding/json"

type (
	Schema map[string]any
	Data   = Schema
)

func (s Schema) Set(key string, value any) {
	s[key] = value
}

func (s Schema) Get(key string) any {
	return s[key]
}

func (s Schema) Json() []byte {
	data, _ := js.Marshal(s)
	return data
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func ErrorResponse(msg string) *Response {
	return &Response{Status: 1, Msg: msg}
}

func SuccessResponse(msg string, data Data) *Response {
	return &Response{Msg: msg, Data: data}
}

func (r *Response) Json() []byte {
	data, _ := js.Marshal(r)
	return data
}

// RawJson create components from amis raw json
func RawJson(input any) any {
	return input
}
