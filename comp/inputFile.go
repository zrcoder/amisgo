package comp

// inputFile 文件上传控件
type inputFile schema

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

// Accept 默认只支持纯文本，要支持其他类型，请配置此属性。建议直接填写文件后缀 如：.txt,.csv多个类型用逗号隔开。
func (fc inputFile) Accept(value string) inputFile {
	return fc.set("accept", value)
}

// AsBase64 如果上传的文件比较小可以设置此选项来简单的把文件 base64 的值给 form 一起提交，目前不支持多选。
func (fc inputFile) AsBase64(value bool) inputFile {
	return fc.set("asBase64", value)
}

// AsBlob 如果不希望 File 组件上传，可以配置 `asBlob` 或者 `asBase64`，采用这种方式后，组件不再自己上传了，而是直接把文件数据作为表单项的值，文件内容会在 Form 表单提交的接口里面一起带上。
func (fc inputFile) AsBlob(value bool) inputFile {
	return fc.set("asBlob", value)
}

// AutoFill 上传后把其他字段同步到表单内部。
func (fc inputFile) AutoFill(value string) inputFile {
	return fc.set("autoFill", value)
}

// AutoUpload 是否自动开始上传
func (fc inputFile) AutoUpload(value bool) inputFile {
	return fc.set("autoUpload", value)
}

// BtnClassName 按钮 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) BtnClassName(value string) inputFile {
	return fc.set("btnClassName", value)
}

// BtnLabel 上传文件按钮说明
func (fc inputFile) BtnLabel(value string) inputFile {
	return fc.set("btnLabel", value)
}

// BtnUploadClassName 上传按钮 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) BtnUploadClassName(value string) inputFile {
	return fc.set("btnUploadClassName", value)
}

// Capture 控制 input 标签的 capture 属性，用于移动端拍照或录像。
func (fc inputFile) Capture(value string) inputFile {
	return fc.set("capture", value)
}

// ChunkApi 默认 `/api/upload/chunk` 想自己存储时才需要关注。 (默认 `/api/upload/chunk` 想自己存储时才需要关注。)
func (fc inputFile) ChunkApi(value string) inputFile {
	return fc.set("chunkApi", value)
}

