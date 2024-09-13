package comp

type inputKV schema

// InputKV 创建一个新的 inputKV 实例，并设置默认的 type
func InputKV() inputKV {
	return make(inputKV).set("type", "input-kv")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputKV) set(key string, value any) inputKV {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i inputKV) AutoFill(value string) inputKV {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i inputKV) ClassName(value string) inputKV {
	return i.set("className", value)
}

// DefaultValue 默认值
func (i inputKV) DefaultValue(value string) inputKV {
	return i.set("defaultValue", value)
}

// Description 表单项描述
func (i inputKV) Description(value string) inputKV {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i inputKV) Disabled(value bool) inputKV {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i inputKV) DisabledOn(value string) inputKV {
	return i.set("disabledOn", value)
}

// Draggable 是否可以拖拽排序
func (i inputKV) Draggable(value bool) inputKV {
	return i.set("draggable", value)
}

// Inline 是否内联
func (i inputKV) Inline(value bool) inputKV {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i inputKV) InputClassName(value string) inputKV {
	return i.set("inputClassName", value)
}

// KeyPlaceholder key 的提示信息
func (i inputKV) KeyPlaceholder(value string) inputKV {
	return i.set("keyPlaceholder", value)
}

// Label 表单项标签
func (i inputKV) Label(value string) inputKV {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (i inputKV) LabelAlign(value string) inputKV {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i inputKV) LabelClassName(value string) inputKV {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i inputKV) LabelRemark(value string) inputKV {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i inputKV) Name(value string) inputKV {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i inputKV) Placeholder(value string) inputKV {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i inputKV) Required(value bool) inputKV {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i inputKV) RequiredOn(value string) inputKV {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i inputKV) SubmitOnChange(value bool) inputKV {
	return i.set("submitOnChange", value)
}

// ValidateApi 表单校验接口
func (i inputKV) ValidateApi(value string) inputKV {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputKV) Validations(value string) inputKV {
	return i.set("validations", value)
}

// Value 表单默认值
func (i inputKV) Value(value any) inputKV {
	return i.set("value", value)
}

// ValuePlaceholder value 的提示信息
func (i inputKV) ValuePlaceholder(value string) inputKV {
	return i.set("valuePlaceholder", value)
}

// ValueType 值类型,  如 input-number
func (i inputKV) ValueType(value string) inputKV {
	return i.set("valueType", value)
}

// Visible 是否可见
func (i inputKV) Visible(value bool) inputKV {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i inputKV) VisibleOn(value string) inputKV {
	return i.set("visibleOn", value)
}
