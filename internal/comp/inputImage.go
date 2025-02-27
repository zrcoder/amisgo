package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
)

// InputImage represents an image upload component
type InputImage schema.Schema

func NewInputImage(mux *http.ServeMux) InputImage {
	return InputImage{"type": "input-image", servemux.Key: mux}
}

func (i InputImage) set(key string, value any) InputImage {
	i[key] = value
	return i
}

// Accept sets the accepted image types
func (i InputImage) Accept(value string) InputImage {
	return i.set("accept", value)
}

// AllowInput allows users to input image URLs
func (i InputImage) AllowInput(value bool) InputImage {
	return i.set("allowInput", value)
}

// AutoFill syncs other fields to the form after upload
func (i InputImage) AutoFill(value string) InputImage {
	return i.set("autoFill", value)
}

// AutoUpload starts upload automatically
func (i InputImage) AutoUpload(value bool) InputImage {
	return i.set("autoUpload", value)
}

// BtnClassName sets the CSS class for the select image button
func (i InputImage) BtnClassName(value string) InputImage {
	return i.set("btnClassName", value)
}

// BtnUploadClassName sets the CSS class for the upload button
func (i InputImage) BtnUploadClassName(value string) InputImage {
	return i.set("btnUploadClassName", value)
}

// Capture configures mobile capture functionality
func (i InputImage) Capture(value string) InputImage {
	return i.set("capture", value)
}

// ClassName sets the container CSS class
func (i InputImage) ClassName(value any) InputImage {
	return i.set("className", value)
}

// ClearValueOnHidden removes the value when the form item is hidden
func (i InputImage) ClearValueOnHidden(value bool) InputImage {
	return i.set("clearValueOnHidden", value)
}

// Compress enables compression
func (i InputImage) Compress(value bool) InputImage {
	return i.set("compress", value)
}

// CompressOptions sets compression options
func (i InputImage) CompressOptions(value ...any) InputImage {
	return i.set("compressOptions", value)
}

// Crop sets crop options
func (i InputImage) Crop(value any) InputImage {
	return i.set("crop", value)
}

// CropFormat sets the format of the cropped image
func (i InputImage) CropFormat(value string) InputImage {
	return i.set("cropFormat", value)
}

// CropQuality sets the quality of the cropped image
func (i InputImage) CropQuality(value string) InputImage {
	return i.set("cropQuality", value)
}

// Delimiter sets the delimiter
func (i InputImage) Delimiter(value string) InputImage {
	return i.set("delimiter", value)
}

// Desc sets the description
func (i InputImage) Desc(value string) InputImage {
	return i.set("desc", value)
}

