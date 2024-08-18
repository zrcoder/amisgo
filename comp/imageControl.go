package comp

// ImageControl 图片上传控件
// PHP 类名: ImageControl
type ImageControl map[string]interface{}

// NewImageControl 创建一个新的 ImageControl 实例
func NewImageControl() ImageControl {
	return make(ImageControl).set("type", "input-image").Receiver("/api/upload")
}

func (i ImageControl) set(key string, value interface{}) ImageControl {
	i[key] = value
	return i
}

// Accept 配置接收的图片类型
func (i ImageControl) Accept(value string) ImageControl {
	return i.set("accept", value)
}

// AllowInput 默认都是通过用户选择图片后上传返回图片地址，如果开启此选项，则可以允许用户图片地址
func (i ImageControl) AllowInput(value bool) ImageControl {
	return i.set("allowInput", value)
}

// AutoFill 上传后把其他字段同步到表单内部
func (i ImageControl) AutoFill(value string) ImageControl {
	return i.set("autoFill", value)
}

// AutoUpload 是否自动开始上传
func (i ImageControl) AutoUpload(value bool) ImageControl {
	return i.set("autoUpload", value)
}

// BtnClassName 选择图片按钮的 CSS 类名
func (i ImageControl) BtnClassName(value string) ImageControl {
	return i.set("btnClassName", value)
}

// BtnUploadClassName 上传按钮的 CSS 类名
func (i ImageControl) BtnUploadClassName(value string) ImageControl {
	return i.set("btnUploadClassName", value)
}

