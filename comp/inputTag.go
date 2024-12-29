package comp

// inputTag 输入框 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tag

type inputTag Schema

func InputTag() inputTag {
	return inputTag{}.set("type", "input-tag")
}

func (t inputTag) set(key string, value any) inputTag {
	t[key] = value
	return t
}

// AddApi 添加时调用的接口
func (t inputTag) AddApi(value string) inputTag {
	return t.set("addApi", value)
}

// AddControls 新增时的表单项。
func (t inputTag) AddControls(value string) inputTag {
	return t.set("addControls", value)
}

// AddDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (t inputTag) AddDialog(value string) inputTag {
	return t.set("addDialog", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t inputTag) AutoFill(value string) inputTag {
	return t.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t inputTag) ClassName(value string) inputTag {
	return t.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t inputTag) ClearValueOnHidden(value bool) inputTag {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t inputTag) Clearable(value bool) inputTag {
	return t.set("clearable", value)
}

// Creatable 是否可以新增
func (t inputTag) Creatable(value bool) inputTag {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t inputTag) CreateBtnLabel(value string) inputTag {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t inputTag) DeferApi(value string) inputTag {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t inputTag) DeferField(value string) inputTag {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t inputTag) DeleteApi(value string) inputTag {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t inputTag) DeleteConfirmText(value string) inputTag {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t inputTag) Delimiter(value string) inputTag {
	return t.set("delimiter", value)
}

// Desc
func (t inputTag) Desc(value string) inputTag {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t inputTag) Description(value string) inputTag {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (t inputTag) DescriptionClassName(value string) inputTag {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t inputTag) Disabled(value bool) inputTag {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputTag) DisabledOn(value string) inputTag {
	return t.set("disabledOn", value)
}

// Dropdown 是否为下拉模式
func (t inputTag) Dropdown(value bool) inputTag {
	return t.set("dropdown", value)
}

// EditApi 编辑时调用的 API
func (t inputTag) EditApi(value string) inputTag {
	return t.set("editApi", value)
}

// EditControls 选项修改的表单项
func (t inputTag) EditControls(value string) inputTag {
	return t.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (t inputTag) EditDialog(value string) inputTag {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t inputTag) Editable(value bool) inputTag {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t inputTag) EditorSetting(value string) inputTag {
	return t.set("editorSetting", value)
}

// EnableBatchAdd 是否开启批量添加模式
func (t inputTag) EnableBatchAdd(value bool) inputTag {
	return t.set("enableBatchAdd", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t inputTag) ExtraName(value string) inputTag {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t inputTag) ExtractValue(value bool) inputTag {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t inputTag) Hidden(value bool) inputTag {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputTag) HiddenOn(value string) inputTag {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t inputTag) Hint(value string) inputTag {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (t inputTag) Horizontal(value string) inputTag {
	return t.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (t inputTag) ID(value string) inputTag {
	return t.set("id", value)
}

// InitAutoFill
func (t inputTag) InitAutoFill(value string) inputTag {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t inputTag) InitFetch(value bool) inputTag {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t inputTag) InitFetchOn(value string) inputTag {
	return t.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t inputTag) Inline(value bool) inputTag {
	return t.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (t inputTag) InputClassName(value string) inputTag {
	return t.set("inputClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t inputTag) JoinValues(value bool) inputTag {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t inputTag) Label(value string) inputTag {
	return t.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (t inputTag) LabelAlign(value string) inputTag {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t inputTag) LabelClassName(value string) inputTag {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (t inputTag) LabelRemark(value string) inputTag {
	return t.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (t inputTag) LabelWidth(value string) inputTag {
	return t.set("labelWidth", value)
}

// Max 允许添加的标签的最大数量
func (t inputTag) Max(value string) inputTag {
	return t.set("max", value)
}

// MaxTagCount 标签的最大展示数量，超出数量后以收纳浮层的方式展示，仅在多选模式开启后生效
func (t inputTag) MaxTagCount(value string) inputTag {
	return t.set("maxTagCount", value)
}

// MaxTagText 标签最大展示数量提示文本，默认“...”
func (t inputTag) MaxTagText(value string) inputTag {
	return t.set("maxTagText", value)
}

// Min 最少选择标签数
func (t inputTag) Min(value string) inputTag {
	return t.set("min", value)
}

// Mode 控制表单项的展示模式 (模式)
func (t inputTag) Mode(value string) inputTag {
	return t.set("mode", value)
}

// Name 表单项 name 属性
func (t inputTag) Name(value string) inputTag {
	return t.set("name", value)
}

// OnEvent 事件处理
func (t inputTag) OnEvent(value map[string]string) inputTag {
	return t.set("onEvent", value)
}

// Options 选项 (选项)
func (t inputTag) Options(value ...any) inputTag {
	return t.set("options", value)
}

// OptionValuePath
func (t inputTag) OptionValuePath(value string) inputTag {
	return t.set("optionValuePath", value)
}

// OptionLabelPath
func (t inputTag) OptionLabelPath(value string) inputTag {
	return t.set("optionLabelPath", value)
}

// OptionClassName
func (t inputTag) OptionClassName(value string) inputTag {
	return t.set("optionClassName", value)
}

// Placeholder 输入框提示信息
func (t inputTag) Placeholder(value string) inputTag {
	return t.set("placeholder", value)
}

// ReadOnly 是否只读
func (t inputTag) ReadOnly(value bool) inputTag {
	return t.set("readOnly", value)
}

// Required 是否必填
func (t inputTag) Required(value bool) inputTag {
	return t.set("required", value)
}

// Rules 表单验证规则
func (t inputTag) Rules(value string) inputTag {
	return t.set("rules", value)
}

// Source 数据源
func (t inputTag) Source(value string) inputTag {
	return t.set("source", value)
}

// TagClassName 标签的 className (标签的 className)
func (t inputTag) TagClassName(value string) inputTag {
	return t.set("tagClassName", value)
}

// TagTpl 自定义标签模板
func (t inputTag) TagTpl(value string) inputTag {
	return t.set("tagTpl", value)
}

// Touchable 是否支持触摸操作 (是否支持触摸操作)
func (t inputTag) Touchable(value bool) inputTag {
	return t.set("touchable", value)
}

// Value 表单项的值
func (t inputTag) Value(value any) inputTag {
	return t.set("value", value)
}

// VerticalLayout
func (t inputTag) VerticalLayout(value bool) inputTag {
	return t.set("verticalLayout", value)
}

// Visible 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (t inputTag) Visible(value bool) inputTag {
	return t.set("visible", value)
}

// ValidationErrors 自定义验证错误信息
func (t inputTag) ValidationErrors(value map[string]string) inputTag {
	return t.set("validationErrors", value)
}

// ValuePath
func (t inputTag) ValuePath(value string) inputTag {
	return t.set("valuePath", value)
}

// ValidateOnChange 失去焦点后是否进行验证
func (t inputTag) ValidateOnChange(value bool) inputTag {
	return t.set("validateOnChange", value)
}

// Validations 设置验证规则
func (t inputTag) Validations(value string) inputTag {
	return t.set("validations", value)
}

// ValuesNoWrap 设置多选模式，值太多时是否避免折行
func (t inputTag) ValuesNoWrap(value bool) inputTag {
	return t.set("valuesNoWrap", value)
}

// VisibleOn 设置是否显示的表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputTag) VisibleOn(value string) inputTag {
	return t.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (t inputTag) Width(value string) inputTag {
	return t.set("width", value)
}
