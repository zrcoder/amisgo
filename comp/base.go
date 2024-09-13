package comp

import js "encoding/json"

type AmisComp interface {
	~map[string]any | ~[]byte | ~string
}

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

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func (r *Response) Json() []byte {
	data, _ := js.Marshal(r)
	return data
}

// RawJson create components from amis raw json
func RawJson[T AmisComp](input T) T {
	return input
}
