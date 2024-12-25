package comp

// options 选择器表单项

type options schema

// Options 创建一个新的 Options 实例
func Options() options {
	return options{}
}

// set 设置字段值
func (o options) set(key string, value any) options {
	o[key] = value
	return o
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (o options) AutoFill(value string) options {
	return o.set("autoFill", value)
}

// ClassName 表单最外层类名
func (o options) ClassName(value string) options {
	return o.set("className", value)
}

// Description 表单项描述
func (o options) Description(value string) options {
	return o.set("description", value)
}

// Disabled 是否禁用
func (o options) Disabled(value bool) options {
	return o.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (o options) DisabledOn(value string) options {
	return o.set("disabledOn", value)
}

// ExtractValue 是否把选中的值从数组中提取出来，只有当 joinValues 为 true 时生效
func (o options) ExtractValue(value bool) options {
	return o.set("extractValue", value)
}

// Inline 是否内联
func (o options) Inline(value bool) options {
	return o.set("inline", value)
}

// InputClassName 表单控制器类名
func (o options) InputClassName(value string) options {
	return o.set("inputClassName", value)
}

// ItemHeight 每个选项的高度，用于虚拟渲染
func (o options) ItemHeight(value string) options {
	return o.set("itemHeight", value)
}

// JoinValues 多选时，是否把选中的值用逗号拼接成字符串
func (o options) JoinValues(value bool) options {
	return o.set("joinValues", value)
}

// Label 表单项标签
func (o options) Label(value string) options {
	return o.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (o options) LabelAlign(value string) options {
	return o.set("labelAlign", value)
}

// LabelClassName label 的类名
func (o options) LabelClassName(value string) options {
	return o.set("labelClassName", value)
}

// LabelField 标识选项中哪个字段是 label 值
func (o options) LabelField(value string) options {
	return o.set("labelField", value)
}

// LabelRemark 表单项标签描述
func (o options) LabelRemark(value string) options {
	return o.set("labelRemark", value)
}

// Multiple 是否支持多选
func (o options) Multiple(value bool) options {
	return o.set("multiple", value)
}

// Name 字段名，指定该表单项提交时的 key
func (o options) Name(value string) options {
	return o.set("name", value)
}

// Options 选项组，供用户选择
func (o options) Options(value ...MOption) options {
	return o.set("options", value)
}

// Required 是否必填
func (o options) Required(value bool) options {
	return o.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (o options) RequiredOn(value string) options {
	return o.set("requiredOn", value)
}

// Source 选项组源，可通过数据映射获取当前数据域变量、或者配置 API 对象
func (o options) Source(value string) options {
	return o.set("source", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (o options) SubmitOnChange(value bool) options {
	return o.set("submitOnChange", value)
}

// ValidateApi 表单校验接口
func (o options) ValidateApi(value string) options {
	return o.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (o options) Validations(value string) options {
	return o.set("validations", value)
}

// Value 表单默认值
func (o options) Value(value string) options {
	return o.set("value", value)
}

// ValueField 标识选项中哪个字段是 value 值
func (o options) ValueField(value string) options {
	return o.set("valueField", value)
}

// ValuesNoWrap 默认情况下多选所有选项都会显示，通过这个可以最多显示一行，超出的部分变成 ...
func (o options) ValuesNoWrap(value bool) options {
	return o.set("valuesNoWrap", value)
}

// VirtualThreshold 在选项数量超过多少时开启虚拟渲染
func (o options) VirtualThreshold(value string) options {
	return o.set("virtualThreshold", value)
}

// Visible 是否可见
func (o options) Visible(value bool) options {
	return o.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (o options) VisibleOn(value string) options {
	return o.set("visibleOn", value)
}
