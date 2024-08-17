package comp

// ChainedSelectControl 链式下拉框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/chained-select
type ChainedSelectControl BaseRenderer

// NewChainedSelectControl 创建一个新的 ChainedSelectControl 实例
func NewChainedSelectControl() ChainedSelectControl {
	c := ChainedSelectControl(make(BaseRenderer))
	c.set("type", "chained-select")
	return c
}

// set 设置属性值
func (c ChainedSelectControl) set(key string, value interface{}) ChainedSelectControl {
	c[key] = value
	return c
}

// AddApi 添加时调用的接口
func (c ChainedSelectControl) AddApi(value string) ChainedSelectControl {
	return c.set("addApi", value)
}

// AddControls 新增时的表单项
func (c ChainedSelectControl) AddControls(value string) ChainedSelectControl {
	return c.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (c ChainedSelectControl) AddDialog(value string) ChainedSelectControl {
	return c.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (c ChainedSelectControl) AutoFill(value string) ChainedSelectControl {
	return c.set("autoFill", value)
}

// ClassName 容器 css 类名
func (c ChainedSelectControl) ClassName(value string) ChainedSelectControl {
	return c.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (c ChainedSelectControl) ClearValueOnHidden(value bool) ChainedSelectControl {
	return c.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (c ChainedSelectControl) Clearable(value bool) ChainedSelectControl {
	return c.set("clearable", value)
}

// Creatable 是否可以新增
func (c ChainedSelectControl) Creatable(value bool) ChainedSelectControl {
	return c.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (c ChainedSelectControl) CreateBtnLabel(value string) ChainedSelectControl {
	return c.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (c ChainedSelectControl) DeferApi(value string) ChainedSelectControl {
	return c.set("deferApi", value)
}

// DeferField 懒加载字段
func (c ChainedSelectControl) DeferField(value string) ChainedSelectControl {
	return c.set("deferField", value)
}

// DeleteApi 选项删除 API
func (c ChainedSelectControl) DeleteApi(value string) ChainedSelectControl {
	return c.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (c ChainedSelectControl) DeleteConfirmText(value string) ChainedSelectControl {
	return c.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (c ChainedSelectControl) Delimiter(value string) ChainedSelectControl {
	return c.set("delimiter", value)
}

// Desc
func (c ChainedSelectControl) Desc(value string) ChainedSelectControl {
	return c.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (c ChainedSelectControl) Description(value string) ChainedSelectControl {
	return c.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (c ChainedSelectControl) DescriptionClassName(value string) ChainedSelectControl {
	return c.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (c ChainedSelectControl) Disabled(value bool) ChainedSelectControl {
	return c.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (c ChainedSelectControl) DisabledOn(value string) ChainedSelectControl {
	return c.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (c ChainedSelectControl) EditApi(value string) ChainedSelectControl {
	return c.set("editApi", value)
}

// EditControls 选项修改的表单项
func (c ChainedSelectControl) EditControls(value string) ChainedSelectControl {
	return c.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (c ChainedSelectControl) EditDialog(value string) ChainedSelectControl {
	return c.set("editDialog", value)
}

// Editable 是否可以编辑
func (c ChainedSelectControl) Editable(value bool) ChainedSelectControl {
	return c.set("editable", value)
}

// EditorSetting 编辑器配置
func (c ChainedSelectControl) EditorSetting(value string) ChainedSelectControl {
	return c.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (c ChainedSelectControl) ExtraName(value string) ChainedSelectControl {
	return c.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (c ChainedSelectControl) ExtractValue(value bool) ChainedSelectControl {
	return c.set("extractValue", value)
}

// Hidden 是否隐藏
func (c ChainedSelectControl) Hidden(value bool) ChainedSelectControl {
	return c.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (c ChainedSelectControl) HiddenOn(value string) ChainedSelectControl {
	return c.set("hiddenOn", value)
}

// Hint 输入提示
func (c ChainedSelectControl) Hint(value string) ChainedSelectControl {
	return c.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (c ChainedSelectControl) Horizontal(value string) ChainedSelectControl {
	return c.set("horizontal", value)
}

// Id 组件唯一 id
func (c ChainedSelectControl) Id(value string) ChainedSelectControl {
	return c.set("id", value)
}

// InitAutoFill
func (c ChainedSelectControl) InitAutoFill(value string) ChainedSelectControl {
	return c.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (c ChainedSelectControl) InitFetch(value bool) ChainedSelectControl {
	return c.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (c ChainedSelectControl) InitFetchOn(value string) ChainedSelectControl {
	return c.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (c ChainedSelectControl) Inline(value bool) ChainedSelectControl {
	return c.set("inline", value)
}

// InputClassName 配置 input className
func (c ChainedSelectControl) InputClassName(value string) ChainedSelectControl {
	return c.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式值处理
func (c ChainedSelectControl) JoinValues(value bool) ChainedSelectControl {
	return c.set("joinValues", value)
}

// Label 描述标题
func (c ChainedSelectControl) Label(value string) ChainedSelectControl {
	return c.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (c ChainedSelectControl) LabelAlign(value string) ChainedSelectControl {
	return c.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (c ChainedSelectControl) LabelClassName(value string) ChainedSelectControl {
	return c.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (c ChainedSelectControl) LabelRemark(value string) ChainedSelectControl {
	return c.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (c ChainedSelectControl) LabelWidth(value string) ChainedSelectControl {
	return c.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (c ChainedSelectControl) Mode(value string) ChainedSelectControl {
	return c.set("mode", value)
}

// Multiple 是否为多选模式
func (c ChainedSelectControl) Multiple(value bool) ChainedSelectControl {
	return c.set("multiple", value)
}

// Name 字段名
func (c ChainedSelectControl) Name(value string) ChainedSelectControl {
	return c.set("name", value)
}

// OnEvent 事件动作配置
func (c ChainedSelectControl) OnEvent(value string) ChainedSelectControl {
	return c.set("onEvent", value)
}

// Options 选项集合
func (c ChainedSelectControl) Options(value string) ChainedSelectControl {
	return c.set("options", value)
}

// Placeholder 占位符
func (c ChainedSelectControl) Placeholder(value string) ChainedSelectControl {
	return c.set("placeholder", value)
}

// ReadOnly 是否只读
func (c ChainedSelectControl) ReadOnly(value bool) ChainedSelectControl {
	return c.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (c ChainedSelectControl) ReadOnlyOn(value string) ChainedSelectControl {
	return c.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (c ChainedSelectControl) Remark(value string) ChainedSelectControl {
	return c.set("remark", value)
}

// Removable 是否可删除
func (c ChainedSelectControl) Removable(value bool) ChainedSelectControl {
	return c.set("removable", value)
}

// Required 是否为必填
func (c ChainedSelectControl) Required(value bool) ChainedSelectControl {
	return c.set("required", value)
}

// ResetValue 点清除按钮时，将表单项设置成当前配置的值
func (c ChainedSelectControl) ResetValue(value string) ChainedSelectControl {
	return c.set("resetValue", value)
}

// Row
func (c ChainedSelectControl) Row(value string) ChainedSelectControl {
	return c.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (c ChainedSelectControl) SaveImmediately(value bool) ChainedSelectControl {
	return c.set("saveImmediately", value)
}

// SelectFirst 默认选择选项第一个值
func (c ChainedSelectControl) SelectFirst(value bool) ChainedSelectControl {
	return c.set("selectFirst", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (c ChainedSelectControl) Size(value string) ChainedSelectControl {
	return c.set("size", value)
}

// Source 可用来通过 API 拉取 options
func (c ChainedSelectControl) Source(value string) ChainedSelectControl {
	return c.set("source", value)
}

// Static 是否静态展示
func (c ChainedSelectControl) Static(value bool) ChainedSelectControl {
	return c.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象)
func (c ChainedSelectControl) StaticClassName(value string) ChainedSelectControl {
	return c.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象)
func (c ChainedSelectControl) StaticInputClassName(value string) ChainedSelectControl {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象)
func (c ChainedSelectControl) StaticLabelClassName(value string) ChainedSelectControl {
	return c.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`)
func (c ChainedSelectControl) StaticOn(value string) ChainedSelectControl {
	return c.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (c ChainedSelectControl) StaticPlaceholder(value string) ChainedSelectControl {
	return c.set("staticPlaceholder", value)
}

// StaticSchema
func (c ChainedSelectControl) StaticSchema(value string) ChainedSelectControl {
	return c.set("staticSchema", value)
}

// Style 组件样式
func (c ChainedSelectControl) Style(value string) ChainedSelectControl {
	return c.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (c ChainedSelectControl) SubmitOnChange(value bool) ChainedSelectControl {
	return c.set("submitOnChange", value)
}

// TestIdBuilder
func (c ChainedSelectControl) TestIdBuilder(value string) ChainedSelectControl {
	return c.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (c ChainedSelectControl) UseMobileUI(value bool) ChainedSelectControl {
	return c.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (c ChainedSelectControl) ValidateApi(value string) ChainedSelectControl {
	return c.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (c ChainedSelectControl) ValidateOnChange(value bool) ChainedSelectControl {
	return c.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (c ChainedSelectControl) ValidationErrors(value string) ChainedSelectControl {
	return c.set("validationErrors", value)
}

// Validations
func (c ChainedSelectControl) Validations(value string) ChainedSelectControl {
	return c.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量
func (c ChainedSelectControl) Value(value string) ChainedSelectControl {
	return c.set("value", value)
}

// ValuesNoWrap 多选模式，值太多时是否避免折行
func (c ChainedSelectControl) ValuesNoWrap(value bool) ChainedSelectControl {
	return c.set("valuesNoWrap", value)
}

// Visible 是否显示
func (c ChainedSelectControl) Visible(value bool) ChainedSelectControl {
	return c.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`)
func (c ChainedSelectControl) VisibleOn(value string) ChainedSelectControl {
	return c.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (c ChainedSelectControl) Width(value string) ChainedSelectControl {
	return c.set("width", value)
}
