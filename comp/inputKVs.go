package comp

// inputKVS 键值对象
//
// @version 6.7.0
type inputKVS schema

// InputKVS 创建一个新的 InputKVS 实例，并设置默认的 type
func InputKVS() inputKVS {
	return make(inputKVS).set("type", "input-kvs")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputKVS) set(key string, value interface{}) inputKVS {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i inputKVS) AutoFill(value string) inputKVS {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i inputKVS) ClassName(value string) inputKVS {
	return i.set("className", value)
}

// Description 表单项描述
func (i inputKVS) Description(value string) inputKVS {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i inputKVS) Disabled(value bool) inputKVS {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i inputKVS) DisabledOn(value string) inputKVS {
	return i.set("disabledOn", value)
}

// Inline 是否内联
func (i inputKVS) Inline(value bool) inputKVS {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i inputKVS) InputClassName(value string) inputKVS {
	return i.set("inputClassName", value)
}

// Label 表单项标签
func (i inputKVS) Label(value string) inputKVS {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (i inputKVS) LabelAlign(value string) inputKVS {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i inputKVS) LabelClassName(value string) inputKVS {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i inputKVS) LabelRemark(value string) inputKVS {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i inputKVS) Name(value string) inputKVS {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i inputKVS) Placeholder(value string) inputKVS {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i inputKVS) Required(value bool) inputKVS {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i inputKVS) RequiredOn(value string) inputKVS {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i inputKVS) SubmitOnChange(value bool) inputKVS {
	return i.set("submitOnChange", value)
}

// ValidateApi 表单校验接口
func (i inputKVS) ValidateApi(value string) inputKVS {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputKVS) Validations(value string) inputKVS {
	return i.set("validations", value)
}

// Value 表单默认值
func (i inputKVS) Value(value string) inputKVS {
	return i.set("value", value)
}

// Visible 是否可见
func (i inputKVS) Visible(value bool) inputKVS {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i inputKVS) VisibleOn(value string) inputKVS {
	return i.set("visibleOn", value)
}
