package comp

// inputDatetimeRange 日期时间范围。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/datetime-range
type inputDatetimeRange Schema

// InputDatetimeRange 创建一个新的 InputDatetimeRange 实例，并设置默认的 type
func InputDatetimeRange() inputDatetimeRange {
	return make(inputDatetimeRange).set("type", "input-datetime-range")
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i inputDatetimeRange) AutoFill(value string) inputDatetimeRange {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i inputDatetimeRange) ClassName(value string) inputDatetimeRange {
	return i.set("className", value)
}

// Clearable 是否可清除
func (i inputDatetimeRange) Clearable(value bool) inputDatetimeRange {
	return i.set("clearable", value)
}

// Description 表单项描述
func (i inputDatetimeRange) Description(value string) inputDatetimeRange {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i inputDatetimeRange) Disabled(value bool) inputDatetimeRange {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i inputDatetimeRange) DisabledOn(value string) inputDatetimeRange {
	return i.set("disabledOn", value)
}

// Format 日期时间选择器值格式
func (i inputDatetimeRange) Format(value string) inputDatetimeRange {
	return i.set("format", value)
}

// Inline 是否内联
func (i inputDatetimeRange) Inline(value bool) inputDatetimeRange {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i inputDatetimeRange) InputClassName(value string) inputDatetimeRange {
	return i.set("inputClassName", value)
}

// InputFormat 日期时间选择器显示格式
func (i inputDatetimeRange) InputFormat(value string) inputDatetimeRange {
	return i.set("inputFormat", value)
}

// Label 表单项标签
func (i inputDatetimeRange) Label(value string) inputDatetimeRange {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i inputDatetimeRange) LabelAlign(value string) inputDatetimeRange {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i inputDatetimeRange) LabelClassName(value string) inputDatetimeRange {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i inputDatetimeRange) LabelRemark(value string) inputDatetimeRange {
	return i.set("labelRemark", value)
}

// MaxDate 限制最大日期时间，用法同限制范围
func (i inputDatetimeRange) MaxDate(value string) inputDatetimeRange {
	return i.set("maxDate", value)
}

// MinDate 限制最小日期时间，用法同限制范围
func (i inputDatetimeRange) MinDate(value string) inputDatetimeRange {
	return i.set("minDate", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i inputDatetimeRange) Name(value string) inputDatetimeRange {
	return i.set("name", value)
}

// Placeholder 表单项描述
func (i inputDatetimeRange) Placeholder(value string) inputDatetimeRange {
	return i.set("placeholder", value)
}

// Ranges 日期范围快捷键
func (i inputDatetimeRange) Ranges(value string) inputDatetimeRange {
	return i.set("ranges", value)
}

// Required 是否必填
func (i inputDatetimeRange) Required(value bool) inputDatetimeRange {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i inputDatetimeRange) RequiredOn(value string) inputDatetimeRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i inputDatetimeRange) SubmitOnChange(value bool) inputDatetimeRange {
	return i.set("submitOnChange", value)
}

// Utc 保存 UTC 值
func (i inputDatetimeRange) Utc(value bool) inputDatetimeRange {
	return i.set("utc", value)
}

// ValidateApi 表单校验接口
func (i inputDatetimeRange) ValidateApi(value string) inputDatetimeRange {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputDatetimeRange) Validations(value string) inputDatetimeRange {
	return i.set("validations", value)
}

// Value 表单默认值
func (i inputDatetimeRange) Value(value string) inputDatetimeRange {
	return i.set("value", value)
}

// Visible 是否显示
func (i inputDatetimeRange) Visible(value bool) inputDatetimeRange {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i inputDatetimeRange) VisibleOn(value string) inputDatetimeRange {
	return i.set("visibleOn", value)
}

func (i inputDatetimeRange) set(key string, value any) inputDatetimeRange {
	i[key] = value
	return i
}
