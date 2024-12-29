package comp

// inputImage 图片上传控件
type inputImage Schema

// InputImage 创建一个新的 InputImage 实例
func InputImage() inputImage {
	return make(inputImage).set("type", "input-image").Receiver("/api/upload")
}

func (i inputImage) set(key string, value any) inputImage {
	i[key] = value
	return i
}

// Accept 配置接收的图片类型
func (i inputImage) Accept(value string) inputImage {
	return i.set("accept", value)
}

// AllowInput 默认都是通过用户选择图片后上传返回图片地址，如果开启此选项，则可以允许用户图片地址
func (i inputImage) AllowInput(value bool) inputImage {
	return i.set("allowInput", value)
}

// AutoFill 上传后把其他字段同步到表单内部
func (i inputImage) AutoFill(value string) inputImage {
	return i.set("autoFill", value)
}

// AutoUpload 是否自动开始上传
func (i inputImage) AutoUpload(value bool) inputImage {
	return i.set("autoUpload", value)
}

// BtnClassName 选择图片按钮的 CSS 类名
func (i inputImage) BtnClassName(value string) inputImage {
	return i.set("btnClassName", value)
}

// BtnUploadClassName 上传按钮的 CSS 类名
func (i inputImage) BtnUploadClassName(value string) inputImage {
	return i.set("btnUploadClassName", value)
}

