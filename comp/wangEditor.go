package comp

import "github.com/zrcoder/amisgo/model"

// Replace with the actual upload package path

// wangEditor

type wangEditor model.Schema

// WangEditor creates a new WangEditor instance
func WangEditor() wangEditor {
	return wangEditor{}.set("type", "custom-wang-editor")
}

func (we wangEditor) set(key string, value any) wangEditor {
	we[key] = value
	return we
}

// AutoFill sets autoFill
func (we wangEditor) AutoFill(value string) wangEditor {
	return we.set("autoFill", value)
}

// AutoFocus sets autoFocus
func (we wangEditor) AutoFocus(value bool) wangEditor {
	return we.set("autoFocus", value)
}

// ClassName sets className
func (we wangEditor) ClassName(value string) wangEditor {
	return we.set("className", value)
}

// ClearValueOnHidden sets clearValueOnHidden
func (we wangEditor) ClearValueOnHidden(value bool) wangEditor {
	return we.set("clearValueOnHidden", value)
}

// Description sets description
func (we wangEditor) Description(value string) wangEditor {
	return we.set("description", value)
}

// Disabled sets disabled
func (we wangEditor) Disabled(value bool) wangEditor {
	return we.set("disabled", value)
}

// DisabledOn sets disabledOn condition
func (we wangEditor) DisabledOn(value string) wangEditor {
	return we.set("disabledOn", value)
}

// ExcludeKeys sets excludeKeys
func (we wangEditor) ExcludeKeys(value string) wangEditor {
	return we.set("excludeKeys", value)
}

// Height sets editor height in px
func (we wangEditor) Height(value string) wangEditor {
	return we.set("height", value)
}

// InsertKeys sets insertKeys
func (we wangEditor) InsertKeys(value string) wangEditor {
	return we.set("insertKeys", value)
}

// Label sets label
func (we wangEditor) Label(value string) wangEditor {
	return we.set("label", value)
}

// LabelAlign sets label alignment, effective only in horizontal mode
func (we wangEditor) LabelAlign(value string) wangEditor {
	return we.set("labelAlign", value)
}

// LabelRemark sets label remark
func (we wangEditor) LabelRemark(value string) wangEditor {
	return we.set("labelRemark", value)
}

// MaxLength sets maxLength
func (we wangEditor) MaxLength(value string) wangEditor {
	return we.set("maxLength", value)
}

// Name sets name
func (we wangEditor) Name(value string) wangEditor {
	return we.set("name", value)
}

// Placeholder sets placeholder
func (we wangEditor) Placeholder(value string) wangEditor {
	return we.set("placeholder", value)
}

// Remark sets remark
func (we wangEditor) Remark(value string) wangEditor {
	return we.set("remark", value)
}

// Required sets required
func (we wangEditor) Required(value bool) wangEditor {
	return we.set("required", value)
}

// RequiredOn sets requiredOn condition
func (we wangEditor) RequiredOn(value string) wangEditor {
	return we.set("requiredOn", value)
}

// Static sets static display
func (we wangEditor) Static(value bool) wangEditor {
	return we.set("static", value)
}

// ToolbarKeys sets toolbarKeys
func (we wangEditor) ToolbarKeys(value string) wangEditor {
	return we.set("toolbarKeys", value)
}

// UploadImageMaxNumber sets max number of images to upload, default is 100
func (we wangEditor) UploadImageMaxNumber(value string) wangEditor {
	return we.set("uploadImageMaxNumber", value)
}

// UploadImageMaxSize sets max size of images to upload in KB, default is 2MB
func (we wangEditor) UploadImageMaxSize(value string) wangEditor {
	return we.set("uploadImageMaxSize", value)
}

// UploadImageServer sets image upload server URL
func (we wangEditor) UploadImageServer(value string) wangEditor {
	return we.set("uploadImageServer", value)
}

// UploadVideoMaxNumber sets max number of videos to upload, default is 5
func (we wangEditor) UploadVideoMaxNumber(value string) wangEditor {
	return we.set("uploadVideoMaxNumber", value)
}

// UploadVideoMaxSize sets max size of videos to upload in KB, default is 10MB
func (we wangEditor) UploadVideoMaxSize(value string) wangEditor {
	return we.set("uploadVideoMaxSize", value)
}

// UploadVideoServer sets video upload server URL
func (we wangEditor) UploadVideoServer(value string) wangEditor {
	return we.set("uploadVideoServer", value)
}

// ValidateApi sets validation API
func (we wangEditor) ValidateApi(value string) wangEditor {
	return we.set("validateApi", value)
}

// ValidationErrors sets validation errors
func (we wangEditor) ValidationErrors(value string) wangEditor {
	return we.set("validationErrors", value)
}

// Validations sets validation rules
func (we wangEditor) Validations(value string) wangEditor {
	return we.set("validations", value)
}

// Value sets default value
func (we wangEditor) Value(value string) wangEditor {
	return we.set("value", value)
}

// Visible sets visibility
func (we wangEditor) Visible(value bool) wangEditor {
	return we.set("visible", value)
}

// VisibleOn sets visibility condition
func (we wangEditor) VisibleOn(value string) wangEditor {
	return we.set("visibleOn", value)
}
