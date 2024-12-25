package comp

// chainedSelect 链式下拉框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chained-select
type chainedSelect schema

// ChainedSelect 创建一个新的 ChainedSelect 实例
func ChainedSelect() chainedSelect {
	return make(chainedSelect).set("type", "chained-select")
}

func (c chainedSelect) set(key string, value any) chainedSelect {
	c[key] = value
	return c
}

// AddApi 添加时调用的接口
func (c chainedSelect) AddApi(value string) chainedSelect {
	return c.set("addApi", value)
}

// AddControls 新增时的表单项
func (c chainedSelect) AddControls(value string) chainedSelect {
	return c.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (c chainedSelect) AddDialog(value string) chainedSelect {
	return c.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (c chainedSelect) AutoFill(value string) chainedSelect {
	return c.set("autoFill", value)
}

// ClassName 容器 css 类名
func (c chainedSelect) ClassName(value string) chainedSelect {
	return c.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (c chainedSelect) ClearValueOnHidden(value bool) chainedSelect {
	return c.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (c chainedSelect) Clearable(value bool) chainedSelect {
	return c.set("clearable", value)
}

// Creatable 是否可以新增
func (c chainedSelect) Creatable(value bool) chainedSelect {
	return c.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (c chainedSelect) CreateBtnLabel(value string) chainedSelect {
	return c.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (c chainedSelect) DeferApi(value string) chainedSelect {
	return c.set("deferApi", value)
}

// DeferField 懒加载字段
func (c chainedSelect) DeferField(value string) chainedSelect {
	return c.set("deferField", value)
}

// DeleteApi 选项删除 API
func (c chainedSelect) DeleteApi(value string) chainedSelect {
	return c.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (c chainedSelect) DeleteConfirmText(value string) chainedSelect {
	return c.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (c chainedSelect) Delimiter(value string) chainedSelect {
	return c.set("delimiter", value)
}

// Desc
func (c chainedSelect) Desc(value string) chainedSelect {
	return c.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (c chainedSelect) Description(value string) chainedSelect {
	return c.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (c chainedSelect) DescriptionClassName(value string) chainedSelect {
	return c.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (c chainedSelect) Disabled(value bool) chainedSelect {
	return c.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (c chainedSelect) DisabledOn(value string) chainedSelect {
	return c.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (c chainedSelect) EditApi(value string) chainedSelect {
	return c.set("editApi", value)
}

// EditControls 选项修改的表单项
func (c chainedSelect) EditControls(value string) chainedSelect {
	return c.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (c chainedSelect) EditDialog(value string) chainedSelect {
	return c.set("editDialog", value)
}

// Editable 是否可以编辑
func (c chainedSelect) Editable(value bool) chainedSelect {
	return c.set("editable", value)
}

// EditorSetting 编辑器配置
func (c chainedSelect) EditorSetting(value string) chainedSelect {
	return c.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (c chainedSelect) ExtraName(value string) chainedSelect {
	return c.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (c chainedSelect) ExtractValue(value bool) chainedSelect {
	return c.set("extractValue", value)
}

// Hidden 是否隐藏
func (c chainedSelect) Hidden(value bool) chainedSelect {
	return c.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (c chainedSelect) HiddenOn(value string) chainedSelect {
	return c.set("hiddenOn", value)
}

// Hint 输入提示
func (c chainedSelect) Hint(value string) chainedSelect {
	return c.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (c chainedSelect) Horizontal(value string) chainedSelect {
	return c.set("horizontal", value)
}

// Id 组件唯一 id
func (c chainedSelect) ID(value string) chainedSelect {
	return c.set("id", value)
}

// InitAutoFill
func (c chainedSelect) InitAutoFill(value string) chainedSelect {
	return c.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (c chainedSelect) InitFetch(value bool) chainedSelect {
	return c.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (c chainedSelect) InitFetchOn(value string) chainedSelect {
	return c.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (c chainedSelect) Inline(value bool) chainedSelect {
	return c.set("inline", value)
}

// InputClassName 配置 input className
func (c chainedSelect) InputClassName(value string) chainedSelect {
	return c.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式值处理
func (c chainedSelect) JoinValues(value bool) chainedSelect {
	return c.set("joinValues", value)
}

// Label 描述标题
func (c chainedSelect) Label(value string) chainedSelect {
	return c.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (c chainedSelect) LabelAlign(value string) chainedSelect {
	return c.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (c chainedSelect) LabelClassName(value string) chainedSelect {
	return c.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (c chainedSelect) LabelRemark(value string) chainedSelect {
	return c.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (c chainedSelect) LabelWidth(value string) chainedSelect {
	return c.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (c chainedSelect) Mode(value string) chainedSelect {
	return c.set("mode", value)
}

// Multiple 是否为多选模式
func (c chainedSelect) Multiple(value bool) chainedSelect {
	return c.set("multiple", value)
}

// Name 字段名
func (c chainedSelect) Name(value string) chainedSelect {
	return c.set("name", value)
}

// OnEvent 事件动作配置
func (c chainedSelect) OnEvent(value any) chainedSelect {
	return c.set("onEvent", value)
}

// Options 选项集合
func (c chainedSelect) Options(value ...MOption) chainedSelect {
	return c.set("options", value)
}

// Placeholder 占位符
func (c chainedSelect) Placeholder(value string) chainedSelect {
	return c.set("placeholder", value)
}

// ReadOnly 是否只读
func (c chainedSelect) ReadOnly(value bool) chainedSelect {
	return c.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (c chainedSelect) ReadOnlyOn(value string) chainedSelect {
	return c.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (c chainedSelect) Remark(value string) chainedSelect {
	return c.set("remark", value)
}

// Removable 是否可删除
func (c chainedSelect) Removable(value bool) chainedSelect {
	return c.set("removable", value)
}

// Required 是否为必填
func (c chainedSelect) Required(value bool) chainedSelect {
	return c.set("required", value)
}

// ResetValue 点清除按钮时，将表单项设置成当前配置的值
func (c chainedSelect) ResetValue(value string) chainedSelect {
	return c.set("resetValue", value)
}

// Row
func (c chainedSelect) Row(value string) chainedSelect {
	return c.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (c chainedSelect) SaveImmediately(value bool) chainedSelect {
	return c.set("saveImmediately", value)
}

// SelectFirst 默认选择选项第一个值
func (c chainedSelect) SelectFirst(value bool) chainedSelect {
	return c.set("selectFirst", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (c chainedSelect) Size(value string) chainedSelect {
	return c.set("size", value)
}

// Source 可用来通过 API 拉取 options
func (c chainedSelect) Source(value string) chainedSelect {
	return c.set("source", value)
}

// Static 是否静态展示
func (c chainedSelect) Static(value bool) chainedSelect {
	return c.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象)
func (c chainedSelect) StaticClassName(value string) chainedSelect {
	return c.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象)
func (c chainedSelect) StaticInputClassName(value string) chainedSelect {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象)
func (c chainedSelect) StaticLabelClassName(value string) chainedSelect {
	return c.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`)
func (c chainedSelect) StaticOn(value string) chainedSelect {
	return c.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (c chainedSelect) StaticPlaceholder(value string) chainedSelect {
	return c.set("staticPlaceholder", value)
}

// StaticSchema
func (c chainedSelect) StaticSchema(value string) chainedSelect {
	return c.set("staticSchema", value)
}

// Style 组件样式
func (c chainedSelect) Style(value any) chainedSelect {
	return c.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (c chainedSelect) SubmitOnChange(value bool) chainedSelect {
	return c.set("submitOnChange", value)
}

// TestIdBuilder
func (c chainedSelect) TestIdBuilder(value string) chainedSelect {
	return c.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (c chainedSelect) UseMobileUI(value bool) chainedSelect {
	return c.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (c chainedSelect) ValidateApi(value string) chainedSelect {
	return c.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (c chainedSelect) ValidateOnChange(value bool) chainedSelect {
	return c.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (c chainedSelect) ValidationErrors(value string) chainedSelect {
	return c.set("validationErrors", value)
}

// Validations
func (c chainedSelect) Validations(value string) chainedSelect {
	return c.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量
func (c chainedSelect) Value(value string) chainedSelect {
	return c.set("value", value)
}

// ValuesNoWrap 多选模式，值太多时是否避免折行
func (c chainedSelect) ValuesNoWrap(value bool) chainedSelect {
	return c.set("valuesNoWrap", value)
}

// Visible 是否显示
func (c chainedSelect) Visible(value bool) chainedSelect {
	return c.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`)
func (c chainedSelect) VisibleOn(value string) chainedSelect {
	return c.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (c chainedSelect) Width(value string) chainedSelect {
	return c.set("width", value)
}
