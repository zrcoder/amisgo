package comp

// TabsTransferPickerControl 穿梭器的弹框形态 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tabs-transfer-picker
//
// @author  slowlyo
// @version 6.7.0
type TabsTransferPickerControl Schema

// NewTabsTransferPickerControl 创建一个新的 TabsTransferPickerControl 实例
func NewTabsTransferPickerControl() TabsTransferPickerControl {
	return TabsTransferPickerControl{}.set("type", "tabs-transfer-picker")
}

// set 设置指定的键值对并返回新的 TabsTransferPickerControl 实例
func (tpc TabsTransferPickerControl) set(key string, value interface{}) TabsTransferPickerControl {
	tpc[key] = value
	return tpc
}

// AddApi 添加时调用的接口
func (tpc TabsTransferPickerControl) AddApi(value string) TabsTransferPickerControl {
	return tpc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tpc TabsTransferPickerControl) AddControls(value string) TabsTransferPickerControl {
	return tpc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (tpc TabsTransferPickerControl) AddDialog(value string) TabsTransferPickerControl {
	return tpc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tpc TabsTransferPickerControl) AutoCheckChildren(value bool) TabsTransferPickerControl {
	return tpc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tpc TabsTransferPickerControl) AutoFill(value string) TabsTransferPickerControl {
	return tpc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tpc TabsTransferPickerControl) ClassName(value string) TabsTransferPickerControl {
	return tpc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tpc TabsTransferPickerControl) ClearValueOnHidden(value bool) TabsTransferPickerControl {
	return tpc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tpc TabsTransferPickerControl) Clearable(value bool) TabsTransferPickerControl {
	return tpc.set("clearable", value)
}

// Columns 当 selectMode 为 table 时定义表格列信息。
func (tpc TabsTransferPickerControl) Columns(value string) TabsTransferPickerControl {
	return tpc.set("columns", value)
}

// Creatable 是否可以新增
func (tpc TabsTransferPickerControl) Creatable(value bool) TabsTransferPickerControl {
	return tpc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tpc TabsTransferPickerControl) CreateBtnLabel(value string) TabsTransferPickerControl {
	return tpc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (tpc TabsTransferPickerControl) DeferApi(value string) TabsTransferPickerControl {
	return tpc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tpc TabsTransferPickerControl) DeferField(value string) TabsTransferPickerControl {
	return tpc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tpc TabsTransferPickerControl) DeleteApi(value string) TabsTransferPickerControl {
	return tpc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tpc TabsTransferPickerControl) DeleteConfirmText(value string) TabsTransferPickerControl {
	return tpc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tpc TabsTransferPickerControl) Delimiter(value string) TabsTransferPickerControl {
	return tpc.set("delimiter", value)
}

// Desc
func (tpc TabsTransferPickerControl) Desc(value string) TabsTransferPickerControl {
	return tpc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tpc TabsTransferPickerControl) Description(value string) TabsTransferPickerControl {
	return tpc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tpc TabsTransferPickerControl) DescriptionClassName(value string) TabsTransferPickerControl {
	return tpc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tpc TabsTransferPickerControl) Disabled(value bool) TabsTransferPickerControl {
	return tpc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc TabsTransferPickerControl) DisabledOn(value string) TabsTransferPickerControl {
	return tpc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tpc TabsTransferPickerControl) EditApi(value string) TabsTransferPickerControl {
	return tpc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tpc TabsTransferPickerControl) EditControls(value string) TabsTransferPickerControl {
	return tpc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (tpc TabsTransferPickerControl) EditDialog(value string) TabsTransferPickerControl {
	return tpc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tpc TabsTransferPickerControl) Editable(value bool) TabsTransferPickerControl {
	return tpc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tpc TabsTransferPickerControl) EditorSetting(value string) TabsTransferPickerControl {
	return tpc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tpc TabsTransferPickerControl) ExtraName(value string) TabsTransferPickerControl {
	return tpc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tpc TabsTransferPickerControl) ExtractValue(value bool) TabsTransferPickerControl {
	return tpc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tpc TabsTransferPickerControl) Hidden(value bool) TabsTransferPickerControl {
	return tpc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tpc TabsTransferPickerControl) HiddenOn(value string) TabsTransferPickerControl {
	return tpc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tpc TabsTransferPickerControl) Hint(value string) TabsTransferPickerControl {
	return tpc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tpc TabsTransferPickerControl) Horizontal(value string) TabsTransferPickerControl {
	return tpc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tpc TabsTransferPickerControl) Id(value string) TabsTransferPickerControl {
	return tpc.set("id", value)
}

// InitAutoFill
func (tpc TabsTransferPickerControl) InitAutoFill(value string) TabsTransferPickerControl {
	return tpc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tpc TabsTransferPickerControl) InitFetch(value bool) TabsTransferPickerControl {
	return tpc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tpc TabsTransferPickerControl) InitFetchOn(value string) TabsTransferPickerControl {
	return tpc.set("initFetchOn", value)
}

// InitiallyOpen 是否默认都展开
func (tpc TabsTransferPickerControl) InitiallyOpen(value bool) TabsTransferPickerControl {
	return tpc.set("initiallyOpen", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tpc TabsTransferPickerControl) Inline(value bool) TabsTransferPickerControl {
	return tpc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tpc TabsTransferPickerControl) InputClassName(value string) TabsTransferPickerControl {
	return tpc.set("inputClassName", value)
}

// ItemHeight 单个选项的高度，主要用于虚拟渲染
func (tpc TabsTransferPickerControl) ItemHeight(value string) TabsTransferPickerControl {
	return tpc.set("itemHeight", value)
}

// JoinValues 设置是否将值与选择项联动
func (tpc TabsTransferPickerControl) JoinValues(value bool) TabsTransferPickerControl {
	return tpc.set("joinValues", value)
}

// Label 标签
func (tpc TabsTransferPickerControl) Label(value string) TabsTransferPickerControl {
	return tpc.set("label", value)
}

// LabelClassName 标签的 CSS 类名
func (tpc TabsTransferPickerControl) LabelClassName(value string) TabsTransferPickerControl {
	return tpc.set("labelClassName", value)
}

// LabelRemark 标签备注
func (tpc TabsTransferPickerControl) LabelRemark(value string) TabsTransferPickerControl {
	return tpc.set("labelRemark", value)
}

// MaxLength 最多选择的数量
func (tpc TabsTransferPickerControl) MaxLength(value int) TabsTransferPickerControl {
	return tpc.set("maxLength", value)
}

// Mode 选择模式
func (tpc TabsTransferPickerControl) Mode(value string) TabsTransferPickerControl {
	return tpc.set("mode", value)
}

// Multiple 是否多选
func (tpc TabsTransferPickerControl) Multiple(value bool) TabsTransferPickerControl {
	return tpc.set("multiple", value)
}

// Name 表单项 name
func (tpc TabsTransferPickerControl) Name(value string) TabsTransferPickerControl {
	return tpc.set("name", value)
}

// OnChange 变化时回调
func (tpc TabsTransferPickerControl) OnChange(value string) TabsTransferPickerControl {
	return tpc.set("onChange", value)
}

// Options 选项
func (tpc TabsTransferPickerControl) Options(value interface{}) TabsTransferPickerControl {
	return tpc.set("options", value)
}

// OptionsSrc 选项数据接口
func (tpc TabsTransferPickerControl) OptionsSrc(value string) TabsTransferPickerControl {
	return tpc.set("optionsSrc", value)
}

// ReadOnly 是否只读
func (tpc TabsTransferPickerControl) ReadOnly(value bool) TabsTransferPickerControl {
	return tpc.set("readOnly", value)
}

// Required 是否必填
func (tpc TabsTransferPickerControl) Required(value bool) TabsTransferPickerControl {
	return tpc.set("required", value)
}

// RequiredApi 提交时必填 API
func (tpc TabsTransferPickerControl) RequiredApi(value string) TabsTransferPickerControl {
	return tpc.set("requiredApi", value)
}

// Source 接口地址
func (tpc TabsTransferPickerControl) Source(value string) TabsTransferPickerControl {
	return tpc.set("source", value)
}

// Tabs 选项卡
func (tpc TabsTransferPickerControl) Tabs(value interface{}) TabsTransferPickerControl {
	return tpc.set("tabs", value)
}

// Tooltip 提示信息
func (tpc TabsTransferPickerControl) Tooltip(value string) TabsTransferPickerControl {
	return tpc.set("tooltip", value)
}

// Value 表单值
func (tpc TabsTransferPickerControl) Value(value interface{}) TabsTransferPickerControl {
	return tpc.set("value", value)
}

// Vertical 列表方向
func (tpc TabsTransferPickerControl) Vertical(value bool) TabsTransferPickerControl {
	return tpc.set("vertical", value)
}

// Visible 是否显示
func (tpc TabsTransferPickerControl) Visible(value bool) TabsTransferPickerControl {
	return tpc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (tpc TabsTransferPickerControl) VisibleOn(value string) TabsTransferPickerControl {
	return tpc.set("visibleOn", value)
}

// Width 宽度
func (tpc TabsTransferPickerControl) Width(value string) TabsTransferPickerControl {
	return tpc.set("width", value)
}

// WidthConfig 配置宽度
func (tpc TabsTransferPickerControl) WidthConfig(value string) TabsTransferPickerControl {
	return tpc.set("widthConfig", value)
}

// Xname 配置 key name
func (tpc TabsTransferPickerControl) Xname(value string) TabsTransferPickerControl {
	return tpc.set("xname", value)
}

// Xtype 配置类型
func (tpc TabsTransferPickerControl) Xtype(value string) TabsTransferPickerControl {
	return tpc.set("xtype", value)
}

// Component 配置自定义组件
func (tpc TabsTransferPickerControl) Component(value string) TabsTransferPickerControl {
	return tpc.set("component", value)
}
