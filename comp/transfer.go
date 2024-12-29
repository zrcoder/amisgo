package comp

// transfer 控件定义

type transfer Schema

// NewTransferControl 创建一个新的 TransferControl 实例
func Transfer() transfer {
	return transfer{}.set("type", "transfer")
}

func (tc transfer) set(key string, value any) transfer {
	tc[key] = value
	return tc
}

// AddApi 添加时调用的接口
func (tc transfer) AddApi(value string) transfer {
	return tc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tc transfer) AddControls(value string) transfer {
	return tc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (tc transfer) AddDialog(value string) transfer {
	return tc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tc transfer) AutoCheckChildren(value bool) transfer {
	return tc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc transfer) AutoFill(value string) transfer {
	return tc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (tc transfer) ClassName(value string) transfer {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc transfer) ClearValueOnHidden(value bool) transfer {
	return tc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tc transfer) Clearable(value bool) transfer {
	return tc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tc transfer) Columns(value ...any) transfer {
	return tc.set("columns", value)
}

// Creatable 是否可以新增
func (tc transfer) Creatable(value bool) transfer {
	return tc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tc transfer) CreateBtnLabel(value string) transfer {
	return tc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tc transfer) DeferApi(value string) transfer {
	return tc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tc transfer) DeferField(value string) transfer {
	return tc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tc transfer) DeleteApi(value string) transfer {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tc transfer) DeleteConfirmText(value string) transfer {
	return tc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tc transfer) Delimiter(value string) transfer {
	return tc.set("delimiter", value)
}

// Desc
func (tc transfer) Desc(value string) transfer {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc transfer) Description(value string) transfer {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc transfer) DescriptionClassName(value string) transfer {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc transfer) Disabled(value bool) transfer {
	return tc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc transfer) DisabledOn(value string) transfer {
	return tc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tc transfer) EditApi(value string) transfer {
	return tc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tc transfer) EditControls(value string) transfer {
	return tc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (tc transfer) EditDialog(value string) transfer {
	return tc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tc transfer) Editable(value bool) transfer {
	return tc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc transfer) EditorSetting(value string) transfer {
	return tc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc transfer) ExtraName(value string) transfer {
	return tc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tc transfer) ExtractValue(value bool) transfer {
	return tc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tc transfer) Hidden(value bool) transfer {
	return tc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tc transfer) HiddenOn(value string) transfer {
	return tc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tc transfer) Hint(value string) transfer {
	return tc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tc transfer) Horizontal(value string) transfer {
	return tc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tc transfer) ID(value string) transfer {
	return tc.set("id", value)
}

// InitAutoFill
func (tc transfer) InitAutoFill(value string) transfer {
	return tc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tc transfer) InitFetch(value bool) transfer {
	return tc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tc transfer) InitFetchOn(value string) transfer {
	return tc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tc transfer) InitiallyOpen(value bool) transfer {
	return tc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tc transfer) Inline(value bool) transfer {
	return tc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tc transfer) InputClassName(value string) transfer {
	return tc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tc transfer) ItemHeight(value string) transfer {
	return tc.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (tc transfer) JoinValues(value bool) transfer {
	return tc.set("joinValues", value)
}

// Label 描述标题
func (tc transfer) Label(value string) transfer {
	return tc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (tc transfer) LabelAlign(value string) transfer {
	return tc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (tc transfer) LabelClassName(value string) transfer {
	return tc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (tc transfer) LabelRemark(value string) transfer {
	return tc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (tc transfer) LabelWidth(value string) transfer {
	return tc.set("labelWidth", value)
}

// LeftMode 当 selectMode 为 associated 时用来定义左侧的选择模式 可选值: tree | list
func (tc transfer) LeftMode(value string) transfer {
	return tc.set("leftMode", value)
}

// Max 可选项的最大值。
func (tc transfer) Max(value int) transfer {
	return tc.set("max", value)
}

// Min 可选项的最小值。
func (tc transfer) Min(value int) transfer {
	return tc.set("min", value)
}

// Name 表单项名称，内部表单渲染时需要
func (tc transfer) Name(value string) transfer {
	return tc.set("name", value)
}

// OnChange 变化回调函数。
func (tc transfer) OnChange(value string) transfer {
	return tc.set("onChange", value)
}

// OnEvent 事件处理
func (tc transfer) OnEvent(value any) transfer {
	return tc.set("onEvent", value)
}

// Option
func (tc transfer) Option(value string) transfer {
	return tc.set("option", value)
}

// Options 选项数据
func (tc transfer) Options(value ...any) transfer {
	return tc.set("options", value)
}

// Placeholder 输入框提示语
func (tc transfer) Placeholder(value string) transfer {
	return tc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tc transfer) ReadOnly(value bool) transfer {
	return tc.set("readOnly", value)
}

// RemoteOptions 远程获取选项
func (tc transfer) RemoteOptions(value ...any) transfer {
	return tc.set("remoteOptions", value)
}

// Render 表单渲染模式，通常在内嵌时用
func (tc transfer) Render(value string) transfer {
	return tc.set("render", value)
}

// ResetValue 清空表单项内容时，设置该项的默认值
func (tc transfer) ResetValue(value string) transfer {
	return tc.set("resetValue", value)
}

// SaveImmediately 表单项内容变更是否立即保存
func (tc transfer) SaveImmediately(value bool) transfer {
	return tc.set("saveImmediately", value)
}

// Source 选项数据来源
func (tc transfer) Source(value string) transfer {
	return tc.set("source", value)
}

// Searchable
func (tc transfer) Searchable(value bool) transfer {
	return tc.set("searchable", value)
}

// SearchApi
func (tc transfer) SearchApi(value string) transfer {
	return tc.set("searchApi", value)
}

// SelectMode
func (tc transfer) SelectMode(value string) transfer {
	return tc.set("selectMode", value)
}

// Sortable
func (tc transfer) Sortable(value bool) transfer {
	return tc.set("sortable", value)
}

// StaticMode 控件是否静态显示
func (tc transfer) StaticMode(value bool) transfer {
	return tc.set("staticMode", value)
}

// StaticClassName 静态模式下的样式类名
func (tc transfer) StaticClassName(value string) transfer {
	return tc.set("staticClassName", value)
}

// StaticLabel 静态模式下的 label 内容
func (tc transfer) StaticLabel(value string) transfer {
	return tc.set("staticLabel", value)
}

// StaticPlaceholder 静态模式下的 placeholder 内容
func (tc transfer) StaticPlaceholder(value string) transfer {
	return tc.set("staticPlaceholder", value)
}

// StrictMode 是否严格模式，开启时只允许选中列表中的选项
func (tc transfer) StrictMode(value bool) transfer {
	return tc.set("strictMode", value)
}

// SubmitOnChange 表单项内容变更时是否触发表单提交
func (tc transfer) SubmitOnChange(value bool) transfer {
	return tc.set("submitOnChange", value)
}

// Sync(true) 设置值时是否同步 (设置值时是否同步)
func (tc transfer) Sync(value bool) transfer {
	return tc.set("sync", value)
}

// TextField 设置选项的展示字段
func (tc transfer) TextField(value string) transfer {
	return tc.set("textField", value)
}

// Value
func (tc transfer) Value(value any) transfer {
	return tc.set("value", value)
}

// Visible 控件显示与否
func (tc transfer) Visible(value bool) transfer {
	return tc.set("visible", value)
}

// VisibleOn 控件显示与否表达式 (表达式，语法 `data.xxx > 5`。)
func (tc transfer) VisibleOn(value string) transfer {
	return tc.set("visibleOn", value)
}

// Width 控件宽度
func (tc transfer) Width(value string) transfer {
	return tc.set("width", value)
}

// Wrappers 包装组件
func (tc transfer) Wrappers(value string) transfer {
	return tc.set("wrappers", value)
}

// AllowMultiple 是否允许多选
func (tc transfer) AllowMultiple(value bool) transfer {
	return tc.set("allowMultiple", value)
}
