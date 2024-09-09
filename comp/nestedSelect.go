package comp

// nestedSelect 嵌套选择控件

type nestedSelect schema

func NestedSelect() nestedSelect {
	return make(nestedSelect).set("type", "nested-select")
}

// set 设置字段值
func (nsc nestedSelect) set(key string, value any) nestedSelect {
	nsc[key] = value
	return nsc
}

// AddApi 添加时调用的接口
func (nsc nestedSelect) AddApi(value string) nestedSelect {
	return nsc.set("addApi", value)
}

// AddControls 新增时的表单项
func (nsc nestedSelect) AddControls(value string) nestedSelect {
	return nsc.set("addControls", value)
}

// AddDialog 控制新增弹框设置项
func (nsc nestedSelect) AddDialog(value string) nestedSelect {
	return nsc.set("addDialog", value)
}

// AutoFill 自动填充
func (nsc nestedSelect) AutoFill(value string) nestedSelect {
	return nsc.set("autoFill", value)
}

// BorderMode 边框模式
func (nsc nestedSelect) BorderMode(value string) nestedSelect {
	return nsc.set("borderMode", value)
}

// Cascade 父子之间是否完全独立
func (nsc nestedSelect) Cascade(value bool) nestedSelect {
	return nsc.set("cascade", value)
}

