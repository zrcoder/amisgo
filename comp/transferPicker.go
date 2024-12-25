package comp

// transferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/transfer-picker

type transferPicker schema

func TransferPicker() transferPicker {
	return transferPicker{}.set("type", "transfer-picker")
}

func (tpc transferPicker) set(key string, value any) transferPicker {
	tpc[key] = value
	return tpc
}

// AddApi 添加时调用的接口
func (tpc transferPicker) AddApi(value string) transferPicker {
	return tpc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tpc transferPicker) AddControls(value string) transferPicker {
	return tpc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (tpc transferPicker) AddDialog(value string) transferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tpc transferPicker) AutoCheckChildren(value bool) transferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tpc transferPicker) AutoFill(value string) transferPicker {
	return tpc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (tpc transferPicker) BorderMode(value string) transferPicker {
	return tpc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tpc transferPicker) ClassName(value string) transferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tpc transferPicker) ClearValueOnHidden(value bool) transferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tpc transferPicker) Clearable(value bool) transferPicker {
	return tpc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tpc transferPicker) Columns(value ...any) transferPicker {
	return tpc.set("columns", value)
}

// Creatable 是否可以新增
func (tpc transferPicker) Creatable(value bool) transferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tpc transferPicker) CreateBtnLabel(value string) transferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tpc transferPicker) DeferApi(value string) transferPicker {
	return tpc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tpc transferPicker) DeferField(value string) transferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tpc transferPicker) DeleteApi(value string) transferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tpc transferPicker) DeleteConfirmText(value string) transferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tpc transferPicker) Delimiter(value string) transferPicker {
	return tpc.set("delimiter", value)
}

// Desc
func (tpc transferPicker) Desc(value string) transferPicker {
	return tpc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tpc transferPicker) Description(value string) transferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tpc transferPicker) DescriptionClassName(value string) transferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tpc transferPicker) Disabled(value bool) transferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc transferPicker) DisabledOn(value string) transferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tpc transferPicker) EditApi(value string) transferPicker {
	return tpc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tpc transferPicker) EditControls(value string) transferPicker {
	return tpc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (tpc transferPicker) EditDialog(value string) transferPicker {
	return tpc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tpc transferPicker) Editable(value bool) transferPicker {
	return tpc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tpc transferPicker) EditorSetting(value string) transferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tpc transferPicker) ExtraName(value string) transferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tpc transferPicker) ExtractValue(value bool) transferPicker {
	return tpc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tpc transferPicker) Hidden(value bool) transferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc transferPicker) HiddenOn(value string) transferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tpc transferPicker) Hint(value string) transferPicker {
	return tpc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tpc transferPicker) Horizontal(value string) transferPicker {
	return tpc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tpc transferPicker) ID(value string) transferPicker {
	return tpc.set("id", value)
}

// InitAutoFill
func (tpc transferPicker) InitAutoFill(value string) transferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tpc transferPicker) InitFetch(value bool) transferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tpc transferPicker) InitFetchOn(value string) transferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tpc transferPicker) InitiallyOpen(value bool) transferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tpc transferPicker) Inline(value bool) transferPicker {
	return tpc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tpc transferPicker) InputClassName(value string) transferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tpc transferPicker) ItemHeight(value string) transferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的选项 value 会以数组形式提交。
func (tpc transferPicker) JoinValues(value bool) transferPicker {
	return tpc.set("joinValues", value)
}

// Label 文字标签
func (tpc transferPicker) Label(value string) transferPicker {
	return tpc.set("label", value)
}

// LabelAlign 标签对齐方式
func (tpc transferPicker) LabelAlign(value string) transferPicker {
	return tpc.set("labelAlign", value)
}

// LabelClassName 标签 css 类名
func (tpc transferPicker) LabelClassName(value string) transferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark 标签补充说明
func (tpc transferPicker) LabelRemark(value string) transferPicker {
	return tpc.set("labelRemark", value)
}

// Max 最大选项数
func (tpc transferPicker) Max(value int) transferPicker {
	return tpc.set("max", value)
}

// MergeDataSource 是否合并 source 数据，默认 false。设置为 true 时，将 source 数据合并到 options 数据中
func (tpc transferPicker) MergeDataSource(value bool) transferPicker {
	return tpc.set("mergeDataSource", value)
}

// Min 最小选项数
func (tpc transferPicker) Min(value int) transferPicker {
	return tpc.set("min", value)
}

// MultiLine 在多行模式下显示 (是否多行显示)
func (tpc transferPicker) MultiLine(value bool) transferPicker {
	return tpc.set("multiLine", value)
}

// Name 组件名字
func (tpc transferPicker) Name(value string) transferPicker {
	return tpc.set("name", value)
}

// Options 数据源配置
func (tpc transferPicker) Options(value ...MOption) transferPicker {
	return tpc.set("options", value)
}

// Placeholder 占位符
func (tpc transferPicker) Placeholder(value string) transferPicker {
	return tpc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tpc transferPicker) ReadOnly(value bool) transferPicker {
	return tpc.set("readOnly", value)
}

// Remark
func (tpc transferPicker) Remark(value string) transferPicker {
	return tpc.set("remark", value)
}

// Removable 是否可移除
func (tpc transferPicker) Removable(value bool) transferPicker {
	return tpc.set("removable", value)
}

// Renderer 组件渲染器类型 (组件渲染器类型)
func (tpc transferPicker) Renderer(value string) transferPicker {
	return tpc.set("renderer", value)
}

// Source 配置数据源 (配置数据源)
func (tpc transferPicker) Source(value string) transferPicker {
	return tpc.set("source", value)
}

// SourceEmptyText 配置为空时文本 (配置为空时文本)
func (tpc transferPicker) SourceEmptyText(value string) transferPicker {
	return tpc.set("sourceEmptyText", value)
}

// Static
func (tpc transferPicker) Static(value bool) transferPicker {
	return tpc.set("static", value)
}

// StaticClassName
func (tpc transferPicker) StaticClassName(value string) transferPicker {
	return tpc.set("staticClassName", value)
}

// Unset 删除某个选项
func (tpc transferPicker) Unset(value string) transferPicker {
	return tpc.set("unset", value)
}

// Value
func (tpc transferPicker) Value(value string) transferPicker {
	return tpc.set("value", value)
}

// ValueField 指定数据源 value 字段，默认值为 value
func (tpc transferPicker) ValueField(value string) transferPicker {
	return tpc.set("valueField", value)
}

// ValueJoin 字段分隔符
func (tpc transferPicker) ValueJoin(value string) transferPicker {
	return tpc.set("valueJoin", value)
}

// Vertical 当配置为垂直布局的时候，用来配置具体的上下分配。
func (tpc transferPicker) Vertical(value string) transferPicker {
	return tpc.set("vertical", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc transferPicker) VisibleOn(value string) transferPicker {
	return tpc.set("visibleOn", value)
}

// Width 宽度
func (tpc transferPicker) Width(value string) transferPicker {
	return tpc.set("width", value)
}

// Wrap 折叠展示的最小宽度
func (tpc transferPicker) Wrap(value string) transferPicker {
	return tpc.set("wrap", value)
}

// ZIndex
func (tpc transferPicker) ZIndex(value int) transferPicker {
	return tpc.set("zIndex", value)
}
