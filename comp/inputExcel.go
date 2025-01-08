package comp

import "github.com/zrcoder/amisgo/model"

// inputExcel parses Excel. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/excel
type inputExcel model.Schema

// InputExcel creates a new InputExcel instance with default type
func InputExcel() inputExcel {
	return make(inputExcel).set("type", "input-excel")
}

// set sets a field value and returns the instance
func (i inputExcel) set(key string, value any) inputExcel {
	i[key] = value
	return i
}

// AllSheets sets whether to parse all sheets
func (i inputExcel) AllSheets(value bool) inputExcel {
	return i.set("allSheets", value)
}

// AutoFill sets data entry configuration
func (i inputExcel) AutoFill(value string) inputExcel {
	return i.set("autoFill", value)
}

// ClassName sets the outermost class name of the form
func (i inputExcel) ClassName(value string) inputExcel {
	return i.set("className", value)
}

// Description sets the form item description
func (i inputExcel) Description(value string) inputExcel {
	return i.set("description", value)
}

// Disabled sets whether the form item is disabled
func (i inputExcel) Disabled(value bool) inputExcel {
	return i.set("disabled", value)
}

// DisabledOn sets the condition for disabling the form item
func (i inputExcel) DisabledOn(value string) inputExcel {
	return i.set("disabledOn", value)
}

// IncludeEmpty sets whether to include empty sheets
func (i inputExcel) IncludeEmpty(value bool) inputExcel {
	return i.set("includeEmpty", value)
}

// Inline sets whether the form item is inline
func (i inputExcel) Inline(value bool) inputExcel {
	return i.set("inline", value)
}

// InputClassName sets the class name of the form controller
func (i inputExcel) InputClassName(value string) inputExcel {
	return i.set("inputClassName", value)
}

// Label sets the form item label
func (i inputExcel) Label(value string) inputExcel {
	return i.set("label", value)
}

// LabelAlign sets the label alignment, effective only when mode is horizontal
func (i inputExcel) LabelAlign(value string) inputExcel {
	return i.set("labelAlign", value)
}

// LabelClassName sets the class name of the label
func (i inputExcel) LabelClassName(value string) inputExcel {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i inputExcel) LabelRemark(value string) inputExcel {
	return i.set("labelRemark", value)
}

// Name sets the field name for form submission
func (i inputExcel) Name(value string) inputExcel {
	return i.set("name", value)
}

// ParseMode sets the parsing mode
func (i inputExcel) ParseMode(value string) inputExcel {
	return i.set("parseMode", value)
}

// Placeholder sets the placeholder text
func (i inputExcel) Placeholder(value string) inputExcel {
	return i.set("placeholder", value)
}

// PlainText sets whether to parse as plain text
func (i inputExcel) PlainText(value bool) inputExcel {
	return i.set("plainText", value)
}

// Required sets whether the form item is required
func (i inputExcel) Required(value bool) inputExcel {
	return i.set("required", value)
}

// RequiredOn sets the condition for the form item to be required
func (i inputExcel) RequiredOn(value string) inputExcel {
	return i.set("requiredOn", value)
}

// SubmitOnChange sets whether to submit the form when the value changes
func (i inputExcel) SubmitOnChange(value bool) inputExcel {
	return i.set("submitOnChange", value)
}

// ValidateApi sets the validation API
func (i inputExcel) ValidateApi(value string) inputExcel {
	return i.set("validateApi", value)
}

// Validations sets the validation rules, multiple rules separated by commas
func (i inputExcel) Validations(value string) inputExcel {
	return i.set("validations", value)
}

// Value sets the default value of the form item
func (i inputExcel) Value(value string) inputExcel {
	return i.set("value", value)
}

// Visible sets whether the form item is visible
func (i inputExcel) Visible(value bool) inputExcel {
	return i.set("visible", value)
}

// VisibleOn sets the condition for the form item to be visible
func (i inputExcel) VisibleOn(value string) inputExcel {
	return i.set("visibleOn", value)
}
