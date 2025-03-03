package comp

import (
	"encoding/json"
	"reflect"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
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

func toSchema(value any) schema.Schema {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Map && v.Type().ConvertibleTo(reflect.TypeOf(schema.Schema{})) {
		converted := v.Convert(reflect.TypeOf(schema.Schema{}))
		return converted.Interface().(schema.Schema)
	}
	return schema.Schema{}
}
