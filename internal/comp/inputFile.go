package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/schema"
)

// InputFile represents a file upload component
type InputFile schema.Schema

func NewInputFile(mux *http.ServeMux) InputFile {
	f := InputFile{}
	f.set("type", "input-file")
	f.set(servemux.Key, mux)
	return f
}

func (f InputFile) set(key string, value any) InputFile {
	f[key] = value
	return f
}

// Accept sets the accepted file types.
func (fc InputFile) Accept(value string) InputFile {
	return fc.set("accept", value)
}

// AsBase64 sets whether to submit the file as base64.
func (fc InputFile) AsBase64(value bool) InputFile {
	return fc.set("asBase64", value)
}

// AsBlob sets whether to submit the file as a blob.
func (fc InputFile) AsBlob(value bool) InputFile {
	return fc.set("asBlob", value)
}

// AutoFill sets whether to auto-fill other fields after upload.
func (fc InputFile) AutoFill(value string) InputFile {
	return fc.set("autoFill", value)
}

// AutoUpload sets whether to start upload automatically.
func (fc InputFile) AutoUpload(value bool) InputFile {
	return fc.set("autoUpload", value)
}

// BtnClassName sets the CSS class name for the button.
func (fc InputFile) BtnClassName(value string) InputFile {
	return fc.set("btnClassName", value)
}

// BtnLabel sets the label for the upload button.
func (fc InputFile) BtnLabel(value string) InputFile {
	return fc.set("btnLabel", value)
}

// BtnUploadClassName sets the CSS class name for the upload button.
func (fc InputFile) BtnUploadClassName(value string) InputFile {
	return fc.set("btnUploadClassName", value)
}

// Capture sets the capture attribute for the input tag.
func (fc InputFile) Capture(value string) InputFile {
	return fc.set("capture", value)
}

// ChunkApi sets the API endpoint for chunk upload.
func (fc InputFile) ChunkApi(value string) InputFile {
	return fc.set("chunkApi", value)
}

// ChunkSize sets the chunk size for upload.
func (fc InputFile) ChunkSize(value string) InputFile {
	return fc.set("chunkSize", value)
}

