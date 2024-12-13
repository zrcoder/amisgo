package comp

import js "encoding/json"

// schema 通用 model，键未指定的 map
type schema = map[string]any

type (
	Schema schema
	Data   schema
)

func (d Data) Set(key string, value any) {
	d[key] = value
}

func (d Data) Get(key string) any {
	return d[key]
}

func (d Data) Json() []byte {
	data, _ := js.Marshal(d)
	return data
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func ErrorResponse(msg string) *Response {
	return &Response{Status: -1, Msg: msg}
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
