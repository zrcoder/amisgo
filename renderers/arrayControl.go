package renderers

// InputArray 数组输入框。 combo 的别名。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-array
type InputArray struct {
	*BaseRenderer
}

// NewArrayControl 创建一个新的 ArrayControl 实例
func NewArrayControl() *InputArray {
	ac := &InputArray{BaseRenderer: NewBaseRenderer()}
	ac.set("type", "input-array")
	return ac
}

func (ac *InputArray) set(key string, value interface{}) *InputArray {
	ac.BaseRenderer.set(key, value)
	return ac
}

// AddButtonClassName 设置新增按钮CSS类名
func (ac *InputArray) AddButtonClassName(value string) *InputArray {
	return ac.set("addButtonClassName", value)
}

// AddButtonText 设置新增按钮文字
func (ac *InputArray) AddButtonText(value string) *InputArray {
	return ac.set("addButtonText", value)
}

// Addable 设置是否可新增
func (ac *InputArray) Addable(value bool) *InputArray {
	return ac.set("addable", value)
}

// Addattop 设置新增按钮是否添加在顶部
func (ac *InputArray) Addattop(value bool) *InputArray {
	return ac.set("addattop", value)
}

// AutoFill 设置自动填充
func (ac *InputArray) AutoFill(value string) *InputArray {
	return ac.set("autoFill", value)
}

// CanAccessSuperData 设置是否可以访问父级数据
func (ac *InputArray) CanAccessSuperData(value bool) *InputArray {
	return ac.set("canAccessSuperData", value)
}

