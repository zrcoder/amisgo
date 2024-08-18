package comp

// option 选项
//
// @version 6.7.0
type option schema

// Option 创建一个新的 Option 实例
func Option() option {
	return option{}
}

// set 设置字段值
func (o option) set(key string, value interface{}) option {
	o[key] = value
	return o
}

// Children 支持嵌套
func (o option) Children(value string) option {
	return o.set("children", value)
}

// Defer 标记后数据延时加载
func (o option) Defer(value bool) option {
	return o.set("defer", value)
}

// DeferApi 如果设置了，优先级更高，不设置走 source 接口加载
func (o option) DeferApi(value string) option {
	return o.set("deferApi", value)
}

// Description 描述，部分控件支持
func (o option) Description(value string) option {
	return o.set("description", value)
}

// Disabled 是否禁用
func (o option) Disabled(value bool) option {
	return o.set("disabled", value)
}

// Hidden 最好不要用！因为有 visible 就够了
func (o option) Hidden(value bool) option {
	return o.set("hidden", value)
}

// Label 用来显示的文字
func (o option) Label(value string) option {
	return o.set("label", value)
}

// Loaded 只有设置了 defer 才有意义，内部字段不可以外部设置
func (o option) Loaded(value bool) option {
	return o.set("loaded", value)
}

// Loading 标记正在加载。只有 defer 为 true 时有意义。内部字段不可以外部设置
func (o option) Loading(value bool) option {
	return o.set("loading", value)
}

// ScopeLabel 可以用来给 Option 标记个范围，让数据展示更清晰。这个只有在数值展示的时候显示
func (o option) ScopeLabel(value string) option {
	return o.set("scopeLabel", value)
}

// Value 请保证数值唯一，多个选项值一致会认为是同一个选项
func (o option) Value(value string) option {
	return o.set("value", value)
}

// Visible 是否可见
func (o option) Visible(value bool) option {
	return o.set("visible", value)
}
