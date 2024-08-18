package comp

// TreeControl 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tree
//
// @author  slowlyo
// @version 6.7.0
type TreeControl Schema

// NewTreeControl 创建一个新的 TreeControl 实例
func NewTreeControl() TreeControl {
	return TreeControl{}.set("type", "input-tree")
}

func (tc TreeControl) set(key string, value interface{}) TreeControl {
	tc[key] = value
	return tc
}

// AddApi 添加时调用的接口
func (tc TreeControl) AddApi(value string) TreeControl {
	return tc.set("addApi", value)
}

// AddControls 新增时的表单项。
func (tc TreeControl) AddControls(value string) TreeControl {
	return tc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (tc TreeControl) AddDialog(value string) TreeControl {
	return tc.set("addDialog", value)
}

// AutoCheckChildren ui级联关系，true代表级联选中，false代表不级联，默认为true
func (tc TreeControl) AutoCheckChildren(value bool) TreeControl {
	return tc.set("autoCheckChildren", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc TreeControl) AutoFill(value string) TreeControl {
	return tc.set("autoFill", value)
}

// Cascade 该属性代表数据级联关系，autoCheckChildren为true时生效，默认为false，具体数据级联关系如下： 1.casacde为false，ui行为为级联选中子节点，子节点禁用；值只包含父节点的值 2.cascade为false，withChildren为true，ui行为为级联选中子节点，子节点禁用；值包含父子节点的值 3.cascade为true，ui行为级联选中子节点，子节点可反选，值包含父子节点的值，此时withChildren属性失效 4.cascade不论为true还是false，onlyChildren为true，ui行为级联选中子节点，子节点可反选，值只包含子节点的值
func (tc TreeControl) Cascade(value bool) TreeControl {
	return tc.set("cascade", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TreeControl) ClassName(value string) TreeControl {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc TreeControl) ClearValueOnHidden(value bool) TreeControl {
	return tc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (tc TreeControl) Clearable(value bool) TreeControl {
	return tc.set("clearable", value)
}

// Creatable 是否可以新增
func (tc TreeControl) Creatable(value bool) TreeControl {
	return tc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tc TreeControl) CreateBtnLabel(value string) TreeControl {
	return tc.set("createBtnLabel", value)
}

// DeferApi 懒加载接口 (懒加载接口)
func (tc TreeControl) DeferApi(value string) TreeControl {
	return tc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tc TreeControl) DeferField(value string) TreeControl {
	return tc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tc TreeControl) DeleteApi(value string) TreeControl {
	return tc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (tc TreeControl) DeleteConfirmText(value string) TreeControl {
	return tc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tc TreeControl) Delimiter(value string) TreeControl {
	return tc.set("delimiter", value)
}

// Desc
func (tc TreeControl) Desc(value string) TreeControl {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc TreeControl) Description(value string) TreeControl {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc TreeControl) DescriptionClassName(value string) TreeControl {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc TreeControl) Disabled(value bool) TreeControl {
	return tc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TreeControl) DisabledOn(value string) TreeControl {
	return tc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tc TreeControl) EditApi(value string) TreeControl {
	return tc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tc TreeControl) EditControls(value string) TreeControl {
	return tc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (tc TreeControl) EditDialog(value string) TreeControl {
	return tc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tc TreeControl) Editable(value bool) TreeControl {
	return tc.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc TreeControl) EditorSetting(value string) TreeControl {
	return tc.set("editorSetting", value)
}

// EnableDefaultIcon 是否为选项添加默认的Icon，默认值为true
func (tc TreeControl) EnableDefaultIcon(value bool) TreeControl {
	return tc.set("enableDefaultIcon", value)
}

// EnableNodePath 是否开启节点路径模式
func (tc TreeControl) EnableNodePath(value bool) TreeControl {
	return tc.set("enableNodePath", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc TreeControl) ExtraName(value string) TreeControl {
	return tc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (tc TreeControl) ExtractValue(value bool) TreeControl {
	return tc.set("extractValue", value)
}

// HeightAuto 高度自动增长？
func (tc TreeControl) HeightAuto(value bool) TreeControl {
	return tc.set("heightAuto", value)
}

// Hidden 是否隐藏
func (tc TreeControl) Hidden(value bool) TreeControl {
	return tc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TreeControl) HiddenOn(value string) TreeControl {
	return tc.set("hiddenOn", value)
}

// HideRoot 是否隐藏顶级
func (tc TreeControl) HideRoot(value bool) TreeControl {
	return tc.set("hideRoot", value)
}

// HighlightTxt 需要高亮的字符串
func (tc TreeControl) HighlightTxt(value string) TreeControl {
	return tc.set("highlightTxt", value)
}

// Hint 输入提示，聚焦的时候显示
func (tc TreeControl) Hint(value string) TreeControl {
	return tc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tc TreeControl) Horizontal(value string) TreeControl {
	return tc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (tc TreeControl) Id(value string) TreeControl {
	return tc.set("id", value)
}

// InitAutoFill
func (tc TreeControl) InitAutoFill(value string) TreeControl {
	return tc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (tc TreeControl) InitFetch(value bool) TreeControl {
	return tc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tc TreeControl) InitFetchOn(value string) TreeControl {
	return tc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tc TreeControl) Inline(value bool) TreeControl {
	return tc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tc TreeControl) InputClassName(value string) TreeControl {
	return tc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (tc TreeControl) JoinValues(value bool) TreeControl {
	return tc.set("joinValues", value)
}

// Label 描述标题
func (tc TreeControl) Label(value string) TreeControl {
	return tc.set("label", value)
}

// LabelClassName 配置 label className (配置 label className)
func (tc TreeControl) LabelClassName(value string) TreeControl {
	return tc.set("labelClassName", value)
}

// LabelRemark
func (tc TreeControl) LabelRemark(value string) TreeControl {
	return tc.set("labelRemark", value)
}

// LabelWidth 配置 label 宽度。
func (tc TreeControl) LabelWidth(value string) TreeControl {
	return tc.set("labelWidth", value)
}

// MaxLength 最大长度。
func (tc TreeControl) MaxLength(value int) TreeControl {
	return tc.set("maxLength", value)
}

// MinLength 最小长度。
func (tc TreeControl) MinLength(value int) TreeControl {
	return tc.set("minLength", value)
}

// Name 表单项 name。
func (tc TreeControl) Name(value string) TreeControl {
	return tc.set("name", value)
}

// NodeValue 选项节点值，表示节点选中后表单项的值。
func (tc TreeControl) NodeValue(value string) TreeControl {
	return tc.set("nodeValue", value)
}

// Options 选项
func (tc TreeControl) Options(value interface{}) TreeControl {
	return tc.set("options", value)
}

// Placeholder 提示文字
func (tc TreeControl) Placeholder(value string) TreeControl {
	return tc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tc TreeControl) ReadOnly(value bool) TreeControl {
	return tc.set("readOnly", value)
}

// RootLabel 顶级 label
func (tc TreeControl) RootLabel(value string) TreeControl {
	return tc.set("rootLabel", value)
}

// SelectMode 选择模式。 支持 multiple，tags 和 single
func (tc TreeControl) SelectMode(value string) TreeControl {
	return tc.set("selectMode", value)
}

// ShowIcon 是否显示图标
func (tc TreeControl) ShowIcon(value bool) TreeControl {
	return tc.set("showIcon", value)
}

// ShowSearch 是否显示搜索框
func (tc TreeControl) ShowSearch(value bool) TreeControl {
	return tc.set("showSearch", value)
}

// ShowTreeIcon 是否显示树形图标
func (tc TreeControl) ShowTreeIcon(value bool) TreeControl {
	return tc.set("showTreeIcon", value)
}

// Size 表单项尺寸
func (tc TreeControl) Size(value string) TreeControl {
	return tc.set("size", value)
}

// Source 选项来源
func (tc TreeControl) Source(value string) TreeControl {
	return tc.set("source", value)
}

// StrictMode 是否开启严格模式。设置为true时将禁用部分功能。
func (tc TreeControl) StrictMode(value bool) TreeControl {
	return tc.set("strictMode", value)
}

// SwitcherIcon 开关图标
func (tc TreeControl) SwitcherIcon(value string) TreeControl {
	return tc.set("switcherIcon", value)
}

// Tips 提示文字
func (tc TreeControl) Tips(value string) TreeControl {
	return tc.set("tips", value)
}

// Value 表单值
func (tc TreeControl) Value(value interface{}) TreeControl {
	return tc.set("value", value)
}

// ValueField 选项值字段
func (tc TreeControl) ValueField(value string) TreeControl {
	return tc.set("valueField", value)
}

// Virtualized 是否启用虚拟列表
func (tc TreeControl) Virtualized(value bool) TreeControl {
	return tc.set("virtualized", value)
}

// Width 组件宽度
func (tc TreeControl) Width(value string) TreeControl {
	return tc.set("width", value)
}

// WidthUnit 组件宽度单位
func (tc TreeControl) WidthUnit(value string) TreeControl {
	return tc.set("widthUnit", value)
}

// StaticClassName 静态展示类名
func (tc TreeControl) StaticClassName(value string) TreeControl {
	return tc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (tc TreeControl) StaticInputClassName(value string) TreeControl {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (tc TreeControl) StaticLabelClassName(value string) TreeControl {
	return tc.set("staticLabelClassName", value)
}

// StaticOn 静态展示表达式
func (tc TreeControl) StaticOn(value string) TreeControl {
	return tc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tc TreeControl) StaticPlaceholder(value string) TreeControl {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema 静态展示表单项Schema
func (tc TreeControl) StaticSchema(value string) TreeControl {
	return tc.set("staticSchema", value)
}

// Style 组件样式
func (tc TreeControl) Style(value string) TreeControl {
	return tc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (tc TreeControl) SubmitOnChange(value bool) TreeControl {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder
func (tc TreeControl) TestIdBuilder(value string) TreeControl {
	return tc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tc TreeControl) UseMobileUI(value bool) TreeControl {
	return tc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (tc TreeControl) ValidateApi(value string) TreeControl {
	return tc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (tc TreeControl) ValidateOnChange(value bool) TreeControl {
	return tc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (tc TreeControl) ValidationErrors(value string) TreeControl {
	return tc.set("validationErrors", value)
}

// Validations
func (tc TreeControl) Validations(value string) TreeControl {
	return tc.set("validations", value)
}

// ValuesNoWrap 多选模式，值太多时是否避免折行
func (tc TreeControl) ValuesNoWrap(value bool) TreeControl {
	return tc.set("valuesNoWrap", value)
}

// Visible 是否显示
func (tc TreeControl) Visible(value bool) TreeControl {
	return tc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (tc TreeControl) VisibleOn(value string) TreeControl {
	return tc.set("visibleOn", value)
}

// WithChildren 选父级的时候是否把子节点的值也包含在内。
func (tc TreeControl) WithChildren(value bool) TreeControl {
	return tc.set("withChildren", value)
}
