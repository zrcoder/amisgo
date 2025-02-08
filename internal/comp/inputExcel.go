package comp

import "github.com/zrcoder/amisgo/model"

// InputExcel parses Excel. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/excel
type InputExcel model.Schema

// NewInputExcel creates a new NewInputExcel instance with default type
func NewInputExcel() InputExcel {
	return InputExcel{"type": "input-excel"}
}

// set sets a field value and returns the instance
func (i InputExcel) set(key string, value any) InputExcel {
	i[key] = value
	return i
}

// AllSheets sets whether to parse all sheets
func (i InputExcel) AllSheets(value bool) InputExcel {
	return i.set("allSheets", value)
}

// AutoFill sets data entry configuration
func (i InputExcel) AutoFill(value string) InputExcel {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form
func (i InputExcel) ClassName(value string) InputExcel {
	return i.set("className", value)
}

// Description sets the form item description
func (i InputExcel) Description(value string) InputExcel {
	return i.set("description", value)
}

// Disabled sets whether the form item is disabled
func (i InputExcel) Disabled(value bool) InputExcel {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the form item
func (i InputExcel) DisabledOn(value string) InputExcel {
	return i.set("disabledOn", value)
}

// IncludeEmpty sets whether to include empty sheets
func (i InputExcel) IncludeEmpty(value bool) InputExcel {
	return i.set("includeEmpty", value)
}

// Inline sets whether the form item is inline
func (i InputExcel) Inline(value bool) InputExcel {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller
func (i InputExcel) InputClassName(value string) InputExcel {
	return i.set("inputClassName", value)
}

// Label sets the form item label
func (i InputExcel) Label(value string) InputExcel {
	return i.set("label", value)
}

// LabelAlign sets the label alignment, effective only when mode is horizontal
func (i InputExcel) LabelAlign(value string) InputExcel {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label
func (i InputExcel) LabelClassName(value string) InputExcel {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i InputExcel) LabelRemark(value string) InputExcel {
	return i.set("labelRemark", value)
}

// Name sets the field name for form submission
func (i InputExcel) Name(value string) InputExcel {
	return i.set("name", value)
}

// ParseMode sets the parsing mode
func (i InputExcel) ParseMode(value string) InputExcel {
	return i.set("parseMode", value)
}

// Placeholder sets the placeholder text
func (i InputExcel) Placeholder(value string) InputExcel {
	return i.set("placeholder", value)
}

// PlainText sets whether to parse as plain text
func (i InputExcel) PlainText(value bool) InputExcel {
	return i.set("plainText", value)
}

// Required sets whether the form item is required
func (i InputExcel) Required(value bool) InputExcel {
	return i.set("required", value)
}

// RequiredOn sets the condition for the form item to be required
func (i InputExcel) RequiredOn(value string) InputExcel {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the value changes
func (i InputExcel) SubmitOnChange(value bool) InputExcel {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (i InputExcel) ValidateApi(value string) InputExcel {
	return i.set("validateApi", value)
}

// Validations sets the validation rules, multiple rules separated by commas
func (i InputExcel) Validations(value string) InputExcel {
	return i.set("validations", value)
}

// Value sets the default value of the form item
func (i InputExcel) Value(value string) InputExcel {
	return i.set("value", value)
}

// Visible sets whether the form item is visible
func (i InputExcel) Visible(value bool) InputExcel {
	return i.set("visible", value)
}

// VisibleOn sets the condition for the form item to be visible
func (i InputExcel) VisibleOn(value string) InputExcel {
	return i.set("visibleOn", value)
}
