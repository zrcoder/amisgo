package comp

// DateRangeControl 表示一个日期范围控件
type DateRangeControl BaseRenderer

// NewDateRangeControl 创建一个新的 DateRangeControl 实例，并设置默认的 type
func NewDateRangeControl() DateRangeControl {
	d := DateRangeControl(make(BaseRenderer))
	return d.set("type", "input-date-range")
}

// set 设置 key-value 对，并返回当前实例
func (d DateRangeControl) set(key string, value interface{}) DateRangeControl {
	d[key] = value
	return d
}

// Animation 设置动画效果
func (d DateRangeControl) Animation(value bool) DateRangeControl {
	return d.set("animation", value)
}

// AutoFill 设置自动填充
func (d DateRangeControl) AutoFill(value string) DateRangeControl {
	return d.set("autoFill", value)
}

// BorderMode 设置边框模式
func (d DateRangeControl) BorderMode(value string) DateRangeControl {
	return d.set("borderMode", value)
}

// ClassName 设置 class 名称
func (d DateRangeControl) ClassName(value string) DateRangeControl {
	return d.set("className", value)
}

// ClearValueOnHidden 设置隐藏时是否清除值
func (d DateRangeControl) ClearValueOnHidden(value bool) DateRangeControl {
	return d.set("clearValueOnHidden", value)
}

// Delimiter 设置分隔符
func (d DateRangeControl) Delimiter(value string) DateRangeControl {
	return d.set("delimiter", value)
}

// Desc 设置描述
func (d DateRangeControl) Desc(value string) DateRangeControl {
	return d.set("desc", value)
}

// Description 设置描述
func (d DateRangeControl) Description(value string) DateRangeControl {
	return d.set("description", value)
}

// DescriptionClassName 设置描述的 class 名称
func (d DateRangeControl) DescriptionClassName(value string) DateRangeControl {
	return d.set("descriptionClassName", value)
}

// Disabled 设置控件是否禁用
func (d DateRangeControl) Disabled(value bool) DateRangeControl {
	return d.set("disabled", value)
}

// DisabledOn 设置禁用条件
func (d DateRangeControl) DisabledOn(value string) DateRangeControl {
	return d.set("disabledOn", value)
}

// DisplayFormat 设置显示格式
func (d DateRangeControl) DisplayFormat(value string) DateRangeControl {
	return d.set("displayFormat", value)
}

// EditorSetting 设置编辑器配置
func (d DateRangeControl) EditorSetting(value string) DateRangeControl {
	return d.set("editorSetting", value)
}

// Embed 设置是否使用内联模式
func (d DateRangeControl) Embed(value bool) DateRangeControl {
	return d.set("embed", value)
}

// EndPlaceholder 设置结束占位符
func (d DateRangeControl) EndPlaceholder(value string) DateRangeControl {
	return d.set("endPlaceholder", value)
}

// ExtraName 设置额外字段名称
func (d DateRangeControl) ExtraName(value string) DateRangeControl {
	return d.set("extraName", value)
}

// Format 设置提交格式
func (d DateRangeControl) Format(value string) DateRangeControl {
	return d.set("format", value)
}

// Hidden 设置控件是否隐藏
func (d DateRangeControl) Hidden(value bool) DateRangeControl {
	return d.set("hidden", value)
}

// HiddenOn 设置隐藏条件
func (d DateRangeControl) HiddenOn(value string) DateRangeControl {
	return d.set("hiddenOn", value)
}

// Hint 设置输入提示
func (d DateRangeControl) Hint(value string) DateRangeControl {
	return d.set("hint", value)
}

// Horizontal 设置水平布局
func (d DateRangeControl) Horizontal(value string) DateRangeControl {
	return d.set("horizontal", value)
}

// ID 设置组件 ID
func (d DateRangeControl) ID(value string) DateRangeControl {
	return d.set("id", value)
}

// InitAutoFill 设置初始化自动填充值
func (d DateRangeControl) InitAutoFill(value string) DateRangeControl {
	return d.set("initAutoFill", value)
}

// Inline 设置是否使用内联模式
func (d DateRangeControl) Inline(value bool) DateRangeControl {
	return d.set("inline", value)
}

