package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// inputFile represents a file upload component
type inputFile model.Schema

func InputFile() inputFile {
	f := make(inputFile)
	f.set("type", "input-file")
	f.Receiver("/api/upload/file")
	f.StartChunkApi("/api/upload/startChunk")
	f.ChunkApi("/api/upload/chunk")
	f.FinishChunkApi("/api/upload/finishChunkApi")
	return f
}

func (f inputFile) set(key string, value any) inputFile {
	f[key] = value
	return f
}

// Accept sets the accepted file types.
func (fc inputFile) Accept(value string) inputFile {
	return fc.set("accept", value)
}

// AsBase64 sets whether to submit the file as base64.
func (fc inputFile) AsBase64(value bool) inputFile {
	return fc.set("asBase64", value)
}

// AsBlob sets whether to submit the file as a blob.
func (fc inputFile) AsBlob(value bool) inputFile {
	return fc.set("asBlob", value)
}

// AutoFill sets whether to auto-fill other fields after upload.
func (fc inputFile) AutoFill(value string) inputFile {
	return fc.set("autoFill", value)
}

// AutoUpload sets whether to start upload automatically.
func (fc inputFile) AutoUpload(value bool) inputFile {
	return fc.set("autoUpload", value)
}

// BtnClassName sets the CSS class name for the button.
func (fc inputFile) BtnClassName(value string) inputFile {
	return fc.set("btnClassName", value)
}

// BtnLabel sets the label for the upload button.
func (fc inputFile) BtnLabel(value string) inputFile {
	return fc.set("btnLabel", value)
}

// BtnUploadClassName sets the CSS class name for the upload button.
func (fc inputFile) BtnUploadClassName(value string) inputFile {
	return fc.set("btnUploadClassName", value)
}

// Capture sets the capture attribute for the input tag.
func (fc inputFile) Capture(value string) inputFile {
	return fc.set("capture", value)
}

// ChunkApi sets the API endpoint for chunk upload.
func (fc inputFile) ChunkApi(value string) inputFile {
	return fc.set("chunkApi", value)
}

// ChunkSize sets the chunk size for upload.
func (fc inputFile) ChunkSize(value string) inputFile {
	return fc.set("chunkSize", value)
}

