package comp

// inputTimeRange 时间范围
//
// @version 6.7.0
type inputTimeRange schema

// InputTimeRange 创建一个新的 InputTimeRange 实例，并设置默认的 type
func InputTimeRange() inputTimeRange {
	return make(inputTimeRange).set("type", "input-time-range")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputTimeRange) set(key string, value interface{}) inputTimeRange {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i inputTimeRange) AutoFill(value string) inputTimeRange {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i inputTimeRange) ClassName(value string) inputTimeRange {
	return i.set("className", value)
}

// Clearable 是否可清除
func (i inputTimeRange) Clearable(value bool) inputTimeRange {
	return i.set("clearable", value)
}

// Description 表单项描述
func (i inputTimeRange) Description(value string) inputTimeRange {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i inputTimeRange) Disabled(value bool) inputTimeRange {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i inputTimeRange) DisabledOn(value string) inputTimeRange {
	return i.set("disabledOn", value)
}

// Embed 是否内联模式
func (i inputTimeRange) Embed(value bool) inputTimeRange {
	return i.set("embed", value)
}

// Format 时间范围选择器值格式
func (i inputTimeRange) Format(value string) inputTimeRange {
	return i.set("format", value)
}

// Inline 是否内联
func (i inputTimeRange) Inline(value bool) inputTimeRange {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i inputTimeRange) InputClassName(value string) inputTimeRange {
	return i.set("inputClassName", value)
}

// InputFormat 时间范围选择器显示格式
func (i inputTimeRange) InputFormat(value string) inputTimeRange {
	return i.set("inputFormat", value)
}

// Label 表单项标签
func (i inputTimeRange) Label(value string) inputTimeRange {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i inputTimeRange) LabelAlign(value string) inputTimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i inputTimeRange) LabelClassName(value string) inputTimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i inputTimeRange) LabelRemark(value string) inputTimeRange {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i inputTimeRange) Name(value string) inputTimeRange {
	return i.set("name", value)
}

// Placeholder 占位文本
func (i inputTimeRange) Placeholder(value string) inputTimeRange {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i inputTimeRange) Required(value bool) inputTimeRange {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i inputTimeRange) RequiredOn(value string) inputTimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i inputTimeRange) SubmitOnChange(value bool) inputTimeRange {
	return i.set("submitOnChange", value)
}

// TimeFormat 时间范围选择器值格式
func (i inputTimeRange) TimeFormat(value string) inputTimeRange {
	return i.set("timeFormat", value)
}

// ValidateApi 表单校验接口
func (i inputTimeRange) ValidateApi(value string) inputTimeRange {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputTimeRange) Validations(value string) inputTimeRange {
	return i.set("validations", value)
}

// Value 表单默认值
func (i inputTimeRange) Value(value string) inputTimeRange {
	return i.set("value", value)
}

// Visible 是否可见
func (i inputTimeRange) Visible(value bool) inputTimeRange {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i inputTimeRange) VisibleOn(value string) inputTimeRange {
	return i.set("visibleOn", value)
}
