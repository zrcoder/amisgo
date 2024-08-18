package comp

// tabsTransferPicker 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer-picker
//
// @version 6.7.0
type tabsTransferPicker schema

func TabsTransferPicker() tabsTransferPicker {
	return tabsTransferPicker{}.set("type", "tabs-transfer-picker")
}

// set 设置指定的键值对并返回新的 TabsTransferPickerControl 实例
func (tpc tabsTransferPicker) set(key string, value interface{}) tabsTransferPicker {
	tpc[key] = value
	return tpc
}

// AddApi 添加时调用的接口
func (tpc tabsTransferPicker) AddApi(value string) tabsTransferPicker {
	return tpc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tpc tabsTransferPicker) AddControls(value string) tabsTransferPicker {
	return tpc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (tpc tabsTransferPicker) AddDialog(value string) tabsTransferPicker {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tpc tabsTransferPicker) AutoCheckChildren(value bool) tabsTransferPicker {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tpc tabsTransferPicker) AutoFill(value string) tabsTransferPicker {
	return tpc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tpc tabsTransferPicker) ClassName(value string) tabsTransferPicker {
	return tpc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tpc tabsTransferPicker) ClearValueOnHidden(value bool) tabsTransferPicker {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tpc tabsTransferPicker) Clearable(value bool) tabsTransferPicker {
	return tpc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tpc tabsTransferPicker) Columns(value string) tabsTransferPicker {
	return tpc.set("columns", value)
}

// Creatable 是否可以新增
func (tpc tabsTransferPicker) Creatable(value bool) tabsTransferPicker {
	return tpc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tpc tabsTransferPicker) CreateBtnLabel(value string) tabsTransferPicker {
	return tpc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tpc tabsTransferPicker) DeferApi(value string) tabsTransferPicker {
	return tpc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tpc tabsTransferPicker) DeferField(value string) tabsTransferPicker {
	return tpc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tpc tabsTransferPicker) DeleteApi(value string) tabsTransferPicker {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tpc tabsTransferPicker) DeleteConfirmText(value string) tabsTransferPicker {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tpc tabsTransferPicker) Delimiter(value string) tabsTransferPicker {
	return tpc.set("delimiter", value)
}

// Desc
func (tpc tabsTransferPicker) Desc(value string) tabsTransferPicker {
	return tpc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tpc tabsTransferPicker) Description(value string) tabsTransferPicker {
	return tpc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tpc tabsTransferPicker) DescriptionClassName(value string) tabsTransferPicker {
	return tpc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tpc tabsTransferPicker) Disabled(value bool) tabsTransferPicker {
	return tpc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc tabsTransferPicker) DisabledOn(value string) tabsTransferPicker {
	return tpc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tpc tabsTransferPicker) EditApi(value string) tabsTransferPicker {
	return tpc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tpc tabsTransferPicker) EditControls(value string) tabsTransferPicker {
	return tpc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (tpc tabsTransferPicker) EditDialog(value string) tabsTransferPicker {
	return tpc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tpc tabsTransferPicker) Editable(value bool) tabsTransferPicker {
	return tpc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tpc tabsTransferPicker) EditorSetting(value string) tabsTransferPicker {
	return tpc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tpc tabsTransferPicker) ExtraName(value string) tabsTransferPicker {
	return tpc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tpc tabsTransferPicker) ExtractValue(value bool) tabsTransferPicker {
	return tpc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tpc tabsTransferPicker) Hidden(value bool) tabsTransferPicker {
	return tpc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc tabsTransferPicker) HiddenOn(value string) tabsTransferPicker {
	return tpc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tpc tabsTransferPicker) Hint(value string) tabsTransferPicker {
	return tpc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tpc tabsTransferPicker) Horizontal(value string) tabsTransferPicker {
	return tpc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tpc tabsTransferPicker) Id(value string) tabsTransferPicker {
	return tpc.set("id", value)
}

// InitAutoFill
func (tpc tabsTransferPicker) InitAutoFill(value string) tabsTransferPicker {
	return tpc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tpc tabsTransferPicker) InitFetch(value bool) tabsTransferPicker {
	return tpc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tpc tabsTransferPicker) InitFetchOn(value string) tabsTransferPicker {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tpc tabsTransferPicker) InitiallyOpen(value bool) tabsTransferPicker {
	return tpc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tpc tabsTransferPicker) Inline(value bool) tabsTransferPicker {
	return tpc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tpc tabsTransferPicker) InputClassName(value string) tabsTransferPicker {
	return tpc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tpc tabsTransferPicker) ItemHeight(value string) tabsTransferPicker {
	return tpc.set("itemHeight", value)
}

// JoinValues 设置是否将值与选择项联动
func (tpc tabsTransferPicker) JoinValues(value bool) tabsTransferPicker {
	return tpc.set("joinValues", value)
}

// Label 标签
func (tpc tabsTransferPicker) Label(value string) tabsTransferPicker {
	return tpc.set("label", value)
}

// LabelClassName 标签的 CSS 类名
func (tpc tabsTransferPicker) LabelClassName(value string) tabsTransferPicker {
	return tpc.set("labelClassName", value)
}

// LabelRemark 标签备注
func (tpc tabsTransferPicker) LabelRemark(value string) tabsTransferPicker {
	return tpc.set("labelRemark", value)
}

// MaxLength 最多选择的数量
func (tpc tabsTransferPicker) MaxLength(value int) tabsTransferPicker {
	return tpc.set("maxLength", value)
}

// Mode 选择模式
func (tpc tabsTransferPicker) Mode(value string) tabsTransferPicker {
	return tpc.set("mode", value)
}

// Multiple 是否多选
func (tpc tabsTransferPicker) Multiple(value bool) tabsTransferPicker {
	return tpc.set("multiple", value)
}

// Name 表单项 name
func (tpc tabsTransferPicker) Name(value string) tabsTransferPicker {
	return tpc.set("name", value)
}

// OnChange 变化时回调
func (tpc tabsTransferPicker) OnChange(value string) tabsTransferPicker {
	return tpc.set("onChange", value)
}

// Options 选项
func (tpc tabsTransferPicker) Options(value interface{}) tabsTransferPicker {
	return tpc.set("options", value)
}

// OptionsSrc 选项数据接口
func (tpc tabsTransferPicker) OptionsSrc(value string) tabsTransferPicker {
	return tpc.set("optionsSrc", value)
}

// ReadOnly 是否只读
func (tpc tabsTransferPicker) ReadOnly(value bool) tabsTransferPicker {
	return tpc.set("readOnly", value)
}

// Required 是否必填
func (tpc tabsTransferPicker) Required(value bool) tabsTransferPicker {
	return tpc.set("required", value)
}

// RequiredApi 提交时必填 API
func (tpc tabsTransferPicker) RequiredApi(value string) tabsTransferPicker {
	return tpc.set("requiredApi", value)
}

// Source 接口地址
func (tpc tabsTransferPicker) Source(value string) tabsTransferPicker {
	return tpc.set("source", value)
}

// Tabs 选项卡
func (tpc tabsTransferPicker) Tabs(value interface{}) tabsTransferPicker {
	return tpc.set("tabs", value)
}

// Tooltip 提示信息
func (tpc tabsTransferPicker) Tooltip(value string) tabsTransferPicker {
	return tpc.set("tooltip", value)
}

// Value 表单值
func (tpc tabsTransferPicker) Value(value interface{}) tabsTransferPicker {
	return tpc.set("value", value)
}

// Vertical 列表方向
func (tpc tabsTransferPicker) Vertical(value bool) tabsTransferPicker {
	return tpc.set("vertical", value)
}

// Visible 是否显示
func (tpc tabsTransferPicker) Visible(value bool) tabsTransferPicker {
	return tpc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (tpc tabsTransferPicker) VisibleOn(value string) tabsTransferPicker {
	return tpc.set("visibleOn", value)
}

// Width 宽度
func (tpc tabsTransferPicker) Width(value string) tabsTransferPicker {
	return tpc.set("width", value)
}

// WidthConfig 配置宽度
func (tpc tabsTransferPicker) WidthConfig(value string) tabsTransferPicker {
	return tpc.set("widthConfig", value)
}

// Xname 配置 key name
func (tpc tabsTransferPicker) Xname(value string) tabsTransferPicker {
	return tpc.set("xname", value)
}

// Xtype 配置类型
func (tpc tabsTransferPicker) Xtype(value string) tabsTransferPicker {
	return tpc.set("xtype", value)
}

// Component 配置自定义组件
func (tpc tabsTransferPicker) Component(value string) tabsTransferPicker {
	return tpc.set("component", value)
}
