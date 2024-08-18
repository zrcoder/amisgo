package comp

// PickerControl 选择器控件
//
// @version 6.7.0
type PickerControl Schema

// NewPickerControl 创建一个新的 PickerControl 实例
func NewPickerControl() PickerControl {
	return PickerControl{}.set("type", "picker")
}

func (pc PickerControl) set(key string, value interface{}) PickerControl {
	pc[key] = value
	return pc
}

// AddApi 添加时调用的接口
func (pc PickerControl) AddApi(value string) PickerControl {
	return pc.set("addApi", value)
}

// AddControls 新增时的表单项
func (pc PickerControl) AddControls(value string) PickerControl {
	return pc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (pc PickerControl) AddDialog(value string) PickerControl {
	return pc.set("addDialog", value)
}

// AutoFill 自动填充
func (pc PickerControl) AutoFill(value string) PickerControl {
	return pc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (pc PickerControl) ClassName(value string) PickerControl {
	return pc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否在当前 Form 中删除掉该表单项值
func (pc PickerControl) ClearValueOnHidden(value bool) PickerControl {
	return pc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (pc PickerControl) Clearable(value bool) PickerControl {
	return pc.set("clearable", value)
}

// Creatable 是否可以新增
func (pc PickerControl) Creatable(value bool) PickerControl {
	return pc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (pc PickerControl) CreateBtnLabel(value string) PickerControl {
	return pc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (pc PickerControl) DeferApi(value string) PickerControl {
	return pc.set("deferApi", value)
}

// DeferField 懒加载字段
func (pc PickerControl) DeferField(value string) PickerControl {
	return pc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (pc PickerControl) DeleteApi(value string) PickerControl {
	return pc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (pc PickerControl) DeleteConfirmText(value string) PickerControl {
	return pc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (pc PickerControl) Delimiter(value string) PickerControl {
	return pc.set("delimiter", value)
}

// Desc 描述
func (pc PickerControl) Desc(value string) PickerControl {
	return pc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (pc PickerControl) Description(value string) PickerControl {
	return pc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (pc PickerControl) DescriptionClassName(value string) PickerControl {
	return pc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (pc PickerControl) Disabled(value bool) PickerControl {
	return pc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (pc PickerControl) DisabledOn(value string) PickerControl {
	return pc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (pc PickerControl) EditApi(value string) PickerControl {
	return pc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (pc PickerControl) EditControls(value string) PickerControl {
	return pc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (pc PickerControl) EditDialog(value string) PickerControl {
	return pc.set("editDialog", value)
}

// Editable 是否可以编辑
func (pc PickerControl) Editable(value bool) PickerControl {
	return pc.set("editable", value)
}

// EditorSetting 编辑器配置
func (pc PickerControl) EditorSetting(value string) PickerControl {
	return pc.set("editorSetting", value)
}

// Embed 内嵌模式
func (pc PickerControl) Embed(value bool) PickerControl {
	return pc.set("embed", value)
}

// ExtraName 额外的字段名
func (pc PickerControl) ExtraName(value string) PickerControl {
	return pc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (pc PickerControl) ExtractValue(value bool) PickerControl {
	return pc.set("extractValue", value)
}

// Hidden 是否隐藏
func (pc PickerControl) Hidden(value bool) PickerControl {
	return pc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (pc PickerControl) HiddenOn(value string) PickerControl {
	return pc.set("hiddenOn", value)
}

// Hint 输入提示
func (pc PickerControl) Hint(value string) PickerControl {
	return pc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (pc PickerControl) Horizontal(value string) PickerControl {
	return pc.set("horizontal", value)
}

// Id 组件唯一 id
func (pc PickerControl) Id(value string) PickerControl {
	return pc.set("id", value)
}

// InitAutoFill 配置 initAutoFill
func (pc PickerControl) InitAutoFill(value string) PickerControl {
	return pc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (pc PickerControl) InitFetch(value bool) PickerControl {
	return pc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (pc PickerControl) InitFetchOn(value string) PickerControl {
	return pc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (pc PickerControl) Inline(value bool) PickerControl {
	return pc.set("inline", value)
}

// InputClassName 配置 input className
func (pc PickerControl) InputClassName(value string) PickerControl {
	return pc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时
func (pc PickerControl) JoinValues(value bool) PickerControl {
	return pc.set("joinValues", value)
}

// Label 描述标题
func (pc PickerControl) Label(value string) PickerControl {
	return pc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (pc PickerControl) LabelAlign(value string) PickerControl {
	return pc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (pc PickerControl) LabelClassName(value string) PickerControl {
	return pc.set("labelClassName", value)
}

// LabelField 建议用 labelTpl 选中一个字段名用来显示
func (pc PickerControl) LabelField(value string) PickerControl {
	return pc.set("labelField", value)
}

// LabelRemark 描述标题备注
func (pc PickerControl) LabelRemark(value string) PickerControl {
	return pc.set("labelRemark", value)
}

// LabelTpl 标签渲染模版
func (pc PickerControl) LabelTpl(value string) PickerControl {
	return pc.set("labelTpl", value)
}

// LabelWidth 描述宽度
func (pc PickerControl) LabelWidth(value int) PickerControl {
	return pc.set("labelWidth", value)
}

// ModalMode 弹框模式
func (pc PickerControl) ModalMode(value string) PickerControl {
	return pc.set("modalMode", value)
}

// ModalTitle 弹框标题
func (pc PickerControl) ModalTitle(value string) PickerControl {
	return pc.set("modalTitle", value)
}

// Mode 组件模式
func (pc PickerControl) Mode(value string) PickerControl {
	return pc.set("mode", value)
}

// Multiple 是否支持多选
func (pc PickerControl) Multiple(value bool) PickerControl {
	return pc.set("multiple", value)
}

// Name 组件名
func (pc PickerControl) Name(value string) PickerControl {
	return pc.set("name", value)
}

// OnEvent 事件
func (pc PickerControl) OnEvent(value string) PickerControl {
	return pc.set("onEvent", value)
}

// Options 选项列表
func (pc PickerControl) Options(value string) PickerControl {
	return pc.set("options", value)
}

// OverflowConfig 控制内容超出后的显示配置
func (pc PickerControl) OverflowConfig(value string) PickerControl {
	return pc.set("overflowConfig", value)
}

// PickerSchema picker 的 schema
func (pc PickerControl) PickerSchema(value string) PickerControl {
	return pc.set("pickerSchema", value)
}

// Placeholder 占位符
func (pc PickerControl) Placeholder(value string) PickerControl {
	return pc.set("placeholder", value)
}

// ReadOnly 是否只读
func (pc PickerControl) ReadOnly(value bool) PickerControl {
	return pc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (pc PickerControl) ReadOnlyOn(value string) PickerControl {
	return pc.set("readOnlyOn", value)
}

// Removable 是否可删除
func (pc PickerControl) Removable(value bool) PickerControl {
	return pc.set("removable", value)
}

// Required 是否必填
func (pc PickerControl) Required(value bool) PickerControl {
	return pc.set("required", value)
}

// ResetValue 重置表单值
func (pc PickerControl) ResetValue(value string) PickerControl {
	return pc.set("resetValue", value)
}

// Row 控制表单项显示的行
func (pc PickerControl) Row(value int) PickerControl {
	return pc.set("row", value)
}

// SaveImmediately 保存表单项
func (pc PickerControl) SaveImmediately(value bool) PickerControl {
	return pc.set("saveImmediately", value)
}

// SelectFirst 是否选中第一个
func (pc PickerControl) SelectFirst(value bool) PickerControl {
	return pc.set("selectFirst", value)
}

// Size 大小
func (pc PickerControl) Size(value string) PickerControl {
	return pc.set("size", value)
}

// Source 数据来源
func (pc PickerControl) Source(value string) PickerControl {
	return pc.set("source", value)
}

// Static 静态模式
func (pc PickerControl) Static(value bool) PickerControl {
	return pc.set("static", value)
}

// StaticClassName 静态模式下的 css className
func (pc PickerControl) StaticClassName(value string) PickerControl {
	return pc.set("staticClassName", value)
}

// StaticInputClassName 静态模式下 input 的 css className
func (pc PickerControl) StaticInputClassName(value string) PickerControl {
	return pc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态模式下 label 的 css className
func (pc PickerControl) StaticLabelClassName(value string) PickerControl {
	return pc.set("staticLabelClassName", value)
}

// StaticOn 是否开启静态模式
func (pc PickerControl) StaticOn(value string) PickerControl {
	return pc.set("staticOn", value)
}

// StaticPlaceholder 静态模式下占位符
func (pc PickerControl) StaticPlaceholder(value string) PickerControl {
	return pc.set("staticPlaceholder", value)
}

// StaticSchema 静态模式下的 schema
func (pc PickerControl) StaticSchema(value string) PickerControl {
	return pc.set("staticSchema", value)
}

// Style 自定义样式
func (pc PickerControl) Style(value string) PickerControl {
	return pc.set("style", value)
}

// SubmitOnChange 是否在改变时提交
func (pc PickerControl) SubmitOnChange(value bool) PickerControl {
	return pc.set("submitOnChange", value)
}

// TestIdBuilder 生成测试 ID 的函数
func (pc PickerControl) TestIdBuilder(value string) PickerControl {
	return pc.set("testIdBuilder", value)
}

// UseMobileUI 是否使用移动端 UI
func (pc PickerControl) UseMobileUI(value bool) PickerControl {
	return pc.set("useMobileUI", value)
}

// ValidateApi 校验 API
func (pc PickerControl) ValidateApi(value string) PickerControl {
	return pc.set("validateApi", value)
}

// ValidateOnChange 是否在改变时校验
func (pc PickerControl) ValidateOnChange(value bool) PickerControl {
	return pc.set("validateOnChange", value)
}

// ValidationErrors 校验错误信息
func (pc PickerControl) ValidationErrors(value string) PickerControl {
	return pc.set("validationErrors", value)
}

// Validations 校验规则
func (pc PickerControl) Validations(value string) PickerControl {
	return pc.set("validations", value)
}

// Value 值
func (pc PickerControl) Value(value string) PickerControl {
	return pc.set("value", value)
}

// ValueField 值字段
func (pc PickerControl) ValueField(value string) PickerControl {
	return pc.set("valueField", value)
}

// ValuesNoWrap 是否不换行
func (pc PickerControl) ValuesNoWrap(value bool) PickerControl {
	return pc.set("valuesNoWrap", value)
}

// Visible 是否可见
func (pc PickerControl) Visible(value bool) PickerControl {
	return pc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (pc PickerControl) VisibleOn(value string) PickerControl {
	return pc.set("visibleOn", value)
}

// Width 宽度
func (pc PickerControl) Width(value int) PickerControl {
	return pc.set("width", value)
}
