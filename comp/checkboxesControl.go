package comp

type CheckboxesControl Schema

func NewCheckboxesControl() CheckboxesControl {
	return make(CheckboxesControl).set("type", "checkboxes")
}

func (c CheckboxesControl) set(key string, value interface{}) CheckboxesControl {
	c[key] = value
	return c
}

// 添加时调用的接口
func (c CheckboxesControl) AddApi(value string) CheckboxesControl {
	return c.set("addApi", value)
}

// 新增时的表单项
func (c CheckboxesControl) AddControls(value string) CheckboxesControl {
	return c.set("addControls", value)
}

// 控制新增弹框设置项
func (c CheckboxesControl) AddDialog(value string) CheckboxesControl {
	return c.set("addDialog", value)
}

// 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (c CheckboxesControl) AutoFill(value string) CheckboxesControl {
	return c.set("autoFill", value)
}

// 是否开启全选功能
func (c CheckboxesControl) CheckAll(value bool) CheckboxesControl {
	return c.set("checkAll", value)
}

// 全选/不选文案
func (c CheckboxesControl) CheckAllText(value string) CheckboxesControl {
	return c.set("checkAllText", value)
}

// 容器 css 类名
func (c CheckboxesControl) ClassName(value string) CheckboxesControl {
	return c.set("className", value)
}

// 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (c CheckboxesControl) ClearValueOnHidden(value bool) CheckboxesControl {
	return c.set("clearValueOnHidden", value)
}

// 是否可清除
func (c CheckboxesControl) Clearable(value bool) CheckboxesControl {
	return c.set("clearable", value)
}

// 每行显示多少个
func (c CheckboxesControl) ColumnsCount(value string) CheckboxesControl {
	return c.set("columnsCount", value)
}

// 是否可以新增
func (c CheckboxesControl) Creatable(value bool) CheckboxesControl {
	return c.set("creatable", value)
}

// 新增文字
func (c CheckboxesControl) CreateBtnLabel(value string) CheckboxesControl {
	return c.set("createBtnLabel", value)
}

// 是否默认全选
func (c CheckboxesControl) DefaultCheckAll(value bool) CheckboxesControl {
	return c.set("defaultCheckAll", value)
}

// 延时加载的 API
func (c CheckboxesControl) DeferApi(value string) CheckboxesControl {
	return c.set("deferApi", value)
}

// 懒加载字段
func (c CheckboxesControl) DeferField(value string) CheckboxesControl {
	return c.set("deferField", value)
}

// 选项删除 API
func (c CheckboxesControl) DeleteApi(value string) CheckboxesControl {
	return c.set("deleteApi", value)
}

// 选项删除提示文字
func (c CheckboxesControl) DeleteConfirmText(value string) CheckboxesControl {
	return c.set("deleteConfirmText", value)
}

// 分割符
func (c CheckboxesControl) Delimiter(value string) CheckboxesControl {
	return c.set("delimiter", value)
}

// 描述内容
func (c CheckboxesControl) Desc(value string) CheckboxesControl {
	return c.set("desc", value)
}

// 描述内容，支持 Html 片段
func (c CheckboxesControl) Description(value string) CheckboxesControl {
	return c.set("description", value)
}

// 配置描述上的 className
func (c CheckboxesControl) DescriptionClassName(value string) CheckboxesControl {
	return c.set("descriptionClassName", value)
}

// 是否禁用
func (c CheckboxesControl) Disabled(value bool) CheckboxesControl {
	return c.set("disabled", value)
}

// 是否禁用表达式
func (c CheckboxesControl) DisabledOn(value string) CheckboxesControl {
	return c.set("disabledOn", value)
}

// 编辑时调用的 API
func (c CheckboxesControl) EditApi(value string) CheckboxesControl {
	return c.set("editApi", value)
}

// 选项修改的表单项
func (c CheckboxesControl) EditControls(value string) CheckboxesControl {
	return c.set("editControls", value)
}

// 控制编辑弹框设置项
func (c CheckboxesControl) EditDialog(value string) CheckboxesControl {
	return c.set("editDialog", value)
}

// 是否可以编辑
func (c CheckboxesControl) Editable(value bool) CheckboxesControl {
	return c.set("editable", value)
}

// 编辑器配置
func (c CheckboxesControl) EditorSetting(value string) CheckboxesControl {
	return c.set("editorSetting", value)
}

