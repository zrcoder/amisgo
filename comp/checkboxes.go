package comp

type checkboxes schema

func Checkboxes() checkboxes {
	return make(checkboxes).set("type", "checkboxes")
}

func (c checkboxes) set(key string, value any) checkboxes {
	c[key] = value
	return c
}

// 添加时调用的接口
func (c checkboxes) AddApi(value string) checkboxes {
	return c.set("addApi", value)
}

// 新增时的表单项
func (c checkboxes) AddControls(value string) checkboxes {
	return c.set("addControls", value)
}

// 控制新增弹框设置项
func (c checkboxes) AddDialog(value string) checkboxes {
	return c.set("addDialog", value)
}

// 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (c checkboxes) AutoFill(value string) checkboxes {
	return c.set("autoFill", value)
}

// 是否开启全选功能
func (c checkboxes) CheckAll(value bool) checkboxes {
	return c.set("checkAll", value)
}

// 全选/不选文案
func (c checkboxes) CheckAllText(value string) checkboxes {
	return c.set("checkAllText", value)
}

// 容器 css 类名
func (c checkboxes) ClassName(value string) checkboxes {
	return c.set("className", value)
}

// 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (c checkboxes) ClearValueOnHidden(value bool) checkboxes {
	return c.set("clearValueOnHidden", value)
}

// 是否可清除
func (c checkboxes) Clearable(value bool) checkboxes {
	return c.set("clearable", value)
}

// 每行显示多少个
func (c checkboxes) ColumnsCount(value string) checkboxes {
	return c.set("columnsCount", value)
}

// 是否可以新增
func (c checkboxes) Creatable(value bool) checkboxes {
	return c.set("creatable", value)
}

// 新增文字
func (c checkboxes) CreateBtnLabel(value string) checkboxes {
	return c.set("createBtnLabel", value)
}

// 是否默认全选
func (c checkboxes) DefaultCheckAll(value bool) checkboxes {
	return c.set("defaultCheckAll", value)
}

// 延时加载的 API
func (c checkboxes) DeferApi(value string) checkboxes {
	return c.set("deferApi", value)
}

// 懒加载字段
func (c checkboxes) DeferField(value string) checkboxes {
	return c.set("deferField", value)
}

// 选项删除 API
func (c checkboxes) DeleteApi(value string) checkboxes {
	return c.set("deleteApi", value)
}

// 选项删除提示文字
func (c checkboxes) DeleteConfirmText(value string) checkboxes {
	return c.set("deleteConfirmText", value)
}

// 分割符
func (c checkboxes) Delimiter(value string) checkboxes {
	return c.set("delimiter", value)
}

// 描述内容
func (c checkboxes) Desc(value string) checkboxes {
	return c.set("desc", value)
}

// 描述内容，支持 Html 片段
func (c checkboxes) Description(value string) checkboxes {
	return c.set("description", value)
}

// 配置描述上的 className
func (c checkboxes) DescriptionClassName(value string) checkboxes {
	return c.set("descriptionClassName", value)
}

// 是否禁用
func (c checkboxes) Disabled(value bool) checkboxes {
	return c.set("disabled", value)
}

// 是否禁用表达式
func (c checkboxes) DisabledOn(value string) checkboxes {
	return c.set("disabledOn", value)
}

// 编辑时调用的 API
func (c checkboxes) EditApi(value string) checkboxes {
	return c.set("editApi", value)
}

// 选项修改的表单项
func (c checkboxes) EditControls(value string) checkboxes {
	return c.set("editControls", value)
}

// 控制编辑弹框设置项
func (c checkboxes) EditDialog(value string) checkboxes {
	return c.set("editDialog", value)
}

// 是否可以编辑
func (c checkboxes) Editable(value bool) checkboxes {
	return c.set("editable", value)
}

// 编辑器配置
func (c checkboxes) EditorSetting(value string) checkboxes {
	return c.set("editorSetting", value)
}