// Description sets the HTML description
func (i InputImage) Description(value string) InputImage {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class for the description
func (i InputImage) DescriptionClassName(value string) InputImage {
	return i.set("descriptionClassName", value)
}

// Disabled disables the component
func (i InputImage) Disabled(value bool) InputImage {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (i InputImage) DisabledOn(value string) InputImage {
	return i.set("disabledOn", value)
}

// Draggable enables drag-and-drop sorting
func (i InputImage) Draggable(value bool) InputImage {
	return i.set("draggable", value)
}

// DraggableTip sets the drag-and-drop tip
func (i InputImage) DraggableTip(value string) InputImage {
	return i.set("draggableTip", value)
}

// DropCrop enables crop mode after upload
func (i InputImage) DropCrop(value bool) InputImage {
	return i.set("dropCrop", value)
}

// EditorSetting sets the editor configuration
func (i InputImage) EditorSetting(value string) InputImage {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name
func (i InputImage) ExtraName(value string) InputImage {
	return i.set("extraName", value)
}

// ExtractValue wraps the selected option values into an array
func (i InputImage) ExtractValue(value bool) InputImage {
	return i.set("extractValue", value)
}

// FixedSize sets a fixed size
func (i InputImage) FixedSize(value any) InputImage {
	return i.set("fixedSize", value)
}

// FixedSizeClassName sets the CSS class for the fixed size
func (i InputImage) FixedSizeClassName(value string) InputImage {
	return i.set("fixedSizeClassName", value)
}

// FrameImage sets the default placeholder image URL
func (i InputImage) FrameImage(value string) InputImage {
	return i.set("frameImage", value)
}

// Hidden hides the component
func (i InputImage) Hidden(value bool) InputImage {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (i InputImage) HiddenOn(value string) InputImage {
	return i.set("hiddenOn", value)
}

// HideUploadButton hides the upload button
func (i InputImage) HideUploadButton(value bool) InputImage {
	return i.set("hideUploadButton", value)
}

// Hint sets the input hint
func (i InputImage) Hint(value string) InputImage {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (i InputImage) Horizontal(value string) InputImage {
	return i.set("horizontal", value)
}

// ID sets the unique component ID
func (i InputImage) ID(value string) InputImage {
	return i.set("id", value)
}

// ImageClassName sets the CSS class for the displayed image
func (i InputImage) ImageClassName(value string) InputImage {
	return i.set("imageClassName", value)
}

// InitAutoFill syncs other fields to the form on initialization
func (i InputImage) InitAutoFill(value bool) InputImage {
	return i.set("initAutoFill", value)
}

// InitCrop enables crop mode on initialization
func (i InputImage) InitCrop(value bool) InputImage {
	return i.set("initCrop", value)
}

// Inline sets the form control to inline mode
func (i InputImage) Inline(value bool) InputImage {
	return i.set("inline", value)
}

// InputClassName sets the CSS class for the input
func (i InputImage) InputClassName(value string) InputImage {
	return i.set("inputClassName", value)
}

// JoinValues handles the value in single or multiple selection mode
func (i InputImage) JoinValues(value bool) InputImage {
	return i.set("joinValues", value)
}

// Label sets the label
func (i InputImage) Label(value string) InputImage {
	return i.set("label", value)
}

// LabelAlign sets the label alignment
func (i InputImage) LabelAlign(value string) InputImage {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class for the label
func (i InputImage) LabelClassName(value string) InputImage {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i InputImage) LabelRemark(value string) InputImage {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (i InputImage) LabelWidth(value string) InputImage {
	return i.set("labelWidth", value)
}

// Limit sets the image size limit
func (i InputImage) Limit(value any) InputImage {
	return i.set("limit", value)
}

// MaxLength sets the maximum number of images
func (i InputImage) MaxLength(value string) InputImage {
	return i.set("maxLength", value)
}

// MaxSize sets the maximum file size for upload
func (i InputImage) MaxSize(value string) InputImage {
	return i.set("maxSize", value)
}

// Mode sets the form item display mode
func (i InputImage) Mode(value string) InputImage {
	return i.set("mode", value)
}

// Multiple enables multiple selection
func (i InputImage) Multiple(value bool) InputImage {
	return i.set("multiple", value)
}

// Name sets the field name
func (i InputImage) Name(value string) InputImage {
	return i.set("name", value)
}

// OnEvent sets the event action configuration
func (i InputImage) OnEvent(value Event) InputImage {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder
func (i InputImage) Placeholder(value string) InputImage {
	return i.set("placeholder", value)
}

// ReCropable allows re-cropping
func (i InputImage) ReCropable(value bool) InputImage {
	return i.set("reCropable", value)
}

// ReadOnly sets the component to read-only
func (i InputImage) ReadOnly(value bool) InputImage {
	return i.set("readOnly", value)
}

// Receiver sets the image upload URL
func (i InputImage) Receiver(value string) InputImage {
	return i.set("receiver", value)
}

// Upload sets the upload handler
func (i InputImage) Upload(maxMemory int64, action func([]byte) (path string, err error)) InputImage {
	return i.Receiver(servemux.ServeUpload(i.mux(), maxMemory, action))
}

// Required sets the component as required
func (i InputImage) Required(value bool) InputImage {
	return i.set("required", value)
}

// Schema sets the custom Schema
func (i InputImage) Schema(value string) InputImage {
	return i.set("schema", value)
}

// ShowTips shows tips
func (i InputImage) ShowTips(value bool) InputImage {
	return i.set("showTips", value)
}

// Size sets the size
func (i InputImage) Size(value string) InputImage {
	return i.set("size", value)
}

// Src sets the default image URL
func (i InputImage) Src(value string) InputImage {
	return i.set("src", value)
}

// Step sets the incremental selection size for upload
func (i InputImage) Step(value string) InputImage {
	return i.set("step", value)
}

// StrictMode enables strict mode
func (i InputImage) StrictMode(value bool) InputImage {
	return i.set("strictMode", value)
}

// SubmitOnChange submits the form on value change
func (i InputImage) SubmitOnChange(value bool) InputImage {
	return i.set("submitOnChange", value)
}

// UploadType sets the upload type
func (i InputImage) UploadType(value string) InputImage {
	return i.set("uploadType", value)
}

// ValueFieldName sets the field name for the component value
func (i InputImage) ValueFieldName(value string) InputImage {
	return i.set("valueFieldName", value)
}

// Vertical sets the component to vertical mode
func (i InputImage) Vertical(value bool) InputImage {
	return i.set("vertical", value)
}

// VisibleOn sets the expression to show the component
func (i InputImage) VisibleOn(value string) InputImage {
	return i.set("visibleOn", value)
}

// Width sets the width
func (i InputImage) Width(value string) InputImage {
	return i.set("width", value)
}

func (i InputImage) mux() *http.ServeMux {
	return i[servemux.Key].(*http.ServeMux)
}

func (i InputImage) MarshalJSON() ([]byte, error) {
	m := make(map[string]any, len(i))
	for k, v := range i {
		pureAmis(k, v, m)
	}
	return marshalMap(m)
}
