package servemux

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/zrcoder/amisgo/schema"
)

var innerApiID int32 = -1

func getInnerApiID() int32 {
	return atomic.AddInt32(&innerApiID, 1)
}

const Key = "__amisgo_mux"

func getRoute() string {
	return fmt.Sprintf("/__amisgo__%d", getInnerApiID())
}

func BindDataRoute(mux *http.ServeMux, callback func(schema.Schema) error) string {
	return handleGenericRequest(mux, func(input []byte) error {
		obj := schema.Schema{}
		if err := json.Unmarshal(input, &obj); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback(obj)
	})
}

func BindRouteTo(mux *http.ServeMux, receiver any, callback func() error) string {
	return handleGenericRequest(mux, func(input []byte) error {
		if err := json.Unmarshal(input, receiver); err != nil {
			return fmt.Errorf("data unmarshal: %w", err)
		}
		return callback()
	})
}

func handleGenericRequest(mux *http.ServeMux, requestProcessor func([]byte) error) string {
	route := getRoute()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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

		rsp := schema.Response{}
		w.Write(rsp.Json())
	})
	return route
}

func ServeData(mux *http.ServeMux, getter func() (any, error)) string {
	route := getRoute()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := getter()
		if err != nil {
			respError(w, err)
			return
		}
		if err = writeJsonResponse(w, input); err != nil {
			respError(w, err)
		}
	})
	return route
}

func writeJsonResponse(w http.ResponseWriter, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	return err
}

func ServeUpload(mux *http.ServeMux, maxMemory int64, action func([]byte) (path string, err error)) string {
	route := getRoute()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
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
		resp := schema.Response{Data: schema.Schema{"value": path}}
		w.Write(resp.Json())
	})
	return route
}

func respError(w http.ResponseWriter, err error) {
	resp := schema.ErrorResponse(err.Error())
	w.Write(resp.Json())
}

func TransformMultiple(mux *http.ServeMux, src []string, transfor func(schema.Schema) (schema.Schema, error)) (route string, data schema.Schema) {
	route = getRoute()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		inputData, err := io.ReadAll(r.Body)
		if err != nil {
			respError(w, err)
			return
		}
		defer r.Body.Close()
		d := schema.Schema{}
		err = json.Unmarshal(inputData, &d)
		if err != nil {
			respError(w, err)
			return
		}
		output, err := transfor(d)
		if err != nil {
			respError(w, err)
			return
		}
		w.Write(output.Json())
	})

	data = make(schema.Schema, len(src))
	for _, s := range src {
		data.Set(s, "${"+s+"}")
	}

	return
}

func ServeQuery(mux *http.ServeMux, handler func(m map[string]string) error, queries ...string) string {
	route := getRoute()
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		m := make(map[string]string, len(query))
		for _, q := range queries {
			m[q] = query.Get(q)
		}
		err := handler(m)
		if err != nil {
			respError(w, err)
			return
		}
		resp := schema.Response{}
		w.Write(resp.Json())
	})
	return route
}
