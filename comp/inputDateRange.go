package comp

// inputDateRange 表示一个日期范围控件
type inputDateRange schema

// InputDateRange 创建一个新的 InputDateRange 实例，并设置默认的 type
func InputDateRange() inputDateRange {
	return make(inputDateRange).set("type", "input-date-range")
}

// set 设置 key-value 对，并返回当前实例
func (d inputDateRange) set(key string, value any) inputDateRange {
	d[key] = value
	return d
}

// Animation 设置动画效果
func (d inputDateRange) Animation(value bool) inputDateRange {
	return d.set("animation", value)
}

// AutoFill 设置自动填充
func (d inputDateRange) AutoFill(value string) inputDateRange {
	return d.set("autoFill", value)
}

// BorderMode 设置边框模式
func (d inputDateRange) BorderMode(value string) inputDateRange {
	return d.set("borderMode", value)
}

// ClassName 设置 class 名称
func (d inputDateRange) ClassName(value string) inputDateRange {
	return d.set("className", value)
}

// ClearValueOnHidden 设置隐藏时是否清除值
func (d inputDateRange) ClearValueOnHidden(value bool) inputDateRange {
	return d.set("clearValueOnHidden", value)
}

// Delimiter 设置分隔符
func (d inputDateRange) Delimiter(value string) inputDateRange {
	return d.set("delimiter", value)
}

// Desc 设置描述
func (d inputDateRange) Desc(value string) inputDateRange {
	return d.set("desc", value)
}

// Description 设置描述
func (d inputDateRange) Description(value string) inputDateRange {
	return d.set("description", value)
}

// DescriptionClassName 设置描述的 class 名称
func (d inputDateRange) DescriptionClassName(value string) inputDateRange {
	return d.set("descriptionClassName", value)
}

// Disabled 设置控件是否禁用
func (d inputDateRange) Disabled(value bool) inputDateRange {
	return d.set("disabled", value)
}

// DisabledOn 设置禁用条件
func (d inputDateRange) DisabledOn(value string) inputDateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat 设置显示格式
func (d inputDateRange) DisplayFormat(value string) inputDateRange {
	return d.set("displayFormat", value)
}

// EditorSetting 设置编辑器配置
func (d inputDateRange) EditorSetting(value string) inputDateRange {
	return d.set("editorSetting", value)
}

// Embed 设置是否使用内联模式
func (d inputDateRange) Embed(value bool) inputDateRange {
	return d.set("embed", value)
}

// EndPlaceholder 设置结束占位符
func (d inputDateRange) EndPlaceholder(value string) inputDateRange {
	return d.set("endPlaceholder", value)
}

// ExtraName 设置额外字段名称
func (d inputDateRange) ExtraName(value string) inputDateRange {
	return d.set("extraName", value)
}

// Format 设置提交格式
func (d inputDateRange) Format(value string) inputDateRange {
	return d.set("format", value)
}

// Hidden 设置控件是否隐藏
func (d inputDateRange) Hidden(value bool) inputDateRange {
	return d.set("hidden", value)
}

// HiddenOn 设置隐藏条件
func (d inputDateRange) HiddenOn(value string) inputDateRange {
	return d.set("hiddenOn", value)
}

// Hint 设置输入提示
func (d inputDateRange) Hint(value string) inputDateRange {
	return d.set("hint", value)
}

// Horizontal 设置水平布局
func (d inputDateRange) Horizontal(value string) inputDateRange {
	return d.set("horizontal", value)
}

// ID 设置组件 ID
func (d inputDateRange) ID(value string) inputDateRange {
	return d.set("id", value)
}

// InitAutoFill 设置初始化自动填充值
func (d inputDateRange) InitAutoFill(value string) inputDateRange {
	return d.set("initAutoFill", value)
}

// Inline 设置是否使用内联模式
func (d inputDateRange) Inline(value bool) inputDateRange {
	return d.set("inline", value)
}

// Value 设置值
func (d inputDateRange) Value(value string) inputDateRange {
	return d.set("value", value)
}

// CloseOnSelect 设置是否在选择后关闭
func (d inputDateRange) CloseOnSelect(value bool) inputDateRange {
	return d.set("closeOnSelect", value)
}

