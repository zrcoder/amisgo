package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

// Traform Action 与注释掉的代码效果一致，甚至不需要写 init() 函数里的内容

func main() {
	index := comp.Page().Body(
		comp.Form().ColumnCount(2).Body(
			comp.Editor().Language("json").Label("json").Name("input").Size("xxl"),
			comp.Editor().Label("yaml").Label("yaml").Name("output").Size("xxl").ReadOnly(true),
		).Actions(
			comp.Action().Label("Converrt").Level("primary").Transform("input", "output", "转换成功", func(input any) any {
				// transform input json to yaml
				output := "age: 27"
				return output
			}),

			// comp.Action().Label("Convert").Level("primary").ActionType("ajax").Api(
			// 	comp.Schema{
			// 		"url":  "/convert",
			// 		"data": comp.Schema{"input": "${input}"},
			// 		"responses": comp.Schema{
			// 			"200": comp.Schema{
			// 				"then": comp.Schema{
			// 					"actionType": "setValue",
			// 					"args": comp.Schema{
			// 						"value": "${response}",
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// ),
		),
	)

	panic(amisgo.ListenAndServe(index))
}

// func init() {
// 	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
// 		input, _ := io.ReadAll(r.Body)
// 		defer r.Body.Close()
// 		m := map[string]any{}
// 		json.Unmarshal(input, &m)
// 		// ...
// 		output := "age: 27"
// 		resp := comp.Response{Msg: "转换成功", Data: comp.Data{"output": output}} // 这里的key值必须是第二个编辑器的 name
// 		data, _ := json.Marshal(resp)
// 		w.Write(data)
// 	})
// }
