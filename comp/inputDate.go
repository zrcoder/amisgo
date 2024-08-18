package comp

// inputDate 日期选择控件
type inputDate schema

func InputDate() inputDate {
	return make(inputDate).set("type", "input-date")
}

func (d inputDate) set(key string, value interface{}) inputDate {
	d[key] = value
	return d
}

// AutoFill 自动填充
func (d inputDate) AutoFill(value string) inputDate {
	return d.set("autoFill", value)
}

// BorderMode 边框模式
func (d inputDate) BorderMode(value string) inputDate {
	return d.set("borderMode", value)
}

// ClassName 容器 css 类名
func (d inputDate) ClassName(value string) inputDate {
	return d.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否清除值
func (d inputDate) ClearValueOnHidden(value bool) inputDate {
	return d.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (d inputDate) Clearable(value bool) inputDate {
	return d.set("clearable", value)
}

// CloseOnSelect 点选日期后是否关闭弹窗
func (d inputDate) CloseOnSelect(value bool) inputDate {
	return d.set("closeOnSelect", value)
}

// Desc 描述内容
func (d inputDate) Desc(value string) inputDate {
	return d.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (d inputDate) Description(value string) inputDate {
	return d.set("description", value)
}

// DescriptionClassName 描述内容的 className
func (d inputDate) DescriptionClassName(value string) inputDate {
	return d.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (d inputDate) Disabled(value bool) inputDate {
	return d.set("disabled", value)
}

// DisabledDate 禁用日期函数
func (d inputDate) DisabledDate(value string) inputDate {
	return d.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式
func (d inputDate) DisabledOn(value string) inputDate {
	return d.set("disabledOn", value)
}

// DisplayFormat 日期展示格式
func (d inputDate) DisplayFormat(value string) inputDate {
	return d.set("displayFormat", value)
}

// EditorSetting 编辑器配置
func (d inputDate) EditorSetting(value string) inputDate {
	return d.set("editorSetting", value)
}

// Embed 是否为内联模式
func (d inputDate) Embed(value bool) inputDate {
	return d.set("embed", value)
}

// ExtraName 额外字段名
func (d inputDate) ExtraName(value string) inputDate {
	return d.set("extraName", value)
}

// Format 日期存储格式
func (d inputDate) Format(value string) inputDate {
	return d.set("format", value)
}

// Hidden 是否隐藏
func (d inputDate) Hidden(value bool) inputDate {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d inputDate) HiddenOn(value string) inputDate {
	return d.set("hiddenOn", value)
}

// Hint 输入提示
func (d inputDate) Hint(value string) inputDate {
	return d.set("hint", value)
}

// Horizontal 水平布局配置
func (d inputDate) Horizontal(value string) inputDate {
	return d.set("horizontal", value)
}

// ID 组件唯一 id
func (d inputDate) ID(value string) inputDate {
	return d.set("id", value)
}

// InitAutoFill 初始化自动填充
func (d inputDate) InitAutoFill(value string) inputDate {
	return d.set("initAutoFill", value)
}

// Inline 是否为 inline 模式
func (d inputDate) Inline(value bool) inputDate {
	return d.set("inline", value)
}

// InputClassName 配置 input className
func (d inputDate) InputClassName(value string) inputDate {
	return d.set("inputClassName", value)
}

// InputFormat 输入格式
func (d inputDate) InputFormat(value string) inputDate {
	return d.set("inputFormat", value)
}

// Label 描述标题
func (d inputDate) Label(value string) inputDate {
	return d.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (d inputDate) LabelAlign(value string) inputDate {
	return d.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (d inputDate) LabelClassName(value string) inputDate {
	return d.set("labelClassName", value)
}

// LabelRemark 显示小图标，鼠标放上去显示提示内容
func (d inputDate) LabelRemark(value string) inputDate {
	return d.set("labelRemark", value)
}

// LabelWidth label 宽度
func (d inputDate) LabelWidth(value string) inputDate {
	return d.set("labelWidth", value)
}

// MaxDate 限制最大日期
func (d inputDate) MaxDate(value string) inputDate {
	return d.set("maxDate", value)
}

// MinDate 限制最小日期
func (d inputDate) MinDate(value string) inputDate {
	return d.set("minDate", value)
}

// Mode 组件展示模式
func (d inputDate) Mode(value string) inputDate {
	return d.set("mode", value)
}

// Name 字段名
func (d inputDate) Name(value string) inputDate {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d inputDate) OnEvent(value string) inputDate {
	return d.set("onEvent", value)
}

// Placeholder 占位符
func (d inputDate) Placeholder(value string) inputDate {
	return d.set("placeholder", value)
}

// PopOverContainerSelector 弹窗容器选择器
func (d inputDate) PopOverContainerSelector(value string) inputDate {
	return d.set("popOverContainerSelector", value)
}

// ReadOnly 是否只读
func (d inputDate) ReadOnly(value bool) inputDate {
	return d.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (d inputDate) ReadOnlyOn(value string) inputDate {
	return d.set("readOnlyOn", value)
}

// Remark 小图标提示内容
func (d inputDate) Remark(value string) inputDate {
	return d.set("remark", value)
}

// Required 是否必填
func (d inputDate) Required(value bool) inputDate {
	return d.set("required", value)
}

// Row 行配置
func (d inputDate) Row(value string) inputDate {
	return d.set("row", value)
}

// SaveImmediately 是否立即保存
func (d inputDate) SaveImmediately(value bool) inputDate {
	return d.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (d inputDate) Shortcuts(value string) inputDate {
	return d.set("shortcuts", value)
}

// Size 表单项大小
func (d inputDate) Size(value string) inputDate {
	return d.set("size", value)
}

// Static 是否静态展示
func (d inputDate) Static(value bool) inputDate {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d inputDate) StaticClassName(value string) inputDate {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d inputDate) StaticInputClassName(value string) inputDate {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d inputDate) StaticLabelClassName(value string) inputDate {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d inputDate) StaticOn(value string) inputDate {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d inputDate) StaticPlaceholder(value string) inputDate {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d inputDate) StaticSchema(value string) inputDate {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d inputDate) Style(value string) inputDate {
	return d.set("style", value)
}

// SubmitOnChange 修改时是否提交表单
func (d inputDate) SubmitOnChange(value bool) inputDate {
	return d.set("submitOnChange", value)
}

// TestIdBuilder 测试 id 构建器
func (d inputDate) TestIdBuilder(value string) inputDate {
	return d.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (d inputDate) UseMobileUI(value bool) inputDate {
	return d.set("useMobileUI", value)
}

// UTC 是否存储 UTC 时间
func (d inputDate) UTC(value bool) inputDate {
	return d.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (d inputDate) ValidateApi(value string) inputDate {
	return d.set("validateApi", value)
}

// ValidateOnChange 是否在修改时触发验证
func (d inputDate) ValidateOnChange(value bool) inputDate {
	return d.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (d inputDate) ValidationErrors(value string) inputDate {
	return d.set("validationErrors", value)
}

// Validations 验证规则
func (d inputDate) Validations(value string) inputDate {
	return d.set("validations", value)
}

// Value 默认值
func (d inputDate) Value(value string) inputDate {
	return d.set("value", value)
}

// ValueFormat 替代 format
func (d inputDate) ValueFormat(value string) inputDate {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d inputDate) Visible(value bool) inputDate {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d inputDate) VisibleOn(value string) inputDate {
	return d.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (d inputDate) Width(value string) inputDate {
	return d.set("width", value)
}
