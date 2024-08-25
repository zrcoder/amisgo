package comp

import js "encoding/json"

type AmisComp any

// schema 通用 model，键未指定的 map
type schema = map[string]any

type (
	Schema schema
	Data   schema
)

func (s Schema) Set(key string, value any) Schema {
	s[key] = value
	return s
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func (r *Response) Json() []byte {
	data, _ := js.Marshal(r)
	return data
}