// Clearable
func (d inputDateRange) Clearable(value bool) inputDateRange {
	return d.set("clearable", value)
}

// InputClassName 设置输入框的 class 名称
func (d inputDateRange) InputClassName(value string) inputDateRange {
	return d.set("inputClassName", value)
}

// InputFormat 设置输入格式
func (d inputDateRange) InputFormat(value string) inputDateRange {
	return d.set("inputFormat", value)
}

// JoinValues 设置是否将值合并为一个字符串
func (d inputDateRange) JoinValues(value bool) inputDateRange {
	return d.set("joinValues", value)
}

// Label 设置控件标签
func (d inputDateRange) Label(value string) inputDateRange {
	return d.set("label", value)
}

// LabelAlign 设置标签对齐方式
func (d inputDateRange) LabelAlign(value string) inputDateRange {
	return d.set("labelAlign", value)
}

// LabelClassName 设置标签的 class 名称
func (d inputDateRange) LabelClassName(value string) inputDateRange {
	return d.set("labelClassName", value)
}

// LabelRemark 设置标签备注
func (d inputDateRange) LabelRemark(value string) inputDateRange {
	return d.set("labelRemark", value)
}

// LabelWidth 设置标签宽度
func (d inputDateRange) LabelWidth(value string) inputDateRange {
	return d.set("labelWidth", value)
}

// MaxDate 设置最大日期限制
func (d inputDateRange) MaxDate(value string) inputDateRange {
	return d.set("maxDate", value)
}

// MaxDuration 设置最大持续时间
func (d inputDateRange) MaxDuration(value string) inputDateRange {
	return d.set("maxDuration", value)
}

// MinDate 设置最小日期限制
func (d inputDateRange) MinDate(value string) inputDateRange {
	return d.set("minDate", value)
}

// MinDuration 设置最小持续时间
func (d inputDateRange) MinDuration(value string) inputDateRange {
	return d.set("minDuration", value)
}

// Mode 设置表单项显示模式
func (d inputDateRange) Mode(value string) inputDateRange {
	return d.set("mode", value)
}

// Name 设置表单提交字段名称
func (d inputDateRange) Name(value string) inputDateRange {
	return d.set("name", value)
}

// OnEvent 设置事件配置
func (d inputDateRange) OnEvent(value any) inputDateRange {
	return d.set("onEvent", value)
}

// Placeholder 设置占位符文本
func (d inputDateRange) Placeholder(value string) inputDateRange {
	return d.set("placeholder", value)
}

// PopOverContainerSelector 设置弹出层容器选择器
func (d inputDateRange) PopOverContainerSelector(value string) inputDateRange {
	return d.set("popOverContainerSelector", value)
}

// Ranges 设置日期范围快捷方式
func (d inputDateRange) Ranges(value string) inputDateRange {
	return d.set("ranges", value)
}

// ReadOnly 设置控件是否为只读
func (d inputDateRange) ReadOnly(value bool) inputDateRange {
	return d.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (d inputDateRange) ReadOnlyOn(value string) inputDateRange {
	return d.set("readOnlyOn", value)
}

// Remark 设置控件备注
func (d inputDateRange) Remark(value string) inputDateRange {
	return d.set("remark", value)
}

// Required 设置控件是否必填
func (d inputDateRange) Required(value bool) inputDateRange {
	return d.set("required", value)
}

// Row 设置行数
func (d inputDateRange) Row(value string) inputDateRange {
	return d.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (d inputDateRange) SaveImmediately(value bool) inputDateRange {
	return d.set("saveImmediately", value)
}

// Shortcuts 设置日期范围快捷方式
func (d inputDateRange) Shortcuts(value string) inputDateRange {
	return d.set("shortcuts", value)
}

// StartPlaceholder 设置开始占位符
func (d inputDateRange) StartPlaceholder(value string) inputDateRange {
	return d.set("startPlaceholder", value)
}

// TimeFormat 设置时间格式
func (d inputDateRange) TimeFormat(value string) inputDateRange {
	return d.set("timeFormat", value)
}

// UseMobileUI 设置是否使用移动端 UI
func (d inputDateRange) UseMobileUI(value bool) inputDateRange {
	return d.set("useMobileUI", value)
}
