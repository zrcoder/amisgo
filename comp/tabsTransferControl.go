package comp

// TabsTransferControl TabsTransfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer
//
// @version 6.7.0
type TabsTransferControl Schema

// NewTabsTransferControl 创建一个新的 TabsTransferControl 实例
func NewTabsTransferControl() TabsTransferControl {
	return TabsTransferControl{}.set("type", "tabs-transfer")
}

func (t TabsTransferControl) set(key string, value interface{}) TabsTransferControl {
	t[key] = value
	return t
}

// AddApi 添加时调用的接口
func (t TabsTransferControl) AddApi(value string) TabsTransferControl {
	return t.set("addApi", value)
}

// AddControls 新增时的表单项。
func (t TabsTransferControl) AddControls(value string) TabsTransferControl {
	return t.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (t TabsTransferControl) AddDialog(value string) TabsTransferControl {
	return t.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (t TabsTransferControl) AutoCheckChildren(value bool) TabsTransferControl {
	return t.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t TabsTransferControl) AutoFill(value string) TabsTransferControl {
	return t.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TabsTransferControl) ClassName(value string) TabsTransferControl {
	return t.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t TabsTransferControl) ClearValueOnHidden(value bool) TabsTransferControl {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t TabsTransferControl) Clearable(value bool) TabsTransferControl {
	return t.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (t TabsTransferControl) Columns(value string) TabsTransferControl {
	return t.set("columns", value)
}

// Creatable 是否可以新增
func (t TabsTransferControl) Creatable(value bool) TabsTransferControl {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t TabsTransferControl) CreateBtnLabel(value string) TabsTransferControl {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t TabsTransferControl) DeferApi(value string) TabsTransferControl {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t TabsTransferControl) DeferField(value string) TabsTransferControl {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t TabsTransferControl) DeleteApi(value string) TabsTransferControl {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t TabsTransferControl) DeleteConfirmText(value string) TabsTransferControl {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t TabsTransferControl) Delimiter(value string) TabsTransferControl {
	return t.set("delimiter", value)
}

// Desc
func (t TabsTransferControl) Desc(value string) TabsTransferControl {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t TabsTransferControl) Description(value string) TabsTransferControl {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (t TabsTransferControl) DescriptionClassName(value string) TabsTransferControl {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t TabsTransferControl) Disabled(value bool) TabsTransferControl {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t TabsTransferControl) DisabledOn(value string) TabsTransferControl {
	return t.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (t TabsTransferControl) EditApi(value string) TabsTransferControl {
	return t.set("editApi", value)
}

// EditControls 选项修改的表单项
func (t TabsTransferControl) EditControls(value string) TabsTransferControl {
	return t.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (t TabsTransferControl) EditDialog(value string) TabsTransferControl {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t TabsTransferControl) Editable(value bool) TabsTransferControl {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t TabsTransferControl) EditorSetting(value string) TabsTransferControl {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t TabsTransferControl) ExtraName(value string) TabsTransferControl {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t TabsTransferControl) ExtractValue(value bool) TabsTransferControl {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t TabsTransferControl) Hidden(value bool) TabsTransferControl {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t TabsTransferControl) HiddenOn(value string) TabsTransferControl {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t TabsTransferControl) Hint(value string) TabsTransferControl {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (t TabsTransferControl) Horizontal(value string) TabsTransferControl {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t TabsTransferControl) Id(value string) TabsTransferControl {
	return t.set("id", value)
}

// InitAutoFill
func (t TabsTransferControl) InitAutoFill(value string) TabsTransferControl {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t TabsTransferControl) InitFetch(value bool) TabsTransferControl {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t TabsTransferControl) InitFetchOn(value string) TabsTransferControl {
	return t.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (t TabsTransferControl) InitiallyOpen(value bool) TabsTransferControl {
	return t.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t TabsTransferControl) Inline(value bool) TabsTransferControl {
	return t.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (t TabsTransferControl) InputClassName(value string) TabsTransferControl {
	return t.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (t TabsTransferControl) ItemHeight(value string) TabsTransferControl {
	return t.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t TabsTransferControl) JoinValues(value bool) TabsTransferControl {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t TabsTransferControl) Label(value string) TabsTransferControl {
	return t.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (t TabsTransferControl) LabelAlign(value string) TabsTransferControl {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t TabsTransferControl) LabelClassName(value string) TabsTransferControl {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (t TabsTransferControl) LabelRemark(value string) TabsTransferControl {
	return t.set("labelRemark", value)
}

// LabelWidth 文字宽度
func (t TabsTransferControl) LabelWidth(value string) TabsTransferControl {
	return t.set("labelWidth", value)
}

// Level
func (t TabsTransferControl) Level(value string) TabsTransferControl {
	return t.set("level", value)
}

// Max 最大值
func (t TabsTransferControl) Max(value int) TabsTransferControl {
	return t.set("max", value)
}

// Min 最小值
func (t TabsTransferControl) Min(value int) TabsTransferControl {
	return t.set("min", value)
}

// Name 组件唯一名称，主要用于表单提交
func (t TabsTransferControl) Name(value string) TabsTransferControl {
	return t.set("name", value)
}

// Nested 是否支持嵌套，默认 false
func (t TabsTransferControl) Nested(value bool) TabsTransferControl {
	return t.set("nested", value)
}

// Options 选项
func (t TabsTransferControl) Options(value string) TabsTransferControl {
	return t.set("options", value)
}

// OptionTextField 选项文本字段
func (t TabsTransferControl) OptionTextField(value string) TabsTransferControl {
	return t.set("optionTextField", value)
}

// OptionValueField 选项值字段
func (t TabsTransferControl) OptionValueField(value string) TabsTransferControl {
	return t.set("optionValueField", value)
}

// PageField 分页字段
func (t TabsTransferControl) PageField(value string) TabsTransferControl {
	return t.set("pageField", value)
}

// Paginate 分页
func (t TabsTransferControl) Paginate(value bool) TabsTransferControl {
	return t.set("paginate", value)
}

// PageSizeField 每页数量字段
func (t TabsTransferControl) PageSizeField(value string) TabsTransferControl {
	return t.set("pageSizeField", value)
}

// Placeholder
func (t TabsTransferControl) Placeholder(value string) TabsTransferControl {
	return t.set("placeholder", value)
}

// ReadOnly 只读
func (t TabsTransferControl) ReadOnly(value bool) TabsTransferControl {
	return t.set("readOnly", value)
}

// ReadOnlyOn 表达式判断只读状态
func (t TabsTransferControl) ReadOnlyOn(value string) TabsTransferControl {
	return t.set("readOnlyOn", value)
}

// RemoveBtnLabel 删除文字
func (t TabsTransferControl) RemoveBtnLabel(value string) TabsTransferControl {
	return t.set("removeBtnLabel", value)
}

// Remote 响应数据接口
func (t TabsTransferControl) Remote(value string) TabsTransferControl {
	return t.set("remote", value)
}

// ResetValue 重置时的默认值
func (t TabsTransferControl) ResetValue(value string) TabsTransferControl {
	return t.set("resetValue", value)
}

// Searchable 是否可搜索
func (t TabsTransferControl) Searchable(value bool) TabsTransferControl {
	return t.set("searchable", value)
}

// ShowBorder 是否显示边框
func (t TabsTransferControl) ShowBorder(value bool) TabsTransferControl {
	return t.set("showBorder", value)
}

// ShowInputPlaceholder 是否展示输入提示
func (t TabsTransferControl) ShowInputPlaceholder(value bool) TabsTransferControl {
	return t.set("showInputPlaceholder", value)
}

// ShowLabel 是否显示 label
func (t TabsTransferControl) ShowLabel(value bool) TabsTransferControl {
	return t.set("showLabel", value)
}

// ShowRemovable 选项是否可以删除 (选项是否可以删除)
func (t TabsTransferControl) ShowRemovable(value bool) TabsTransferControl {
	return t.set("showRemovable", value)
}

// ShowSearch 选项是否可以搜索 (选项是否可以搜索)
func (t TabsTransferControl) ShowSearch(value bool) TabsTransferControl {
	return t.set("showSearch", value)
}

// Size 尺寸
func (t TabsTransferControl) Size(value string) TabsTransferControl {
	return t.set("size", value)
}

// Source 获取选项的接口
func (t TabsTransferControl) Source(value string) TabsTransferControl {
	return t.set("source", value)
}

// StaticClassName 配置静态部分的 className (配置静态部分的 className)
func (t TabsTransferControl) StaticClassName(value string) TabsTransferControl {
	return t.set("staticClassName", value)
}

// Step 步长
func (t TabsTransferControl) Step(value int) TabsTransferControl {
	return t.set("step", value)
}

// TabField Tab 分页字段
func (t TabsTransferControl) TabField(value string) TabsTransferControl {
	return t.set("tabField", value)
}

// Tabs Tab 页签的内容，支持模板，表达式等。
func (t TabsTransferControl) Tabs(value string) TabsTransferControl {
	return t.set("tabs", value)
}

// Toggable 是否可切换
func (t TabsTransferControl) Toggable(value bool) TabsTransferControl {
	return t.set("toggable", value)
}

// TreeMode 是否为树模式
func (t TabsTransferControl) TreeMode(value bool) TabsTransferControl {
	return t.set("treeMode", value)
}

// Value 选中的值
func (t TabsTransferControl) Value(value string) TabsTransferControl {
	return t.set("value", value)
}

// Visible 是否可见
func (t TabsTransferControl) Visible(value bool) TabsTransferControl {
	return t.set("visible", value)
}

// VisibleOn 表达式判断是否可见
func (t TabsTransferControl) VisibleOn(value string) TabsTransferControl {
	return t.set("visibleOn", value)
}

// Width 宽度
func (t TabsTransferControl) Width(value string) TabsTransferControl {
	return t.set("width", value)
}

// WrapperClassName 配置 wrapper className
func (t TabsTransferControl) WrapperClassName(value string) TabsTransferControl {
	return t.set("wrapperClassName", value)
}

// WrapperStyle 配置 wrapper style
func (t TabsTransferControl) WrapperStyle(value string) TabsTransferControl {
	return t.set("wrapperStyle", value)
}
