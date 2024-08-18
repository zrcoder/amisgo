package comp

// SelectControl 下拉选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/select
//
// @version 6.7.0
type SelectControl Schema

// NewSelectControl 创建一个新的 SelectControl 实例
func NewSelectControl() SelectControl {
	return SelectControl{}.set("type", "select")
}

func (sc SelectControl) set(key string, value interface{}) SelectControl {
	sc[key] = value
	return sc
}

// addApi 添加时调用的接口
func (sc SelectControl) AddApi(value string) SelectControl {
	return sc.set("addApi", value)
}

// addControls 新增时的表单项。
func (sc SelectControl) AddControls(value string) SelectControl {
	return sc.set("addControls", value)
}

// addDialog 控制新增弹框设置项 (控制新增弹框设置项)
func (sc SelectControl) AddDialog(value string) SelectControl {
	return sc.set("addDialog", value)
}

// autoCheckChildren 是否自动选中子节点
func (sc SelectControl) AutoCheckChildren(value string) SelectControl {
	return sc.set("autoCheckChildren", value)
}

// autoComplete 自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。 (自动完成 API，当输入部分文字的时候，会将这些文字通过 ${term} 可以取到，发送给接口。 接口可以返回匹配到的选项，帮助用户输入。)
func (sc SelectControl) AutoComplete(value string) SelectControl {
	return sc.set("autoComplete", value)
}

// autoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (sc SelectControl) AutoFill(value string) SelectControl {
	return sc.set("autoFill", value)
}

// borderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (sc SelectControl) BorderMode(value string) SelectControl {
	return sc.set("borderMode", value)
}

// checkAll 可多选条件下，是否可全选
func (sc SelectControl) CheckAll(value bool) SelectControl {
	return sc.set("checkAll", value)
}

// checkAllLabel 可多选条件下，全选项文案，默认 ”全选“
func (sc SelectControl) CheckAllLabel(value string) SelectControl {
	return sc.set("checkAllLabel", value)
}

// className 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (sc SelectControl) ClassName(value string) SelectControl {
	return sc.set("className", value)
}

// clearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (sc SelectControl) ClearValueOnHidden(value bool) SelectControl {
	return sc.set("clearValueOnHidden", value)
}

// clearable 是否可清除。
func (sc SelectControl) Clearable(value bool) SelectControl {
	return sc.set("clearable", value)
}

// columns 当 selectMode 为 table 时定义表格列信息。
func (sc SelectControl) Columns(value string) SelectControl {
	return sc.set("columns", value)
}

// creatable 是否可以新增
func (sc SelectControl) Creatable(value bool) SelectControl {
	return sc.set("creatable", value)
}

// createBtnLabel 新增文字
func (sc SelectControl) CreateBtnLabel(value string) SelectControl {
	return sc.set("createBtnLabel", value)
}

// defaultCheckAll 可多选条件下，是否默认全选中所有值
func (sc SelectControl) DefaultCheckAll(value bool) SelectControl {
	return sc.set("defaultCheckAll", value)
}

// deferApi 延时加载的 API，当选项中有 defer: true 的选项时，点开会通过此接口扩充。
func (sc SelectControl) DeferApi(value string) SelectControl {
	return sc.set("deferApi", value)
}

// deferField 懒加载字段
func (sc SelectControl) DeferField(value string) SelectControl {
	return sc.set("deferField", value)
}

// deleteApi 选项删除 API
func (sc SelectControl) DeleteApi(value string) SelectControl {
	return sc.set("deleteApi", value)
}

// deleteConfirmText 选项删除提示文字。
func (sc SelectControl) DeleteConfirmText(value string) SelectControl {
	return sc.set("deleteConfirmText", value)
}

// delimiter 分割符
func (sc SelectControl) Delimiter(value string) SelectControl {
	return sc.set("delimiter", value)
}

// desc
func (sc SelectControl) Desc(value string) SelectControl {
	return sc.set("desc", value)
}