// ClassName 容器 css 类名
func (nsc nestedSelect) ClassName(value string) nestedSelect {
	return nsc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (nsc nestedSelect) ClearValueOnHidden(value bool) nestedSelect {
	return nsc.set("clearValueOnHidden", value)
}

// Clearable 是否可清除
func (nsc nestedSelect) Clearable(value bool) nestedSelect {
	return nsc.set("clearable", value)
}

// Creatable 是否可以新增
func (nsc nestedSelect) Creatable(value bool) nestedSelect {
	return nsc.set("creatable", value)
}

// CreateBtnLabel 新增文字
func (nsc nestedSelect) CreateBtnLabel(value string) nestedSelect {
	return nsc.set("createBtnLabel", value)
}

// DeferApi 延时加载的 API
func (nsc nestedSelect) DeferApi(value string) nestedSelect {
	return nsc.set("deferApi", value)
}

// DeferField 懒加载字段
func (nsc nestedSelect) DeferField(value string) nestedSelect {
	return nsc.set("deferField", value)
}

// DeleteApi 选项删除 API
func (nsc nestedSelect) DeleteApi(value string) nestedSelect {
	return nsc.set("deleteApi", value)
}

// DeleteConfirmText 选项删除提示文字
func (nsc nestedSelect) DeleteConfirmText(value string) nestedSelect {
	return nsc.set("deleteConfirmText", value)
}

// Delimiter 分割符
func (nsc nestedSelect) Delimiter(value string) nestedSelect {
	return nsc.set("delimiter", value)
}

// Desc 描述
func (nsc nestedSelect) Desc(value string) nestedSelect {
	return nsc.set("desc", value)
}

// Description 描述内容
func (nsc nestedSelect) Description(value string) nestedSelect {
	return nsc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (nsc nestedSelect) DescriptionClassName(value string) nestedSelect {
	return nsc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (nsc nestedSelect) Disabled(value bool) nestedSelect {
	return nsc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (nsc nestedSelect) DisabledOn(value string) nestedSelect {
	return nsc.set("disabledOn", value)
}

// EditApi 编辑时调用的 API
func (nsc nestedSelect) EditApi(value string) nestedSelect {
	return nsc.set("editApi", value)
}

// EditControls 选项修改的表单项
func (nsc nestedSelect) EditControls(value string) nestedSelect {
	return nsc.set("editControls", value)
}

// EditDialog 控制编辑弹框设置项
func (nsc nestedSelect) EditDialog(value string) nestedSelect {
	return nsc.set("editDialog", value)
}

// Editable 是否可以编辑
func (nsc nestedSelect) Editable(value bool) nestedSelect {
	return nsc.set("editable", value)
}

// EditorSetting 编辑器配置
func (nsc nestedSelect) EditorSetting(value string) nestedSelect {
	return nsc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (nsc nestedSelect) ExtraName(value string) nestedSelect {
	return nsc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项 value 的值封装为数组
func (nsc nestedSelect) ExtractValue(value bool) nestedSelect {
	return nsc.set("extractValue", value)
}

// Hidden 是否隐藏
func (nsc nestedSelect) Hidden(value bool) nestedSelect {
	return nsc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (nsc nestedSelect) HiddenOn(value string) nestedSelect {
	return nsc.set("hiddenOn", value)
}

// HideNodePathLabel 是否隐藏选择框中已选中节点的祖先节点的文本信息
func (nsc nestedSelect) HideNodePathLabel(value bool) nestedSelect {
	return nsc.set("hideNodePathLabel", value)
}

// Hint 输入提示
func (nsc nestedSelect) Hint(value string) nestedSelect {
	return nsc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候
func (nsc nestedSelect) Horizontal(value string) nestedSelect {
	return nsc.set("horizontal", value)
}

// Id 组件唯一 id
func (nsc nestedSelect) Id(value string) nestedSelect {
	return nsc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (nsc nestedSelect) InitAutoFill(value string) nestedSelect {
	return nsc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉不拉取
func (nsc nestedSelect) InitFetch(value bool) nestedSelect {
	return nsc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (nsc nestedSelect) InitFetchOn(value string) nestedSelect {
	return nsc.set("initFetchOn", value)
}

// Inline 表单 control 是否为 inline 模式
func (nsc nestedSelect) Inline(value bool) nestedSelect {
	return nsc.set("inline", value)
}

// InputClassName 配置 input className
func (nsc nestedSelect) InputClassName(value string) nestedSelect {
	return nsc.set("inputClassName", value)
}

// JoinValues 单选模式或多选模式
func (nsc nestedSelect) JoinValues(value bool) nestedSelect {
	return nsc.set("joinValues", value)
}

// Label 描述标题
func (nsc nestedSelect) Label(value string) nestedSelect {
	return nsc.set("label", value)
}

// LabelAlign 描述标题对齐
func (nsc nestedSelect) LabelAlign(value string) nestedSelect {
	return nsc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (nsc nestedSelect) LabelClassName(value string) nestedSelect {
	return nsc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标
func (nsc nestedSelect) LabelRemark(value string) nestedSelect {
	return nsc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (nsc nestedSelect) LabelWidth(value string) nestedSelect {
	return nsc.set("labelWidth", value)
}

// MaxTagCount 标签的最大展示数量
func (nsc nestedSelect) MaxTagCount(value string) nestedSelect {
	return nsc.set("maxTagCount", value)
}

// MenuClassName 弹框的 css 类
func (nsc nestedSelect) MenuClassName(value string) nestedSelect {
	return nsc.set("menuClassName", value)
}

// Mode 配置当前表单项展示模式
func (nsc nestedSelect) Mode(value string) nestedSelect {
	return nsc.set("mode", value)
}

// Multiple 是否为多选模式
func (nsc nestedSelect) Multiple(value bool) nestedSelect {
	return nsc.set("multiple", value)
}

// Name 字段名
func (nsc nestedSelect) Name(value string) nestedSelect {
	return nsc.set("name", value)
}

// OnEvent 事件动作配置
func (nsc nestedSelect) OnEvent(value any) nestedSelect {
	return nsc.set("onEvent", value)
}

// OnlyChildren 选父级的时候是否只把子节点的值包含在内
func (nsc nestedSelect) OnlyChildren(value bool) nestedSelect {
	return nsc.set("onlyChildren", value)
}

// Options 选项
func (nsc nestedSelect) Options(value any) nestedSelect {
	return nsc.set("options", value)
}

// OptionTpl 选项模板
func (nsc nestedSelect) OptionTpl(value string) nestedSelect {
	return nsc.set("optionTpl", value)
}

// Placeholder 占位提示
func (nsc nestedSelect) Placeholder(value string) nestedSelect {
	return nsc.set("placeholder", value)
}

// Remark 备注
func (nsc nestedSelect) Remark(value string) nestedSelect {
	return nsc.set("remark", value)
}

// Render 模板渲染方式
func (nsc nestedSelect) Render(value string) nestedSelect {
	return nsc.set("render", value)
}

// SearchAble 是否可搜索
func (nsc nestedSelect) SearchAble(value bool) nestedSelect {
	return nsc.set("searchAble", value)
}

// ShowDesc 显示描述
func (nsc nestedSelect) ShowDesc(value bool) nestedSelect {
	return nsc.set("showDesc", value)
}

// Size 大小
func (nsc nestedSelect) Size(value string) nestedSelect {
	return nsc.set("size", value)
}

// Static 动态属性
func (nsc nestedSelect) Static(value bool) nestedSelect {
	return nsc.set("static", value)
}

// StaticLabel 静态展示内容
func (nsc nestedSelect) StaticLabel(value string) nestedSelect {
	return nsc.set("staticLabel", value)
}

// Style 控件样式
func (nsc nestedSelect) Style(value any) nestedSelect {
	return nsc.set("style", value)
}

// Tips 提示文字
func (nsc nestedSelect) Tips(value string) nestedSelect {
	return nsc.set("tips", value)
}

// Tree 子节点是否为树形结构
func (nsc nestedSelect) Tree(value bool) nestedSelect {
	return nsc.set("tree", value)
}

// Value 选中值
func (nsc nestedSelect) Value(value any) nestedSelect {
	return nsc.set("value", value)
}

// ValueField 表单项值字段
func (nsc nestedSelect) ValueField(value string) nestedSelect {
	return nsc.set("valueField", value)
}

// Visible 控件是否可见
func (nsc nestedSelect) Visible(value bool) nestedSelect {
	return nsc.set("visible", value)
}

// VisibleOn 是否可见表达式
func (nsc nestedSelect) VisibleOn(value string) nestedSelect {
	return nsc.set("visibleOn", value)
}

// WithChildren 是否包含子节点
func (nsc nestedSelect) WithChildren(value bool) nestedSelect {
	return nsc.set("withChildren", value)
}

// WithLabelChildren 显示标签的子节点
func (nsc nestedSelect) WithLabelChildren(value bool) nestedSelect {
	return nsc.set("withLabelChildren", value)
}

// WithSelect 用于选择节点是否选中
func (nsc nestedSelect) WithSelect(value bool) nestedSelect {
	return nsc.set("withSelect", value)
}
