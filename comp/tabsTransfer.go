package comp

// tabsTransfer tabsTransfer 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer

type tabsTransfer Schema

func TabsTransfer() tabsTransfer {
	return tabsTransfer{}.set("type", "tabs-transfer")
}

func (t tabsTransfer) set(key string, value any) tabsTransfer {
	t[key] = value
	return t
}

// AddApi 添加时调用的接口
func (t tabsTransfer) AddApi(value string) tabsTransfer {
	return t.set("addApi", value)
}

// Adds 新增时的表单项。
func (t tabsTransfer) Adds(value string) tabsTransfer {
	return t.set("adds", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (t tabsTransfer) AddDialog(value string) tabsTransfer {
	return t.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (t tabsTransfer) AutoCheckChildren(value bool) tabsTransfer {
	return t.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t tabsTransfer) AutoFill(value string) tabsTransfer {
	return t.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t tabsTransfer) ClassName(value string) tabsTransfer {
	return t.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t tabsTransfer) ClearValueOnHidden(value bool) tabsTransfer {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t tabsTransfer) Clearable(value bool) tabsTransfer {
	return t.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (t tabsTransfer) Columns(value ...any) tabsTransfer {
	return t.set("columns", value)
}

// Creatable 是否可以新增
func (t tabsTransfer) Creatable(value bool) tabsTransfer {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t tabsTransfer) CreateBtnLabel(value string) tabsTransfer {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t tabsTransfer) DeferApi(value string) tabsTransfer {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t tabsTransfer) DeferField(value string) tabsTransfer {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t tabsTransfer) DeleteApi(value string) tabsTransfer {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t tabsTransfer) DeleteConfirmText(value string) tabsTransfer {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t tabsTransfer) Delimiter(value string) tabsTransfer {
	return t.set("delimiter", value)
}

// Desc
func (t tabsTransfer) Desc(value string) tabsTransfer {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t tabsTransfer) Description(value string) tabsTransfer {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (t tabsTransfer) DescriptionClassName(value string) tabsTransfer {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t tabsTransfer) Disabled(value bool) tabsTransfer {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t tabsTransfer) DisabledOn(value string) tabsTransfer {
	return t.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (t tabsTransfer) EditApi(value string) tabsTransfer {
	return t.set("editApi", value)
}

// Edits 选项修改的表单项
func (t tabsTransfer) Edits(value string) tabsTransfer {
	return t.set("edits", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (t tabsTransfer) EditDialog(value string) tabsTransfer {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t tabsTransfer) Editable(value bool) tabsTransfer {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t tabsTransfer) EditorSetting(value string) tabsTransfer {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t tabsTransfer) ExtraName(value string) tabsTransfer {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t tabsTransfer) ExtractValue(value bool) tabsTransfer {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t tabsTransfer) Hidden(value bool) tabsTransfer {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t tabsTransfer) HiddenOn(value string) tabsTransfer {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t tabsTransfer) Hint(value string) tabsTransfer {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (t tabsTransfer) Horizontal(value string) tabsTransfer {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t tabsTransfer) ID(value string) tabsTransfer {
	return t.set("id", value)
}

// InitAutoFill
func (t tabsTransfer) InitAutoFill(value string) tabsTransfer {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t tabsTransfer) InitFetch(value bool) tabsTransfer {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t tabsTransfer) InitFetchOn(value string) tabsTransfer {
	return t.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (t tabsTransfer) InitiallyOpen(value bool) tabsTransfer {
	return t.set("initiallyOpen", value)
}

// Inline 表单  是否为 inline 模式。
func (t tabsTransfer) Inline(value bool) tabsTransfer {
	return t.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (t tabsTransfer) InputClassName(value string) tabsTransfer {
	return t.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (t tabsTransfer) ItemHeight(value string) tabsTransfer {
	return t.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t tabsTransfer) JoinValues(value bool) tabsTransfer {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t tabsTransfer) Label(value string) tabsTransfer {
	return t.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (t tabsTransfer) LabelAlign(value string) tabsTransfer {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t tabsTransfer) LabelClassName(value string) tabsTransfer {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (t tabsTransfer) LabelRemark(value string) tabsTransfer {
	return t.set("labelRemark", value)
}

// LabelWidth 文字宽度
func (t tabsTransfer) LabelWidth(value string) tabsTransfer {
	return t.set("labelWidth", value)
}

// Level
func (t tabsTransfer) Level(value string) tabsTransfer {
	return t.set("level", value)
}

// Max 最大值
func (t tabsTransfer) Max(value int) tabsTransfer {
	return t.set("max", value)
}

// Min 最小值
func (t tabsTransfer) Min(value int) tabsTransfer {
	return t.set("min", value)
}

// Name 组件唯一名称，主要用于表单提交
func (t tabsTransfer) Name(value string) tabsTransfer {
	return t.set("name", value)
}

// Nested 是否支持嵌套，默认 false
func (t tabsTransfer) Nested(value bool) tabsTransfer {
	return t.set("nested", value)
}

// Options 选项
func (t tabsTransfer) Options(value ...MOption) tabsTransfer {
	return t.set("options", value)
}

// OptionTextField 选项文本字段
func (t tabsTransfer) OptionTextField(value string) tabsTransfer {
	return t.set("optionTextField", value)
}

// OptionValueField 选项值字段
func (t tabsTransfer) OptionValueField(value string) tabsTransfer {
	return t.set("optionValueField", value)
}

// PageField 分页字段
func (t tabsTransfer) PageField(value string) tabsTransfer {
	return t.set("pageField", value)
}

// Paginate 分页
func (t tabsTransfer) Paginate(value bool) tabsTransfer {
	return t.set("paginate", value)
}

// PageSizeField 每页数量字段
func (t tabsTransfer) PageSizeField(value string) tabsTransfer {
	return t.set("pageSizeField", value)
}

// Placeholder
func (t tabsTransfer) Placeholder(value string) tabsTransfer {
	return t.set("placeholder", value)
}

// ReadOnly 只读
func (t tabsTransfer) ReadOnly(value bool) tabsTransfer {
	return t.set("readOnly", value)
}

// ReadOnlyOn 表达式判断只读状态
func (t tabsTransfer) ReadOnlyOn(value string) tabsTransfer {
	return t.set("readOnlyOn", value)
}

// RemoveBtnLabel 删除文字
func (t tabsTransfer) RemoveBtnLabel(value string) tabsTransfer {
	return t.set("removeBtnLabel", value)
}

// Remote 响应数据接口
func (t tabsTransfer) Remote(value string) tabsTransfer {
	return t.set("remote", value)
}

// ResetValue 重置时的默认值
func (t tabsTransfer) ResetValue(value string) tabsTransfer {
	return t.set("resetValue", value)
}

// Searchable 是否可搜索
func (t tabsTransfer) Searchable(value bool) tabsTransfer {
	return t.set("searchable", value)
}

// ShowBorder 是否显示边框
func (t tabsTransfer) ShowBorder(value bool) tabsTransfer {
	return t.set("showBorder", value)
}

// ShowInputPlaceholder 是否展示输入提示
func (t tabsTransfer) ShowInputPlaceholder(value bool) tabsTransfer {
	return t.set("showInputPlaceholder", value)
}

// ShowLabel 是否显示 label
func (t tabsTransfer) ShowLabel(value bool) tabsTransfer {
	return t.set("showLabel", value)
}

// ShowRemovable 选项是否可以删除 (选项是否可以删除)
func (t tabsTransfer) ShowRemovable(value bool) tabsTransfer {
	return t.set("showRemovable", value)
}

// ShowSearch 选项是否可以搜索 (选项是否可以搜索)
func (t tabsTransfer) ShowSearch(value bool) tabsTransfer {
	return t.set("showSearch", value)
}

// Size 尺寸
func (t tabsTransfer) Size(value string) tabsTransfer {
	return t.set("size", value)
}

// Source 获取选项的接口
func (t tabsTransfer) Source(value string) tabsTransfer {
	return t.set("source", value)
}

// StaticClassName 配置静态部分的 className (配置静态部分的 className)
func (t tabsTransfer) StaticClassName(value string) tabsTransfer {
	return t.set("staticClassName", value)
}

// Step 步长
func (t tabsTransfer) Step(value int) tabsTransfer {
	return t.set("step", value)
}

// TabField Tab 分页字段
func (t tabsTransfer) TabField(value string) tabsTransfer {
	return t.set("tabField", value)
}

// Tabs Tab 页签的内容，支持模板，表达式等。
func (t tabsTransfer) Tabs(value string) tabsTransfer {
	return t.set("tabs", value)
}

// Toggable 是否可切换
func (t tabsTransfer) Toggable(value bool) tabsTransfer {
	return t.set("toggable", value)
}

// TreeMode 是否为树模式
func (t tabsTransfer) TreeMode(value bool) tabsTransfer {
	return t.set("treeMode", value)
}

// Value 选中的值
func (t tabsTransfer) Value(value string) tabsTransfer {
	return t.set("value", value)
}

// Visible 是否可见
func (t tabsTransfer) Visible(value bool) tabsTransfer {
	return t.set("visible", value)
}

// VisibleOn 表达式判断是否可见
func (t tabsTransfer) VisibleOn(value string) tabsTransfer {
	return t.set("visibleOn", value)
}

// Width 宽度
func (t tabsTransfer) Width(value string) tabsTransfer {
	return t.set("width", value)
}

// WrapperClassName 配置 wrapper className
func (t tabsTransfer) WrapperClassName(value string) tabsTransfer {
	return t.set("wrapperClassName", value)
}

// WrapperStyle 配置 wrapper style
func (t tabsTransfer) WrapperStyle(value any) tabsTransfer {
	return t.set("wrapperStyle", value)
}
