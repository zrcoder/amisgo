package comp

// TransferControl 控件定义
//
// @version 6.7.0
type TransferControl Schema

// NewTransferControl 创建一个新的 TransferControl 实例
func NewTransferControl() TransferControl {
	return TransferControl{}.set("type", "transfer")
}

func (tc TransferControl) set(key string, value interface{}) TransferControl {
	tc[key] = value
	return tc
}

// AddApi 添加时调用的接口
func (tc TransferControl) AddApi(value string) TransferControl {
	return tc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tc TransferControl) AddControls(value string) TransferControl {
	return tc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (tc TransferControl) AddDialog(value string) TransferControl {
	return tc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tc TransferControl) AutoCheckChildren(value bool) TransferControl {
	return tc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc TransferControl) AutoFill(value string) TransferControl {
	return tc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (tc TransferControl) ClassName(value string) TransferControl {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc TransferControl) ClearValueOnHidden(value bool) TransferControl {
	return tc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tc TransferControl) Clearable(value bool) TransferControl {
	return tc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tc TransferControl) Columns(value string) TransferControl {
	return tc.set("columns", value)
}

// Creatable 是否可以新增
func (tc TransferControl) Creatable(value bool) TransferControl {
	return tc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tc TransferControl) CreateBtnLabel(value string) TransferControl {
	return tc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tc TransferControl) DeferApi(value string) TransferControl {
	return tc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tc TransferControl) DeferField(value string) TransferControl {
	return tc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tc TransferControl) DeleteApi(value string) TransferControl {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tc TransferControl) DeleteConfirmText(value string) TransferControl {
	return tc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tc TransferControl) Delimiter(value string) TransferControl {
	return tc.set("delimiter", value)
}

// Desc
func (tc TransferControl) Desc(value string) TransferControl {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc TransferControl) Description(value string) TransferControl {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc TransferControl) DescriptionClassName(value string) TransferControl {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc TransferControl) Disabled(value bool) TransferControl {
	return tc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TransferControl) DisabledOn(value string) TransferControl {
	return tc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tc TransferControl) EditApi(value string) TransferControl {
	return tc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tc TransferControl) EditControls(value string) TransferControl {
	return tc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (tc TransferControl) EditDialog(value string) TransferControl {
	return tc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tc TransferControl) Editable(value bool) TransferControl {
	return tc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc TransferControl) EditorSetting(value string) TransferControl {
	return tc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc TransferControl) ExtraName(value string) TransferControl {
	return tc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tc TransferControl) ExtractValue(value bool) TransferControl {
	return tc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tc TransferControl) Hidden(value bool) TransferControl {
	return tc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TransferControl) HiddenOn(value string) TransferControl {
	return tc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tc TransferControl) Hint(value string) TransferControl {
	return tc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tc TransferControl) Horizontal(value string) TransferControl {
	return tc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tc TransferControl) Id(value string) TransferControl {
	return tc.set("id", value)
}

// InitAutoFill
func (tc TransferControl) InitAutoFill(value string) TransferControl {
	return tc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tc TransferControl) InitFetch(value bool) TransferControl {
	return tc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tc TransferControl) InitFetchOn(value string) TransferControl {
	return tc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tc TransferControl) InitiallyOpen(value bool) TransferControl {
	return tc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tc TransferControl) Inline(value bool) TransferControl {
	return tc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tc TransferControl) InputClassName(value string) TransferControl {
	return tc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tc TransferControl) ItemHeight(value string) TransferControl {
	return tc.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (tc TransferControl) JoinValues(value bool) TransferControl {
	return tc.set("joinValues", value)
}

// Label 描述标题
func (tc TransferControl) Label(value string) TransferControl {
	return tc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (tc TransferControl) LabelAlign(value string) TransferControl {
	return tc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (tc TransferControl) LabelClassName(value string) TransferControl {
	return tc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (tc TransferControl) LabelRemark(value string) TransferControl {
	return tc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (tc TransferControl) LabelWidth(value string) TransferControl {
	return tc.set("labelWidth", value)
}

// LeftMode 当 selectMode 为 associated 时用来定义左侧的选择模式 可选值: tree | list
func (tc TransferControl) LeftMode(value string) TransferControl {
	return tc.set("leftMode", value)
}

// Max 可选项的最大值。
func (tc TransferControl) Max(value int) TransferControl {
	return tc.set("max", value)
}

// Min 可选项的最小值。
func (tc TransferControl) Min(value int) TransferControl {
	return tc.set("min", value)
}

// Name 表单项名称，内部表单渲染时需要
func (tc TransferControl) Name(value string) TransferControl {
	return tc.set("name", value)
}

// OnChange 变化回调函数。
func (tc TransferControl) OnChange(value string) TransferControl {
	return tc.set("onChange", value)
}

// OnEvent 事件处理
func (tc TransferControl) OnEvent(value string) TransferControl {
	return tc.set("onEvent", value)
}

// Option
func (tc TransferControl) Option(value string) TransferControl {
	return tc.set("option", value)
}

// Options 选项数据
func (tc TransferControl) Options(value string) TransferControl {
	return tc.set("options", value)
}

// Placeholder 输入框提示语
func (tc TransferControl) Placeholder(value string) TransferControl {
	return tc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tc TransferControl) ReadOnly(value bool) TransferControl {
	return tc.set("readOnly", value)
}

// RemoteOptions 远程获取选项
func (tc TransferControl) RemoteOptions(value string) TransferControl {
	return tc.set("remoteOptions", value)
}

// Render 表单渲染模式，通常在内嵌时用
func (tc TransferControl) Render(value string) TransferControl {
	return tc.set("render", value)
}

// ResetValue 清空表单项内容时，设置该项的默认值
func (tc TransferControl) ResetValue(value string) TransferControl {
	return tc.set("resetValue", value)
}

// SaveImmediately 表单项内容变更是否立即保存
func (tc TransferControl) SaveImmediately(value bool) TransferControl {
	return tc.set("saveImmediately", value)
}

// Source 选项数据来源
func (tc TransferControl) Source(value string) TransferControl {
	return tc.set("source", value)
}

// StaticMode 控件是否静态显示
func (tc TransferControl) StaticMode(value bool) TransferControl {
	return tc.set("staticMode", value)
}

// StaticClassName 静态模式下的样式类名
func (tc TransferControl) StaticClassName(value string) TransferControl {
	return tc.set("staticClassName", value)
}

// StaticLabel 静态模式下的 label 内容
func (tc TransferControl) StaticLabel(value string) TransferControl {
	return tc.set("staticLabel", value)
}

// StaticPlaceholder 静态模式下的 placeholder 内容
func (tc TransferControl) StaticPlaceholder(value string) TransferControl {
	return tc.set("staticPlaceholder", value)
}

// StrictMode 是否严格模式，开启时只允许选中列表中的选项
func (tc TransferControl) StrictMode(value bool) TransferControl {
	return tc.set("strictMode", value)
}

// SubmitOnChange 表单项内容变更时是否触发表单提交
func (tc TransferControl) SubmitOnChange(value bool) TransferControl {
	return tc.set("submitOnChange", value)
}

// Sync(true) 设置值时是否同步 (设置值时是否同步)
func (tc TransferControl) Sync(value bool) TransferControl {
	return tc.set("sync", value)
}

// TextField 设置选项的展示字段
func (tc TransferControl) TextField(value string) TransferControl {
	return tc.set("textField", value)
}

// Value
func (tc TransferControl) Value(value interface{}) TransferControl {
	return tc.set("value", value)
}

// Visible 控件显示与否
func (tc TransferControl) Visible(value bool) TransferControl {
	return tc.set("visible", value)
}

// VisibleOn 控件显示与否表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TransferControl) VisibleOn(value string) TransferControl {
	return tc.set("visibleOn", value)
}

// Width 控件宽度
func (tc TransferControl) Width(value string) TransferControl {
	return tc.set("width", value)
}

// Wrappers 包装组件
func (tc TransferControl) Wrappers(value string) TransferControl {
	return tc.set("wrappers", value)
}

// AllowMultiple 是否允许多选
func (tc TransferControl) AllowMultiple(value bool) TransferControl {
	return tc.set("allowMultiple", value)
}
