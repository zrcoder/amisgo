package comp

// InputTimeRange 时间范围
//
// @version 6.7.0
type InputTimeRange Schema

// NewInputTimeRange 创建一个新的 InputTimeRange 实例，并设置默认的 type
func NewInputTimeRange() InputTimeRange {
	return make(InputTimeRange).set("type", "input-time-range")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i InputTimeRange) set(key string, value interface{}) InputTimeRange {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i InputTimeRange) AutoFill(value string) InputTimeRange {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i InputTimeRange) ClassName(value string) InputTimeRange {
	return i.set("className", value)
}

// Clearable 是否可清除
func (i InputTimeRange) Clearable(value bool) InputTimeRange {
	return i.set("clearable", value)
}

// Description 表单项描述
func (i InputTimeRange) Description(value string) InputTimeRange {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i InputTimeRange) Disabled(value bool) InputTimeRange {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i InputTimeRange) DisabledOn(value string) InputTimeRange {
	return i.set("disabledOn", value)
}

// Embed 是否内联模式
func (i InputTimeRange) Embed(value bool) InputTimeRange {
	return i.set("embed", value)
}

// Format 时间范围选择器值格式
func (i InputTimeRange) Format(value string) InputTimeRange {
	return i.set("format", value)
}

// Inline 是否内联
func (i InputTimeRange) Inline(value bool) InputTimeRange {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i InputTimeRange) InputClassName(value string) InputTimeRange {
	return i.set("inputClassName", value)
}

// InputFormat 时间范围选择器显示格式
func (i InputTimeRange) InputFormat(value string) InputTimeRange {
	return i.set("inputFormat", value)
}

// Label 表单项标签
func (i InputTimeRange) Label(value string) InputTimeRange {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i InputTimeRange) LabelAlign(value string) InputTimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i InputTimeRange) LabelClassName(value string) InputTimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i InputTimeRange) LabelRemark(value string) InputTimeRange {
	return i.set("labelRemark", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i InputTimeRange) Name(value string) InputTimeRange {
	return i.set("name", value)
}

// Placeholder 占位文本
func (i InputTimeRange) Placeholder(value string) InputTimeRange {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i InputTimeRange) Required(value bool) InputTimeRange {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i InputTimeRange) RequiredOn(value string) InputTimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i InputTimeRange) SubmitOnChange(value bool) InputTimeRange {
	return i.set("submitOnChange", value)
}

// TimeFormat 时间范围选择器值格式
func (i InputTimeRange) TimeFormat(value string) InputTimeRange {
	return i.set("timeFormat", value)
}

// ValidateApi 表单校验接口
func (i InputTimeRange) ValidateApi(value string) InputTimeRange {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i InputTimeRange) Validations(value string) InputTimeRange {
	return i.set("validations", value)
}

// Value 表单默认值
func (i InputTimeRange) Value(value string) InputTimeRange {
	return i.set("value", value)
}

// Visible 是否可见
func (i InputTimeRange) Visible(value bool) InputTimeRange {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i InputTimeRange) VisibleOn(value string) InputTimeRange {
	return i.set("visibleOn", value)
}
