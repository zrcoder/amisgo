package comp

// 替换为实际的 upload 包路径

// WangEditor
//
// @version 6.7.0
type WangEditor Schema

// NewWangEditor 创建一个新的 WangEditor 实例
func NewWangEditor() WangEditor {
	return WangEditor{}.set("type", "custom-wang-editor")
}

func (we WangEditor) set(key string, value interface{}) WangEditor {
	we[key] = value
	return we
}

// autoFill 自动填充
func (we WangEditor) AutoFill(value string) WangEditor {
	return we.set("autoFill", value)
}

// autoFocus 自动获取焦点
func (we WangEditor) AutoFocus(value bool) WangEditor {
	return we.set("autoFocus", value)
}

// className 设置类名
func (we WangEditor) ClassName(value string) WangEditor {
	return we.set("className", value)
}

// clearValueOnHidden 当前表单项隐藏时是否清除值
func (we WangEditor) ClearValueOnHidden(value bool) WangEditor {
	return we.set("clearValueOnHidden", value)
}

// description 描述
func (we WangEditor) Description(value string) WangEditor {
	return we.set("description", value)
}

// disabled 是否禁用
func (we WangEditor) Disabled(value bool) WangEditor {
	return we.set("disabled", value)
}

// disabledOn 当前表单项是否禁用的条件
func (we WangEditor) DisabledOn(value string) WangEditor {
	return we.set("disabledOn", value)
}

// excludeKeys 排除按钮
func (we WangEditor) ExcludeKeys(value string) WangEditor {
	return we.set("excludeKeys", value)
}

// height 编辑器高度 px
func (we WangEditor) Height(value string) WangEditor {
	return we.set("height", value)
}

// insertKeys 插入按钮
func (we WangEditor) InsertKeys(value string) WangEditor {
	return we.set("insertKeys", value)
}

// label 字段标签
func (we WangEditor) Label(value string) WangEditor {
	return we.set("label", value)
}

// labelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (we WangEditor) LabelAlign(value string) WangEditor {
	return we.set("labelAlign", value)
}

// labelRemark 表单项标签描述
func (we WangEditor) LabelRemark(value string) WangEditor {
	return we.set("labelRemark", value)
}

// maxLength 最大长度
func (we WangEditor) MaxLength(value string) WangEditor {
	return we.set("maxLength", value)
}

// name 字段名
func (we WangEditor) Name(value string) WangEditor {
	return we.set("name", value)
}

// placeholder 占位符
func (we WangEditor) Placeholder(value string) WangEditor {
	return we.set("placeholder", value)
}

// remark 备注
func (we WangEditor) Remark(value string) WangEditor {
	return we.set("remark", value)
}

// required 是否必填
func (we WangEditor) Required(value bool) WangEditor {
	return we.set("required", value)
}

// requiredOn 是否必填的条件
func (we WangEditor) RequiredOn(value string) WangEditor {
	return we.set("requiredOn", value)
}

// static 是否静态展示
func (we WangEditor) Static(value bool) WangEditor {
	return we.set("static", value)
}

// toolbarKeys 工具栏按钮
func (we WangEditor) ToolbarKeys(value string) WangEditor {
	return we.set("toolbarKeys", value)
}

// type 指定为 wang-editor 渲染器。
func (we WangEditor) Type(value string) WangEditor {
	return we.set("type", value)
}

// uploadImageMaxNumber 上传图片的最大数量 默认 100
func (we WangEditor) UploadImageMaxNumber(value string) WangEditor {
	return we.set("uploadImageMaxNumber", value)
}

// uploadImageMaxSize 上传图片的最大大小 单位: K 默认 2M
func (we WangEditor) UploadImageMaxSize(value string) WangEditor {
	return we.set("uploadImageMaxSize", value)
}

// uploadImageServer 上传图片的服务器地址
func (we WangEditor) UploadImageServer(value string) WangEditor {
	return we.set("uploadImageServer", value)
}

// uploadVideoMaxNumber 上传视频的最大数量 默认 5
func (we WangEditor) UploadVideoMaxNumber(value string) WangEditor {
	return we.set("uploadVideoMaxNumber", value)
}

// uploadVideoMaxSize 上传视频的最大大小 单位: K 默认 10M
func (we WangEditor) UploadVideoMaxSize(value string) WangEditor {
	return we.set("uploadVideoMaxSize", value)
}

// uploadVideoServer 上传视频的服务器地址
func (we WangEditor) UploadVideoServer(value string) WangEditor {
	return we.set("uploadVideoServer", value)
}

// validateApi 校验接口
func (we WangEditor) ValidateApi(value string) WangEditor {
	return we.set("validateApi", value)
}

// validationErrors 校验错误信息
func (we WangEditor) ValidationErrors(value string) WangEditor {
	return we.set("validationErrors", value)
}

// validations 校验规则
func (we WangEditor) Validations(value string) WangEditor {
	return we.set("validations", value)
}

// value 默认值
func (we WangEditor) Value(value string) WangEditor {
	return we.set("value", value)
}

// visible 是否显示
func (we WangEditor) Visible(value bool) WangEditor {
	return we.set("visible", value)
}

// visibleOn 当前表单项是否显示的条件
func (we WangEditor) VisibleOn(value string) WangEditor {
	return we.set("visibleOn", value)
}
