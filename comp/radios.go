package comp

// radios 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

type radios Schema

// Radios 创建一个新的 RadiosControl 实例
func Radios() radios {
	return radios{}.set("type", "radios")
}

func (rc radios) set(key string, value any) radios {
	rc[key] = value
	return rc
}

// AddApi 添加时调用的接口
func (rc radios) AddApi(value string) radios {
	return rc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (rc radios) AddControls(value string) radios {
	return rc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (rc radios) AddDialog(value string) radios {
	return rc.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc radios) AutoFill(value string) radios {
	return rc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc radios) ClassName(value string) radios {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc radios) ClearValueOnHidden(value bool) radios {
	return rc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (rc radios) Clearable(value bool) radios {
	return rc.set("clearable", value)
}

// ColumnsCount 每行显示多少个
func (rc radios) ColumnsCount(value string) radios {
	return rc.set("columnsCount", value)
}

// Creatable 是否可以新增
func (rc radios) Creatable(value bool) radios {
	return rc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (rc radios) CreateBtnLabel(value string) radios {
	return rc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (rc radios) DeferApi(value string) radios {
	return rc.set("deferApi", value)
}

// DeferField 懒加载字段
func (rc radios) DeferField(value string) radios {
	return rc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (rc radios) DeleteApi(value string) radios {
	return rc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (rc radios) DeleteConfirmText(value string) radios {
	return rc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (rc radios) Delimiter(value string) radios {
	return rc.set("delimiter", value)
}

// Desc
func (rc radios) Desc(value string) radios {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc radios) Description(value string) radios {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc radios) DescriptionClassName(value string) radios {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc radios) Disabled(value bool) radios {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radios) DisabledOn(value string) radios {
	return rc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (rc radios) EditApi(value string) radios {
	return rc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (rc radios) EditControls(value string) radios {
	return rc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (rc radios) EditDialog(value string) radios {
	return rc.set("editDialog", value)
}

// Editable 是否可以编辑
func (rc radios) Editable(value bool) radios {
	return rc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc radios) EditorSetting(value string) radios {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc radios) ExtraName(value string) radios {
	return rc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (rc radios) ExtractValue(value bool) radios {
	return rc.set("extractValue", value)
}

// Hidden 是否隐藏
func (rc radios) Hidden(value bool) radios {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radios) HiddenOn(value string) radios {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc radios) Hint(value string) radios {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (rc radios) Horizontal(value string) radios {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc radios) ID(value string) radios {
	return rc.set("id", value)
}

// InitAutoFill
func (rc radios) InitAutoFill(value string) radios {
	return rc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (rc radios) InitFetch(value bool) radios {
	return rc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (rc radios) InitFetchOn(value string) radios {
	return rc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc radios) Inline(value bool) radios {
	return rc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (rc radios) InputClassName(value string) radios {
	return rc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (rc radios) JoinValues(value bool) radios {
	return rc.set("joinValues", value)
}

// Label 描述标题
func (rc radios) Label(value string) radios {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc radios) LabelAlign(value string) radios {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc radios) LabelClassName(value string) radios {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (rc radios) LabelRemark(value string) radios {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc radios) LabelWidth(value string) radios {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc radios) Mode(value string) radios {
	return rc.set("mode", value)
}

// Multiple 是否为多选模式
func (rc radios) Multiple(value bool) radios {
	return rc.set("multiple", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc radios) Name(value string) radios {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc radios) OnEvent(value any) radios {
	return rc.set("onEvent", value)
}

// Option 配置单选框的选项。 (配置单选框的选项。)
func (rc radios) Option(value any) radios {
	return rc.set("option", value)
}

// Options 配置选项列表
func (rc radios) Options(value ...any) radios {
	return rc.set("options", value)
}

// Placeholder 占位符
func (rc radios) Placeholder(value string) radios {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc radios) ReadOnly(value bool) radios {
	return rc.set("readOnly", value)
}

// Removable 是否可移除
func (rc radios) Removable(value bool) radios {
	return rc.set("removable", value)
}

// RenderMode 渲染模式
func (rc radios) RenderMode(value string) radios {
	return rc.set("renderMode", value)
}

// Required 是否必填
func (rc radios) Required(value bool) radios {
	return rc.set("required", value)
}

// ShowIcon 是否显示图标
func (rc radios) ShowIcon(value bool) radios {
	return rc.set("showIcon", value)
}

// ShowRadio 是否显示单选框
func (rc radios) ShowRadio(value bool) radios {
	return rc.set("showRadio", value)
}

// ShowText 是否显示文本
func (rc radios) ShowText(value bool) radios {
	return rc.set("showText", value)
}

// Size 控件大小 (配置控件大小。)
func (rc radios) Size(value string) radios {
	return rc.set("size", value)
}

// Source 数据源接口
func (rc radios) Source(value string) radios {
	return rc.set("source", value)
}

// SourceApi 自定义 source 接口
func (rc radios) SourceApi(value string) radios {
	return rc.set("sourceApi", value)
}

// Toggle 是否为切换模式
func (rc radios) Toggle(value bool) radios {
	return rc.set("toggle", value)
}

// Value 当前值
func (rc radios) Value(value any) radios {
	return rc.set("value", value)
}

// Visible 控件是否可见
func (rc radios) Visible(value bool) radios {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (rc radios) VisibleOn(value string) radios {
	return rc.set("visibleOn", value)
}

// Width 控件宽度
func (rc radios) Width(value string) radios {
	return rc.set("width", value)
}

// WrapperClassName 设置包裹元素的 CSS 类名
func (rc radios) WrapperClassName(value string) radios {
	return rc.set("wrapperClassName", value)
}

// ValueField 数据字段
func (rc radios) ValueField(value string) radios {
	return rc.set("valueField", value)
}

// LabelField 标签字段
func (rc radios) LabelField(value string) radios {
	return rc.set("labelField", value)
}

// DescriptionField 描述字段
func (rc radios) DescriptionField(value string) radios {
	return rc.set("descriptionField", value)
}

// OptionValueField 选项值字段
func (rc radios) OptionValueField(value string) radios {
	return rc.set("optionValueField", value)
}

// OptionLabelField 选项标签字段
func (rc radios) OptionLabelField(value string) radios {
	return rc.set("optionLabelField", value)
}

// OptionDescriptionField 选项描述字段
func (rc radios) OptionDescriptionField(value string) radios {
	return rc.set("optionDescriptionField", value)
}