// Capture 可配置移动端的拍照功能
func (i inputImage) Capture(value string) inputImage {
	return i.set("capture", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (i inputImage) ClassName(value any) inputImage {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (i inputImage) ClearValueOnHidden(value bool) inputImage {
	return i.set("clearValueOnHidden", value)
}

// Compress 启用压缩
func (i inputImage) Compress(value bool) inputImage {
	return i.set("compress", value)
}

// CompressOptions 压缩选项
func (i inputImage) CompressOptions(value ...any) inputImage {
	return i.set("compressOptions", value)
}

// Crop 裁剪选项
func (i inputImage) Crop(value any) inputImage {
	return i.set("crop", value)
}

// CropFormat 裁剪后的图片类型
func (i inputImage) CropFormat(value string) inputImage {
	return i.set("cropFormat", value)
}

// CropQuality 裁剪后的质量
func (i inputImage) CropQuality(value string) inputImage {
	return i.set("cropQuality", value)
}

// Delimiter 分割符
func (i inputImage) Delimiter(value string) inputImage {
	return i.set("delimiter", value)
}

// Desc 描述内容
func (i inputImage) Desc(value string) inputImage {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (i inputImage) Description(value string) inputImage {
	return i.set("description", value)
}

// DescriptionClassName 描述上的 className
func (i inputImage) DescriptionClassName(value string) inputImage {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i inputImage) Disabled(value bool) inputImage {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i inputImage) DisabledOn(value string) inputImage {
	return i.set("disabledOn", value)
}

// Draggable 是否可拖拽排序
func (i inputImage) Draggable(value bool) inputImage {
	return i.set("draggable", value)
}

// DraggableTip 可拖拽排序的提示信息
func (i inputImage) DraggableTip(value string) inputImage {
	return i.set("draggableTip", value)
}

// DropCrop 图片上传完毕是否进入裁剪模式
func (i inputImage) DropCrop(value bool) inputImage {
	return i.set("dropCrop", value)
}

// EditorSetting 编辑器配置
func (i inputImage) EditorSetting(value string) inputImage {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (i inputImage) ExtraName(value string) inputImage {
	return i.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (i inputImage) ExtractValue(value bool) inputImage {
	return i.set("extractValue", value)
}

// FixedSize 固定尺寸
func (i inputImage) FixedSize(value any) inputImage {
	return i.set("fixedSize", value)
}

// FixedSizeClassName 固定尺寸的 CSS 类名
func (i inputImage) FixedSizeClassName(value string) inputImage {
	return i.set("fixedSizeClassName", value)
}

// FrameImage 默认占位图图片地址
func (i inputImage) FrameImage(value string) inputImage {
	return i.set("frameImage", value)
}

// Hidden 是否隐藏
func (i inputImage) Hidden(value bool) inputImage {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i inputImage) HiddenOn(value string) inputImage {
	return i.set("hiddenOn", value)
}

// HideUploadButton 是否隐藏上传按钮
func (i inputImage) HideUploadButton(value bool) inputImage {
	return i.set("hideUploadButton", value)
}

// Hint 输入提示
func (i inputImage) Hint(value string) inputImage {
	return i.set("hint", value)
}

// Horizontal 配置具体的左右分配
func (i inputImage) Horizontal(value string) inputImage {
	return i.set("horizontal", value)
}

// ID 组件唯一 id
func (i inputImage) ID(value string) inputImage {
	return i.set("id", value)
}

// ImageClassName 默认展示图片的类名
func (i inputImage) ImageClassName(value string) inputImage {
	return i.set("imageClassName", value)
}

// InitAutoFill 初始化时是否把其他字段同步到表单内部
func (i inputImage) InitAutoFill(value bool) inputImage {
	return i.set("initAutoFill", value)
}

// InitCrop 初始化时是否打开裁剪模式
func (i inputImage) InitCrop(value bool) inputImage {
	return i.set("initCrop", value)
}

// Inline 表单 control 是否为 inline 模式
func (i inputImage) Inline(value bool) inputImage {
	return i.set("inline", value)
}

// InputClassName 配置 input className
func (i inputImage) InputClassName(value string) inputImage {
	return i.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式的值处理
func (i inputImage) JoinValues(value bool) inputImage {
	return i.set("joinValues", value)
}

// Label 描述标题
func (i inputImage) Label(value string) inputImage {
	return i.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (i inputImage) LabelAlign(value string) inputImage {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i inputImage) LabelClassName(value string) inputImage {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标并显示提示内容
func (i inputImage) LabelRemark(value string) inputImage {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (i inputImage) LabelWidth(value string) inputImage {
	return i.set("labelWidth", value)
}

// Limit 限制图片大小
func (i inputImage) Limit(value any) inputImage {
	return i.set("limit", value)
}

// MaxLength 最多的个数
func (i inputImage) MaxLength(value string) inputImage {
	return i.set("maxLength", value)
}

// MaxSize 默认没有限制，设置后，文件大小大于此值将不允许上传
func (i inputImage) MaxSize(value string) inputImage {
	return i.set("maxSize", value)
}

// Mode 表单项展示模式
func (i inputImage) Mode(value string) inputImage {
	return i.set("mode", value)
}

// Multiple 是否为多选
func (i inputImage) Multiple(value bool) inputImage {
	return i.set("multiple", value)
}

// Name 字段名
func (i inputImage) Name(value string) inputImage {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i inputImage) OnEvent(value any) inputImage {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i inputImage) Placeholder(value string) inputImage {
	return i.set("placeholder", value)
}

// ReCropable 是否允许二次裁剪
func (i inputImage) ReCropable(value bool) inputImage {
	return i.set("reCropable", value)
}

// ReadOnly 是否只读
func (i inputImage) ReadOnly(value bool) inputImage {
	return i.set("readOnly", value)
}

// Receiver 图片上传接口
func (i inputImage) Receiver(value string) inputImage {
	return i.set("receiver", value)
}

// Upload
func (i inputImage) Upload(maxMemory int64, action func([]byte) (path string, err error)) inputImage {
	return i.Receiver(serveUpload(maxMemory, action))
}

// Required 是否必填
func (i inputImage) Required(value bool) inputImage {
	return i.set("required", value)
}

// Schema 自定义 schema
func (i inputImage) Schema(value string) inputImage {
	return i.set("schema", value)
}

// ShowTips 是否显示提示
func (i inputImage) ShowTips(value bool) inputImage {
	return i.set("showTips", value)
}

// Size 尺寸
func (i inputImage) Size(value string) inputImage {
	return i.set("size", value)
}

// Src 默认图片地址
func (i inputImage) Src(value string) inputImage {
	return i.set("src", value)
}

// Step 上传时增量选择的大小
func (i inputImage) Step(value string) inputImage {
	return i.set("step", value)
}

// StrictMode 是否严格模式
func (i inputImage) StrictMode(value bool) inputImage {
	return i.set("strictMode", value)
}

// SubmitOnChange 是否在值变化时提交表单
func (i inputImage) SubmitOnChange(value bool) inputImage {
	return i.set("submitOnChange", value)
}

// UploadType 上传类型
func (i inputImage) UploadType(value string) inputImage {
	return i.set("uploadType", value)
}

// ValueFieldName 控件值的字段名
func (i inputImage) ValueFieldName(value string) inputImage {
	return i.set("valueFieldName", value)
}

// Vertical 设置为垂直模式
func (i inputImage) Vertical(value bool) inputImage {
	return i.set("vertical", value)
}

// VisibleOn 是否显示表达式
func (i inputImage) VisibleOn(value string) inputImage {
	return i.set("visibleOn", value)
}

// Width 宽度
func (i inputImage) Width(value string) inputImage {
	return i.set("width", value)
}
