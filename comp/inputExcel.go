package comp

// inputExcel 解析 Excel。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/excel
type inputExcel schema

// InputExcel 创建一个新的 InputExcel 实例，并设置默认的 type
func InputExcel() inputExcel {
	return make(inputExcel).set("type", "input-excel")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputExcel) set(key string, value any) inputExcel {
	i[key] = value
	return i
}

// AllSheets 设置是否解析所有 sheet
func (i inputExcel) AllSheets(value bool) inputExcel {
	return i.set("allSheets", value)
}

// AutoFill 设置数据录入配置，自动填充或者参照录入
func (i inputExcel) AutoFill(value string) inputExcel {
	return i.set("autoFill", value)
}

// ClassName 设置表单最外层类名
func (i inputExcel) ClassName(value string) inputExcel {
	return i.set("className", value)
}

// Description 设置表单项描述
func (i inputExcel) Description(value string) inputExcel {
	return i.set("description", value)
}

// Disabled 设置是否禁用
func (i inputExcel) Disabled(value bool) inputExcel {
	return i.set("disabled", value)
}

// DisabledOn 设置当前表单项是否禁用的条件
func (i inputExcel) DisabledOn(value string) inputExcel {
	return i.set("disabledOn", value)
}

// IncludeEmpty 设置是否解析所有 sheet
func (i inputExcel) IncludeEmpty(value bool) inputExcel {
	return i.set("includeEmpty", value)
}

// Inline 设置是否内联
func (i inputExcel) Inline(value bool) inputExcel {
	return i.set("inline", value)
}

// InputClassName 设置表单控制器类名
func (i inputExcel) InputClassName(value string) inputExcel {
	return i.set("inputClassName", value)
}

// Label 设置表单项标签
func (i inputExcel) Label(value string) inputExcel {
	return i.set("label", value)
}

// LabelAlign 设置表单项标签对齐方式，默认右对齐，仅在 mode 为 horizontal 时生效
func (i inputExcel) LabelAlign(value string) inputExcel {
	return i.set("labelAlign", value)
}

// LabelClassName 设置 label 的类名
func (i inputExcel) LabelClassName(value string) inputExcel {
	return i.set("labelClassName", value)
}

// LabelRemark 设置表单项标签描述
func (i inputExcel) LabelRemark(value string) inputExcel {
	return i.set("labelRemark", value)
}

// Name 设置字段名，指定该表单项提交时的 key
func (i inputExcel) Name(value string) inputExcel {
	return i.set("name", value)
}

// ParseMode 设置解析模式
func (i inputExcel) ParseMode(value string) inputExcel {
	return i.set("parseMode", value)
}

// Placeholder 设置表单项描述
func (i inputExcel) Placeholder(value string) inputExcel {
	return i.set("placeholder", value)
}

// PlainText 设置是否解析为纯文本
func (i inputExcel) PlainText(value bool) inputExcel {
	return i.set("plainText", value)
}

// Required 设置是否必填
func (i inputExcel) Required(value bool) inputExcel {
	return i.set("required", value)
}

// RequiredOn 设置通过表达式来配置当前表单项是否为必填
func (i inputExcel) RequiredOn(value string) inputExcel {
	return i.set("requiredOn", value)
}

// SubmitOnChange 设置是否该表单项值发生变化时就提交当前表单
func (i inputExcel) SubmitOnChange(value bool) inputExcel {
	return i.set("submitOnChange", value)
}

// ValidateApi 设置表单校验接口
func (i inputExcel) ValidateApi(value string) inputExcel {
	return i.set("validateApi", value)
}

// Validations 设置表单项值格式验证，支持设置多个，多个规则用英文逗号隔开
func (i inputExcel) Validations(value string) inputExcel {
	return i.set("validations", value)
}

// Value 设置表单默认值
func (i inputExcel) Value(value string) inputExcel {
	return i.set("value", value)
}

// Visible 设置是否可见
func (i inputExcel) Visible(value bool) inputExcel {
	return i.set("visible", value)
}

// VisibleOn 设置当前表单项是否禁用的条件
func (i inputExcel) VisibleOn(value string) inputExcel {
	return i.set("visibleOn", value)
}
