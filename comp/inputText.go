package comp

// inputText 文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/text

type inputText schema

func InputText() inputText {
	return inputText{}.set("type", "input-text")
}

func InputEmail() inputText {
	return inputText{}.set("type", "input-email")
}

func InputUrl() inputText {
	return inputText{}.set("type", "input-url")
}

func InputPassword() inputText {
	return inputText{}.set("type", "input-password")
}

// AddApi 添加时调用的接口
func (t inputText) AddApi(value string) inputText {
	return t.set("addApi", value)
}

// AddControls 新增时的表单项。
func (t inputText) AddControls(value string) inputText {
	return t.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (t inputText) AddDialog(value string) inputText {
	return t.set("addDialog", value)
}

// AddOn
func (t inputText) AddOn(value any) inputText {
	return t.set("addOn", value)
}

// AutoComplete 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。
func (t inputText) AutoComplete(value string) inputText {
	return t.set("autoComplete", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t inputText) AutoFill(value string) inputText {
	return t.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (t inputText) BorderMode(value string) inputText {
	return t.set("borderMode", value)
}

// ClassName 容器 css 类名
func (t inputText) ClassName(value string) inputText {
	return t.set("className", value)
}

// ColumnClassName
func (t inputText) ColumnClassName(value string) inputText {
	return t.set("columnClassName", value)
}

// ClearValueOnEmpty 在内容为空的时候清除值
func (t inputText) ClearValueOnEmpty(value bool) inputText {
	return t.set("clearValueOnEmpty", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t inputText) ClearValueOnHidden(value bool) inputText {
	return t.set("clearValueOnHidden", value)
}

// Clearable 是否可清除。
func (t inputText) Clearable(value bool) inputText {
	return t.set("clearable", value)
}

// Creatable 是否可以新增
func (t inputText) Creatable(value bool) inputText {
	return t.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (t inputText) CreateBtnLabel(value string) inputText {
	return t.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (t inputText) DeferApi(value string) inputText {
	return t.set("deferApi", value)
}

// DeferField 懒加载字段
func (t inputText) DeferField(value string) inputText {
	return t.set("deferField", value)
}

// DeleteApi 选项删除 API
func (t inputText) DeleteApi(value string) inputText {
	return t.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字。
func (t inputText) DeleteConfirmText(value string) inputText {
	return t.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (t inputText) Delimiter(value string) inputText {
	return t.set("delimiter", value)
}

// Desc 描述内容
func (t inputText) Desc(value string) inputText {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t inputText) Description(value string) inputText {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (t inputText) DescriptionClassName(value string) inputText {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t inputText) Disabled(value bool) inputText {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputText) DisabledOn(value string) inputText {
	return t.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (t inputText) EditApi(value string) inputText {
	return t.set("editApi", value)
}

// EditControls 选项修改的表单项
func (t inputText) EditControls(value string) inputText {
	return t.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (t inputText) EditDialog(value string) inputText {
	return t.set("editDialog", value)
}

// Editable 是否可以编辑
func (t inputText) Editable(value bool) inputText {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t inputText) EditorSetting(value string) inputText {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t inputText) ExtraName(value string) inputText {
	return t.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (t inputText) ExtractValue(value bool) inputText {
	return t.set("extractValue", value)
}

// Hidden 是否隐藏
func (t inputText) Hidden(value bool) inputText {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputText) HiddenOn(value string) inputText {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t inputText) Hint(value string) inputText {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (t inputText) Horizontal(value string) inputText {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t inputText) ID(value string) inputText {
	return t.set("id", value)
}

// InitAutoFill
func (t inputText) InitAutoFill(value string) inputText {
	return t.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取。
func (t inputText) InitFetch(value bool) inputText {
	return t.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (t inputText) InitFetchOn(value string) inputText {
	return t.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t inputText) Inline(value bool) inputText {
	return t.set("inline", value)
}

// InputClassName 配置 input className
func (t inputText) InputClassName(value string) inputText {
	return t.set("inputClassName", value)
}

// InputControlClassName control节点的CSS类名
func (t inputText) InputControlClassName(value string) inputText {
	return t.set("inputControlClassName", value)
}

// JoinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (t inputText) JoinValues(value bool) inputText {
	return t.set("joinValues", value)
}

// Label 描述标题
func (t inputText) Label(value string) inputText {
	return t.set("label", value)
}

// Size 尺寸 'xs' | 'sm' | 'md' | 'lg' | 'full'
func (t inputText) Size(value string) inputText {
	return t.set("size", value)
}

// Remark string or remark
func (t inputText) Remark(value any) inputText {
	return t.set("remark", value)
}

// LabelAlign 描述标题 可选值: right | left | top | inherit
func (t inputText) LabelAlign(value string) inputText {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t inputText) LabelClassName(value string) inputText {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (t inputText) LabelRemark(value string) inputText {
	return t.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (t inputText) LabelWidth(value string) inputText {
	return t.set("labelWidth", value)
}

// MaxLength 限制文字最大输入个数
func (t inputText) MaxLength(value string) inputText {
	return t.set("maxLength", value)
}

// MinLength 限制文字最小输入个数
func (t inputText) MinLength(value string) inputText {
	return t.set("minLength", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (t inputText) Mode(value string) inputText {
	return t.set("mode", value)
}

// Multiple 是否为多选模式
func (t inputText) Multiple(value bool) inputText {
	return t.set("multiple", value)
}

// Name 表单项名称
func (t inputText) Name(value string) inputText {
	return t.set("name", value)
}

// Required 是否必须输入
func (t inputText) Required(value bool) inputText {
	return t.set("required", value)
}

// Note
func (t inputText) Note(value string) inputText {
	return t.set("note", value)
}

// OnChange 触发的事件
func (t inputText) OnChange(value string) inputText {
	return t.set("onChange", value)
}

// OnFocus 触发的事件
func (t inputText) OnFocus(value string) inputText {
	return t.set("onFocus", value)
}

// OnBlur 触发的事件
func (t inputText) OnBlur(value string) inputText {
	return t.set("onBlur", value)
}

// OnInit 触发的事件
func (t inputText) OnInit(value string) inputText {
	return t.set("onInit", value)
}

// OnSearch 触发的事件
func (t inputText) OnSearch(value string) inputText {
	return t.set("onSearch", value)
}

// Option
func (t inputText) Option(value string) inputText {
	return t.set("option", value)
}

// Options 选项列表
func (t inputText) Options(value ...string) inputText {
	return t.set("options", value)
}

// OptionLabel 在多选模式下，选项的 label 位置。可选值: top | left | right
func (t inputText) OptionLabel(value string) inputText {
	return t.set("optionLabel", value)
}

// Placeholder 提示文字
func (t inputText) Placeholder(value string) inputText {
	return t.set("placeholder", value)
}

// ReadOnly 是否为只读模式
func (t inputText) ReadOnly(value bool) inputText {
	return t.set("readOnly", value)
}

// RefixApi 配置选择项的接口
func (t inputText) RefixApi(value string) inputText {
	return t.set("refixApi", value)
}

// ResetValue
func (t inputText) ResetValue(value string) inputText {
	return t.set("resetValue", value)
}

// Rules 校验规则
func (t inputText) Rules(value string) inputText {
	return t.set("rules", value)
}

// SearchApi 搜索接口
func (t inputText) SearchApi(value string) inputText {
	return t.set("searchApi", value)
}

// Source 数据来源
func (t inputText) Source(value string) inputText {
	return t.set("source", value)
}

// StaticClassName 用于静态展示的类名
func (t inputText) StaticClassName(value string) inputText {
	return t.set("staticClassName", value)
}

// StaticLabel 静态显示的标题
func (t inputText) StaticLabel(value string) inputText {
	return t.set("staticLabel", value)
}

// StaticOn 静态显示的条件
func (t inputText) StaticOn(value string) inputText {
	return t.set("staticOn", value)
}

// StaticValue 静态值
func (t inputText) StaticValue(value string) inputText {
	return t.set("staticValue", value)
}

// Tips 提示文本
func (t inputText) Tips(value string) inputText {
	return t.set("tips", value)
}

// ValidationErrors 校验不通过时的提示信息
func (t inputText) ValidationErrors(value string) inputText {
	return t.set("validationErrors", value)
}

// Value 表单项的值
func (t inputText) Value(value any) inputText {
	return t.set("value", value)
}

// Visible 控制组件是否可见
func (t inputText) Visible(value bool) inputText {
	return t.set("visible", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (t inputText) VisibleOn(value string) inputText {
	return t.set("visibleOn", value)
}

// Width 控件宽度
func (t inputText) Width(value string) inputText {
	return t.set("width", value)
}

// OnInput 触发的事件
func (t inputText) OnInput(value string) inputText {
	return t.set("onInput", value)
}

// ValueType
func (t inputText) ValueType(value string) inputText {
	return t.set("valueType", value)
}

// ValueOnEvent
func (t inputText) ValueOnEvent(value any) inputText {
	return t.set("valueOnEvent", value)
}

// set 设置属性
func (t inputText) set(key string, value any) inputText {
	t[key] = value
	return t
}
