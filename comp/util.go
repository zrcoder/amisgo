package comp

import (
	js "encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/zrcoder/amisgo/internal/servermux"
)

var innerApiID int32 = -1

func getInnerApiID() int32 {
	return atomic.AddInt32(&innerApiID, 1)
}

func getRoute() string {
	return fmt.Sprintf("/__amisgo__%d", getInnerApiID())
}

func bindDataRoute(callback func(Data) error) string {
	return handleGenericRequest(func(input []byte) error {
		obj := Data{}
		if err := js.Unmarshal(input, &obj); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback(obj)
	})
}

func bindRouteTo(receiver any, callback func(any) error) string {
	return handleGenericRequest(func(input []byte) error {
		if err := js.Unmarshal(input, receiver); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback(receiver)
	})
}

func handleGenericRequest(requestProcessor func([]byte) error) string {
	route := getRoute()
	servermux.Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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

		rsp := Response{}
		w.Write(rsp.Json())
	})
	return route
}

func serveData(getter func() (any, error)) string {
	route := getRoute()
	servermux.Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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

func serveUpload(maxMemory int64, action func([]byte) (path string, err error)) string {
	route := getRoute()
	servermux.Mux().HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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
	resp := ErrorResponse(err.Error())
	w.Write(resp.Json())
}
