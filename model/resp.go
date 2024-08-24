package model

import "encoding/json"

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Data   `json:"data"`
}

func (r *Response) Json() []byte {
	data, _ := json.Marshal(r)
	return data
}
