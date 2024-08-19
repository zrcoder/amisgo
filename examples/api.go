package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ID       string `json:"id"`
	Engine   string `json:"engine"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Grade    string `json:"grade"`
}

const (
	itemsRouter = "/items"
	trendRouter = "/trend"
)

func ServeApi() {
	http.HandleFunc(itemsRouter, func(w http.ResponseWriter, r *http.Request) {
		items := []Item{
			{Engine: "e1", Browser: "chrome", Platform: "windows", Version: "1.0"},
			{Browser: "safri", Platform: "macOS", Version: "2.0"},
		}
		data, _ := json.Marshal(map[string]interface{}{"data": items})
		w.Write(data)
	})
	http.HandleFunc(trendRouter, func(w http.ResponseWriter, r *http.Request) {
		m := map[string]map[string]interface{}{"data": {"config": trendConfig}}
		resp, _ := json.Marshal(m)
		w.Write(resp)
	})
}

var (
	trendConfig = map[string]interface{}{
		"title":  titleMap,
		"xAxis":  xAxisMap,
		"series": series,
	}
	titleMap = map[string]string{"text": "消费趋势"}
	xAxisMap = map[string]interface{}{
		"type":        "category",
		"boundaryGap": false,
		"data":        xData,
	}
	series = []map[string]interface{}{
		{
			"name":      "销量",
			"type":      "line",
			"areaStyle": areaStyle,
			"data":      data,
		},
	}
	xData     = []string{"一月", "二月", "三月", "四月", "五月", "六月"}
	areaStyle = map[string]interface{}{
		"color": colorMap,
	}
	data     = []int{5, 20, 36, 10, 10, 20}
	colorMap = map[string]interface{}{
		"type":       "linear",
		"x":          0,
		"y":          0,
		"x2":         0,
		"y2":         1,
		"colorStops": colorStops,
		"global":     false,
	}
	colorStops = []map[string]interface{}{
		{
			"offset": 0,
			"color":  "rgba(84, 112, 197, 1)",
		},
		{
			"offset": 1,
			"color":  "rgba(84, 112, 197, 0)",
		},
	}
)
