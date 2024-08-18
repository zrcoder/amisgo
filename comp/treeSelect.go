package comp

// treeSelect 下拉选择框
//
// @version 6.7.0
type treeSelect schema

func TreeSelect() treeSelect {
	return treeSelect{}.set("type", "tree-select")
}

func (tsc treeSelect) set(key string, value interface{}) treeSelect {
	tsc[key] = value
	return tsc
}

// AddApi 添加时调用的接口
func (tsc treeSelect) AddApi(value string) treeSelect {
	return tsc.set("addApi", value)
}

// AddControls 新增时的表单项
func (tsc treeSelect) AddControls(value string) treeSelect {
	return tsc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (tsc treeSelect) AddDialog(value string) treeSelect {
	return tsc.set("addDialog", value)
}

// AutoCheckChildren 是否自动选中子节点
func (tsc treeSelect) AutoCheckChildren(value string) treeSelect {
	return tsc.set("autoCheckChildren", value)
}

// AutoFill 自动填充
func (tsc treeSelect) AutoFill(value string) treeSelect {
	return tsc.set("autoFill", value)
}

// Cascade 父子之间是否完全独立
func (tsc treeSelect) Cascade(value bool) treeSelect {
	return tsc.set("cascade", value)
}

// ClassName 容器 css 类名
func (tsc treeSelect) ClassName(value string) treeSelect {
	return tsc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (tsc treeSelect) ClearValueOnHidden(value bool) treeSelect {
	return tsc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (tsc treeSelect) Clearable(value bool) treeSelect {
	return tsc.set("clearable", value)
}

// Creatable 是否可以新增
func (tsc treeSelect) Creatable(value bool) treeSelect {
	return tsc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (tsc treeSelect) CreateBtnLabel(value string) treeSelect {
	return tsc.set("createBtnLabel", value)
}

// DeferApi 懒加载接口
func (tsc treeSelect) DeferApi(value string) treeSelect {
	return tsc.set("deferApi", value)
}

// DeferField 懒加载字段
func (tsc treeSelect) DeferField(value string) treeSelect {
	return tsc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (tsc treeSelect) DeleteApi(value string) treeSelect {
	return tsc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (tsc treeSelect) DeleteConfirmText(value string) treeSelect {
	return tsc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (tsc treeSelect) Delimiter(value string) treeSelect {
	return tsc.set("delimiter", value)
}

// Desc
func (tsc treeSelect) Desc(value string) treeSelect {
	return tsc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (tsc treeSelect) Description(value string) treeSelect {
	return tsc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (tsc treeSelect) DescriptionClassName(value string) treeSelect {
	return tsc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tsc treeSelect) Disabled(value bool) treeSelect {
	return tsc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (tsc treeSelect) DisabledOn(value string) treeSelect {
	return tsc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (tsc treeSelect) EditApi(value string) treeSelect {
	return tsc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (tsc treeSelect) EditControls(value string) treeSelect {
	return tsc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (tsc treeSelect) EditDialog(value string) treeSelect {
	return tsc.set("editDialog", value)
}

// Editable 是否可以编辑
func (tsc treeSelect) Editable(value bool) treeSelect {
	return tsc.set("editable", value)
}

// EditorSetting 编辑器配置
func (tsc treeSelect) EditorSetting(value string) treeSelect {
	return tsc.set("editorSetting", value)
}

// EnableDefaultIcon 是否为选项添加默认的Icon
func (tsc treeSelect) EnableDefaultIcon(value bool) treeSelect {
	return tsc.set("enableDefaultIcon", value)
}

// EnableNodePath 是否开启节点路径模式
func (tsc treeSelect) EnableNodePath(value bool) treeSelect {
	return tsc.set("enableNodePath", value)
}

// ExtraName 额外的字段名
func (tsc treeSelect) ExtraName(value string) treeSelect {
	return tsc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (tsc treeSelect) ExtractValue(value bool) treeSelect {
	return tsc.set("extractValue", value)
}

// Hidden 是否隐藏
func (tsc treeSelect) Hidden(value bool) treeSelect {
	return tsc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (tsc treeSelect) HiddenOn(value string) treeSelect {
	return tsc.set("hiddenOn", value)
}

// HideNodePathLabel 是否隐藏选择框中已选中节点的祖先节点的文本信息
func (tsc treeSelect) HideNodePathLabel(value bool) treeSelect {
	return tsc.set("hideNodePathLabel", value)
}

// HideRoot 是否隐藏顶级
func (tsc treeSelect) HideRoot(value bool) treeSelect {
	return tsc.set("hideRoot", value)
}

// Hint 输入提示
func (tsc treeSelect) Hint(value string) treeSelect {
	return tsc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (tsc treeSelect) Horizontal(value string) treeSelect {
	return tsc.set("horizontal", value)
}

// Id 组件唯一 id
func (tsc treeSelect) Id(value string) treeSelect {
	return tsc.set("id", value)
}

// InitAutoFill
func (tsc treeSelect) InitAutoFill(value string) treeSelect {
	return tsc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (tsc treeSelect) InitFetch(value bool) treeSelect {
	return tsc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (tsc treeSelect) InitFetchOn(value string) treeSelect {
	return tsc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (tsc treeSelect) Inline(value bool) treeSelect {
	return tsc.set("inline", value)
}

// InputClassName 配置 input className
func (tsc treeSelect) InputClassName(value string) treeSelect {
	return tsc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交
func (tsc treeSelect) JoinValues(value bool) treeSelect {
	return tsc.set("joinValues", value)
}

// Label 描述标题
func (tsc treeSelect) Label(value string) treeSelect {
	return tsc.set("label", value)
}

// LabelAlign 描述标题
func (tsc treeSelect) LabelAlign(value string) treeSelect {
	return tsc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (tsc treeSelect) LabelClassName(value string) treeSelect {
	return tsc.set("labelClassName", value)
}

// LabelField 设置label字段
func (tsc treeSelect) LabelField(value string) treeSelect {
	return tsc.set("labelField", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (tsc treeSelect) LabelRemark(value string) treeSelect {
	return tsc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (tsc treeSelect) LabelWidth(value string) treeSelect {
	return tsc.set("labelWidth", value)
}

// MaxTagCount 标签的最大展示数量
func (tsc treeSelect) MaxTagCount(value string) treeSelect {
	return tsc.set("maxTagCount", value)
}

// MenuTpl 自定义选项
func (tsc treeSelect) MenuTpl(value string) treeSelect {
	return tsc.set("menuTpl", value)
}

// Multiple 是否多选
func (tsc treeSelect) Multiple(value bool) treeSelect {
	return tsc.set("multiple", value)
}

// Name 表单字段名
func (tsc treeSelect) Name(value string) treeSelect {
	return tsc.set("name", value)
}

// Options 选项
func (tsc treeSelect) Options(value interface{}) treeSelect {
	return tsc.set("options", value)
}

// OptionValues
func (tsc treeSelect) OptionValues(value string) treeSelect {
	return tsc.set("optionValues", value)
}

// OptionLabel
func (tsc treeSelect) OptionLabel(value string) treeSelect {
	return tsc.set("optionLabel", value)
}

// OptionName
func (tsc treeSelect) OptionName(value string) treeSelect {
	return tsc.set("optionName", value)
}

// OptionValue
func (tsc treeSelect) OptionValue(value string) treeSelect {
	return tsc.set("optionValue", value)
}

// OptionIcon
func (tsc treeSelect) OptionIcon(value string) treeSelect {
	return tsc.set("optionIcon", value)
}

// Placeholder 提示文案
func (tsc treeSelect) Placeholder(value string) treeSelect {
	return tsc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tsc treeSelect) ReadOnly(value bool) treeSelect {
	return tsc.set("readOnly", value)
}

// ReBuildApi 下拉框的重建接口
func (tsc treeSelect) ReBuildApi(value string) treeSelect {
	return tsc.set("reBuildApi", value)
}

// Reload 是否可以通过 reloadApi 来刷新选项
func (tsc treeSelect) Reload(value bool) treeSelect {
	return tsc.set("reload", value)
}

// ReloadApi 刷新选项的 API
func (tsc treeSelect) ReloadApi(value string) treeSelect {
	return tsc.set("reloadApi", value)
}

// ResetValue 表单重置的值
func (tsc treeSelect) ResetValue(value interface{}) treeSelect {
	return tsc.set("resetValue", value)
}

// RootLabel 根节点文本
func (tsc treeSelect) RootLabel(value string) treeSelect {
	return tsc.set("rootLabel", value)
}

// Selectable 用来控制下拉框中项是否可选
func (tsc treeSelect) Selectable(value string) treeSelect {
	return tsc.set("selectable", value)
}

// ShowIcon 是否显示图标
func (tsc treeSelect) ShowIcon(value bool) treeSelect {
	return tsc.set("showIcon", value)
}

// ShowRadio 是否展示单选框
func (tsc treeSelect) ShowRadio(value bool) treeSelect {
	return tsc.set("showRadio", value)
}

// ShowSearch 是否展示搜索框
func (tsc treeSelect) ShowSearch(value bool) treeSelect {
	return tsc.set("showSearch", value)
}

// Source 数据源
func (tsc treeSelect) Source(value string) treeSelect {
	return tsc.set("source", value)
}

// StaticClassName 静态 className
func (tsc treeSelect) StaticClassName(value string) treeSelect {
	return tsc.set("staticClassName", value)
}

// Static 是否静态展示
func (tsc treeSelect) Static(value bool) treeSelect {
	return tsc.set("static", value)
}

// StaticLabel 静态展示文案
func (tsc treeSelect) StaticLabel(value string) treeSelect {
	return tsc.set("staticLabel", value)
}

// Taggable 是否可以添加 tag
func (tsc treeSelect) Taggable(value bool) treeSelect {
	return tsc.set("taggable", value)
}

// Tags 是否启用 tag
func (tsc treeSelect) Tags(value bool) treeSelect {
	return tsc.set("tags", value)
}

// Tip 提示
func (tsc treeSelect) Tip(value string) treeSelect {
	return tsc.set("tip", value)
}

// Value 表单项的值
func (tsc treeSelect) Value(value interface{}) treeSelect {
	return tsc.set("value", value)
}

// ValueField 选项的值字段
func (tsc treeSelect) ValueField(value string) treeSelect {
	return tsc.set("valueField", value)
}

// Vertical
func (tsc treeSelect) Vertical(value string) treeSelect {
	return tsc.set("vertical", value)
}

// Visible 是否可见
func (tsc treeSelect) Visible(value bool) treeSelect {
	return tsc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (tsc treeSelect) VisibleOn(value string) treeSelect {
	return tsc.set("visibleOn", value)
}

// Width 控件宽度
func (tsc treeSelect) Width(value string) treeSelect {
	return tsc.set("width", value)
}

// Wrapping 是否换行
func (tsc treeSelect) Wrapping(value bool) treeSelect {
	return tsc.set("wrapping", value)
}

// 自定义其他字段
func (tsc treeSelect) Other(key string, value interface{}) treeSelect {
	return tsc.set(key, value)
}

// RootCreatable 顶级节点是否可以创建子节点
func (tsc treeSelect) RootCreatable(value bool) treeSelect {
	return tsc.set("rootCreatable", value)
}

// RootValue 顶级选项的值
func (tsc treeSelect) RootValue(value string) treeSelect {
	return tsc.set("rootValue", value)
}

// Row
func (tsc treeSelect) Row(value string) treeSelect {
	return tsc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (tsc treeSelect) SaveImmediately(value bool) treeSelect {
	return tsc.set("saveImmediately", value)
}

// Searchable 是否可搜索
func (tsc treeSelect) Searchable(value bool) treeSelect {
	return tsc.set("searchable", value)
}

// SelectFirst 默认选择选项第一个值
func (tsc treeSelect) SelectFirst(value bool) treeSelect {
	return tsc.set("selectFirst", value)
}

// ShowOutline 是否显示展开线
func (tsc treeSelect) ShowOutline(value bool) treeSelect {
	return tsc.set("showOutline", value)
}

// Size 表单项大小
func (tsc treeSelect) Size(value string) treeSelect {
	return tsc.set("size", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (tsc treeSelect) StaticInputClassName(value string) treeSelect {
	return tsc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (tsc treeSelect) StaticLabelClassName(value string) treeSelect {
	return tsc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (tsc treeSelect) StaticOn(value string) treeSelect {
	return tsc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tsc treeSelect) StaticPlaceholder(value string) treeSelect {
	return tsc.set("staticPlaceholder", value)
}

// StaticSchema
func (tsc treeSelect) StaticSchema(value string) treeSelect {
	return tsc.set("staticSchema", value)
}

// Style 组件样式
func (tsc treeSelect) Style(value string) treeSelect {
	return tsc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (tsc treeSelect) SubmitOnChange(value bool) treeSelect {
	return tsc.set("submitOnChange", value)
}

// TestIdBuilder
func (tsc treeSelect) TestIdBuilder(value string) treeSelect {
	return tsc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tsc treeSelect) UseMobileUI(value bool) treeSelect {
	return tsc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (tsc treeSelect) ValidateApi(value string) treeSelect {
	return tsc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (tsc treeSelect) ValidateOnChange(value bool) treeSelect {
	return tsc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (tsc treeSelect) ValidationErrors(value string) treeSelect {
	return tsc.set("validationErrors", value)
}

// Validations
func (tsc treeSelect) Validations(value string) treeSelect {
	return tsc.set("validations", value)
}

// ValuesNoWrap 多选模式，值太多时是否避免折行
func (tsc treeSelect) ValuesNoWrap(value bool) treeSelect {
	return tsc.set("valuesNoWrap", value)
}

// WithChildren 选父级的时候是否把子节点的值也包含在内
func (tsc treeSelect) WithChildren(value bool) treeSelect {
	return tsc.set("withChildren", value)
}