// description 描述内容，支持 Html 片段。
func (sc SelectControl) Description(value string) SelectControl {
	return sc.set("description", value)
}

// descriptionClassName 配置描述上的 className (配置描述上的 className)
func (sc SelectControl) DescriptionClassName(value string) SelectControl {
	return sc.set("descriptionClassName", value)
}

// disabled 是否禁用
func (sc SelectControl) Disabled(value bool) SelectControl {
	return sc.set("disabled", value)
}

// disabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (sc SelectControl) DisabledOn(value string) SelectControl {
	return sc.set("disabledOn", value)
}

// editApi 编辑时调用的 API
func (sc SelectControl) EditApi(value string) SelectControl {
	return sc.set("editApi", value)
}

// editControls 选项修改的表单项
func (sc SelectControl) EditControls(value string) SelectControl {
	return sc.set("editControls", value)
}

// editDialog 控制编辑弹框设置项 (控制编辑弹框设置项)
func (sc SelectControl) EditDialog(value string) SelectControl {
	return sc.set("editDialog", value)
}

// editable 是否可以编辑
func (sc SelectControl) Editable(value bool) SelectControl {
	return sc.set("editable", value)
}

// editorSetting 编辑器配置，运行时可以忽略
func (sc SelectControl) EditorSetting(value string) SelectControl {
	return sc.set("editorSetting", value)
}

// extraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (sc SelectControl) ExtraName(value string) SelectControl {
	return sc.set("extraName", value)
}

// extractValue 开启后将选中的选项 value 的值封装为数组，作为当前表单项的值。
func (sc SelectControl) ExtractValue(value bool) SelectControl {
	return sc.set("extractValue", value)
}

// hidden 是否隐藏
func (sc SelectControl) Hidden(value bool) SelectControl {
	return sc.set("hidden", value)
}

// hiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (sc SelectControl) HiddenOn(value string) SelectControl {
	return sc.set("hiddenOn", value)
}

// hint 输入提示，聚焦的时候显示
func (sc SelectControl) Hint(value string) SelectControl {
	return sc.set("hint", value)
}

// horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (sc SelectControl) Horizontal(value string) SelectControl {
	return sc.set("horizontal", value)
}

// id 组件唯一 id，主要用于日志采集
func (sc SelectControl) Id(value string) SelectControl {
	return sc.set("id", value)
}

// initAutoFill
func (sc SelectControl) InitAutoFill(value string) SelectControl {
	return sc.set("initAutoFill", value)
}

// initFetch 配置 source 接口初始拉不拉取。
func (sc SelectControl) InitFetch(value bool) SelectControl {
	return sc.set("initFetch", value)
}

// initFetchOn 用表达式来配置 source 接口初始要不要拉取
func (sc SelectControl) InitFetchOn(value string) SelectControl {
	return sc.set("initFetchOn", value)
}

// inline 表单 control 是否为 inline 模式。
func (sc SelectControl) Inline(value bool) SelectControl {
	return sc.set("inline", value)
}

// inputClassName 配置 input className (配置 input className)
func (sc SelectControl) InputClassName(value string) SelectControl {
	return sc.set("inputClassName", value)
}

// itemHeight 单个选项的高度，主要用于虚拟渲染
func (sc SelectControl) ItemHeight(value string) SelectControl {
	return sc.set("itemHeight", value)
}

// joinValues 单选模式：当用户选中某个选项时，选项中的 value 将被作为该表单项的值提交，否则，整个选项对象都会作为该表单项的值提交。 多选模式：选中的多个选项的 `value` 会通过 `delimiter` 连接起来，否则直接将以数组的形式提交值。
func (sc SelectControl) JoinValues(value bool) SelectControl {
	return sc.set("joinValues", value)
}

// label 标签文本
func (sc SelectControl) Label(value string) SelectControl {
	return sc.set("label", value)
}

