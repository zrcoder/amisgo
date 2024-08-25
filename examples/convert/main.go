package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func init() {
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		input, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		m := map[string]any{}
		json.Unmarshal(input, &m)
		resp := comp.Response{Msg: "转换成功", Data: comp.Data{"output": m["input"]}} // 这里的key值必须是第二个编辑器的name
		data, _ := json.Marshal(resp)
		w.Write(data)
	})
}

func main() {
	index := comp.Page().Body(
		comp.Form().ColumnCount(2).Body(
			comp.Editor().Language("json").Label("json").Name("input").Size("xxl"),
			comp.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
		).Actions(
			comp.Action().Label("Convert").ActionType("ajax").Level("primary").Api(
				comp.Schema{
					"url":  "/convert",
					"data": comp.Schema{"input": "${input}"},
					"responses": comp.Schema{
						"200": comp.Schema{
							"then": comp.Schema{
								"actionType": "setValue",
								"args": comp.Schema{
									"value": "${response}",
								},
							},
						},
					},
				},
			),
		),
	)

	panic(amisgo.ListenAndServe(index))
}
