package comp

// inputYearRange 年份范围
//
// @version 6.7.0
type inputYearRange schema

// InputYearRange 创建一个新的 InputYearRange 实例，并设置默认的 type
func InputYearRange() inputYearRange {
	return make(inputYearRange).set("type", "input-year-range")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputYearRange) set(key string, value any) inputYearRange {
	i[key] = value
	return i
}

// AutoFill 数据录入配置，自动填充或者参照录入
func (i inputYearRange) AutoFill(value string) inputYearRange {
	return i.set("autoFill", value)
}

// ClassName 表单最外层类名
func (i inputYearRange) ClassName(value string) inputYearRange {
	return i.set("className", value)
}

// Clearable 是否可清除
func (i inputYearRange) Clearable(value bool) inputYearRange {
	return i.set("clearable", value)
}

// Description 表单项描述
func (i inputYearRange) Description(value string) inputYearRange {
	return i.set("description", value)
}

// Disabled 是否禁用
func (i inputYearRange) Disabled(value bool) inputYearRange {
	return i.set("disabled", value)
}

// DisabledOn 当前表单项是否禁用的条件
func (i inputYearRange) DisabledOn(value string) inputYearRange {
	return i.set("disabledOn", value)
}

// Embed 是否内联模式
func (i inputYearRange) Embed(value bool) inputYearRange {
	return i.set("embed", value)
}

// Format 年份选择器值格式
func (i inputYearRange) Format(value string) inputYearRange {
	return i.set("format", value)
}

// Inline 是否内联
func (i inputYearRange) Inline(value bool) inputYearRange {
	return i.set("inline", value)
}

// InputClassName 表单控制器类名
func (i inputYearRange) InputClassName(value string) inputYearRange {
	return i.set("inputClassName", value)
}

// InputFormat 年份选择器显示格式
func (i inputYearRange) InputFormat(value string) inputYearRange {
	return i.set("inputFormat", value)
}

// Label 表单项标签
func (i inputYearRange) Label(value string) inputYearRange {
	return i.set("label", value)
}

// LabelAlign 表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i inputYearRange) LabelAlign(value string) inputYearRange {
	return i.set("labelAlign", value)
}

// LabelClassName label 的类名
func (i inputYearRange) LabelClassName(value string) inputYearRange {
	return i.set("labelClassName", value)
}

// LabelRemark 表单项标签描述
func (i inputYearRange) LabelRemark(value string) inputYearRange {
	return i.set("labelRemark", value)
}

// MaxDate 限制最大日期，用法同限制范围
func (i inputYearRange) MaxDate(value string) inputYearRange {
	return i.set("maxDate", value)
}

// MaxDuration 限制最大跨度，如：4year
func (i inputYearRange) MaxDuration(value string) inputYearRange {
	return i.set("maxDuration", value)
}

// MinDate 限制最小日期，用法同限制范围
func (i inputYearRange) MinDate(value string) inputYearRange {
	return i.set("minDate", value)
}

// MinDuration 限制最小跨度，如：2year
func (i inputYearRange) MinDuration(value string) inputYearRange {
	return i.set("minDuration", value)
}

// Name 字段名，指定该表单项提交时的 key
func (i inputYearRange) Name(value string) inputYearRange {
	return i.set("name", value)
}

// Placeholder 占位文本
func (i inputYearRange) Placeholder(value string) inputYearRange {
	return i.set("placeholder", value)
}

// Required 是否必填
func (i inputYearRange) Required(value bool) inputYearRange {
	return i.set("required", value)
}

// RequiredOn 通过表达式来配置当前表单项是否为必填
func (i inputYearRange) RequiredOn(value string) inputYearRange {
	return i.set("requiredOn", value)
}

// SubmitOnChange 是否该表单项值发生变化时就提交当前表单
func (i inputYearRange) SubmitOnChange(value bool) inputYearRange {
	return i.set("submitOnChange", value)
}

// Utc 保存 UTC 值
func (i inputYearRange) Utc(value bool) inputYearRange {
	return i.set("utc", value)
}

// ValidateApi 表单校验接口
func (i inputYearRange) ValidateApi(value string) inputYearRange {
	return i.set("validateApi", value)
}

// Validations 表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputYearRange) Validations(value string) inputYearRange {
	return i.set("validations", value)
}

// Value 表单默认值
func (i inputYearRange) Value(value string) inputYearRange {
	return i.set("value", value)
}

// Visible 是否可见
func (i inputYearRange) Visible(value bool) inputYearRange {
	return i.set("visible", value)
}

// VisibleOn 当前表单项是否禁用的条件
func (i inputYearRange) VisibleOn(value string) inputYearRange {
	return i.set("visibleOn", value)
}