// ClassName sets the CSS class name for the container.
func (fc InputFile) ClassName(value string) InputFile {
	return fc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (fc InputFile) ClearValueOnHidden(value bool) InputFile {
	return fc.set("clearValueOnHidden", value)
}

// Concurrency sets the concurrency for chunk upload.
func (fc InputFile) Concurrency(value string) InputFile {
	return fc.set("concurrency", value)
}

// Delimiter sets the delimiter for the file.
func (fc InputFile) Delimiter(value string) InputFile {
	return fc.set("delimiter", value)
}

// Desc sets the description.
func (fc InputFile) Desc(value string) InputFile {
	return fc.set("desc", value)
}

// Description sets the description content.
func (fc InputFile) Description(value string) InputFile {
	return fc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (fc InputFile) DescriptionClassName(value string) InputFile {
	return fc.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (fc InputFile) Disabled(value bool) InputFile {
	return fc.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (fc InputFile) DisabledOn(value string) InputFile {
	return fc.set("disabledOn", value)
}

// DocumentLink sets the link to the documentation.
func (fc InputFile) DocumentLink(value string) InputFile {
	return fc.set("documentLink", value)
}

// Documentation sets the documentation content.
func (fc InputFile) Documentation(value string) InputFile {
	return fc.set("documentation", value)
}

// DownloadUrl sets the download URL.
func (fc InputFile) DownloadUrl(value string) InputFile {
	return fc.set("downloadUrl", value)
}

// UseChunk sets whether to use chunk upload. true | false | "auto"
func (fc InputFile) UseChunk(value any) InputFile {
	return fc.set("useChunk", value)
}

// Drag sets whether drag-and-drop upload is enabled.
func (fc InputFile) Drag(value bool) InputFile {
	return fc.set("drag", value)
}

// EditorSetting sets the editor configuration.
func (fc InputFile) EditorSetting(value string) InputFile {
	return fc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (fc InputFile) ExtraName(value string) InputFile {
	return fc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (fc InputFile) ExtractValue(value bool) InputFile {
	return fc.set("extractValue", value)
}

// FileField sets the file field name.
func (fc InputFile) FileField(value string) InputFile {
	return fc.set("fileField", value)
}

// StartChunkApi sets the API endpoint to start chunk upload.
func (fc InputFile) StartChunkApi(value string) InputFile {
	return fc.set("startChunkApi", value)
}

// FinishChunkApi sets the API endpoint to finish chunk upload.
func (fc InputFile) FinishChunkApi(value string) InputFile {
	return fc.set("finishChunkApi", value)
}

// Hidden sets whether the component is hidden.
func (fc InputFile) Hidden(value bool) InputFile {
	return fc.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (fc InputFile) HiddenOn(value string) InputFile {
	return fc.set("hiddenOn", value)
}

// HideUploadButton sets whether to hide the upload button.
func (fc InputFile) HideUploadButton(value bool) InputFile {
	return fc.set("hideUploadButton", value)
}

// Hint sets the input hint.
func (fc InputFile) Hint(value string) InputFile {
	return fc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (fc InputFile) Horizontal(value string) InputFile {
	return fc.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (fc InputFile) ID(value string) InputFile {
	return fc.set("id", value)
}

// InitAutoFill sets whether to auto-fill other fields on initialization.
func (fc InputFile) InitAutoFill(value string) InputFile {
	return fc.set("initAutoFill", value)
}

// InputClassName sets the CSS class name for the input.
func (fc InputFile) InputClassName(value string) InputFile {
	return fc.set("inputClassName", value)
}

// InputFileClassName sets the CSS class name for the file input.
func (fc InputFile) InputFileClassName(value string) InputFile {
	return fc.set("inputFileClassName", value)
}

// Label sets the label content.
func (fc InputFile) Label(value string) InputFile {
	return fc.set("label", value)
}

// LabelClassName sets the CSS class name for the label.
func (fc InputFile) LabelClassName(value string) InputFile {
	return fc.set("labelClassName", value)
}

// MaxSize sets the maximum file size in KB.
func (fc InputFile) MaxSize(value int) InputFile {
	return fc.set("maxSize", value)
}

// MaxSizeMessage sets the message for the maximum file size.
func (fc InputFile) MaxSizeMessage(value string) InputFile {
	return fc.set("maxSizeMessage", value)
}

// MinSize sets the minimum file size in KB.
func (fc InputFile) MinSize(value int) InputFile {
	return fc.set("minSize", value)
}

// MinSizeMessage sets the message for the minimum file size.
func (fc InputFile) MinSizeMessage(value string) InputFile {
	return fc.set("minSizeMessage", value)
}

// Multiple sets whether multiple file selection is allowed.
func (fc InputFile) Multiple(value bool) InputFile {
	return fc.set("multiple", value)
}

// Name sets the name attribute for the form field.
func (fc InputFile) Name(value string) InputFile {
	return fc.set("name", value)
}

// ProgressClassName sets the CSS class name for the progress bar.
func (fc InputFile) ProgressClassName(value string) InputFile {
	return fc.set("progressClassName", value)
}

// Receiver sets the default upload URL.
func (fc InputFile) Receiver(value string) InputFile {
	return fc.set("receiver", value)
}

// Upload sets the upload function.
func (i InputFile) Upload(maxMemory int64, action func([]byte) (path string, err error)) InputFile {
	return i.Receiver(servemux.ServeUpload(i.mux(), maxMemory, action))
}

// RenderType sets the default render type.
func (fc InputFile) RenderType(value string) InputFile {
	return fc.set("renderType", value)
}

// ResizeMaxWidth sets the maximum width for image resizing.
func (fc InputFile) ResizeMaxWidth(value int) InputFile {
	return fc.set("resizeMaxWidth", value)
}

// ResizeMaxHeight sets the maximum height for image resizing.
func (fc InputFile) ResizeMaxHeight(value int) InputFile {
	return fc.set("resizeMaxHeight", value)
}

// ShowTips sets whether to show tips during upload.
func (fc InputFile) ShowTips(value bool) InputFile {
	return fc.set("showTips", value)
}

// Size sets the file display size.
func (fc InputFile) Size(value string) InputFile {
	return fc.set("size", value)
}

// ShowDownload sets whether to show the download button.
func (fc InputFile) ShowDownload(value bool) InputFile {
	return fc.set("showDownload", value)
}

// ShowPreview sets whether to show the preview icon.
func (fc InputFile) ShowPreview(value bool) InputFile {
	return fc.set("showPreview", value)
}

// SizeMessage sets the message for the file size.
func (fc InputFile) SizeMessage(value string) InputFile {
	return fc.set("sizeMessage", value)
}

// Source sets the data source for read-only mode.
func (fc InputFile) Source(value string) InputFile {
	return fc.set("source", value)
}

// Static sets whether to display the component statically.
func (fc InputFile) Static(value bool) InputFile {
	return fc.set("static", value)
}

// Status sets the upload status.
func (fc InputFile) Status(value string) InputFile {
	return fc.set("status", value)
}

// SubmitOnChange sets whether to submit the form on value change.
func (fc InputFile) SubmitOnChange(value bool) InputFile {
	return fc.set("submitOnChange", value)
}

// UploadBtnClassName sets the CSS class name for the upload button.
func (fc InputFile) UploadBtnClassName(value string) InputFile {
	return fc.set("uploadBtnClassName", value)
}

// Url sets the file download URL.
func (fc InputFile) Url(value string) InputFile {
	return fc.set("url", value)
}

// ValidateType sets the file validation type.
func (fc InputFile) ValidateType(value string) InputFile {
	return fc.set("validateType", value)
}

// Value sets the component value.
func (fc InputFile) Value(value string) InputFile {
	return fc.set("value", value)
}

// Width sets the component width.
func (fc InputFile) Width(value string) InputFile {
	return fc.set("width", value)
}

// Mode sets the display mode for the form item.
func (fc InputFile) Mode(value string) InputFile {
	return fc.set("mode", value)
}

// ShowLabel sets whether to show the label.
func (fc InputFile) ShowLabel(value bool) InputFile {
	return fc.set("showLabel", value)
}

// ShowError sets whether to show error messages.
func (fc InputFile) ShowError(value bool) InputFile {
	return fc.set("showError", value)
}

func (f InputFile) mux() *http.ServeMux {
	return f[servemux.Key].(*http.ServeMux)
}

func (f InputFile) MarshalJSON() ([]byte, error) {
	m := make(map[string]any, len(f))
	for k, v := range f {
		pureAmis(k, v, m)
	}
	return marshalMap(m)
}