// 额外的字段名
func (c CheckboxesControl) ExtraName(value string) CheckboxesControl {
	return c.set("extraName", value)
}

// 开启后将选中的选项 value 的值封装为数组
func (c CheckboxesControl) ExtractValue(value bool) CheckboxesControl {
	return c.set("extractValue", value)
}

// 是否隐藏
func (c CheckboxesControl) Hidden(value bool) CheckboxesControl {
	return c.set("hidden", value)
}

// 是否隐藏表达式
func (c CheckboxesControl) HiddenOn(value string) CheckboxesControl {
	return c.set("hiddenOn", value)
}

// 输入提示
func (c CheckboxesControl) Hint(value string) CheckboxesControl {
	return c.set("hint", value)
}

// 当配置为水平布局的时候，用来配置具体的左右分配
func (c CheckboxesControl) Horizontal(value string) CheckboxesControl {
	return c.set("horizontal", value)
}

// 组件唯一 id
func (c CheckboxesControl) Id(value string) CheckboxesControl {
	return c.set("id", value)
}

// 配置 source 接口初始拉不拉取
func (c CheckboxesControl) InitFetch(value bool) CheckboxesControl {
	return c.set("initFetch", value)
}

// 用表达式来配置 source 接口初始要不要拉取
func (c CheckboxesControl) InitFetchOn(value string) CheckboxesControl {
	return c.set("initFetchOn", value)
}

// 表单 control 是否为 inline 模式
func (c CheckboxesControl) Inline(value bool) CheckboxesControl {
	return c.set("inline", value)
}

// 配置 input className
func (c CheckboxesControl) InputClassName(value string) CheckboxesControl {
	return c.set("inputClassName", value)
}

// 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 value 会通过 delimiter 连接起来，否则直接将以数组的形式提交值
func (c CheckboxesControl) JoinValues(value bool) CheckboxesControl {
	return c.set("joinValues", value)
}

// 描述标题
func (c CheckboxesControl) Label(value string) CheckboxesControl {
	return c.set("label", value)
}

// 描述标题 可选值: right | left | top | inherit
func (c CheckboxesControl) LabelAlign(value string) CheckboxesControl {
	return c.set("labelAlign", value)
}

// 配置 label className
func (c CheckboxesControl) LabelClassName(value string) CheckboxesControl {
	return c.set("labelClassName", value)
}

// 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (c CheckboxesControl) LabelRemark(value string) CheckboxesControl {
	return c.set("labelRemark", value)
}

// label自定义宽度，默认单位为px
func (c CheckboxesControl) LabelWidth(value string) CheckboxesControl {
	return c.set("labelWidth", value)
}

// 自定义选项展示
func (c CheckboxesControl) MenuTpl(value string) CheckboxesControl {
	return c.set("menuTpl", value)
}

// 配置当前表单项展示模式
func (c CheckboxesControl) Mode(value string) CheckboxesControl {
	return c.set("mode", value)
}

// 是否为多选模式
func (c CheckboxesControl) Multiple(value bool) CheckboxesControl {
	return c.set("multiple", value)
}

// 字段名，表单提交时的 key
func (c CheckboxesControl) Name(value string) CheckboxesControl {
	return c.set("name", value)
}

// 事件动作配置
func (c CheckboxesControl) OnEvent(value string) CheckboxesControl {
	return c.set("onEvent", value)
}

// 选项集合
func (c CheckboxesControl) Options(value string) CheckboxesControl {
	return c.set("options", value)
}

// 选项值字段
func (c CheckboxesControl) OptionValue(value string) CheckboxesControl {
	return c.set("optionValue", value)
}

// 是否为单选模式
func (c CheckboxesControl) Radio(value bool) CheckboxesControl {
	return c.set("radio", value)
}

// 刷新时是否重新加载选项
func (c CheckboxesControl) Reload(value bool) CheckboxesControl {
	return c.set("reload", value)
}

// 自定义配置的值字段名
func (c CheckboxesControl) Source(value string) CheckboxesControl {
	return c.set("source", value)
}

// 默认值
func (c CheckboxesControl) Value(value string) CheckboxesControl {
	return c.set("value", value)
}

// 初始值
func (c CheckboxesControl) DefaultValue(value string) CheckboxesControl {
	return c.set("defaultValue", value)
}
