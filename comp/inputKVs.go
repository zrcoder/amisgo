package comp

// InputKVS 键值对象
//
// @version 6.7.0
type InputKVS Schema

// NewInputKVS 创建一个新的 InputKVS 实例，并设置默认的 type
func NewInputKVS() InputKVS {
	return make(InputKVS).set("type", "input-kvs")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i InputKVS) set(key string, value interface{}) InputKVS {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i InputKVS) AutoFill(value string) InputKVS {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i InputKVS) ClassName(value string) InputKVS {
	return i.set("className", value)
}

// Description 表单项描述
func (i InputKVS) Description(value string) InputKVS {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i InputKVS) Disabled(value bool) InputKVS {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i InputKVS) DisabledOn(value string) InputKVS {
	return i.set("disabledOn", value)
}

// Inline 是否内联
func (i InputKVS) Inline(value bool) InputKVS {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i InputKVS) InputClassName(value string) InputKVS {
	return i.set("inputClassName", value)
}

// Label 表单项标签
func (i InputKVS) Label(value string) InputKVS {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (i InputKVS) LabelAlign(value string) InputKVS {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i InputKVS) LabelClassName(value string) InputKVS {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i InputKVS) LabelRemark(value string) InputKVS {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i InputKVS) Name(value string) InputKVS {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i InputKVS) Placeholder(value string) InputKVS {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i InputKVS) Required(value bool) InputKVS {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i InputKVS) RequiredOn(value string) InputKVS {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i InputKVS) SubmitOnChange(value bool) InputKVS {
	return i.set("submitOnChange", value)
}

// ValidateApi 表单校验接口
func (i InputKVS) ValidateApi(value string) InputKVS {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i InputKVS) Validations(value string) InputKVS {
	return i.set("validations", value)
}

// Value 表单默认值
func (i InputKVS) Value(value string) InputKVS {
	return i.set("value", value)
}

// Visible 是否可见
func (i InputKVS) Visible(value bool) InputKVS {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i InputKVS) VisibleOn(value string) InputKVS {
	return i.set("visibleOn", value)
}