// 额外的字段名
func (c checkboxes) ExtraName(value string) checkboxes {
	return c.set("extraName", value)
}

// 开启后将选中的选项 value 的值封装为数组
func (c checkboxes) ExtractValue(value bool) checkboxes {
	return c.set("extractValue", value)
}

// 是否隐藏
func (c checkboxes) Hidden(value bool) checkboxes {
	return c.set("hidden", value)
}

// 是否隐藏表达式
func (c checkboxes) HiddenOn(value string) checkboxes {
	return c.set("hiddenOn", value)
}

// 输入提示
func (c checkboxes) Hint(value string) checkboxes {
	return c.set("hint", value)
}

// 当配置为水平布局的时候，用来配置具体的左右分配
func (c checkboxes) Horizontal(value string) checkboxes {
	return c.set("horizontal", value)
}

// 组件唯一 id
func (c checkboxes) ID(value string) checkboxes {
	return c.set("id", value)
}

// 配置 source 接口初始拉不拉取
func (c checkboxes) InitFetch(value bool) checkboxes {
	return c.set("initFetch", value)
}

// 用表达式来配置 source 接口初始要不要拉取
func (c checkboxes) InitFetchOn(value string) checkboxes {
	return c.set("initFetchOn", value)
}

// 表单 control 是否为 inline 模式
func (c checkboxes) Inline(value bool) checkboxes {
	return c.set("inline", value)
}

// 配置 input className
func (c checkboxes) InputClassName(value string) checkboxes {
	return c.set("inputClassName", value)
}

// 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 value 会通过 delimiter 连接起来，否则直接将以数组的形式提交值
func (c checkboxes) JoinValues(value bool) checkboxes {
	return c.set("joinValues", value)
}

// 描述标题
func (c checkboxes) Label(value string) checkboxes {
	return c.set("label", value)
}

// 描述标题 可选值: right | left | top | inherit
func (c checkboxes) LabelAlign(value string) checkboxes {
	return c.set("labelAlign", value)
}

// 配置 label className
func (c checkboxes) LabelClassName(value string) checkboxes {
	return c.set("labelClassName", value)
}

// 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (c checkboxes) LabelRemark(value string) checkboxes {
	return c.set("labelRemark", value)
}

// label自定义宽度，默认单位为px
func (c checkboxes) LabelWidth(value string) checkboxes {
	return c.set("labelWidth", value)
}

// 自定义选项展示
func (c checkboxes) MenuTpl(value string) checkboxes {
	return c.set("menuTpl", value)
}

// 配置当前表单项展示模式
func (c checkboxes) Mode(value string) checkboxes {
	return c.set("mode", value)
}

// 是否为多选模式
func (c checkboxes) Multiple(value bool) checkboxes {
	return c.set("multiple", value)
}

// 字段名，表单提交时的 key
func (c checkboxes) Name(value string) checkboxes {
	return c.set("name", value)
}

// 事件动作配置
func (c checkboxes) OnEvent(value any) checkboxes {
	return c.set("onEvent", value)
}

// 选项集合
func (c checkboxes) Options(value ...MOption) checkboxes {
	return c.set("options", value)
}

// 选项值字段
func (c checkboxes) OptionValue(value string) checkboxes {
	return c.set("optionValue", value)
}

// 是否为单选模式
func (c checkboxes) Radio(value bool) checkboxes {
	return c.set("radio", value)
}

// 刷新时是否重新加载选项
func (c checkboxes) Reload(value bool) checkboxes {
	return c.set("reload", value)
}

// 自定义配置的值字段名
func (c checkboxes) Source(value string) checkboxes {
	return c.set("source", value)
}

// 默认值
func (c checkboxes) Value(value any) checkboxes {
	return c.set("value", value)
}

// 初始值
func (c checkboxes) DefaultValue(value string) checkboxes {
	return c.set("defaultValue", value)
}
