package comp

// InputExcel 解析 Excel。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/excel
type InputExcel Schema

// NewInputExcel 创建一个新的 InputExcel 实例，并设置默认的 type
func NewInputExcel() InputExcel {
	return make(InputExcel).set("type", "input-excel")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i InputExcel) set(key string, value interface{}) InputExcel {
	i[key] = value
	return i
}

// AllSheets 设置是否解析所有 sheet
func (i InputExcel) AllSheets(value bool) InputExcel {
	return i.set("allSheets", value)
}

// AutoFill 设置数据录入配置，自动填充或者参照录入
func (i InputExcel) AutoFill(value string) InputExcel {
	return i.set("autoFill", value)
}

// ClassName 设置表单最外层类名
func (i InputExcel) ClassName(value string) InputExcel {
	return i.set("className", value)
}

// Description 设置表单项描述
func (i InputExcel) Description(value string) InputExcel {
	return i.set("description", value)
}

// Disabled 设置是否禁用
func (i InputExcel) Disabled(value bool) InputExcel {
	return i.set("disabled", value)
}

// DisabledOn 设置当前表单项是否禁用的条件
func (i InputExcel) DisabledOn(value string) InputExcel {
	return i.set("disabledOn", value)
}

// IncludeEmpty 设置是否解析所有 sheet
func (i InputExcel) IncludeEmpty(value bool) InputExcel {
	return i.set("includeEmpty", value)
}

// Inline 设置是否内联
func (i InputExcel) Inline(value bool) InputExcel {
	return i.set("inline", value)
}

// InputClassName 设置表单控制器类名
func (i InputExcel) InputClassName(value string) InputExcel {
	return i.set("inputClassName", value)
}

// Label 设置表单项标签
func (i InputExcel) Label(value string) InputExcel {
	return i.set("label", value)
}

// LabelAlign 设置表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i InputExcel) LabelAlign(value string) InputExcel {
	return i.set("labelAlign", value)
}

// LabelClassName 设置 label 的类名
func (i InputExcel) LabelClassName(value string) InputExcel {
	return i.set("labelClassName", value)
}

// LabelRemark 设置表单项标签描述
func (i InputExcel) LabelRemark(value string) InputExcel {
	return i.set("labelRemark", value)
}

// Name 设置字段名，指定该表单项提交时的 key
func (i InputExcel) Name(value string) InputExcel {
	return i.set("name", value)
}

// ParseMode 设置解析模式
func (i InputExcel) ParseMode(value string) InputExcel {
	return i.set("parseMode", value)
}

// Placeholder 设置表单项描述
func (i InputExcel) Placeholder(value string) InputExcel {
	return i.set("placeholder", value)
}

// PlainText 设置是否解析为纯文本
func (i InputExcel) PlainText(value bool) InputExcel {
	return i.set("plainText", value)
}

// Required 设置是否必填
func (i InputExcel) Required(value bool) InputExcel {
	return i.set("required", value)
}

// RequiredOn 设置通过表达式来配置当前表单项是否为必填
func (i InputExcel) RequiredOn(value string) InputExcel {
	return i.set("requiredOn", value)
}

// SubmitOnChange 设置是否该表单项值发生变化时就提交当前表单
func (i InputExcel) SubmitOnChange(value bool) InputExcel {
	return i.set("submitOnChange", value)
}

// ValidateApi 设置表单校验接口
func (i InputExcel) ValidateApi(value string) InputExcel {
	return i.set("validateApi", value)
}

// Validations 设置表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i InputExcel) Validations(value string) InputExcel {
	return i.set("validations", value)
}

// Value 设置表单默认值
func (i InputExcel) Value(value string) InputExcel {
	return i.set("value", value)
}

// Visible 设置是否可见
func (i InputExcel) Visible(value bool) InputExcel {
	return i.set("visible", value)
}

// VisibleOn 设置当前表单项是否禁用的条件
func (i InputExcel) VisibleOn(value string) InputExcel {
	return i.set("visibleOn", value)
}