// labelClassName 配置标签上的 className (配置标签上的 className)
func (sc SelectControl) LabelClassName(value string) SelectControl {
	return sc.set("labelClassName", value)
}

// labelRemark 标签备注
func (sc SelectControl) LabelRemark(value string) SelectControl {
	return sc.set("labelRemark", value)
}

// matchObj
func (sc SelectControl) MatchObj(value string) SelectControl {
	return sc.set("matchObj", value)
}

// multiple 是否支持多选
func (sc SelectControl) Multiple(value bool) SelectControl {
	return sc.set("multiple", value)
}

// name 表单项的名称
func (sc SelectControl) Name(value string) SelectControl {
	return sc.set("name", value)
}

// noneOptions 对于 `autoComplete` API 提供的值不被选择项本身支持的时候，用于配置没有值时的提示文字
func (sc SelectControl) NoneOptions(value string) SelectControl {
	return sc.set("noneOptions", value)
}

// onChange 变更回调
func (sc SelectControl) OnChange(value string) SelectControl {
	return sc.set("onChange", value)
}

// onDialogConfirm 点击弹框确认时的回调
func (sc SelectControl) OnDialogConfirm(value string) SelectControl {
	return sc.set("onDialogConfirm", value)
}

// options 选项集合，主要用来配置下拉列表项。
func (sc SelectControl) Options(value string) SelectControl {
	return sc.set("options", value)
}

// popOverConfig PopOver 配置。
func (sc SelectControl) PopOverConfig(value string) SelectControl {
	return sc.set("popOverConfig", value)
}

// reload 时重新加载数据
func (sc SelectControl) Reload(value string) SelectControl {
	return sc.set("reload", value)
}

// resetValue 重置时的值
func (sc SelectControl) ResetValue(value string) SelectControl {
	return sc.set("resetValue", value)
}

// rootClassName 组件外层的 className (组件外层的 className)
func (sc SelectControl) RootClassName(value string) SelectControl {
	return sc.set("rootClassName", value)
}

// searchable 是否支持搜索
func (sc SelectControl) Searchable(value bool) SelectControl {
	return sc.set("searchable", value)
}

// source 数据源
func (sc SelectControl) Source(value string) SelectControl {
	return sc.set("source", value)
}

// sourceClassName 额外的 className 配置到源 (额外的 className 配置到源)
func (sc SelectControl) SourceClassName(value string) SelectControl {
	return sc.set("sourceClassName", value)
}

// sourceApi source 数据源的 API
func (sc SelectControl) SourceApi(value string) SelectControl {
	return sc.set("sourceApi", value)
}

// sourceMessage source 数据源的提示信息
func (sc SelectControl) SourceMessage(value string) SelectControl {
	return sc.set("sourceMessage", value)
}

// staticClassName 类名 (类名)
func (sc SelectControl) StaticClassName(value string) SelectControl {
	return sc.set("staticClassName", value)
}

// submitOnChange 是否在值变更后自动提交
func (sc SelectControl) SubmitOnChange(value bool) SelectControl {
	return sc.set("submitOnChange", value)
}

// type 组件类型
func (sc SelectControl) Type_(value string) SelectControl {
	return sc.set("type", value)
}

// value 绑定值
func (sc SelectControl) Value(value string) SelectControl {
	return sc.set("value", value)
}

// visible 显示条件 (显示条件，语法 `data.xxx > 5`。)
func (sc SelectControl) Visible(value bool) SelectControl {
	return sc.set("visible", value)
}

// visibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (sc SelectControl) VisibleOn(value string) SelectControl {
	return sc.set("visibleOn", value)
}

// virtual 是否虚拟渲染
func (sc SelectControl) Virtual(value bool) SelectControl {
	return sc.set("virtual", value)
}

// width 组件宽度
func (sc SelectControl) Width(value string) SelectControl {
	return sc.set("width", value)
}
