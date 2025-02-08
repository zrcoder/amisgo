package comp

import "github.com/zrcoder/amisgo/model"

// WangEditor
type WangEditor model.Schema

func NewWangEditor() WangEditor {
	return WangEditor{"type": "custom-wang-editor"}
}

func (we WangEditor) set(key string, value any) WangEditor {
	we[key] = value
	return we
}

// AutoFill sets autoFill
func (we WangEditor) AutoFill(value string) WangEditor {
	return we.set("autoFill", value)
}

// AutoFocus sets autoFocus
func (we WangEditor) AutoFocus(value bool) WangEditor {
	return we.set("autoFocus", value)
}

// ClassName sets className
func (we WangEditor) ClassName(value string) WangEditor {
	return we.set("className", value)
}

// ClearValueOnHidden sets clearValueOnHidden
func (we WangEditor) ClearValueOnHidden(value bool) WangEditor {
	return we.set("clearValueOnHidden", value)
}

// Description sets description
func (we WangEditor) Description(value string) WangEditor {
	return we.set("description", value)
}

// Disabled sets disabled
func (we WangEditor) Disabled(value bool) WangEditor {
	return we.set("disabled", value)
}

// DisabledOn sets disabledOn condition
func (we WangEditor) DisabledOn(value string) WangEditor {
	return we.set("disabledOn", value)
}

// ExcludeKeys sets excludeKeys
func (we WangEditor) ExcludeKeys(value string) WangEditor {
	return we.set("excludeKeys", value)
}

// Height sets editor height in px
func (we WangEditor) Height(value string) WangEditor {
	return we.set("height", value)
}

// InsertKeys sets insertKeys
func (we WangEditor) InsertKeys(value string) WangEditor {
	return we.set("insertKeys", value)
}

// Label sets label
func (we WangEditor) Label(value string) WangEditor {
	return we.set("label", value)
}

// LabelAlign sets label alignment, effective only in horizontal mode
func (we WangEditor) LabelAlign(value string) WangEditor {
	return we.set("labelAlign", value)
}

// LabelRemark sets label remark
func (we WangEditor) LabelRemark(value string) WangEditor {
	return we.set("labelRemark", value)
}

// MaxLength sets maxLength
func (we WangEditor) MaxLength(value string) WangEditor {
	return we.set("maxLength", value)
}

// Name sets name
func (we WangEditor) Name(value string) WangEditor {
	return we.set("name", value)
}

// Placeholder sets placeholder
func (we WangEditor) Placeholder(value string) WangEditor {
	return we.set("placeholder", value)
}

// Remark sets remark
func (we WangEditor) Remark(value string) WangEditor {
	return we.set("remark", value)
}

// Required sets required
func (we WangEditor) Required(value bool) WangEditor {
	return we.set("required", value)
}

// RequiredOn sets requiredOn condition
func (we WangEditor) RequiredOn(value string) WangEditor {
	return we.set("requiredOn", value)
}

// Static sets static display
func (we WangEditor) Static(value bool) WangEditor {
	return we.set("static", value)
}

// ToolbarKeys sets toolbarKeys
func (we WangEditor) ToolbarKeys(value string) WangEditor {
	return we.set("toolbarKeys", value)
}

// UploadImageMaxNumber sets max number of images to upload, default is 100
func (we WangEditor) UploadImageMaxNumber(value string) WangEditor {
	return we.set("uploadImageMaxNumber", value)
}

// UploadImageMaxSize sets max size of images to upload in KB, default is 2MB
func (we WangEditor) UploadImageMaxSize(value string) WangEditor {
	return we.set("uploadImageMaxSize", value)
}

// UploadImageServer sets image upload server URL
func (we WangEditor) UploadImageServer(value string) WangEditor {
	return we.set("uploadImageServer", value)
}

// UploadVideoMaxNumber sets max number of videos to upload, default is 5
func (we WangEditor) UploadVideoMaxNumber(value string) WangEditor {
	return we.set("uploadVideoMaxNumber", value)
}

// UploadVideoMaxSize sets max size of videos to upload in KB, default is 10MB
func (we WangEditor) UploadVideoMaxSize(value string) WangEditor {
	return we.set("uploadVideoMaxSize", value)
}

// UploadVideoServer sets video upload server URL
func (we WangEditor) UploadVideoServer(value string) WangEditor {
	return we.set("uploadVideoServer", value)
}

// ValidateApi sets validation API
func (we WangEditor) ValidateApi(value string) WangEditor {
	return we.set("validateApi", value)
}

// ValidationErrors sets validation errors
func (we WangEditor) ValidationErrors(value string) WangEditor {
	return we.set("validationErrors", value)
}

// Validations sets validation rules
func (we WangEditor) Validations(value string) WangEditor {
	return we.set("validations", value)
}

// Value sets default value
func (we WangEditor) Value(value string) WangEditor {
	return we.set("value", value)
}

// Visible sets visibility
func (we WangEditor) Visible(value bool) WangEditor {
	return we.set("visible", value)
}

// VisibleOn sets visibility condition
func (we WangEditor) VisibleOn(value string) WangEditor {
	return we.set("visibleOn", value)
}
