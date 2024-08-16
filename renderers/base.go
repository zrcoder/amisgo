package renderers

import (
	"encoding/json"
	"reflect"
)

// BaseRenderer 基础渲染器
type BaseRenderer struct {
	Type       string                 `json:"type"`
	AmisSchema map[string]interface{} `json:"-"`
}

// NewBaseRenderer 创建一个新的 BaseRenderer 实例
func NewBaseRenderer() *BaseRenderer {
	return &BaseRenderer{
		AmisSchema: make(map[string]interface{}),
	}
}

// Set 设置属性值
func (b *BaseRenderer) set(name string, value interface{}) *BaseRenderer {
	if name == "map" {
		if reflect.TypeOf(value).Kind() == reflect.Map {
			m := value.(map[string]interface{})
			allKeysAreStrings := true
			for k := range m {
				if reflect.TypeOf(k).Kind() != reflect.String {
					allKeysAreStrings = false
					break
				}
			}
			if allKeysAreStrings {
				value = m
			}
		}
	}

	b.AmisSchema[name] = value
	return b
}

// ToJSON 返回 JSON 字符串
func (b *BaseRenderer) ToJSON() string {
	data, _ := json.Marshal(b.AmisSchema)
	return string(data)
}