// ClassName sets the CSS class name for the container.
func (fc inputFile) ClassName(value string) inputFile {
	return fc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (fc inputFile) ClearValueOnHidden(value bool) inputFile {
	return fc.set("clearValueOnHidden", value)
}

// Concurrency sets the concurrency for chunk upload.
func (fc inputFile) Concurrency(value string) inputFile {
	return fc.set("concurrency", value)
}

// Delimiter sets the delimiter for the file.
func (fc inputFile) Delimiter(value string) inputFile {
	return fc.set("delimiter", value)
}

// Desc sets the description.
func (fc inputFile) Desc(value string) inputFile {
	return fc.set("desc", value)
}

// Description sets the description content.
func (fc inputFile) Description(value string) inputFile {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (fc inputFile) DescriptionClassName(value string) inputFile {
	return fc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (fc inputFile) Disabled(value bool) inputFile {
	return fc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (fc inputFile) DisabledOn(value string) inputFile {
	return fc.set("disabledOn", value)
}

// DocumentLink sets the link to the documentation.
func (fc inputFile) DocumentLink(value string) inputFile {
	return fc.set("documentLink", value)
}

// Documentation sets the documentation content.
func (fc inputFile) Documentation(value string) inputFile {
	return fc.set("documentation", value)
}

// DownloadUrl sets the download URL.
func (fc inputFile) DownloadUrl(value string) inputFile {
	return fc.set("downloadUrl", value)
}

// Drag sets whether drag-and-drop upload is enabled.
func (fc inputFile) Drag(value bool) inputFile {
	return fc.set("drag", value)
}

// EditorSetting sets the editor configuration.
func (fc inputFile) EditorSetting(value string) inputFile {
	return fc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (fc inputFile) ExtraName(value string) inputFile {
	return fc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (fc inputFile) ExtractValue(value bool) inputFile {
	return fc.set("extractValue", value)
}

// FileField sets the file field name.
func (fc inputFile) FileField(value string) inputFile {
	return fc.set("fileField", value)
}

// StartChunkApi sets the API endpoint to start chunk upload.
func (fc inputFile) StartChunkApi(value string) inputFile {
	return fc.set("startChunkApi", value)
}

// FinishChunkApi sets the API endpoint to finish chunk upload.
func (fc inputFile) FinishChunkApi(value string) inputFile {
	return fc.set("finishChunkApi", value)
}

// Hidden sets whether the component is hidden.
func (fc inputFile) Hidden(value bool) inputFile {
	return fc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (fc inputFile) HiddenOn(value string) inputFile {
	return fc.set("hiddenOn", value)
}

// HideUploadButton sets whether to hide the upload button.
func (fc inputFile) HideUploadButton(value bool) inputFile {
	return fc.set("hideUploadButton", value)
}

// Hint sets the input hint.
func (fc inputFile) Hint(value string) inputFile {
	return fc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (fc inputFile) Horizontal(value string) inputFile {
	return fc.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (fc inputFile) ID(value string) inputFile {
	return fc.set("id", value)
}

// InitAutoFill sets whether to auto-fill other fields on initialization.
func (fc inputFile) InitAutoFill(value string) inputFile {
	return fc.set("initAutoFill", value)
}

// InputClassName sets the CSS class name for the input.
func (fc inputFile) InputClassName(value string) inputFile {
	return fc.set("inputClassName", value)
}

// InputFileClassName sets the CSS class name for the file input.
func (fc inputFile) InputFileClassName(value string) inputFile {
	return fc.set("inputFileClassName", value)
}

// Label sets the label content.
func (fc inputFile) Label(value string) inputFile {
	return fc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (fc inputFile) LabelClassName(value string) inputFile {
	return fc.set("labelClassName", value)
}

// MaxSize sets the maximum file size in KB.
func (fc inputFile) MaxSize(value int) inputFile {
	return fc.set("maxSize", value)
}

// MaxSizeMessage sets the message for the maximum file size.
func (fc inputFile) MaxSizeMessage(value string) inputFile {
	return fc.set("maxSizeMessage", value)
}

// MinSize sets the minimum file size in KB.
func (fc inputFile) MinSize(value int) inputFile {
	return fc.set("minSize", value)
}

// MinSizeMessage sets the message for the minimum file size.
func (fc inputFile) MinSizeMessage(value string) inputFile {
	return fc.set("minSizeMessage", value)
}

// Multiple sets whether multiple file selection is allowed.
func (fc inputFile) Multiple(value bool) inputFile {
	return fc.set("multiple", value)
}

// Name sets the name attribute for the form field.
func (fc inputFile) Name(value string) inputFile {
	return fc.set("name", value)
}

// ProgressClassName sets the CSS class name for the progress bar.
func (fc inputFile) ProgressClassName(value string) inputFile {
	return fc.set("progressClassName", value)
}

// Receiver sets the default upload URL.
func (fc inputFile) Receiver(value string) inputFile {
	return fc.set("receiver", value)
}

// Upload sets the upload function.
func (i inputFile) Upload(maxMemory int64, action func([]byte) (path string, err error)) inputFile {
	return i.Receiver(servermux.ServeUpload(maxMemory, action))
}

// RenderType sets the default render type.
func (fc inputFile) RenderType(value string) inputFile {
	return fc.set("renderType", value)
}

// ResizeMaxWidth sets the maximum width for image resizing.
func (fc inputFile) ResizeMaxWidth(value int) inputFile {
	return fc.set("resizeMaxWidth", value)
}

// ResizeMaxHeight sets the maximum height for image resizing.
func (fc inputFile) ResizeMaxHeight(value int) inputFile {
	return fc.set("resizeMaxHeight", value)
}

// ShowTips sets whether to show tips during upload.
func (fc inputFile) ShowTips(value bool) inputFile {
	return fc.set("showTips", value)
}

// Size sets the file display size.
func (fc inputFile) Size(value string) inputFile {
	return fc.set("size", value)
}

// ShowDownload sets whether to show the download button.
func (fc inputFile) ShowDownload(value bool) inputFile {
	return fc.set("showDownload", value)
}

// ShowPreview sets whether to show the preview icon.
func (fc inputFile) ShowPreview(value bool) inputFile {
	return fc.set("showPreview", value)
}

// SizeMessage sets the message for the file size.
func (fc inputFile) SizeMessage(value string) inputFile {
	return fc.set("sizeMessage", value)
}

// Source sets the data source for read-only mode.
func (fc inputFile) Source(value string) inputFile {
	return fc.set("source", value)
}

// Static sets whether to display the component statically.
func (fc inputFile) Static(value bool) inputFile {
	return fc.set("static", value)
}

// Status sets the upload status.
func (fc inputFile) Status(value string) inputFile {
	return fc.set("status", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (fc inputFile) SubmitOnChange(value bool) inputFile {
	return fc.set("submitOnChange", value)
}

// UploadBtnClassName sets the CSS class name for the upload button.
func (fc inputFile) UploadBtnClassName(value string) inputFile {
	return fc.set("uploadBtnClassName", value)
}

// Url sets the file download URL.
func (fc inputFile) Url(value string) inputFile {
	return fc.set("url", value)
}

// ValidateType sets the file validation type.
func (fc inputFile) ValidateType(value string) inputFile {
	return fc.set("validateType", value)
}

// Value sets the component value.
func (fc inputFile) Value(value string) inputFile {
	return fc.set("value", value)
}

// Width sets the component width.
func (fc inputFile) Width(value string) inputFile {
	return fc.set("width", value)
}

// Mode sets the display mode for the form item.
func (fc inputFile) Mode(value string) inputFile {
	return fc.set("mode", value)
}

// ShowLabel sets whether to show the label.
func (fc inputFile) ShowLabel(value bool) inputFile {
	return fc.set("showLabel", value)
}

// ShowError sets whether to show error messages.
func (fc inputFile) ShowError(value bool) inputFile {
	return fc.set("showError", value)
}
