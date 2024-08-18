package comp

// RadiosControl 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios
//
// @version 6.7.0
type RadiosControl Schema

// NewRadiosControl 创建一个新的 RadiosControl 实例
func NewRadiosControl() RadiosControl {
	return RadiosControl{}.set("type", "radios")
}

func (rc RadiosControl) set(key string, value interface{}) RadiosControl {
	rc[key] = value
	return rc
}

// AddApi 添加时调用的接口
func (rc RadiosControl) AddApi(value string) RadiosControl {
	return rc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (rc RadiosControl) AddControls(value string) RadiosControl {
	return rc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (rc RadiosControl) AddDialog(value string) RadiosControl {
	return rc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc RadiosControl) AutoFill(value string) RadiosControl {
	return rc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc RadiosControl) ClassName(value string) RadiosControl {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc RadiosControl) ClearValueOnHidden(value bool) RadiosControl {
	return rc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (rc RadiosControl) Clearable(value bool) RadiosControl {
	return rc.set("clearable", value)
}

// ColumnsCount 每行显示多少个
func (rc RadiosControl) ColumnsCount(value string) RadiosControl {
	return rc.set("columnsCount", value)
}

// Creatable 是否可以新增
func (rc RadiosControl) Creatable(value bool) RadiosControl {
	return rc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (rc RadiosControl) CreateBtnLabel(value string) RadiosControl {
	return rc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (rc RadiosControl) DeferApi(value string) RadiosControl {
	return rc.set("deferApi", value)
}

// DeferField 懒加载字段
func (rc RadiosControl) DeferField(value string) RadiosControl {
	return rc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (rc RadiosControl) DeleteApi(value string) RadiosControl {
	return rc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (rc RadiosControl) DeleteConfirmText(value string) RadiosControl {
	return rc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (rc RadiosControl) Delimiter(value string) RadiosControl {
	return rc.set("delimiter", value)
}

// Desc
func (rc RadiosControl) Desc(value string) RadiosControl {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc RadiosControl) Description(value string) RadiosControl {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc RadiosControl) DescriptionClassName(value string) RadiosControl {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc RadiosControl) Disabled(value bool) RadiosControl {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadiosControl) DisabledOn(value string) RadiosControl {
	return rc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (rc RadiosControl) EditApi(value string) RadiosControl {
	return rc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (rc RadiosControl) EditControls(value string) RadiosControl {
	return rc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (rc RadiosControl) EditDialog(value string) RadiosControl {
	return rc.set("editDialog", value)
}

// Editable 是否可以编辑
func (rc RadiosControl) Editable(value bool) RadiosControl {
	return rc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc RadiosControl) EditorSetting(value string) RadiosControl {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc RadiosControl) ExtraName(value string) RadiosControl {
	return rc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (rc RadiosControl) ExtractValue(value bool) RadiosControl {
	return rc.set("extractValue", value)
}

// Hidden 是否隐藏
func (rc RadiosControl) Hidden(value bool) RadiosControl {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadiosControl) HiddenOn(value string) RadiosControl {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc RadiosControl) Hint(value string) RadiosControl {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (rc RadiosControl) Horizontal(value string) RadiosControl {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc RadiosControl) Id(value string) RadiosControl {
	return rc.set("id", value)
}

// InitAutoFill
func (rc RadiosControl) InitAutoFill(value string) RadiosControl {
	return rc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (rc RadiosControl) InitFetch(value bool) RadiosControl {
	return rc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (rc RadiosControl) InitFetchOn(value string) RadiosControl {
	return rc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc RadiosControl) Inline(value bool) RadiosControl {
	return rc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (rc RadiosControl) InputClassName(value string) RadiosControl {
	return rc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (rc RadiosControl) JoinValues(value bool) RadiosControl {
	return rc.set("joinValues", value)
}

// Label 描述标题
func (rc RadiosControl) Label(value string) RadiosControl {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc RadiosControl) LabelAlign(value string) RadiosControl {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc RadiosControl) LabelClassName(value string) RadiosControl {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (rc RadiosControl) LabelRemark(value string) RadiosControl {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc RadiosControl) LabelWidth(value string) RadiosControl {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc RadiosControl) Mode(value string) RadiosControl {
	return rc.set("mode", value)
}

// Multiple 是否为多选模式
func (rc RadiosControl) Multiple(value bool) RadiosControl {
	return rc.set("multiple", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc RadiosControl) Name(value string) RadiosControl {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc RadiosControl) OnEvent(value interface{}) RadiosControl {
	return rc.set("onEvent", value)
}

// Option 配置单选框的选项。 (配置单选框的选项。)
func (rc RadiosControl) Option(value interface{}) RadiosControl {
	return rc.set("option", value)
}

// Options 配置选项列表
func (rc RadiosControl) Options(value interface{}) RadiosControl {
	return rc.set("options", value)
}

// Placeholder 占位符
func (rc RadiosControl) Placeholder(value string) RadiosControl {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc RadiosControl) ReadOnly(value bool) RadiosControl {
	return rc.set("readOnly", value)
}

// Removable 是否可移除
func (rc RadiosControl) Removable(value bool) RadiosControl {
	return rc.set("removable", value)
}

// RenderMode 渲染模式
func (rc RadiosControl) RenderMode(value string) RadiosControl {
	return rc.set("renderMode", value)
}

// Required 是否必填
func (rc RadiosControl) Required(value bool) RadiosControl {
	return rc.set("required", value)
}

// ShowIcon 是否显示图标
func (rc RadiosControl) ShowIcon(value bool) RadiosControl {
	return rc.set("showIcon", value)
}

// ShowRadio 是否显示单选框
func (rc RadiosControl) ShowRadio(value bool) RadiosControl {
	return rc.set("showRadio", value)
}

// ShowText 是否显示文本
func (rc RadiosControl) ShowText(value bool) RadiosControl {
	return rc.set("showText", value)
}

// Size 控件大小 (配置控件大小。)
func (rc RadiosControl) Size(value string) RadiosControl {
	return rc.set("size", value)
}

// Source 数据源接口
func (rc RadiosControl) Source(value string) RadiosControl {
	return rc.set("source", value)
}

// SourceApi 自定义 source 接口
func (rc RadiosControl) SourceApi(value string) RadiosControl {
	return rc.set("sourceApi", value)
}

// Toggle 是否为切换模式
func (rc RadiosControl) Toggle(value bool) RadiosControl {
	return rc.set("toggle", value)
}

// Value 当前值
func (rc RadiosControl) Value(value interface{}) RadiosControl {
	return rc.set("value", value)
}

// Visible 控件是否可见
func (rc RadiosControl) Visible(value bool) RadiosControl {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (rc RadiosControl) VisibleOn(value string) RadiosControl {
	return rc.set("visibleOn", value)
}

// Width 控件宽度
func (rc RadiosControl) Width(value string) RadiosControl {
	return rc.set("width", value)
}

// WrapperClassName 设置包裹元素的 CSS 类名
func (rc RadiosControl) WrapperClassName(value string) RadiosControl {
	return rc.set("wrapperClassName", value)
}

// ValueField 数据字段
func (rc RadiosControl) ValueField(value string) RadiosControl {
	return rc.set("valueField", value)
}

// LabelField 标签字段
func (rc RadiosControl) LabelField(value string) RadiosControl {
	return rc.set("labelField", value)
}

// DescriptionField 描述字段
func (rc RadiosControl) DescriptionField(value string) RadiosControl {
	return rc.set("descriptionField", value)
}

// OptionValueField 选项值字段
func (rc RadiosControl) OptionValueField(value string) RadiosControl {
	return rc.set("optionValueField", value)
}

// OptionLabelField 选项标签字段
func (rc RadiosControl) OptionLabelField(value string) RadiosControl {
	return rc.set("optionLabelField", value)
}

// OptionDescriptionField 选项描述字段
func (rc RadiosControl) OptionDescriptionField(value string) RadiosControl {
	return rc.set("optionDescriptionField", value)
}
