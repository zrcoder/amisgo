package comp

// DateControl 日期选择控件
type DateControl Schema

// NewDateControl 创建一个新的 DateControl 实例，并设置默认的 type
func NewDateControl() DateControl {
	return make(DateControl).set("type", "input-date")
}

func (d DateControl) set(key string, value interface{}) DateControl {
	d[key] = value
	return d
}

// AutoFill 自动填充
func (d DateControl) AutoFill(value string) DateControl {
	return d.set("autoFill", value)
}

// BorderMode 边框模式
func (d DateControl) BorderMode(value string) DateControl {
	return d.set("borderMode", value)
}

// ClassName 容器 css 类名
func (d DateControl) ClassName(value string) DateControl {
	return d.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否清除值
func (d DateControl) ClearValueOnHidden(value bool) DateControl {
	return d.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (d DateControl) Clearable(value bool) DateControl {
	return d.set("clearable", value)
}

// CloseOnSelect 点选日期后是否关闭弹窗
func (d DateControl) CloseOnSelect(value bool) DateControl {
	return d.set("closeOnSelect", value)
}

// Desc 描述内容
func (d DateControl) Desc(value string) DateControl {
	return d.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (d DateControl) Description(value string) DateControl {
	return d.set("description", value)
}

// DescriptionClassName 描述内容的 className
func (d DateControl) DescriptionClassName(value string) DateControl {
	return d.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (d DateControl) Disabled(value bool) DateControl {
	return d.set("disabled", value)
}

// DisabledDate 禁用日期函数
func (d DateControl) DisabledDate(value string) DateControl {
	return d.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式
func (d DateControl) DisabledOn(value string) DateControl {
	return d.set("disabledOn", value)
}

// DisplayFormat 日期展示格式
func (d DateControl) DisplayFormat(value string) DateControl {
	return d.set("displayFormat", value)
}

// EditorSetting 编辑器配置
func (d DateControl) EditorSetting(value string) DateControl {
	return d.set("editorSetting", value)
}

// Embed 是否为内联模式
func (d DateControl) Embed(value bool) DateControl {
	return d.set("embed", value)
}

// ExtraName 额外字段名
func (d DateControl) ExtraName(value string) DateControl {
	return d.set("extraName", value)
}

// Format 日期存储格式
func (d DateControl) Format(value string) DateControl {
	return d.set("format", value)
}

// Hidden 是否隐藏
func (d DateControl) Hidden(value bool) DateControl {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d DateControl) HiddenOn(value string) DateControl {
	return d.set("hiddenOn", value)
}

// Hint 输入提示
func (d DateControl) Hint(value string) DateControl {
	return d.set("hint", value)
}

// Horizontal 水平布局配置
func (d DateControl) Horizontal(value string) DateControl {
	return d.set("horizontal", value)
}

// ID 组件唯一 id
func (d DateControl) ID(value string) DateControl {
	return d.set("id", value)
}

// InitAutoFill 初始化自动填充
func (d DateControl) InitAutoFill(value string) DateControl {
	return d.set("initAutoFill", value)
}

// Inline 是否为 inline 模式
func (d DateControl) Inline(value bool) DateControl {
	return d.set("inline", value)
}

// InputClassName 配置 input className
func (d DateControl) InputClassName(value string) DateControl {
	return d.set("inputClassName", value)
}

// InputFormat 输入格式
func (d DateControl) InputFormat(value string) DateControl {
	return d.set("inputFormat", value)
}

// Label 描述标题
func (d DateControl) Label(value string) DateControl {
	return d.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (d DateControl) LabelAlign(value string) DateControl {
	return d.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (d DateControl) LabelClassName(value string) DateControl {
	return d.set("labelClassName", value)
}

// LabelRemark 显示小图标，鼠标放上去显示提示内容
func (d DateControl) LabelRemark(value string) DateControl {
	return d.set("labelRemark", value)
}

// LabelWidth label 宽度
func (d DateControl) LabelWidth(value string) DateControl {
	return d.set("labelWidth", value)
}

// MaxDate 限制最大日期
func (d DateControl) MaxDate(value string) DateControl {
	return d.set("maxDate", value)
}

// MinDate 限制最小日期
func (d DateControl) MinDate(value string) DateControl {
	return d.set("minDate", value)
}

// Mode 组件展示模式
func (d DateControl) Mode(value string) DateControl {
	return d.set("mode", value)
}

// Name 字段名
func (d DateControl) Name(value string) DateControl {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d DateControl) OnEvent(value string) DateControl {
	return d.set("onEvent", value)
}

// Placeholder 占位符
func (d DateControl) Placeholder(value string) DateControl {
	return d.set("placeholder", value)
}

// PopOverContainerSelector 弹窗容器选择器
func (d DateControl) PopOverContainerSelector(value string) DateControl {
	return d.set("popOverContainerSelector", value)
}

// ReadOnly 是否只读
func (d DateControl) ReadOnly(value bool) DateControl {
	return d.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (d DateControl) ReadOnlyOn(value string) DateControl {
	return d.set("readOnlyOn", value)
}

// Remark 小图标提示内容
func (d DateControl) Remark(value string) DateControl {
	return d.set("remark", value)
}

// Required 是否必填
func (d DateControl) Required(value bool) DateControl {
	return d.set("required", value)
}

// Row 行配置
func (d DateControl) Row(value string) DateControl {
	return d.set("row", value)
}

// SaveImmediately 是否立即保存
func (d DateControl) SaveImmediately(value bool) DateControl {
	return d.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (d DateControl) Shortcuts(value string) DateControl {
	return d.set("shortcuts", value)
}

// Size 表单项大小
func (d DateControl) Size(value string) DateControl {
	return d.set("size", value)
}

// Static 是否静态展示
func (d DateControl) Static(value bool) DateControl {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d DateControl) StaticClassName(value string) DateControl {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d DateControl) StaticInputClassName(value string) DateControl {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d DateControl) StaticLabelClassName(value string) DateControl {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d DateControl) StaticOn(value string) DateControl {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d DateControl) StaticPlaceholder(value string) DateControl {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d DateControl) StaticSchema(value string) DateControl {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d DateControl) Style(value string) DateControl {
	return d.set("style", value)
}

// SubmitOnChange 修改时是否提交表单
func (d DateControl) SubmitOnChange(value bool) DateControl {
	return d.set("submitOnChange", value)
}

// TestIdBuilder 测试 id 构建器
func (d DateControl) TestIdBuilder(value string) DateControl {
	return d.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (d DateControl) UseMobileUI(value bool) DateControl {
	return d.set("useMobileUI", value)
}

// UTC 是否存储 UTC 时间
func (d DateControl) UTC(value bool) DateControl {
	return d.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (d DateControl) ValidateApi(value string) DateControl {
	return d.set("validateApi", value)
}

// ValidateOnChange 是否在修改时触发验证
func (d DateControl) ValidateOnChange(value bool) DateControl {
	return d.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (d DateControl) ValidationErrors(value string) DateControl {
	return d.set("validationErrors", value)
}

// Validations 验证规则
func (d DateControl) Validations(value string) DateControl {
	return d.set("validations", value)
}

// Value 默认值
func (d DateControl) Value(value string) DateControl {
	return d.set("value", value)
}

// ValueFormat 替代 format
func (d DateControl) ValueFormat(value string) DateControl {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d DateControl) Visible(value bool) DateControl {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d DateControl) VisibleOn(value string) DateControl {
	return d.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (d DateControl) Width(value string) DateControl {
	return d.set("width", value)
}
