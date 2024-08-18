package comp

// TreeSelectControl 下拉选择框
//
// @version 6.7.0
type TreeSelectControl Schema

// NewTreeSelectControl 创建一个新的 TreeSelectControl 实例
func NewTreeSelectControl() TreeSelectControl {
	return TreeSelectControl{}.set("type", "tree-select")
}

func (tsc TreeSelectControl) set(key string, value interface{}) TreeSelectControl {
	tsc[key] = value
	return tsc
}

// AddApi 添加时调用的接口
func (tsc TreeSelectControl) AddApi(value string) TreeSelectControl {
	return tsc.set("addApi", value)
}

// AddControls 新增时的表单项
func (tsc TreeSelectControl) AddControls(value string) TreeSelectControl {
	return tsc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (tsc TreeSelectControl) AddDialog(value string) TreeSelectControl {
	return tsc.set("addDialog", value)
}

// AutoCheckChildren 是否自动选中子节点
func (tsc TreeSelectControl) AutoCheckChildren(value string) TreeSelectControl {
	return tsc.set("autoCheckChildren", value)
}

// AutoFill 自动填充
func (tsc TreeSelectControl) AutoFill(value string) TreeSelectControl {
	return tsc.set("autoFill", value)
}

// Cascade 父子之间是否完全独立
func (tsc TreeSelectControl) Cascade(value bool) TreeSelectControl {
	return tsc.set("cascade", value)
}

// ClassName 容器 css 类名
func (tsc TreeSelectControl) ClassName(value string) TreeSelectControl {
	return tsc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (tsc TreeSelectControl) ClearValueOnHidden(value bool) TreeSelectControl {
	return tsc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (tsc TreeSelectControl) Clearable(value bool) TreeSelectControl {
	return tsc.set("clearable", value)
}

// Creatable 是否可以新增
func (tsc TreeSelectControl) Creatable(value bool) TreeSelectControl {
	return tsc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tsc TreeSelectControl) CreateBtnLabel(value string) TreeSelectControl {
	return tsc.set("createBtnLabel", value)
}

// DeferApi 懒加载接口
func (tsc TreeSelectControl) DeferApi(value string) TreeSelectControl {
	return tsc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tsc TreeSelectControl) DeferField(value string) TreeSelectControl {
	return tsc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tsc TreeSelectControl) DeleteApi(value string) TreeSelectControl {
	return tsc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (tsc TreeSelectControl) DeleteConfirmText(value string) TreeSelectControl {
	return tsc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tsc TreeSelectControl) Delimiter(value string) TreeSelectControl {
	return tsc.set("delimiter", value)
}

// Desc
func (tsc TreeSelectControl) Desc(value string) TreeSelectControl {
	return tsc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (tsc TreeSelectControl) Description(value string) TreeSelectControl {
	return tsc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (tsc TreeSelectControl) DescriptionClassName(value string) TreeSelectControl {
	return tsc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tsc TreeSelectControl) Disabled(value bool) TreeSelectControl {
	return tsc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (tsc TreeSelectControl) DisabledOn(value string) TreeSelectControl {
	return tsc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tsc TreeSelectControl) EditApi(value string) TreeSelectControl {
	return tsc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tsc TreeSelectControl) EditControls(value string) TreeSelectControl {
	return tsc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (tsc TreeSelectControl) EditDialog(value string) TreeSelectControl {
	return tsc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tsc TreeSelectControl) Editable(value bool) TreeSelectControl {
	return tsc.set("editable", value)
}

// EditorSetting 编辑器配置
func (tsc TreeSelectControl) EditorSetting(value string) TreeSelectControl {
	return tsc.set("editorSetting", value)
}

// EnableDefaultIcon 是否为选项添加默认的Icon
func (tsc TreeSelectControl) EnableDefaultIcon(value bool) TreeSelectControl {
	return tsc.set("enableDefaultIcon", value)
}

// EnableNodePath 是否开启节点路径模式
func (tsc TreeSelectControl) EnableNodePath(value bool) TreeSelectControl {
	return tsc.set("enableNodePath", value)
}

// ExtraName 额外的字段名
func (tsc TreeSelectControl) ExtraName(value string) TreeSelectControl {
	return tsc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (tsc TreeSelectControl) ExtractValue(value bool) TreeSelectControl {
	return tsc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tsc TreeSelectControl) Hidden(value bool) TreeSelectControl {
	return tsc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (tsc TreeSelectControl) HiddenOn(value string) TreeSelectControl {
	return tsc.set("hiddenOn", value)
}

// HideNodePathLabel 是否隐藏选择框中已选中节点的祖先节点的文本信息
func (tsc TreeSelectControl) HideNodePathLabel(value bool) TreeSelectControl {
	return tsc.set("hideNodePathLabel", value)
}

// HideRoot 是否隐藏顶级
func (tsc TreeSelectControl) HideRoot(value bool) TreeSelectControl {
	return tsc.set("hideRoot", value)
}

// Hint 输入提示
func (tsc TreeSelectControl) Hint(value string) TreeSelectControl {
	return tsc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (tsc TreeSelectControl) Horizontal(value string) TreeSelectControl {
	return tsc.set("horizontal", value)
}

// Id 组件唯一 id
func (tsc TreeSelectControl) Id(value string) TreeSelectControl {
	return tsc.set("id", value)
}

// InitAutoFill
func (tsc TreeSelectControl) InitAutoFill(value string) TreeSelectControl {
	return tsc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (tsc TreeSelectControl) InitFetch(value bool) TreeSelectControl {
	return tsc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tsc TreeSelectControl) InitFetchOn(value string) TreeSelectControl {
	return tsc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (tsc TreeSelectControl) Inline(value bool) TreeSelectControl {
	return tsc.set("inline", value)
}

// InputClassName 配置 input className
func (tsc TreeSelectControl) InputClassName(value string) TreeSelectControl {
	return tsc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交
func (tsc TreeSelectControl) JoinValues(value bool) TreeSelectControl {
	return tsc.set("joinValues", value)
}

// Label 描述标题
func (tsc TreeSelectControl) Label(value string) TreeSelectControl {
	return tsc.set("label", value)
}

// LabelAlign 描述标题
func (tsc TreeSelectControl) LabelAlign(value string) TreeSelectControl {
	return tsc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (tsc TreeSelectControl) LabelClassName(value string) TreeSelectControl {
	return tsc.set("labelClassName", value)
}

// LabelField 设置label字段
func (tsc TreeSelectControl) LabelField(value string) TreeSelectControl {
	return tsc.set("labelField", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (tsc TreeSelectControl) LabelRemark(value string) TreeSelectControl {
	return tsc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (tsc TreeSelectControl) LabelWidth(value string) TreeSelectControl {
	return tsc.set("labelWidth", value)
}

// MaxTagCount 标签的最大展示数量
func (tsc TreeSelectControl) MaxTagCount(value string) TreeSelectControl {
	return tsc.set("maxTagCount", value)
}

// MenuTpl 自定义选项
func (tsc TreeSelectControl) MenuTpl(value string) TreeSelectControl {
	return tsc.set("menuTpl", value)
}

// Multiple 是否多选
func (tsc TreeSelectControl) Multiple(value bool) TreeSelectControl {
	return tsc.set("multiple", value)
}

// Name 表单字段名
func (tsc TreeSelectControl) Name(value string) TreeSelectControl {
	return tsc.set("name", value)
}

// Options 选项
func (tsc TreeSelectControl) Options(value interface{}) TreeSelectControl {
	return tsc.set("options", value)
}

// OptionValues
func (tsc TreeSelectControl) OptionValues(value string) TreeSelectControl {
	return tsc.set("optionValues", value)
}

// OptionLabel
func (tsc TreeSelectControl) OptionLabel(value string) TreeSelectControl {
	return tsc.set("optionLabel", value)
}

// OptionName
func (tsc TreeSelectControl) OptionName(value string) TreeSelectControl {
	return tsc.set("optionName", value)
}

// OptionValue
func (tsc TreeSelectControl) OptionValue(value string) TreeSelectControl {
	return tsc.set("optionValue", value)
}

// OptionIcon
func (tsc TreeSelectControl) OptionIcon(value string) TreeSelectControl {
	return tsc.set("optionIcon", value)
}

// Placeholder 提示文案
func (tsc TreeSelectControl) Placeholder(value string) TreeSelectControl {
	return tsc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tsc TreeSelectControl) ReadOnly(value bool) TreeSelectControl {
	return tsc.set("readOnly", value)
}

// ReBuildApi 下拉框的重建接口
func (tsc TreeSelectControl) ReBuildApi(value string) TreeSelectControl {
	return tsc.set("reBuildApi", value)
}

// Reload 是否可以通过 reloadApi 来刷新选项
func (tsc TreeSelectControl) Reload(value bool) TreeSelectControl {
	return tsc.set("reload", value)
}

// ReloadApi 刷新选项的 API
func (tsc TreeSelectControl) ReloadApi(value string) TreeSelectControl {
	return tsc.set("reloadApi", value)
}

// ResetValue 表单重置的值
func (tsc TreeSelectControl) ResetValue(value interface{}) TreeSelectControl {
	return tsc.set("resetValue", value)
}

// RootLabel 根节点文本
func (tsc TreeSelectControl) RootLabel(value string) TreeSelectControl {
	return tsc.set("rootLabel", value)
}

// Selectable 用来控制下拉框中项是否可选
func (tsc TreeSelectControl) Selectable(value string) TreeSelectControl {
	return tsc.set("selectable", value)
}

// ShowIcon 是否显示图标
func (tsc TreeSelectControl) ShowIcon(value bool) TreeSelectControl {
	return tsc.set("showIcon", value)
}

// ShowRadio 是否展示单选框
func (tsc TreeSelectControl) ShowRadio(value bool) TreeSelectControl {
	return tsc.set("showRadio", value)
}

// ShowSearch 是否展示搜索框
func (tsc TreeSelectControl) ShowSearch(value bool) TreeSelectControl {
	return tsc.set("showSearch", value)
}

// Source 数据源
func (tsc TreeSelectControl) Source(value string) TreeSelectControl {
	return tsc.set("source", value)
}

// StaticClassName 静态 className
func (tsc TreeSelectControl) StaticClassName(value string) TreeSelectControl {
	return tsc.set("staticClassName", value)
}

// Static 是否静态展示
func (tsc TreeSelectControl) Static(value bool) TreeSelectControl {
	return tsc.set("static", value)
}

// StaticLabel 静态展示文案
func (tsc TreeSelectControl) StaticLabel(value string) TreeSelectControl {
	return tsc.set("staticLabel", value)
}

// Taggable 是否可以添加 tag
func (tsc TreeSelectControl) Taggable(value bool) TreeSelectControl {
	return tsc.set("taggable", value)
}

// Tags 是否启用 tag
func (tsc TreeSelectControl) Tags(value bool) TreeSelectControl {
	return tsc.set("tags", value)
}

// Tip 提示
func (tsc TreeSelectControl) Tip(value string) TreeSelectControl {
	return tsc.set("tip", value)
}

// Type 组件类型
func (tsc TreeSelectControl) Type(value string) TreeSelectControl {
	return tsc.set("type", value)
}

// Value 表单项的值
func (tsc TreeSelectControl) Value(value interface{}) TreeSelectControl {
	return tsc.set("value", value)
}

// ValueField 选项的值字段
func (tsc TreeSelectControl) ValueField(value string) TreeSelectControl {
	return tsc.set("valueField", value)
}

// Vertical
func (tsc TreeSelectControl) Vertical(value string) TreeSelectControl {
	return tsc.set("vertical", value)
}

// Visible 是否可见
func (tsc TreeSelectControl) Visible(value bool) TreeSelectControl {
	return tsc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (tsc TreeSelectControl) VisibleOn(value string) TreeSelectControl {
	return tsc.set("visibleOn", value)
}

// Width 控件宽度
func (tsc TreeSelectControl) Width(value string) TreeSelectControl {
	return tsc.set("width", value)
}

// Wrapping 是否换行
func (tsc TreeSelectControl) Wrapping(value bool) TreeSelectControl {
	return tsc.set("wrapping", value)
}

// 自定义其他字段
func (tsc TreeSelectControl) Other(key string, value interface{}) TreeSelectControl {
	return tsc.set(key, value)
}

// RootCreatable 顶级节点是否可以创建子节点
func (tsc TreeSelectControl) RootCreatable(value bool) TreeSelectControl {
	return tsc.set("rootCreatable", value)
}

// RootValue 顶级选项的值
func (tsc TreeSelectControl) RootValue(value string) TreeSelectControl {
	return tsc.set("rootValue", value)
}

// Row
func (tsc TreeSelectControl) Row(value string) TreeSelectControl {
	return tsc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (tsc TreeSelectControl) SaveImmediately(value bool) TreeSelectControl {
	return tsc.set("saveImmediately", value)
}

// Searchable 是否可搜索
func (tsc TreeSelectControl) Searchable(value bool) TreeSelectControl {
	return tsc.set("searchable", value)
}

// SelectFirst 默认选择选项第一个值
func (tsc TreeSelectControl) SelectFirst(value bool) TreeSelectControl {
	return tsc.set("selectFirst", value)
}

// ShowOutline 是否显示展开线
func (tsc TreeSelectControl) ShowOutline(value bool) TreeSelectControl {
	return tsc.set("showOutline", value)
}

// Size 表单项大小
func (tsc TreeSelectControl) Size(value string) TreeSelectControl {
	return tsc.set("size", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (tsc TreeSelectControl) StaticInputClassName(value string) TreeSelectControl {
	return tsc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (tsc TreeSelectControl) StaticLabelClassName(value string) TreeSelectControl {
	return tsc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (tsc TreeSelectControl) StaticOn(value string) TreeSelectControl {
	return tsc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tsc TreeSelectControl) StaticPlaceholder(value string) TreeSelectControl {
	return tsc.set("staticPlaceholder", value)
}

// StaticSchema
func (tsc TreeSelectControl) StaticSchema(value string) TreeSelectControl {
	return tsc.set("staticSchema", value)
}

// Style 组件样式
func (tsc TreeSelectControl) Style(value string) TreeSelectControl {
	return tsc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (tsc TreeSelectControl) SubmitOnChange(value bool) TreeSelectControl {
	return tsc.set("submitOnChange", value)
}

// TestIdBuilder
func (tsc TreeSelectControl) TestIdBuilder(value string) TreeSelectControl {
	return tsc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tsc TreeSelectControl) UseMobileUI(value bool) TreeSelectControl {
	return tsc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (tsc TreeSelectControl) ValidateApi(value string) TreeSelectControl {
	return tsc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (tsc TreeSelectControl) ValidateOnChange(value bool) TreeSelectControl {
	return tsc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (tsc TreeSelectControl) ValidationErrors(value string) TreeSelectControl {
	return tsc.set("validationErrors", value)
}

// Validations
func (tsc TreeSelectControl) Validations(value string) TreeSelectControl {
	return tsc.set("validations", value)
}

// ValuesNoWrap 多选模式，值太多时是否避免折行
func (tsc TreeSelectControl) ValuesNoWrap(value bool) TreeSelectControl {
	return tsc.set("valuesNoWrap", value)
}

// WithChildren 选父级的时候是否把子节点的值也包含在内
func (tsc TreeSelectControl) WithChildren(value bool) TreeSelectControl {
	return tsc.set("withChildren", value)
}