// ClassName 设置容器CSS类名
func (ac *InputArray) ClassName(value string) *InputArray {
	return ac.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除值
func (ac *InputArray) ClearValueOnHidden(value bool) *InputArray {
	return ac.set("clearValueOnHidden", value)
}

// DeleteApi 设置删除时调用的API
func (ac *InputArray) DeleteApi(value string) *InputArray {
	return ac.set("deleteApi", value)
}

// DeleteConfirmText 设置确认删除时的提示
func (ac *InputArray) DeleteConfirmText(value string) *InputArray {
	return ac.set("deleteConfirmText", value)
}

// Delimiter 设置扁平化时的分隔符
func (ac *InputArray) Delimiter(value string) *InputArray {
	return ac.set("delimiter", value)
}

// Desc 设置描述
func (ac *InputArray) Desc(value string) *InputArray {
	return ac.set("desc", value)
}

// Description 设置描述内容
func (ac *InputArray) Description(value string) *InputArray {
	return ac.set("description", value)
}

// DescriptionClassName 设置描述上的className
func (ac *InputArray) DescriptionClassName(value string) *InputArray {
	return ac.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (ac *InputArray) Disabled(value bool) *InputArray {
	return ac.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (ac *InputArray) DisabledOn(value string) *InputArray {
	return ac.set("disabledOn", value)
}

// Draggable 设置是否可拖拽排序
func (ac *InputArray) Draggable(value bool) *InputArray {
	return ac.set("draggable", value)
}

// DraggableTip 设置拖拽提示信息
func (ac *InputArray) DraggableTip(value string) *InputArray {
	return ac.set("draggableTip", value)
}

// EditorSetting 设置编辑器配置
func (ac *InputArray) EditorSetting(value string) *InputArray {
	return ac.set("editorSetting", value)
}

// ExtraName 设置额外字段名
func (ac *InputArray) ExtraName(value string) *InputArray {
	return ac.set("extraName", value)
}

// Flat 设置是否将结果扁平化
func (ac *InputArray) Flat(value bool) *InputArray {
	return ac.set("flat", value)
}

// FormClassName 设置内部单组表单项的类名
func (ac *InputArray) FormClassName(value string) *InputArray {
	return ac.set("formClassName", value)
}

// Hidden 设置是否隐藏
func (ac *InputArray) Hidden(value bool) *InputArray {
	return ac.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (ac *InputArray) HiddenOn(value string) *InputArray {
	return ac.set("hiddenOn", value)
}

// Hint 设置输入提示
func (ac *InputArray) Hint(value string) *InputArray {
	return ac.set("hint", value)
}

// Horizontal 设置水平布局
func (ac *InputArray) Horizontal(value string) *InputArray {
	return ac.set("horizontal", value)
}

// ID 设置组件唯一ID
func (ac *InputArray) ID(value string) *InputArray {
	return ac.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (ac *InputArray) InitAutoFill(value string) *InputArray {
	return ac.set("initAutoFill", value)
}

// Inline 设置是否为inline模式
func (ac *InputArray) Inline(value bool) *InputArray {
	return ac.set("inline", value)
}

// InputClassName 设置input className
func (ac *InputArray) InputClassName(value string) *InputArray {
	return ac.set("inputClassName", value)
}

// Items 设置成员渲染器配置
func (ac *InputArray) Items(value string) *InputArray {
	return ac.set("items", value)
}

// JoinValues 设置是否用分隔符发送给后端
func (ac *InputArray) JoinValues(value bool) *InputArray {
	return ac.set("joinValues", value)
}

// Label 设置描述标题
func (ac *InputArray) Label(value string) *InputArray {
	return ac.set("label", value)
}

// LabelAlign 设置描述标题对齐
func (ac *InputArray) LabelAlign(value string) *InputArray {
	return ac.set("labelAlign", value)
}

// LabelClassName 设置label className
func (ac *InputArray) LabelClassName(value string) *InputArray {
	return ac.set("labelClassName", value)
}

// LabelRemark 设置label备注
func (ac *InputArray) LabelRemark(value string) *InputArray {
	return ac.set("labelRemark", value)
}

// LabelWidth 设置label宽度
func (ac *InputArray) LabelWidth(value string) *InputArray {
	return ac.set("labelWidth", value)
}

// LazyLoad 设置是否开启懒加载
func (ac *InputArray) LazyLoad(value bool) *InputArray {
	return ac.set("lazyLoad", value)
}

// MaxLength 设置限制最大个数
func (ac *InputArray) MaxLength(value string) *InputArray {
	return ac.set("maxLength", value)
}

// Messages 设置提示信息
func (ac *InputArray) Messages(value string) *InputArray {
	return ac.set("messages", value)
}

// MinLength 设置限制最小个数
func (ac *InputArray) MinLength(value string) *InputArray {
	return ac.set("minLength", value)
}

// Mode 设置展示模式
func (ac *InputArray) Mode(value string) *InputArray {
	return ac.set("mode", value)
}

// MultiLine 设置是否多行模式
func (ac *InputArray) MultiLine(value bool) *InputArray {
	return ac.set("multiLine", value)
}

// Multiple 设置是否可多选
func (ac *InputArray) Multiple(value bool) *InputArray {
	return ac.set("multiple", value)
}

// Name 设置字段名
func (ac *InputArray) Name(value string) *InputArray {
	return ac.set("name", value)
}

// NoBorder 设置是否含有边框
func (ac *InputArray) NoBorder(value bool) *InputArray {
	return ac.set("noBorder", value)
}

// Nullable 设置是否允许为空
func (ac *InputArray) Nullable(value bool) *InputArray {
	return ac.set("nullable", value)
}

// OnEvent 设置事件动作配置
func (ac *InputArray) OnEvent(value string) *InputArray {
	return ac.set("onEvent", value)
}

// Placeholder 设置没有成员时的显示内容
func (ac *InputArray) Placeholder(value string) *InputArray {
	return ac.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (ac *InputArray) ReadOnly(value bool) *InputArray {
	return ac.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (ac *InputArray) ReadOnlyOn(value string) *InputArray {
	return ac.set("readOnlyOn", value)
}

// Remark 设置备注
func (ac *InputArray) Remark(value string) *InputArray {
	return ac.set("remark", value)
}

// Removable 设置是否可删除
func (ac *InputArray) Removable(value bool) *InputArray {
	return ac.set("removable", value)
}

// Required 设置是否为必填
func (ac *InputArray) Required(value bool) *InputArray {
	return ac.set("required", value)
}

// Row 设置行数
func (ac *InputArray) Row(value string) *InputArray {
	return ac.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (ac *InputArray) SaveImmediately(value bool) *InputArray {
	return ac.set("saveImmediately", value)
}

// Scaffold 设置新增成员时的默认值
func (ac *InputArray) Scaffold(value string) *InputArray {
	return ac.set("scaffold", value)
}

// Size 设置表单项大小
func (ac *InputArray) Size(value string) *InputArray {
	return ac.set("size", value)
}

// Static 设置是否静态展示
func (ac *InputArray) Static(value bool) *InputArray {
	return ac.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (ac *InputArray) StaticClassName(value string) *InputArray {
	return ac.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (ac *InputArray) StaticInputClassName(value string) *InputArray {
	return ac.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (ac *InputArray) StaticLabelClassName(value string) *InputArray {
	return ac.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (ac *InputArray) StaticOn(value string) *InputArray {
	return ac.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (ac *InputArray) StaticPlaceholder(value string) *InputArray {
	return ac.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (ac *InputArray) StaticSchema(value string) *InputArray {
	return ac.set("staticSchema", value)
}

// StrictMode 设置严格模式
func (ac *InputArray) StrictMode(value bool) *InputArray {
	return ac.set("strictMode", value)
}

// Style 设置组件样式
func (ac *InputArray) Style(value string) *InputArray {
	return ac.set("style", value)
}

// SubFormHorizontal 设置子表单水平排版宽度占比
func (ac *InputArray) SubFormHorizontal(value string) *InputArray {
	return ac.set("subFormHorizontal", value)
}

// SubFormMode 设置子表单模式
func (ac *InputArray) SubFormMode(value string) *InputArray {
	return ac.set("subFormMode", value)
}

// SubmitOnChange 设置是否在修改后提交表单
func (ac *InputArray) SubmitOnChange(value bool) *InputArray {
	return ac.set("submitOnChange", value)
}

// SyncFields 配置同步字段
func (ac *InputArray) SyncFields(value string) *InputArray {
	return ac.set("syncFields", value)
}

// TabsLabelTpl 设置选项卡标题的生成模板
func (ac *InputArray) TabsLabelTpl(value string) *InputArray {
	return ac.set("tabsLabelTpl", value)
}

// TabsMode 设置是否使用 Tabs 展示方式
func (ac *InputArray) TabsMode(value bool) *InputArray {
	return ac.set("tabsMode", value)
}

// TabsStyle 设置 Tabs 的展示模式
func (ac *InputArray) TabsStyle(value string) *InputArray {
	return ac.set("tabsStyle", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (ac *InputArray) TestIdBuilder(value string) *InputArray {
	return ac.set("testIdBuilder", value)
}

// TypeSwitchable 设置是否可切换条件
func (ac *InputArray) TypeSwitchable(value bool) *InputArray {
	return ac.set("typeSwitchable", value)
}

// UpdatePristineAfterStoreDataReInit 设置是否在数据重新初始化后更新原始值
func (ac *InputArray) UpdatePristineAfterStoreDataReInit(value bool) *InputArray {
	return ac.set("updatePristineAfterStoreDataReInit", value)
}

// UseMobileUI 设置是否使用移动端样式
func (ac *InputArray) UseMobileUI(value bool) *InputArray {
	return ac.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (ac *InputArray) ValidateApi(value string) *InputArray {
	return ac.set("validateApi", value)
}

// ValidateOnChange 设置是否在每次修改时触发验证
func (ac *InputArray) ValidateOnChange(value bool) *InputArray {
	return ac.set("validateOnChange", value)
}

// ValidationErrors 设置验证失败的提示信息
func (ac *InputArray) ValidationErrors(value string) *InputArray {
	return ac.set("validationErrors", value)
}

// Validations 设置验证规则
func (ac *InputArray) Validations(value string) *InputArray {
	return ac.set("validations", value)
}

// Value 设置默认值
func (ac *InputArray) Value(value string) *InputArray {
	return ac.set("value", value)
}

// Visible 设置是否显示
func (ac *InputArray) Visible(value bool) *InputArray {
	return ac.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (ac *InputArray) VisibleOn(value string) *InputArray {
	return ac.set("visibleOn", value)
}

// Width 设置在 Table 中的宽度
func (ac *InputArray) Width(value string) *InputArray {
	return ac.set("width", value)
}
