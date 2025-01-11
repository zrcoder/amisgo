package servermux

import (
	js "encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/zrcoder/amisgo/model"
)

var innerApiID int32 = -1

func getInnerApiID() int32 {
	return atomic.AddInt32(&innerApiID, 1)
}

func getRoute() string {
	return fmt.Sprintf("/__amisgo__%d", getInnerApiID())
}

func BindDataRoute(callback func(model.Schema) error) string {
	return handleGenericRequest(func(input []byte) error {
		obj := model.Schema{}
		if err := js.Unmarshal(input, &obj); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback(obj)
	})
}

func BindRouteTo(receiver any, callback func(any) error) string {
	return handleGenericRequest(func(input []byte) error {
		if err := js.Unmarshal(input, receiver); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback(receiver)
	})
}

func handleGenericRequest(requestProcessor func([]byte) error) string {
	route := getRoute()
	Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := io.ReadAll(r.Body)
		if err != nil {
			respError(w, fmt.Errorf("failed to read request body: %w", err))
			return
		}
		defer r.Body.Close()

		if err = requestProcessor(input); err != nil {
			respError(w, fmt.Errorf("request processing failed: %w", err))
			return
		}

		rsp := model.Response{}
		w.Write(rsp.Json())
	})
	return route
}

func ServeData(getter func() (any, error)) string {
	route := getRoute()
	Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := getter()
		if err != nil {
			respError(w, err)
			return
		}
		if err := writeJsonResponse(w, input); err != nil {
			respError(w, err)
		}
	})
	return route
}

func writeJsonResponse(w http.ResponseWriter, data any) error {
	jsonData, err := js.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	return err
}

func ServeUpload(maxMemory int64, action func([]byte) (path string, err error)) string {
	route := getRoute()
	Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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
		resp := model.Response{Data: model.Schema{"value": path}}
		w.Write(resp.Json())
	})
	return route
}

func respError(w http.ResponseWriter, err error) {
	resp := model.ErrorResponse(err.Error())
	w.Write(resp.Json())
}

func TransformMultiple(src []string, transfor func(model.Schema) (model.Schema, error)) (route string, data model.Schema) {
	route = getRoute()
	Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		inputData, err := io.ReadAll(r.Body)
		if err != nil {
			respError(w, err)
			return
		}
		defer r.Body.Close()
		d := model.Schema{}
		err = js.Unmarshal(inputData, &d)
		if err != nil {
			respError(w, err)
			return
		}
		output, err := transfor(d)
		if err != nil {
			respError(w, err)
			return
		}
		resp := model.SuccessResponse(" ", output) // " " for empty msg
		w.Write(resp.Json())
	})

	data = make(model.Schema, len(src))
	for _, s := range src {
		data.Set(s, "${"+s+"}")
	}

	return
}
