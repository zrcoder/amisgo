package comp

// inputImage represents an image upload component
type inputImage Schema

// InputImage creates a new InputImage instance
func InputImage() inputImage {
	return make(inputImage).set("type", "input-image").Receiver("/api/upload")
}

func (i inputImage) set(key string, value any) inputImage {
	i[key] = value
	return i
}

// Accept sets the accepted image types
func (i inputImage) Accept(value string) inputImage {
	return i.set("accept", value)
}

// AllowInput allows users to input image URLs
func (i inputImage) AllowInput(value bool) inputImage {
	return i.set("allowInput", value)
}

// AutoFill syncs other fields to the form after upload
func (i inputImage) AutoFill(value string) inputImage {
	return i.set("autoFill", value)
}

// AutoUpload starts upload automatically
func (i inputImage) AutoUpload(value bool) inputImage {
	return i.set("autoUpload", value)
}

// BtnClassName sets the CSS class for the select image button
func (i inputImage) BtnClassName(value string) inputImage {
	return i.set("btnClassName", value)
}

// BtnUploadClassName sets the CSS class for the upload button
func (i inputImage) BtnUploadClassName(value string) inputImage {
	return i.set("btnUploadClassName", value)
}

// Capture configures mobile capture functionality
func (i inputImage) Capture(value string) inputImage {
	return i.set("capture", value)
}

// ClassName sets the container CSS class
func (i inputImage) ClassName(value any) inputImage {
	return i.set("className", value)
}

// ClearValueOnHidden removes the value when the form item is hidden
func (i inputImage) ClearValueOnHidden(value bool) inputImage {
	return i.set("clearValueOnHidden", value)
}

// Compress enables compression
func (i inputImage) Compress(value bool) inputImage {
	return i.set("compress", value)
}

// CompressOptions sets compression options
func (i inputImage) CompressOptions(value ...any) inputImage {
	return i.set("compressOptions", value)
}

// Crop sets crop options
func (i inputImage) Crop(value any) inputImage {
	return i.set("crop", value)
}

// CropFormat sets the format of the cropped image
func (i inputImage) CropFormat(value string) inputImage {
	return i.set("cropFormat", value)
}

// CropQuality sets the quality of the cropped image
func (i inputImage) CropQuality(value string) inputImage {
	return i.set("cropQuality", value)
}

// Delimiter sets the delimiter
func (i inputImage) Delimiter(value string) inputImage {
	return i.set("delimiter", value)
}

// Desc sets the description
func (i inputImage) Desc(value string) inputImage {
	return i.set("desc", value)
}

