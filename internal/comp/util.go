package comp

import (
	"encoding/json"

	"github.com/zrcoder/amisgo/internal/servemux"
)

func pureAmis(key string, value any, m map[string]any) {
	if key == servemux.Key {
		return
	}
	m[key] = value
}

func marshalMap(m map[string]any) ([]byte, error) {
	return json.Marshal(m)
}
