package comp

// picker 选择器控件

type picker schema

// Picker 创建一个新的 PickerControl 实例
func Picker() picker {
	return picker{}.set("type", "picker")
}

func (pc picker) set(key string, value any) picker {
	pc[key] = value
	return pc
}

// AddApi 添加时调用的接口
func (pc picker) AddApi(value string) picker {
	return pc.set("addApi", value)
}

// AddControls 新增时的表单项
func (pc picker) AddControls(value string) picker {
	return pc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (pc picker) AddDialog(value string) picker {
	return pc.set("addDialog", value)
}

// AutoFill 自动填充
func (pc picker) AutoFill(value string) picker {
	return pc.set("autoFill", value)
}

// ClassName 容器 css 类名
func (pc picker) ClassName(value string) picker {
	return pc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否在当前 Form 中删除掉该表单项值
func (pc picker) ClearValueOnHidden(value bool) picker {
	return pc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (pc picker) Clearable(value bool) picker {
	return pc.set("clearable", value)
}

// Creatable 是否可以新增
func (pc picker) Creatable(value bool) picker {
	return pc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (pc picker) CreateBtnLabel(value string) picker {
	return pc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (pc picker) DeferApi(value string) picker {
	return pc.set("deferApi", value)
}

// DeferField 懒加载字段
func (pc picker) DeferField(value string) picker {
	return pc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (pc picker) DeleteApi(value string) picker {
	return pc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (pc picker) DeleteConfirmText(value string) picker {
	return pc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (pc picker) Delimiter(value string) picker {
	return pc.set("delimiter", value)
}

// Desc 描述
func (pc picker) Desc(value string) picker {
	return pc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (pc picker) Description(value string) picker {
	return pc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (pc picker) DescriptionClassName(value string) picker {
	return pc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (pc picker) Disabled(value bool) picker {
	return pc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (pc picker) DisabledOn(value string) picker {
	return pc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (pc picker) EditApi(value string) picker {
	return pc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (pc picker) EditControls(value string) picker {
	return pc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (pc picker) EditDialog(value string) picker {
	return pc.set("editDialog", value)
}

// Editable 是否可以编辑
func (pc picker) Editable(value bool) picker {
	return pc.set("editable", value)
}

// EditorSetting 编辑器配置
func (pc picker) EditorSetting(value string) picker {
	return pc.set("editorSetting", value)
}

// Embed 内嵌模式
func (pc picker) Embed(value bool) picker {
	return pc.set("embed", value)
}

// ExtraName 额外的字段名
func (pc picker) ExtraName(value string) picker {
	return pc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (pc picker) ExtractValue(value bool) picker {
	return pc.set("extractValue", value)
}

// Hidden 是否隐藏
func (pc picker) Hidden(value bool) picker {
	return pc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (pc picker) HiddenOn(value string) picker {
	return pc.set("hiddenOn", value)
}

// Hint 输入提示
func (pc picker) Hint(value string) picker {
	return pc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (pc picker) Horizontal(value string) picker {
	return pc.set("horizontal", value)
}

// Id 组件唯一 id
func (pc picker) ID(value string) picker {
	return pc.set("id", value)
}

// InitAutoFill 配置 initAutoFill
func (pc picker) InitAutoFill(value string) picker {
	return pc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (pc picker) InitFetch(value bool) picker {
	return pc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (pc picker) InitFetchOn(value string) picker {
	return pc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (pc picker) Inline(value bool) picker {
	return pc.set("inline", value)
}

// InputClassName 配置 input className
func (pc picker) InputClassName(value string) picker {
	return pc.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时
func (pc picker) JoinValues(value bool) picker {
	return pc.set("joinValues", value)
}

// Label 描述标题
func (pc picker) Label(value string) picker {
	return pc.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (pc picker) LabelAlign(value string) picker {
	return pc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (pc picker) LabelClassName(value string) picker {
	return pc.set("labelClassName", value)
}

// LabelField 建议用 labelTpl 选中一个字段名用来显示
func (pc picker) LabelField(value string) picker {
	return pc.set("labelField", value)
}

// LabelRemark 描述标题备注
func (pc picker) LabelRemark(value string) picker {
	return pc.set("labelRemark", value)
}

// LabelTpl 标签渲染模版
func (pc picker) LabelTpl(value string) picker {
	return pc.set("labelTpl", value)
}

// LabelWidth 描述宽度
func (pc picker) LabelWidth(value int) picker {
	return pc.set("labelWidth", value)
}

// ModalMode 弹框模式
func (pc picker) ModalMode(value string) picker {
	return pc.set("modalMode", value)
}

// ModalTitle 弹框标题
func (pc picker) ModalTitle(value any) picker {
	return pc.set("modalTitle", value)
}

// Mode 组件模式
func (pc picker) Mode(value string) picker {
	return pc.set("mode", value)
}

// Multiple 是否支持多选
func (pc picker) Multiple(value bool) picker {
	return pc.set("multiple", value)
}

// Name 组件名
func (pc picker) Name(value string) picker {
	return pc.set("name", value)
}

// OnEvent 事件
func (pc picker) OnEvent(value any) picker {
	return pc.set("onEvent", value)
}

// Options 选项列表
func (pc picker) Options(value ...option) picker {
	return pc.set("options", value)
}

// OverflowConfig 控制内容超出后的显示配置
func (pc picker) OverflowConfig(value string) picker {
	return pc.set("overflowConfig", value)
}

// PickerSchema picker 的 schema
func (pc picker) PickerSchema(value string) picker {
	return pc.set("pickerSchema", value)
}

// Placeholder 占位符
func (pc picker) Placeholder(value string) picker {
	return pc.set("placeholder", value)
}

// ReadOnly 是否只读
func (pc picker) ReadOnly(value bool) picker {
	return pc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (pc picker) ReadOnlyOn(value string) picker {
	return pc.set("readOnlyOn", value)
}

// Removable 是否可删除
func (pc picker) Removable(value bool) picker {
	return pc.set("removable", value)
}

// Required 是否必填
func (pc picker) Required(value bool) picker {
	return pc.set("required", value)
}

// ResetValue 重置表单值
func (pc picker) ResetValue(value string) picker {
	return pc.set("resetValue", value)
}

// Row 控制表单项显示的行
func (pc picker) Row(value int) picker {
	return pc.set("row", value)
}

// SaveImmediately 保存表单项
func (pc picker) SaveImmediately(value bool) picker {
	return pc.set("saveImmediately", value)
}

// SelectFirst 是否选中第一个
func (pc picker) SelectFirst(value bool) picker {
	return pc.set("selectFirst", value)
}

// Size 大小
func (pc picker) Size(value string) picker {
	return pc.set("size", value)
}

// Source 数据来源
func (pc picker) Source(value string) picker {
	return pc.set("source", value)
}

// Static 静态模式
func (pc picker) Static(value bool) picker {
	return pc.set("static", value)
}

// StaticClassName 静态模式下的 css className
func (pc picker) StaticClassName(value string) picker {
	return pc.set("staticClassName", value)
}

// StaticInputClassName 静态模式下 input 的 css className
func (pc picker) StaticInputClassName(value string) picker {
	return pc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态模式下 label 的 css className
func (pc picker) StaticLabelClassName(value string) picker {
	return pc.set("staticLabelClassName", value)
}

// StaticOn 是否开启静态模式
func (pc picker) StaticOn(value string) picker {
	return pc.set("staticOn", value)
}

// StaticPlaceholder 静态模式下占位符
func (pc picker) StaticPlaceholder(value string) picker {
	return pc.set("staticPlaceholder", value)
}

// StaticSchema 静态模式下的 schema
func (pc picker) StaticSchema(value string) picker {
	return pc.set("staticSchema", value)
}

// Style 自定义样式
func (pc picker) Style(value any) picker {
	return pc.set("style", value)
}

// SubmitOnChange 是否在改变时提交
func (pc picker) SubmitOnChange(value bool) picker {
	return pc.set("submitOnChange", value)
}

// TestIdBuilder 生成测试 ID 的函数
func (pc picker) TestIdBuilder(value string) picker {
	return pc.set("testIdBuilder", value)
}

// UseMobileUI 是否使用移动端 UI
func (pc picker) UseMobileUI(value bool) picker {
	return pc.set("useMobileUI", value)
}

// ValidateApi 校验 API
func (pc picker) ValidateApi(value string) picker {
	return pc.set("validateApi", value)
}

// ValidateOnChange 是否在改变时校验
func (pc picker) ValidateOnChange(value bool) picker {
	return pc.set("validateOnChange", value)
}

// ValidationErrors 校验错误信息
func (pc picker) ValidationErrors(value string) picker {
	return pc.set("validationErrors", value)
}

// Validations 校验规则
func (pc picker) Validations(value string) picker {
	return pc.set("validations", value)
}

// Value 值
func (pc picker) Value(value string) picker {
	return pc.set("value", value)
}

// ValueField 值字段
func (pc picker) ValueField(value string) picker {
	return pc.set("valueField", value)
}

// ValuesNoWrap 是否不换行
func (pc picker) ValuesNoWrap(value bool) picker {
	return pc.set("valuesNoWrap", value)
}

// Visible 是否可见
func (pc picker) Visible(value bool) picker {
	return pc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (pc picker) VisibleOn(value string) picker {
	return pc.set("visibleOn", value)
}

// Width 宽度
func (pc picker) Width(value int) picker {
	return pc.set("width", value)
}