// Capture 可配置移动端的拍照功能
func (i ImageControl) Capture(value string) ImageControl {
	return i.set("capture", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (i ImageControl) ClassName(value interface{}) ImageControl {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (i ImageControl) ClearValueOnHidden(value bool) ImageControl {
	return i.set("clearValueOnHidden", value)
}

// Compress 启用压缩
func (i ImageControl) Compress(value bool) ImageControl {
	return i.set("compress", value)
}

// CompressOptions 压缩选项
func (i ImageControl) CompressOptions(value string) ImageControl {
	return i.set("compressOptions", value)
}

// Crop 裁剪选项
func (i ImageControl) Crop(value string) ImageControl {
	return i.set("crop", value)
}

// CropFormat 裁剪后的图片类型
func (i ImageControl) CropFormat(value string) ImageControl {
	return i.set("cropFormat", value)
}

// CropQuality 裁剪后的质量
func (i ImageControl) CropQuality(value string) ImageControl {
	return i.set("cropQuality", value)
}

// Delimiter 分割符
func (i ImageControl) Delimiter(value string) ImageControl {
	return i.set("delimiter", value)
}

// Desc 描述内容
func (i ImageControl) Desc(value string) ImageControl {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (i ImageControl) Description(value string) ImageControl {
	return i.set("description", value)
}

// DescriptionClassName 描述上的 className
func (i ImageControl) DescriptionClassName(value string) ImageControl {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i ImageControl) Disabled(value bool) ImageControl {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i ImageControl) DisabledOn(value string) ImageControl {
	return i.set("disabledOn", value)
}

// Draggable 是否可拖拽排序
func (i ImageControl) Draggable(value bool) ImageControl {
	return i.set("draggable", value)
}

// DraggableTip 可拖拽排序的提示信息
func (i ImageControl) DraggableTip(value string) ImageControl {
	return i.set("draggableTip", value)
}

// DropCrop 图片上传完毕是否进入裁剪模式
func (i ImageControl) DropCrop(value bool) ImageControl {
	return i.set("dropCrop", value)
}

// EditorSetting 编辑器配置
func (i ImageControl) EditorSetting(value string) ImageControl {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (i ImageControl) ExtraName(value string) ImageControl {
	return i.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (i ImageControl) ExtractValue(value bool) ImageControl {
	return i.set("extractValue", value)
}

// FixedSize 是否开启固定尺寸
func (i ImageControl) FixedSize(value bool) ImageControl {
	return i.set("fixedSize", value)
}

// FixedSizeClassName 固定尺寸的 CSS 类名
func (i ImageControl) FixedSizeClassName(value string) ImageControl {
	return i.set("fixedSizeClassName", value)
}

// FrameImage 默认占位图图片地址
func (i ImageControl) FrameImage(value string) ImageControl {
	return i.set("frameImage", value)
}

// Hidden 是否隐藏
func (i ImageControl) Hidden(value bool) ImageControl {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i ImageControl) HiddenOn(value string) ImageControl {
	return i.set("hiddenOn", value)
}

// HideUploadButton 是否隐藏上传按钮
func (i ImageControl) HideUploadButton(value bool) ImageControl {
	return i.set("hideUploadButton", value)
}

// Hint 输入提示
func (i ImageControl) Hint(value string) ImageControl {
	return i.set("hint", value)
}

// Horizontal 配置具体的左右分配
func (i ImageControl) Horizontal(value string) ImageControl {
	return i.set("horizontal", value)
}

// ID 组件唯一 id
func (i ImageControl) ID(value string) ImageControl {
	return i.set("id", value)
}

// ImageClassName 默认展示图片的类名
func (i ImageControl) ImageClassName(value string) ImageControl {
	return i.set("imageClassName", value)
}

// InitAutoFill 初始化时是否把其他字段同步到表单内部
func (i ImageControl) InitAutoFill(value bool) ImageControl {
	return i.set("initAutoFill", value)
}

// InitCrop 初始化时是否打开裁剪模式
func (i ImageControl) InitCrop(value bool) ImageControl {
	return i.set("initCrop", value)
}

// Inline 表单 control 是否为 inline 模式
func (i ImageControl) Inline(value bool) ImageControl {
	return i.set("inline", value)
}

// InputClassName 配置 input className
func (i ImageControl) InputClassName(value string) ImageControl {
	return i.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式的值处理
func (i ImageControl) JoinValues(value bool) ImageControl {
	return i.set("joinValues", value)
}

// Label 描述标题
func (i ImageControl) Label(value string) ImageControl {
	return i.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (i ImageControl) LabelAlign(value string) ImageControl {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i ImageControl) LabelClassName(value string) ImageControl {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标并显示提示内容
func (i ImageControl) LabelRemark(value string) ImageControl {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (i ImageControl) LabelWidth(value string) ImageControl {
	return i.set("labelWidth", value)
}

// Limit 限制图片大小
func (i ImageControl) Limit(value string) ImageControl {
	return i.set("limit", value)
}

// MaxLength 最多的个数
func (i ImageControl) MaxLength(value string) ImageControl {
	return i.set("maxLength", value)
}

// MaxSize 默认没有限制，设置后，文件大小大于此值将不允许上传
func (i ImageControl) MaxSize(value string) ImageControl {
	return i.set("maxSize", value)
}

// Mode 表单项展示模式
func (i ImageControl) Mode(value string) ImageControl {
	return i.set("mode", value)
}

// Multiple 是否为多选
func (i ImageControl) Multiple(value bool) ImageControl {
	return i.set("multiple", value)
}

// Name 字段名
func (i ImageControl) Name(value string) ImageControl {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i ImageControl) OnEvent(value string) ImageControl {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i ImageControl) Placeholder(value string) ImageControl {
	return i.set("placeholder", value)
}

// ReCropable 是否允许二次裁剪
func (i ImageControl) ReCropable(value bool) ImageControl {
	return i.set("reCropable", value)
}

// ReadOnly 是否只读
func (i ImageControl) ReadOnly(value bool) ImageControl {
	return i.set("readOnly", value)
}

// Receiver 图片上传接口
func (i ImageControl) Receiver(value string) ImageControl {
	return i.set("receiver", value)
}

// Required 是否必填
func (i ImageControl) Required(value bool) ImageControl {
	return i.set("required", value)
}

// Schema 自定义 schema
func (i ImageControl) Schema(value string) ImageControl {
	return i.set("schema", value)
}

// ShowTips 是否显示提示
func (i ImageControl) ShowTips(value bool) ImageControl {
	return i.set("showTips", value)
}

// Size 尺寸
func (i ImageControl) Size(value string) ImageControl {
	return i.set("size", value)
}

// Src 默认图片地址
func (i ImageControl) Src(value string) ImageControl {
	return i.set("src", value)
}

// Step 上传时增量选择的大小
func (i ImageControl) Step(value string) ImageControl {
	return i.set("step", value)
}

// StrictMode 是否严格模式
func (i ImageControl) StrictMode(value bool) ImageControl {
	return i.set("strictMode", value)
}

// SubmitOnChange 是否在值变化时提交表单
func (i ImageControl) SubmitOnChange(value bool) ImageControl {
	return i.set("submitOnChange", value)
}

// UploadType 上传类型
func (i ImageControl) UploadType(value string) ImageControl {
	return i.set("uploadType", value)
}

// ValueFieldName 控件值的字段名
func (i ImageControl) ValueFieldName(value string) ImageControl {
	return i.set("valueFieldName", value)
}

// Vertical 设置为垂直模式
func (i ImageControl) Vertical(value bool) ImageControl {
	return i.set("vertical", value)
}

// VisibleOn 是否显示表达式
func (i ImageControl) VisibleOn(value string) ImageControl {
	return i.set("visibleOn", value)
}

// Width 宽度
func (i ImageControl) Width(value string) ImageControl {
	return i.set("width", value)
}
