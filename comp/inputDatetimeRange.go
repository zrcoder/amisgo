package comp

// InputDatetimeRange 日期时间范围。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/datetime-range
type InputDatetimeRange Schema

// NewInputDatetimeRange 创建一个新的 InputDatetimeRange 实例，并设置默认的 type
func NewInputDatetimeRange() InputDatetimeRange {
	return make(InputDatetimeRange).set("type", "input-datetime-range")
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i InputDatetimeRange) AutoFill(value string) InputDatetimeRange {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i InputDatetimeRange) ClassName(value string) InputDatetimeRange {
	return i.set("className", value)
}

// Clearable 是否可清除
func (i InputDatetimeRange) Clearable(value bool) InputDatetimeRange {
	return i.set("clearable", value)
}

// Description 表单项描述
func (i InputDatetimeRange) Description(value string) InputDatetimeRange {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i InputDatetimeRange) Disabled(value bool) InputDatetimeRange {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i InputDatetimeRange) DisabledOn(value string) InputDatetimeRange {
	return i.set("disabledOn", value)
}

// Format 日期时间选择器值格式
func (i InputDatetimeRange) Format(value string) InputDatetimeRange {
	return i.set("format", value)
}

// Inline 是否内联
func (i InputDatetimeRange) Inline(value bool) InputDatetimeRange {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i InputDatetimeRange) InputClassName(value string) InputDatetimeRange {
	return i.set("inputClassName", value)
}

// InputFormat 日期时间选择器显示格式
func (i InputDatetimeRange) InputFormat(value string) InputDatetimeRange {
	return i.set("inputFormat", value)
}

// Label 表单项标签
func (i InputDatetimeRange) Label(value string) InputDatetimeRange {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i InputDatetimeRange) LabelAlign(value string) InputDatetimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i InputDatetimeRange) LabelClassName(value string) InputDatetimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i InputDatetimeRange) LabelRemark(value string) InputDatetimeRange {
	return i.set("labelRemark", value)
}

// MaxDate 限制最大日期时间，用法同限制范围
func (i InputDatetimeRange) MaxDate(value string) InputDatetimeRange {
	return i.set("maxDate", value)
}

// MinDate 限制最小日期时间，用法同限制范围
func (i InputDatetimeRange) MinDate(value string) InputDatetimeRange {
	return i.set("minDate", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i InputDatetimeRange) Name(value string) InputDatetimeRange {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i InputDatetimeRange) Placeholder(value string) InputDatetimeRange {
	return i.set("placeholder", value)
}

// Ranges 日期范围快捷键
func (i InputDatetimeRange) Ranges(value string) InputDatetimeRange {
	return i.set("ranges", value)
}

// Required 是否必填
func (i InputDatetimeRange) Required(value bool) InputDatetimeRange {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i InputDatetimeRange) RequiredOn(value string) InputDatetimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i InputDatetimeRange) SubmitOnChange(value bool) InputDatetimeRange {
	return i.set("submitOnChange", value)
}

// Utc 保存 UTC 值
func (i InputDatetimeRange) Utc(value bool) InputDatetimeRange {
	return i.set("utc", value)
}

// ValidateApi 表单校验接口
func (i InputDatetimeRange) ValidateApi(value string) InputDatetimeRange {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i InputDatetimeRange) Validations(value string) InputDatetimeRange {
	return i.set("validations", value)
}

// Value 表单默认值
func (i InputDatetimeRange) Value(value string) InputDatetimeRange {
	return i.set("value", value)
}

// Visible 是否显示
func (i InputDatetimeRange) Visible(value bool) InputDatetimeRange {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i InputDatetimeRange) VisibleOn(value string) InputDatetimeRange {
	return i.set("visibleOn", value)
}

func (i InputDatetimeRange) set(key string, value interface{}) InputDatetimeRange {
	i[key] = value
	return i
}
