package comp

// TransferPickerControl 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer-picker
//
// @author  slowlyo
// @version 6.7.0
type TransferPickerControl Schema

// NewTransferPickerControl 创建一个新的 TransferPickerControl 实例
func NewTransferPickerControl() TransferPickerControl {
	return TransferPickerControl{}.set("type", "transfer-picker")
}

func (tpc TransferPickerControl) set(key string, value interface{}) TransferPickerControl {
	tpc[key] = value
	return tpc
}

// AddApi 添加时调用的接口
func (tpc TransferPickerControl) AddApi(value string) TransferPickerControl {
	return tpc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tpc TransferPickerControl) AddControls(value string) TransferPickerControl {
	return tpc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (tpc TransferPickerControl) AddDialog(value string) TransferPickerControl {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tpc TransferPickerControl) AutoCheckChildren(value bool) TransferPickerControl {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tpc TransferPickerControl) AutoFill(value string) TransferPickerControl {
	return tpc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (tpc TransferPickerControl) BorderMode(value string) TransferPickerControl {
	return tpc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tpc TransferPickerControl) ClassName(value string) TransferPickerControl {
	return tpc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tpc TransferPickerControl) ClearValueOnHidden(value bool) TransferPickerControl {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tpc TransferPickerControl) Clearable(value bool) TransferPickerControl {
	return tpc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tpc TransferPickerControl) Columns(value string) TransferPickerControl {
	return tpc.set("columns", value)
}

// Creatable 是否可以新增
func (tpc TransferPickerControl) Creatable(value bool) TransferPickerControl {
	return tpc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tpc TransferPickerControl) CreateBtnLabel(value string) TransferPickerControl {
	return tpc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tpc TransferPickerControl) DeferApi(value string) TransferPickerControl {
	return tpc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tpc TransferPickerControl) DeferField(value string) TransferPickerControl {
	return tpc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tpc TransferPickerControl) DeleteApi(value string) TransferPickerControl {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tpc TransferPickerControl) DeleteConfirmText(value string) TransferPickerControl {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tpc TransferPickerControl) Delimiter(value string) TransferPickerControl {
	return tpc.set("delimiter", value)
}

// Desc
func (tpc TransferPickerControl) Desc(value string) TransferPickerControl {
	return tpc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tpc TransferPickerControl) Description(value string) TransferPickerControl {
	return tpc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tpc TransferPickerControl) DescriptionClassName(value string) TransferPickerControl {
	return tpc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tpc TransferPickerControl) Disabled(value bool) TransferPickerControl {
	return tpc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc TransferPickerControl) DisabledOn(value string) TransferPickerControl {
	return tpc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tpc TransferPickerControl) EditApi(value string) TransferPickerControl {
	return tpc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tpc TransferPickerControl) EditControls(value string) TransferPickerControl {
	return tpc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (tpc TransferPickerControl) EditDialog(value string) TransferPickerControl {
	return tpc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tpc TransferPickerControl) Editable(value bool) TransferPickerControl {
	return tpc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tpc TransferPickerControl) EditorSetting(value string) TransferPickerControl {
	return tpc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tpc TransferPickerControl) ExtraName(value string) TransferPickerControl {
	return tpc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tpc TransferPickerControl) ExtractValue(value bool) TransferPickerControl {
	return tpc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tpc TransferPickerControl) Hidden(value bool) TransferPickerControl {
	return tpc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc TransferPickerControl) HiddenOn(value string) TransferPickerControl {
	return tpc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tpc TransferPickerControl) Hint(value string) TransferPickerControl {
	return tpc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tpc TransferPickerControl) Horizontal(value string) TransferPickerControl {
	return tpc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tpc TransferPickerControl) Id(value string) TransferPickerControl {
	return tpc.set("id", value)
}

// InitAutoFill
func (tpc TransferPickerControl) InitAutoFill(value string) TransferPickerControl {
	return tpc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tpc TransferPickerControl) InitFetch(value bool) TransferPickerControl {
	return tpc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tpc TransferPickerControl) InitFetchOn(value string) TransferPickerControl {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tpc TransferPickerControl) InitiallyOpen(value bool) TransferPickerControl {
	return tpc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tpc TransferPickerControl) Inline(value bool) TransferPickerControl {
	return tpc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tpc TransferPickerControl) InputClassName(value string) TransferPickerControl {
	return tpc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tpc TransferPickerControl) ItemHeight(value string) TransferPickerControl {
	return tpc.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的选项 value 会以数组形式提交。
func (tpc TransferPickerControl) JoinValues(value bool) TransferPickerControl {
	return tpc.set("joinValues", value)
}

// Label 文字标签
func (tpc TransferPickerControl) Label(value string) TransferPickerControl {
	return tpc.set("label", value)
}

// LabelAlign 标签对齐方式
func (tpc TransferPickerControl) LabelAlign(value string) TransferPickerControl {
	return tpc.set("labelAlign", value)
}

// LabelClassName 标签 css 类名
func (tpc TransferPickerControl) LabelClassName(value string) TransferPickerControl {
	return tpc.set("labelClassName", value)
}

// LabelRemark 标签补充说明
func (tpc TransferPickerControl) LabelRemark(value string) TransferPickerControl {
	return tpc.set("labelRemark", value)
}

// Max 最大选项数
func (tpc TransferPickerControl) Max(value int) TransferPickerControl {
	return tpc.set("max", value)
}

// MergeDataSource 是否合并 source 数据，默认 false。设置为 true 时，将 source 数据合并到 options 数据中
func (tpc TransferPickerControl) MergeDataSource(value bool) TransferPickerControl {
	return tpc.set("mergeDataSource", value)
}

// Min 最小选项数
func (tpc TransferPickerControl) Min(value int) TransferPickerControl {
	return tpc.set("min", value)
}

// MultiLine 在多行模式下显示 (是否多行显示)
func (tpc TransferPickerControl) MultiLine(value bool) TransferPickerControl {
	return tpc.set("multiLine", value)
}

// Name 组件名字
func (tpc TransferPickerControl) Name(value string) TransferPickerControl {
	return tpc.set("name", value)
}

// Options 数据源配置
func (tpc TransferPickerControl) Options(value string) TransferPickerControl {
	return tpc.set("options", value)
}

// Placeholder 占位符
func (tpc TransferPickerControl) Placeholder(value string) TransferPickerControl {
	return tpc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tpc TransferPickerControl) ReadOnly(value bool) TransferPickerControl {
	return tpc.set("readOnly", value)
}

// Remark
func (tpc TransferPickerControl) Remark(value string) TransferPickerControl {
	return tpc.set("remark", value)
}

// Removable 是否可移除
func (tpc TransferPickerControl) Removable(value bool) TransferPickerControl {
	return tpc.set("removable", value)
}

// Renderer 组件渲染器类型 (组件渲染器类型)
func (tpc TransferPickerControl) Renderer(value string) TransferPickerControl {
	return tpc.set("renderer", value)
}

// Source 配置数据源 (配置数据源)
func (tpc TransferPickerControl) Source(value string) TransferPickerControl {
	return tpc.set("source", value)
}

// SourceEmptyText 配置为空时文本 (配置为空时文本)
func (tpc TransferPickerControl) SourceEmptyText(value string) TransferPickerControl {
	return tpc.set("sourceEmptyText", value)
}

// Static
func (tpc TransferPickerControl) Static(value bool) TransferPickerControl {
	return tpc.set("static", value)
}

// StaticClassName
func (tpc TransferPickerControl) StaticClassName(value string) TransferPickerControl {
	return tpc.set("staticClassName", value)
}

// Type 组件类型
func (tpc TransferPickerControl) Type(value string) TransferPickerControl {
	return tpc.set("type", value)
}

// Unset 删除某个选项
func (tpc TransferPickerControl) Unset(value string) TransferPickerControl {
	return tpc.set("unset", value)
}

// Value
func (tpc TransferPickerControl) Value(value string) TransferPickerControl {
	return tpc.set("value", value)
}

// ValueField 指定数据源 value 字段，默认值为 value
func (tpc TransferPickerControl) ValueField(value string) TransferPickerControl {
	return tpc.set("valueField", value)
}

// ValueJoin 字段分隔符
func (tpc TransferPickerControl) ValueJoin(value string) TransferPickerControl {
	return tpc.set("valueJoin", value)
}

// Vertical 当配置为垂直布局的时候，用来配置具体的上下分配。
func (tpc TransferPickerControl) Vertical(value string) TransferPickerControl {
	return tpc.set("vertical", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc TransferPickerControl) VisibleOn(value string) TransferPickerControl {
	return tpc.set("visibleOn", value)
}

// Width 宽度
func (tpc TransferPickerControl) Width(value string) TransferPickerControl {
	return tpc.set("width", value)
}

// Wrap 折叠展示的最小宽度
func (tpc TransferPickerControl) Wrap(value string) TransferPickerControl {
	return tpc.set("wrap", value)
}

// ZIndex
func (tpc TransferPickerControl) ZIndex(value int) TransferPickerControl {
	return tpc.set("zIndex", value)
}
