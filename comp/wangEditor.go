package comp

// 替换为实际的 upload 包路径

// wangEditor

type wangEditor schema

// WangEditor 创建一个新的 WangEditor 实例
func WangEditor() wangEditor {
	return wangEditor{}.set("type", "custom-wang-editor")
}

func (we wangEditor) set(key string, value any) wangEditor {
	we[key] = value
	return we
}

// autoFill 自动填充
func (we wangEditor) AutoFill(value string) wangEditor {
	return we.set("autoFill", value)
}

// autoFocus 自动获取焦点
func (we wangEditor) AutoFocus(value bool) wangEditor {
	return we.set("autoFocus", value)
}

// className 设置类名
func (we wangEditor) ClassName(value string) wangEditor {
	return we.set("className", value)
}

// clearValueOnHidden 当前表单项隐藏时是否清除值
func (we wangEditor) ClearValueOnHidden(value bool) wangEditor {
	return we.set("clearValueOnHidden", value)
}

// description 描述
func (we wangEditor) Description(value string) wangEditor {
	return we.set("description", value)
}

// disabled 是否禁用
func (we wangEditor) Disabled(value bool) wangEditor {
	return we.set("disabled", value)
}

// disabledOn 当前表单项是否禁用的条件
func (we wangEditor) DisabledOn(value string) wangEditor {
	return we.set("disabledOn", value)
}

// excludeKeys 排除按钮
func (we wangEditor) ExcludeKeys(value string) wangEditor {
	return we.set("excludeKeys", value)
}

// height 编辑器高度 px
func (we wangEditor) Height(value string) wangEditor {
	return we.set("height", value)
}

// insertKeys 插入按钮
func (we wangEditor) InsertKeys(value string) wangEditor {
	return we.set("insertKeys", value)
}

// label 字段标签
func (we wangEditor) Label(value string) wangEditor {
	return we.set("label", value)
}

// labelAlign 表单项标签对齐方式，默认右对齐，仅在 mode为horizontal 时生效
func (we wangEditor) LabelAlign(value string) wangEditor {
	return we.set("labelAlign", value)
}

// labelRemark 表单项标签描述
func (we wangEditor) LabelRemark(value string) wangEditor {
	return we.set("labelRemark", value)
}

// maxLength 最大长度
func (we wangEditor) MaxLength(value string) wangEditor {
	return we.set("maxLength", value)
}

// name 字段名
func (we wangEditor) Name(value string) wangEditor {
	return we.set("name", value)
}

// placeholder 占位符
func (we wangEditor) Placeholder(value string) wangEditor {
	return we.set("placeholder", value)
}

// remark 备注
func (we wangEditor) Remark(value string) wangEditor {
	return we.set("remark", value)
}

// required 是否必填
func (we wangEditor) Required(value bool) wangEditor {
	return we.set("required", value)
}

// requiredOn 是否必填的条件
func (we wangEditor) RequiredOn(value string) wangEditor {
	return we.set("requiredOn", value)
}

// static 是否静态展示
func (we wangEditor) Static(value bool) wangEditor {
	return we.set("static", value)
}

// toolbarKeys 工具栏按钮
func (we wangEditor) ToolbarKeys(value string) wangEditor {
	return we.set("toolbarKeys", value)
}

// uploadImageMaxNumber 上传图片的最大数量 默认 100
func (we wangEditor) UploadImageMaxNumber(value string) wangEditor {
	return we.set("uploadImageMaxNumber", value)
}

// uploadImageMaxSize 上传图片的最大大小 单位: K 默认 2M
func (we wangEditor) UploadImageMaxSize(value string) wangEditor {
	return we.set("uploadImageMaxSize", value)
}

// uploadImageServer 上传图片的服务器地址
func (we wangEditor) UploadImageServer(value string) wangEditor {
	return we.set("uploadImageServer", value)
}

// uploadVideoMaxNumber 上传视频的最大数量 默认 5
func (we wangEditor) UploadVideoMaxNumber(value string) wangEditor {
	return we.set("uploadVideoMaxNumber", value)
}

// uploadVideoMaxSize 上传视频的最大大小 单位: K 默认 10M
func (we wangEditor) UploadVideoMaxSize(value string) wangEditor {
	return we.set("uploadVideoMaxSize", value)
}

// uploadVideoServer 上传视频的服务器地址
func (we wangEditor) UploadVideoServer(value string) wangEditor {
	return we.set("uploadVideoServer", value)
}

// validateApi 校验接口
func (we wangEditor) ValidateApi(value string) wangEditor {
	return we.set("validateApi", value)
}

// validationErrors 校验错误信息
func (we wangEditor) ValidationErrors(value string) wangEditor {
	return we.set("validationErrors", value)
}

// validations 校验规则
func (we wangEditor) Validations(value string) wangEditor {
	return we.set("validations", value)
}

// value 默认值
func (we wangEditor) Value(value string) wangEditor {
	return we.set("value", value)
}

// visible 是否显示
func (we wangEditor) Visible(value bool) wangEditor {
	return we.set("visible", value)
}

// visibleOn 当前表单项是否显示的条件
func (we wangEditor) VisibleOn(value string) wangEditor {
	return we.set("visibleOn", value)
}
