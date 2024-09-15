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
		rsp := Response{}
		w.Write(rsp.Json())
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

func serveUpload(maxMemory int64, action func([]byte) (path string, err error)) string {
	route := getRoute()
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		// Parse the form to get file data
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			respError(w, err)
			return
		}

		// Get the file from the form
		file, _, err := r.FormFile("file")
		if err != nil {
			respError(w, err)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			respError(w, err)
			return
		}

		path, err := action(data)
		if err != nil {
			respError(w, err)
			return
		}
		resp := Response{Data: Data{"value": path}}
		w.Write(resp.Json())
	})
	return route
}

func respError(w http.ResponseWriter, err error) {
	resp := Response{Status: 1, Msg: err.Error()}
	w.Write(resp.Json())
}
