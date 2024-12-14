package comp

// inputCity 城市选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/city
type inputCity schema

func InputCity() inputCity {
	return make(inputCity).set("type", "input-city")
}

func (i inputCity) set(key string, value any) inputCity {
	i[key] = value
	return i
}

// AllowCity 允许选择城市？
func (i inputCity) AllowCity(value bool) inputCity {
	return i.set("allowCity", value)
}

// AllowDistrict 允许选择地区？
func (i inputCity) AllowDistrict(value bool) inputCity {
	return i.set("allowDistrict", value)
}

// AllowStreet 允许选择街道？
func (i inputCity) AllowStreet(value bool) inputCity {
	return i.set("allowStreet", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (i inputCity) AutoFill(value string) inputCity {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputCity) ClassName(value string) inputCity {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (i inputCity) ClearValueOnHidden(value bool) inputCity {
	return i.set("clearValueOnHidden", value)
}

// Delimiter 拼接的符号是啥？
func (i inputCity) Delimiter(value string) inputCity {
	return i.set("delimiter", value)
}

// Desc
func (i inputCity) Desc(value string) inputCity {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (i inputCity) Description(value string) inputCity {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (i inputCity) DescriptionClassName(value string) inputCity {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i inputCity) Disabled(value bool) inputCity {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputCity) DisabledOn(value string) inputCity {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i inputCity) EditorSetting(value string) inputCity {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i inputCity) ExtraName(value string) inputCity {
	return i.set("extraName", value)
}

// ExtractValue 开启后只会存城市的 code 信息
func (i inputCity) ExtractValue(value bool) inputCity {
	return i.set("extractValue", value)
}

// Hidden 是否隐藏
func (i inputCity) Hidden(value bool) inputCity {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputCity) HiddenOn(value string) inputCity {
	return i.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (i inputCity) Hint(value string) inputCity {
	return i.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (i inputCity) Horizontal(value string) inputCity {
	return i.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i inputCity) ID(value string) inputCity {
	return i.set("id", value)
}

// InitAutoFill
func (i inputCity) InitAutoFill(value string) inputCity {
	return i.set("initAutoFill", value)
}

// Inline 表单  是否为 inline 模式。
func (i inputCity) Inline(value bool) inputCity {
	return i.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (i inputCity) InputClassName(value string) inputCity {
	return i.set("inputClassName", value)
}

// ItemClassName 下拉框className
func (i inputCity) ItemClassName(value string) inputCity {
	return i.set("itemClassName", value)
}

// JoinValues 是否将各个信息拼接成字符串。
func (i inputCity) JoinValues(value bool) inputCity {
	return i.set("joinValues", value)
}

// Label 描述标题
func (i inputCity) Label(value string) inputCity {
	return i.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (i inputCity) LabelAlign(value string) inputCity {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i inputCity) LabelClassName(value string) inputCity {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (i inputCity) LabelRemark(value string) inputCity {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (i inputCity) LabelWidth(value string) inputCity {
	return i.set("labelWidth", value)
}

// LoadingConfig
func (i inputCity) LoadingConfig(value string) inputCity {
	return i.set("loadingConfig", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (i inputCity) Mode(value string) inputCity {
	return i.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (i inputCity) Name(value string) inputCity {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i inputCity) OnEvent(value any) inputCity {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i inputCity) Placeholder(value string) inputCity {
	return i.set("placeholder", value)
}

// ReadOnly 是否只读
func (i inputCity) ReadOnly(value bool) inputCity {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i inputCity) ReadOnlyOn(value string) inputCity {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (i inputCity) Remark(value string) inputCity {
	return i.set("remark", value)
}

// Required 是否为必填
func (i inputCity) Required(value bool) inputCity {
	return i.set("required", value)
}

// Row
func (i inputCity) Row(value string) inputCity {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (i inputCity) SaveImmediately(value bool) inputCity {
	return i.set("saveImmediately", value)
}

// Searchable 是否显示搜索框
func (i inputCity) Searchable(value bool) inputCity {
	return i.set("searchable", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (i inputCity) Size(value string) inputCity {
	return i.set("size", value)
}

// Static 是否静态展示
func (i inputCity) Static(value bool) inputCity {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputCity) StaticClassName(value string) inputCity {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputCity) StaticInputClassName(value string) inputCity {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputCity) StaticLabelClassName(value string) inputCity {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputCity) StaticOn(value string) inputCity {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i inputCity) StaticPlaceholder(value string) inputCity {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i inputCity) StaticSchema(value string) inputCity {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i inputCity) Style(value any) inputCity {
	return i.set("style", value)
}

// SubmitOnChange 当值变化后是否自动提交
func (i inputCity) SubmitOnChange(value bool) inputCity {
	return i.set("submitOnChange", value)
}

// TestIdBuilder
func (i inputCity) TestIdBuilder(value string) inputCity {
	return i.set("testIdBuilder", value)
}

// UseMobile 是否使用移动端适配
func (i inputCity) UseMobile(value bool) inputCity {
	return i.set("useMobile", value)
}

// ValidateApi 远端校验表单项接口
func (i inputCity) ValidateApi(value string) inputCity {
	return i.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (i inputCity) ValidateOnChange(value bool) inputCity {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i inputCity) ValidationErrors(value string) inputCity {
	return i.set("validationErrors", value)
}

// Validations
func (i inputCity) Validations(value string) inputCity {
	return i.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (i inputCity) Value(value any) inputCity {
	return i.set("value", value)
}

// Visible 是否显示
func (i inputCity) Visible(value bool) inputCity {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputCity) VisibleOn(value string) inputCity {
	return i.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (i inputCity) Width(value string) inputCity {
	return i.set("width", value)
}