// ChunkSize 分块大小，默认为 5M.
func (fc inputFile) ChunkSize(value string) inputFile {
	return fc.set("chunkSize", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) ClassName(value string) inputFile {
	return fc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (fc inputFile) ClearValueOnHidden(value bool) inputFile {
	return fc.set("clearValueOnHidden", value)
}

// Concurrency 分块上传的并发数
func (fc inputFile) Concurrency(value string) inputFile {
	return fc.set("concurrency", value)
}

// Delimiter 分割符
func (fc inputFile) Delimiter(value string) inputFile {
	return fc.set("delimiter", value)
}

// Desc
func (fc inputFile) Desc(value string) inputFile {
	return fc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (fc inputFile) Description(value string) inputFile {
	return fc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (fc inputFile) DescriptionClassName(value string) inputFile {
	return fc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (fc inputFile) Disabled(value bool) inputFile {
	return fc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (fc inputFile) DisabledOn(value string) inputFile {
	return fc.set("disabledOn", value)
}

// DocumentLink 说明文档链接配置
func (fc inputFile) DocumentLink(value string) inputFile {
	return fc.set("documentLink", value)
}

// Documentation 说明文档内容配置
func (fc inputFile) Documentation(value string) inputFile {
	return fc.set("documentation", value)
}

// DownloadUrl 默认显示文件路径的时候会支持直接下载， 可以支持加前缀如：`http://xx.dom/filename=` ， 如果不希望这样，可以把当前配置项设置为 `false`。1.1.6 版本开始将支持变量 ${xxx} 来自己拼凑个下载地址，并且支持配置成 post. (默认显示文件路径的时候会支持直接下载， 可以支持加前缀如：`http://xx.dom/filename=` ， 如果不希望这样，可以把当前配置项设置为 `false`。1.1.6 版本开始将支持变量 ${xxx} 来自己拼凑个下载地址，并且支持配置成 post.)
func (fc inputFile) DownloadUrl(value string) inputFile {
	return fc.set("downloadUrl", value)
}

// Drag 是否为拖拽上传
func (fc inputFile) Drag(value bool) inputFile {
	return fc.set("drag", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (fc inputFile) EditorSetting(value string) inputFile {
	return fc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (fc inputFile) ExtraName(value string) inputFile {
	return fc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (fc inputFile) ExtractValue(value bool) inputFile {
	return fc.set("extractValue", value)
}

// FileField 默认 `file`, 如果你不想自己存储，则可以忽略此属性。
func (fc inputFile) FileField(value string) inputFile {
	return fc.set("fileField", value)
}

// StartChunkApi 默认 "/api/upload/startChunk"
func (fc inputFile) StartChunkApi(value string) inputFile {
	return fc.set("startChunkApi", value)
}

// FinishChunkApi 默认 `/api/upload/finishChunkApi` 想自己存储时才需要关注。
func (fc inputFile) FinishChunkApi(value string) inputFile {
	return fc.set("finishChunkApi", value)
}

// Hidden 是否隐藏
func (fc inputFile) Hidden(value bool) inputFile {
	return fc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (fc inputFile) HiddenOn(value string) inputFile {
	return fc.set("hiddenOn", value)
}

// HideUploadButton 是否隐藏上传按钮
func (fc inputFile) HideUploadButton(value bool) inputFile {
	return fc.set("hideUploadButton", value)
}

// Hint 输入提示，聚焦的时候显示
func (fc inputFile) Hint(value string) inputFile {
	return fc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (fc inputFile) Horizontal(value string) inputFile {
	return fc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (fc inputFile) Id(value string) inputFile {
	return fc.set("id", value)
}

// InitAutoFill 初始化时是否把其他字段的值同步到表单内部
func (fc inputFile) InitAutoFill(value string) inputFile {
	return fc.set("initAutoFill", value)
}

// InputClassName 上传控件的 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) InputClassName(value string) inputFile {
	return fc.set("inputClassName", value)
}

// InputFileClassName 文件输入框 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) InputFileClassName(value string) inputFile {
	return fc.set("inputFileClassName", value)
}

// Label 标签内容
func (fc inputFile) Label(value string) inputFile {
	return fc.set("label", value)
}

// LabelClassName 标签 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) LabelClassName(value string) inputFile {
	return fc.set("labelClassName", value)
}

// MaxSize 文件最大体积，单位为 KB
func (fc inputFile) MaxSize(value int) inputFile {
	return fc.set("maxSize", value)
}

// MaxSizeMessage 文件最大体积，单位为 KB 的提示消息。
func (fc inputFile) MaxSizeMessage(value string) inputFile {
	return fc.set("maxSizeMessage", value)
}

// MinSize 文件最小体积，单位为 KB
func (fc inputFile) MinSize(value int) inputFile {
	return fc.set("minSize", value)
}

// MinSizeMessage 文件最小体积，单位为 KB 的提示消息。
func (fc inputFile) MinSizeMessage(value string) inputFile {
	return fc.set("minSizeMessage", value)
}

// Multiple 是否允许多选 (是否允许多选)
func (fc inputFile) Multiple(value bool) inputFile {
	return fc.set("multiple", value)
}

// Name 表单项的 name 属性，配置到表单中提交时的字段名称。
func (fc inputFile) Name(value string) inputFile {
	return fc.set("name", value)
}

// ProgressClassName 进度条 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) ProgressClassName(value string) inputFile {
	return fc.set("progressClassName", value)
}

// Receiver 默认上传地址，支持后端 url 或者函数
func (fc inputFile) Receiver(value string) inputFile {
	return fc.set("receiver", value)
}

// RenderType 默认类型，值为： `input-file`， 表示普通文件上传控件 (值为： `input-file`， 表示普通文件上传控件)
func (fc inputFile) RenderType(value string) inputFile {
	return fc.set("renderType", value)
}

// Resize 图片文件上传时，支持设置缩略图的最大宽度（像素）
func (fc inputFile) ResizeMaxWidth(value int) inputFile {
	return fc.set("resizeMaxWidth", value)
}

// ResizeMaxHeight 图片文件上传时，支持设置缩略图的最大高度（像素）
func (fc inputFile) ResizeMaxHeight(value int) inputFile {
	return fc.set("resizeMaxHeight", value)
}

// ShowTips 上传过程中是否显示提示
func (fc inputFile) ShowTips(value bool) inputFile {
	return fc.set("showTips", value)
}

// Size 文件显示大小 (size)：仅支持“K”或“M”。
func (fc inputFile) Size(value string) inputFile {
	return fc.set("size", value)
}

// ShowDownload 是否显示下载按钮
func (fc inputFile) ShowDownload(value bool) inputFile {
	return fc.set("showDownload", value)
}

// ShowPreview 是否显示预览图标
func (fc inputFile) ShowPreview(value bool) inputFile {
	return fc.set("showPreview", value)
}

// SizeMessage 文件显示大小的提示消息。
func (fc inputFile) SizeMessage(value string) inputFile {
	return fc.set("sizeMessage", value)
}

// Source 只读模式下，上传文件显示的数据源配置 (只读模式下，上传文件显示的数据源配置)
func (fc inputFile) Source(value string) inputFile {
	return fc.set("source", value)
}

// Static 是否静态显示 (静态显示不支持上传、拖拽、删除等操作。)
func (fc inputFile) Static(value bool) inputFile {
	return fc.set("static", value)
}

// Status 上传状态
func (fc inputFile) Status(value string) inputFile {
	return fc.set("status", value)
}

// SubmitOnChange 表单项值变化时是否立即提交
func (fc inputFile) SubmitOnChange(value bool) inputFile {
	return fc.set("submitOnChange", value)
}

// UploadBtnClassName 上传按钮 CSS 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (fc inputFile) UploadBtnClassName(value string) inputFile {
	return fc.set("uploadBtnClassName", value)
}

// Url 文件下载地址，若值为 `null` 则表示不显示下载链接。
func (fc inputFile) Url(value string) inputFile {
	return fc.set("url", value)
}

// ValidateType 文件验证类型。默认支持 `text`、`image`、`media`、`audio`、`video`、`application` 等类型，支持使用 **逗号分隔的字符串** 来定义更多类型。
func (fc inputFile) ValidateType(value string) inputFile {
	return fc.set("validateType", value)
}

// Value 组件值
func (fc inputFile) Value(value string) inputFile {
	return fc.set("value", value)
}

// Width 宽度
func (fc inputFile) Width(value string) inputFile {
	return fc.set("width", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (fc inputFile) Mode(value string) inputFile {
	return fc.set("mode", value)
}

// ShowLabel 是否显示标签 (是否显示标签)
func (fc inputFile) ShowLabel(value bool) inputFile {
	return fc.set("showLabel", value)
}

// ShowError 失败提示的开关
func (fc inputFile) ShowError(value bool) inputFile {
	return fc.set("showError", value)
}
