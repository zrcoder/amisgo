package comp

// InputKV 键值对
//
// @version 6.7.0
type InputKV schema

// NewInputKV 创建一个新的 InputKV 实例，并设置默认的 type
func NewInputKV() InputKV {
	return make(InputKV).set("type", "input-kv")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i InputKV) set(key string, value any) InputKV {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i InputKV) AutoFill(value string) InputKV {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i InputKV) ClassName(value string) InputKV {
	return i.set("className", value)
}

// DefaultValue 默认值
func (i InputKV) DefaultValue(value string) InputKV {
	return i.set("defaultValue", value)
}

// Description 表单项描述
func (i InputKV) Description(value string) InputKV {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i InputKV) Disabled(value bool) InputKV {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i InputKV) DisabledOn(value string) InputKV {
	return i.set("disabledOn", value)
}

// Draggable 是否可以拖拽排序
func (i InputKV) Draggable(value bool) InputKV {
	return i.set("draggable", value)
}

// Inline 是否内联
func (i InputKV) Inline(value bool) InputKV {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i InputKV) InputClassName(value string) InputKV {
	return i.set("inputClassName", value)
}

// KeyPlaceholder key 的提示信息
func (i InputKV) KeyPlaceholder(value string) InputKV {
	return i.set("keyPlaceholder", value)
}

// Label 表单项标签
func (i InputKV) Label(value string) InputKV {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (i InputKV) LabelAlign(value string) InputKV {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i InputKV) LabelClassName(value string) InputKV {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i InputKV) LabelRemark(value string) InputKV {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i InputKV) Name(value string) InputKV {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i InputKV) Placeholder(value string) InputKV {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i InputKV) Required(value bool) InputKV {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i InputKV) RequiredOn(value string) InputKV {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i InputKV) SubmitOnChange(value bool) InputKV {
	return i.set("submitOnChange", value)
}

// ValidateApi 表单校验接口
func (i InputKV) ValidateApi(value string) InputKV {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i InputKV) Validations(value string) InputKV {
	return i.set("validations", value)
}

// Value 表单默认值
func (i InputKV) Value(value string) InputKV {
	return i.set("value", value)
}

// ValuePlaceholder value 的提示信息
func (i InputKV) ValuePlaceholder(value string) InputKV {
	return i.set("valuePlaceholder", value)
}

// ValueType 值类型
func (i InputKV) ValueType(value string) InputKV {
	return i.set("valueType", value)
}

// Visible 是否可见
func (i InputKV) Visible(value bool) InputKV {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i InputKV) VisibleOn(value string) InputKV {
	return i.set("visibleOn", value)
}
