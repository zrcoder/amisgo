package comp

// inputArray 数组输入框。 combo 的别名。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-array
type inputArray schema

// InputArray 创建一个新的 InputArray 实例
func InputArray() inputArray {
	return make(inputArray).set("type", "input-array")
}

func (ac inputArray) set(key string, value interface{}) inputArray {
	ac[key] = value
	return ac
}

// AddButtonClassName 设置新增按钮CSS类名
func (ac inputArray) AddButtonClassName(value string) inputArray {
	return ac.set("addButtonClassName", value)
}

// AddButtonText 设置新增按钮文字
func (ac inputArray) AddButtonText(value string) inputArray {
	return ac.set("addButtonText", value)
}

// Addable 设置是否可新增
func (ac inputArray) Addable(value bool) inputArray {
	return ac.set("addable", value)
}

// Addattop 设置新增按钮是否添加在顶部
func (ac inputArray) Addattop(value bool) inputArray {
	return ac.set("addattop", value)
}

// AutoFill 设置自动填充
func (ac inputArray) AutoFill(value string) inputArray {
	return ac.set("autoFill", value)
}

// CanAccessSuperData 设置是否可以访问父级数据
func (ac inputArray) CanAccessSuperData(value bool) inputArray {
	return ac.set("canAccessSuperData", value)
}