// Description sets the HTML description
func (i inputImage) Description(value string) inputImage {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class for the description
func (i inputImage) DescriptionClassName(value string) inputImage {
	return i.set("descriptionClassName", value)
}

// Disabled disables the component
func (i inputImage) Disabled(value bool) inputImage {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the component
func (i inputImage) DisabledOn(value string) inputImage {
	return i.set("disabledOn", value)
}

// Draggable enables drag-and-drop sorting
func (i inputImage) Draggable(value bool) inputImage {
	return i.set("draggable", value)
}

// DraggableTip sets the drag-and-drop tip
func (i inputImage) DraggableTip(value string) inputImage {
	return i.set("draggableTip", value)
}

// DropCrop enables crop mode after upload
func (i inputImage) DropCrop(value bool) inputImage {
	return i.set("dropCrop", value)
}

// EditorSetting sets the editor configuration
func (i inputImage) EditorSetting(value string) inputImage {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name
func (i inputImage) ExtraName(value string) inputImage {
	return i.set("extraName", value)
}

// ExtractValue wraps the selected option values into an array
func (i inputImage) ExtractValue(value bool) inputImage {
	return i.set("extractValue", value)
}

// FixedSize sets a fixed size
func (i inputImage) FixedSize(value any) inputImage {
	return i.set("fixedSize", value)
}

// FixedSizeClassName sets the CSS class for the fixed size
func (i inputImage) FixedSizeClassName(value string) inputImage {
	return i.set("fixedSizeClassName", value)
}

// FrameImage sets the default placeholder image URL
func (i inputImage) FrameImage(value string) inputImage {
	return i.set("frameImage", value)
}

// Hidden hides the component
func (i inputImage) Hidden(value bool) inputImage {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the component
func (i inputImage) HiddenOn(value string) inputImage {
	return i.set("hiddenOn", value)
}

// HideUploadButton hides the upload button
func (i inputImage) HideUploadButton(value bool) inputImage {
	return i.set("hideUploadButton", value)
}

// Hint sets the input hint
func (i inputImage) Hint(value string) inputImage {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (i inputImage) Horizontal(value string) inputImage {
	return i.set("horizontal", value)
}

// ID sets the unique component ID
func (i inputImage) ID(value string) inputImage {
	return i.set("id", value)
}

// ImageClassName sets the CSS class for the displayed image
func (i inputImage) ImageClassName(value string) inputImage {
	return i.set("imageClassName", value)
}

// InitAutoFill syncs other fields to the form on initialization
func (i inputImage) InitAutoFill(value bool) inputImage {
	return i.set("initAutoFill", value)
}

// InitCrop enables crop mode on initialization
func (i inputImage) InitCrop(value bool) inputImage {
	return i.set("initCrop", value)
}

// Inline sets the form control to inline mode
func (i inputImage) Inline(value bool) inputImage {
	return i.set("inline", value)
}

// InputClassName sets the CSS class for the input
func (i inputImage) InputClassName(value string) inputImage {
	return i.set("inputClassName", value)
}

// JoinValues handles the value in single or multiple selection mode
func (i inputImage) JoinValues(value bool) inputImage {
	return i.set("joinValues", value)
}

// Label sets the label
func (i inputImage) Label(value string) inputImage {
	return i.set("label", value)
}

// LabelAlign sets the label alignment
func (i inputImage) LabelAlign(value string) inputImage {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class for the label
func (i inputImage) LabelClassName(value string) inputImage {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (i inputImage) LabelRemark(value string) inputImage {
	return i.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (i inputImage) LabelWidth(value string) inputImage {
	return i.set("labelWidth", value)
}

// Limit sets the image size limit
func (i inputImage) Limit(value any) inputImage {
	return i.set("limit", value)
}

// MaxLength sets the maximum number of images
func (i inputImage) MaxLength(value string) inputImage {
	return i.set("maxLength", value)
}

// MaxSize sets the maximum file size for upload
func (i inputImage) MaxSize(value string) inputImage {
	return i.set("maxSize", value)
}

// Mode sets the form item display mode
func (i inputImage) Mode(value string) inputImage {
	return i.set("mode", value)
}

// Multiple enables multiple selection
func (i inputImage) Multiple(value bool) inputImage {
	return i.set("multiple", value)
}

// Name sets the field name
func (i inputImage) Name(value string) inputImage {
	return i.set("name", value)
}

// OnEvent sets the event action configuration
func (i inputImage) OnEvent(value any) inputImage {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder
func (i inputImage) Placeholder(value string) inputImage {
	return i.set("placeholder", value)
}

// ReCropable allows re-cropping
func (i inputImage) ReCropable(value bool) inputImage {
	return i.set("reCropable", value)
}

// ReadOnly sets the component to read-only
func (i inputImage) ReadOnly(value bool) inputImage {
	return i.set("readOnly", value)
}

// Receiver sets the image upload URL
func (i inputImage) Receiver(value string) inputImage {
	return i.set("receiver", value)
}

// Upload sets the upload handler
func (i inputImage) Upload(maxMemory int64, action func([]byte) (path string, err error)) inputImage {
	return i.Receiver(serveUpload(maxMemory, action))
}

// Required sets the component as required
func (i inputImage) Required(value bool) inputImage {
	return i.set("required", value)
}

// Schema sets the custom schema
func (i inputImage) Schema(value string) inputImage {
	return i.set("schema", value)
}

// ShowTips shows tips
func (i inputImage) ShowTips(value bool) inputImage {
	return i.set("showTips", value)
}

// Size sets the size
func (i inputImage) Size(value string) inputImage {
	return i.set("size", value)
}

// Src sets the default image URL
func (i inputImage) Src(value string) inputImage {
	return i.set("src", value)
}

// Step sets the incremental selection size for upload
func (i inputImage) Step(value string) inputImage {
	return i.set("step", value)
}

// StrictMode enables strict mode
func (i inputImage) StrictMode(value bool) inputImage {
	return i.set("strictMode", value)
}

// SubmitOnChange submits the form on value change
func (i inputImage) SubmitOnChange(value bool) inputImage {
	return i.set("submitOnChange", value)
}

// UploadType sets the upload type
func (i inputImage) UploadType(value string) inputImage {
	return i.set("uploadType", value)
}

// ValueFieldName sets the field name for the component value
func (i inputImage) ValueFieldName(value string) inputImage {
	return i.set("valueFieldName", value)
}

// Vertical sets the component to vertical mode
func (i inputImage) Vertical(value bool) inputImage {
	return i.set("vertical", value)
}

// VisibleOn sets the expression to show the component
func (i inputImage) VisibleOn(value string) inputImage {
	return i.set("visibleOn", value)
}

// Width sets the width
func (i inputImage) Width(value string) inputImage {
	return i.set("width", value)
}
