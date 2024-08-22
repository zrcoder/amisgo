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

func serveApi(action func(map[string]any)) string {
	route := fmt.Sprintf("/__amisgo_inner_%d", getInnerApiID())
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		input, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		m := map[string]any{}
		err = js.Unmarshal(input, &m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		action(m)
	})
	return route
}