// ClassName 设置容器CSS类名
func (ac inputArray) ClassName(value string) inputArray {
	return ac.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除值
func (ac inputArray) ClearValueOnHidden(value bool) inputArray {
	return ac.set("clearValueOnHidden", value)
}

// DeleteApi 设置删除时调用的API
func (ac inputArray) DeleteApi(value string) inputArray {
	return ac.set("deleteApi", value)
}

// DeleteConfirmText 设置确认删除时的提示
func (ac inputArray) DeleteConfirmText(value string) inputArray {
	return ac.set("deleteConfirmText", value)
}

// Delimiter 设置扁平化时的分隔符
func (ac inputArray) Delimiter(value string) inputArray {
	return ac.set("delimiter", value)
}

// Desc 设置描述
func (ac inputArray) Desc(value string) inputArray {
	return ac.set("desc", value)
}

// Description 设置描述内容
func (ac inputArray) Description(value string) inputArray {
	return ac.set("description", value)
}

// DescriptionClassName 设置描述上的className
func (ac inputArray) DescriptionClassName(value string) inputArray {
	return ac.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (ac inputArray) Disabled(value bool) inputArray {
	return ac.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (ac inputArray) DisabledOn(value string) inputArray {
	return ac.set("disabledOn", value)
}

// Draggable 设置是否可拖拽排序
func (ac inputArray) Draggable(value bool) inputArray {
	return ac.set("draggable", value)
}

// DraggableTip 设置拖拽提示信息
func (ac inputArray) DraggableTip(value string) inputArray {
	return ac.set("draggableTip", value)
}

// EditorSetting 设置编辑器配置
func (ac inputArray) EditorSetting(value string) inputArray {
	return ac.set("editorSetting", value)
}

// ExtraName 设置额外字段名
func (ac inputArray) ExtraName(value string) inputArray {
	return ac.set("extraName", value)
}

// Flat 设置是否将结果扁平化
func (ac inputArray) Flat(value bool) inputArray {
	return ac.set("flat", value)
}

// FormClassName 设置内部单组表单项的类名
func (ac inputArray) FormClassName(value string) inputArray {
	return ac.set("formClassName", value)
}

// Hidden 设置是否隐藏
func (ac inputArray) Hidden(value bool) inputArray {
	return ac.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (ac inputArray) HiddenOn(value string) inputArray {
	return ac.set("hiddenOn", value)
}

// Hint 设置输入提示
func (ac inputArray) Hint(value string) inputArray {
	return ac.set("hint", value)
}

// Horizontal 设置水平布局
func (ac inputArray) Horizontal(value string) inputArray {
	return ac.set("horizontal", value)
}

// ID 设置组件唯一ID
func (ac inputArray) ID(value string) inputArray {
	return ac.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (ac inputArray) InitAutoFill(value string) inputArray {
	return ac.set("initAutoFill", value)
}

// Inline 设置是否为inline模式
func (ac inputArray) Inline(value bool) inputArray {
	return ac.set("inline", value)
}

// InputClassName 设置input className
func (ac inputArray) InputClassName(value string) inputArray {
	return ac.set("inputClassName", value)
}

// Items 设置成员渲染器配置
func (ac inputArray) Items(value string) inputArray {
	return ac.set("items", value)
}

// JoinValues 设置是否用分隔符发送给后端
func (ac inputArray) JoinValues(value bool) inputArray {
	return ac.set("joinValues", value)
}

// Label 设置描述标题
func (ac inputArray) Label(value string) inputArray {
	return ac.set("label", value)
}

// LabelAlign 设置描述标题对齐
func (ac inputArray) LabelAlign(value string) inputArray {
	return ac.set("labelAlign", value)
}

// LabelClassName 设置label className
func (ac inputArray) LabelClassName(value string) inputArray {
	return ac.set("labelClassName", value)
}

// LabelRemark 设置label备注
func (ac inputArray) LabelRemark(value string) inputArray {
	return ac.set("labelRemark", value)
}

// LabelWidth 设置label宽度
func (ac inputArray) LabelWidth(value string) inputArray {
	return ac.set("labelWidth", value)
}

// LazyLoad 设置是否开启懒加载
func (ac inputArray) LazyLoad(value bool) inputArray {
	return ac.set("lazyLoad", value)
}

// MaxLength 设置限制最大个数
func (ac inputArray) MaxLength(value string) inputArray {
	return ac.set("maxLength", value)
}

// Messages 设置提示信息
func (ac inputArray) Messages(value string) inputArray {
	return ac.set("messages", value)
}

// MinLength 设置限制最小个数
func (ac inputArray) MinLength(value string) inputArray {
	return ac.set("minLength", value)
}

// Mode 设置展示模式
func (ac inputArray) Mode(value string) inputArray {
	return ac.set("mode", value)
}

// MultiLine 设置是否多行模式
func (ac inputArray) MultiLine(value bool) inputArray {
	return ac.set("multiLine", value)
}

// Multiple 设置是否可多选
func (ac inputArray) Multiple(value bool) inputArray {
	return ac.set("multiple", value)
}

// Name 设置字段名
func (ac inputArray) Name(value string) inputArray {
	return ac.set("name", value)
}

// NoBorder 设置是否含有边框
func (ac inputArray) NoBorder(value bool) inputArray {
	return ac.set("noBorder", value)
}

// Nullable 设置是否允许为空
func (ac inputArray) Nullable(value bool) inputArray {
	return ac.set("nullable", value)
}

// OnEvent 设置事件动作配置
func (ac inputArray) OnEvent(value string) inputArray {
	return ac.set("onEvent", value)
}

// Placeholder 设置没有成员时的显示内容
func (ac inputArray) Placeholder(value string) inputArray {
	return ac.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (ac inputArray) ReadOnly(value bool) inputArray {
	return ac.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (ac inputArray) ReadOnlyOn(value string) inputArray {
	return ac.set("readOnlyOn", value)
}

// Remark 设置备注
func (ac inputArray) Remark(value string) inputArray {
	return ac.set("remark", value)
}

// Removable 设置是否可删除
func (ac inputArray) Removable(value bool) inputArray {
	return ac.set("removable", value)
}

// Required 设置是否为必填
func (ac inputArray) Required(value bool) inputArray {
	return ac.set("required", value)
}

// Row 设置行数
func (ac inputArray) Row(value string) inputArray {
	return ac.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (ac inputArray) SaveImmediately(value bool) inputArray {
	return ac.set("saveImmediately", value)
}

// Scaffold 设置新增成员时的默认值
func (ac inputArray) Scaffold(value string) inputArray {
	return ac.set("scaffold", value)
}

// Size 设置表单项大小
func (ac inputArray) Size(value string) inputArray {
	return ac.set("size", value)
}

// Static 设置是否静态展示
func (ac inputArray) Static(value bool) inputArray {
	return ac.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (ac inputArray) StaticClassName(value string) inputArray {
	return ac.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (ac inputArray) StaticInputClassName(value string) inputArray {
	return ac.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (ac inputArray) StaticLabelClassName(value string) inputArray {
	return ac.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (ac inputArray) StaticOn(value string) inputArray {
	return ac.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (ac inputArray) StaticPlaceholder(value string) inputArray {
	return ac.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (ac inputArray) StaticSchema(value string) inputArray {
	return ac.set("staticSchema", value)
}

// StrictMode 设置严格模式
func (ac inputArray) StrictMode(value bool) inputArray {
	return ac.set("strictMode", value)
}

// Style 设置组件样式
func (ac inputArray) Style(value string) inputArray {
	return ac.set("style", value)
}

// SubFormHorizontal 设置子表单水平排版宽度占比
func (ac inputArray) SubFormHorizontal(value string) inputArray {
	return ac.set("subFormHorizontal", value)
}

// SubFormMode 设置子表单模式
func (ac inputArray) SubFormMode(value string) inputArray {
	return ac.set("subFormMode", value)
}

// SubmitOnChange 设置是否在修改后提交表单
func (ac inputArray) SubmitOnChange(value bool) inputArray {
	return ac.set("submitOnChange", value)
}

// SyncFields 配置同步字段
func (ac inputArray) SyncFields(value string) inputArray {
	return ac.set("syncFields", value)
}

// TabsLabelTpl 设置选项卡标题的生成模板
func (ac inputArray) TabsLabelTpl(value string) inputArray {
	return ac.set("tabsLabelTpl", value)
}

// TabsMode 设置是否使用 Tabs 展示方式
func (ac inputArray) TabsMode(value bool) inputArray {
	return ac.set("tabsMode", value)
}

// TabsStyle 设置 Tabs 的展示模式
func (ac inputArray) TabsStyle(value string) inputArray {
	return ac.set("tabsStyle", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (ac inputArray) TestIdBuilder(value string) inputArray {
	return ac.set("testIdBuilder", value)
}

// TypeSwitchable 设置是否可切换条件
func (ac inputArray) TypeSwitchable(value bool) inputArray {
	return ac.set("typeSwitchable", value)
}

// UpdatePristineAfterStoreDataReInit 设置是否在数据重新初始化后更新原始值
func (ac inputArray) UpdatePristineAfterStoreDataReInit(value bool) inputArray {
	return ac.set("updatePristineAfterStoreDataReInit", value)
}

// UseMobileUI 设置是否使用移动端样式
func (ac inputArray) UseMobileUI(value bool) inputArray {
	return ac.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (ac inputArray) ValidateApi(value string) inputArray {
	return ac.set("validateApi", value)
}

// ValidateOnChange 设置是否在每次修改时触发验证
func (ac inputArray) ValidateOnChange(value bool) inputArray {
	return ac.set("validateOnChange", value)
}

// ValidationErrors 设置验证失败的提示信息
func (ac inputArray) ValidationErrors(value string) inputArray {
	return ac.set("validationErrors", value)
}

// Validations 设置验证规则
func (ac inputArray) Validations(value string) inputArray {
	return ac.set("validations", value)
}

// Value 设置默认值
func (ac inputArray) Value(value string) inputArray {
	return ac.set("value", value)
}

// Visible 设置是否显示
func (ac inputArray) Visible(value bool) inputArray {
	return ac.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (ac inputArray) VisibleOn(value string) inputArray {
	return ac.set("visibleOn", value)
}

// Width 设置在 Table 中的宽度
func (ac inputArray) Width(value string) inputArray {
	return ac.set("width", value)
}
