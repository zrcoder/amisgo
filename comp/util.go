package comp

import (
	js "encoding/json"
	"fmt"
	"io"
	"net/http"
)

var innerApiID = -1

func getInnerApiID() int {
	innerApiID++
	return innerApiID
}

func getRoute() string {
	return fmt.Sprintf("/__amisgo__%d", getInnerApiID())
}

func serveApi(action func(Data) error) string {
	route := getRoute()
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := io.ReadAll(r.Body)
		if err != nil {
			respError(w, err)
			return
		}
		defer r.Body.Close()
		m := Data{}
		if err = js.Unmarshal(input, &m); err != nil {
			respError(w, err)
			return
		}
		if err = action(m); err != nil {
			respError(w, err)
			return
		}
	})
	return route
}

func serveData(getter func() (any, error)) string {
	route := getRoute()
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := getter()
		if err != nil {
			respError(w, err)
			return
		}
		data, err := js.Marshal(input)
		if err != nil {
			respError(w, err)
			return
		}
		w.Write(data)
	})
	return route
}

func respError(w http.ResponseWriter, err error) {
	resp := Response{Status: 1, Msg: err.Error()}
	w.Write(resp.Json())
}