// InputClassName 设置输入框的 class 名称
func (d DateRangeControl) InputClassName(value string) DateRangeControl {
	return d.set("inputClassName", value)
}

// InputFormat 设置输入格式
func (d DateRangeControl) InputFormat(value string) DateRangeControl {
	return d.set("inputFormat", value)
}

// JoinValues 设置是否将值合并为一个字符串
func (d DateRangeControl) JoinValues(value bool) DateRangeControl {
	return d.set("joinValues", value)
}

// Label 设置控件标签
func (d DateRangeControl) Label(value string) DateRangeControl {
	return d.set("label", value)
}

// LabelAlign 设置标签对齐方式
func (d DateRangeControl) LabelAlign(value string) DateRangeControl {
	return d.set("labelAlign", value)
}

// LabelClassName 设置标签的 class 名称
func (d DateRangeControl) LabelClassName(value string) DateRangeControl {
	return d.set("labelClassName", value)
}

// LabelRemark 设置标签备注
func (d DateRangeControl) LabelRemark(value string) DateRangeControl {
	return d.set("labelRemark", value)
}

// LabelWidth 设置标签宽度
func (d DateRangeControl) LabelWidth(value string) DateRangeControl {
	return d.set("labelWidth", value)
}

// MaxDate 设置最大日期限制
func (d DateRangeControl) MaxDate(value string) DateRangeControl {
	return d.set("maxDate", value)
}

// MaxDuration 设置最大持续时间
func (d DateRangeControl) MaxDuration(value string) DateRangeControl {
	return d.set("maxDuration", value)
}

// MinDate 设置最小日期限制
func (d DateRangeControl) MinDate(value string) DateRangeControl {
	return d.set("minDate", value)
}

// MinDuration 设置最小持续时间
func (d DateRangeControl) MinDuration(value string) DateRangeControl {
	return d.set("minDuration", value)
}

// Mode 设置表单项显示模式
func (d DateRangeControl) Mode(value string) DateRangeControl {
	return d.set("mode", value)
}

// Name 设置表单提交字段名称
func (d DateRangeControl) Name(value string) DateRangeControl {
	return d.set("name", value)
}

// OnEvent 设置事件配置
func (d DateRangeControl) OnEvent(value string) DateRangeControl {
	return d.set("onEvent", value)
}

// Placeholder 设置占位符文本
func (d DateRangeControl) Placeholder(value string) DateRangeControl {
	return d.set("placeholder", value)
}

// PopOverContainerSelector 设置弹出层容器选择器
func (d DateRangeControl) PopOverContainerSelector(value string) DateRangeControl {
	return d.set("popOverContainerSelector", value)
}

// Ranges 设置日期范围快捷方式
func (d DateRangeControl) Ranges(value string) DateRangeControl {
	return d.set("ranges", value)
}

// ReadOnly 设置控件是否为只读
func (d DateRangeControl) ReadOnly(value bool) DateRangeControl {
	return d.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (d DateRangeControl) ReadOnlyOn(value string) DateRangeControl {
	return d.set("readOnlyOn", value)
}

// Remark 设置控件备注
func (d DateRangeControl) Remark(value string) DateRangeControl {
	return d.set("remark", value)
}

// Required 设置控件是否必填
func (d DateRangeControl) Required(value bool) DateRangeControl {
	return d.set("required", value)
}

// Row 设置行数
func (d DateRangeControl) Row(value string) DateRangeControl {
	return d.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (d DateRangeControl) SaveImmediately(value bool) DateRangeControl {
	return d.set("saveImmediately", value)
}

// Shortcuts 设置日期范围快捷方式
func (d DateRangeControl) Shortcuts(value string) DateRangeControl {
	return d.set("shortcuts", value)
}

// StartPlaceholder 设置开始占位符
func (d DateRangeControl) StartPlaceholder(value string) DateRangeControl {
	return d.set("startPlaceholder", value)
}

// TimeFormat 设置时间格式
func (d DateRangeControl) TimeFormat(value string) DateRangeControl {
	return d.set("timeFormat", value)
}

// UseMobileUI 设置是否使用移动端 UI
func (d DateRangeControl) UseMobileUI(value bool) DateRangeControl {
	return d.set("useMobileUI", value)
}
