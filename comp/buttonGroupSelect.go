package comp

// buttonGroupSelect 代表按钮组控件渲染器
type buttonGroupSelect schema

// ButtonGroupSelect 创建一个新的 ButtonGroupSelect 实例
func ButtonGroupSelect() buttonGroupSelect {
	return make(buttonGroupSelect).set("type", "button-group-select")
}

func (bgc buttonGroupSelect) set(key string, value any) buttonGroupSelect {
	bgc[key] = value
	return bgc
}

// AddApi 设置添加时调用的接口
func (bgc buttonGroupSelect) AddApi(value string) buttonGroupSelect {
	return bgc.set("addApi", value)
}

// AddControls 设置新增时的表单项
func (bgc buttonGroupSelect) AddControls(value string) buttonGroupSelect {
	return bgc.set("addControls", value)
}

// AddDialog 设置控制新增弹框设置项
func (bgc buttonGroupSelect) AddDialog(value string) buttonGroupSelect {
	return bgc.set("addDialog", value)
}

// AutoFill 自动填充
func (bgc buttonGroupSelect) AutoFill(value string) buttonGroupSelect {
	return bgc.set("autoFill", value)
}

// BtnActiveClassName 设置按钮激活状态的类名
func (bgc buttonGroupSelect) BtnActiveClassName(value string) buttonGroupSelect {
	return bgc.set("btnActiveClassName", value)
}

// BtnActiveLevel 设置按钮选中的样式级别
func (bgc buttonGroupSelect) BtnActiveLevel(value string) buttonGroupSelect {
	return bgc.set("btnActiveLevel", value)
}

// BtnClassName 设置按钮的 CSS 类名
func (bgc buttonGroupSelect) BtnClassName(value string) buttonGroupSelect {
	return bgc.set("btnClassName", value)
}

// BtnLevel 设置按钮样式级别
func (bgc buttonGroupSelect) BtnLevel(value string) buttonGroupSelect {
	return bgc.set("btnLevel", value)
}

// Buttons 设置按钮集合
func (bgc buttonGroupSelect) Buttons(value ...any) buttonGroupSelect {
	return bgc.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bgc buttonGroupSelect) ClassName(value string) buttonGroupSelect {
	return bgc.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除掉该表单项值
func (bgc buttonGroupSelect) ClearValueOnHidden(value bool) buttonGroupSelect {
	return bgc.set("clearValueOnHidden", value)
}

// Clearable 设置是否可清除
func (bgc buttonGroupSelect) Clearable(value bool) buttonGroupSelect {
	return bgc.set("clearable", value)
}

// Creatable 设置是否可以新增
func (bgc buttonGroupSelect) Creatable(value bool) buttonGroupSelect {
	return bgc.set("creatable", value)
}

// CreateBtnLabel 设置新增文字
func (bgc buttonGroupSelect) CreateBtnLabel(value string) buttonGroupSelect {
	return bgc.set("createBtnLabel", value)
}

// DeferApi 设置延时加载的 API
func (bgc buttonGroupSelect) DeferApi(value string) buttonGroupSelect {
	return bgc.set("deferApi", value)
}

// DeferField 设置懒加载字段
func (bgc buttonGroupSelect) DeferField(value string) buttonGroupSelect {
	return bgc.set("deferField", value)
}

// DeleteApi 设置选项删除 API
func (bgc buttonGroupSelect) DeleteApi(value string) buttonGroupSelect {
	return bgc.set("deleteApi", value)
}

// DeleteConfirmText 设置选项删除提示文字
func (bgc buttonGroupSelect) DeleteConfirmText(value string) buttonGroupSelect {
	return bgc.set("deleteConfirmText", value)
}

// Delimiter 设置分割符
func (bgc buttonGroupSelect) Delimiter(value string) buttonGroupSelect {
	return bgc.set("delimiter", value)
}

// Desc 设置描述
func (bgc buttonGroupSelect) Desc(value string) buttonGroupSelect {
	return bgc.set("desc", value)
}

// Description 设置描述内容
func (bgc buttonGroupSelect) Description(value string) buttonGroupSelect {
	return bgc.set("description", value)
}

// DescriptionClassName 设置描述上的 className
func (bgc buttonGroupSelect) DescriptionClassName(value string) buttonGroupSelect {
	return bgc.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (bgc buttonGroupSelect) Disabled(value bool) buttonGroupSelect {
	return bgc.set("disabled", value)
}

// DisabledOn 设置通过 JS 表达式来配置禁用状态
func (bgc buttonGroupSelect) DisabledOn(value string) buttonGroupSelect {
	return bgc.set("disabledOn", value)
}

// EditApi 设置编辑时调用的 API
func (bgc buttonGroupSelect) EditApi(value string) buttonGroupSelect {
	return bgc.set("editApi", value)
}

// EditControls 设置选项修改的表单项
func (bgc buttonGroupSelect) EditControls(value string) buttonGroupSelect {
	return bgc.set("editControls", value)
}

// EditDialog 设置控制编辑弹框设置项
func (bgc buttonGroupSelect) EditDialog(value string) buttonGroupSelect {
	return bgc.set("editDialog", value)
}

// Editable 设置是否可以编辑
func (bgc buttonGroupSelect) Editable(value bool) buttonGroupSelect {
	return bgc.set("editable", value)
}

// EditorSetting 设置编辑器配置
func (bgc buttonGroupSelect) EditorSetting(value string) buttonGroupSelect {
	return bgc.set("editorSetting", value)
}

// ExtraName 设置额外的字段名
func (bgc buttonGroupSelect) ExtraName(value string) buttonGroupSelect {
	return bgc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项值封装为数组
func (bgc buttonGroupSelect) ExtractValue(value bool) buttonGroupSelect {
	return bgc.set("extractValue", value)
}

// Hidden 设置是否隐藏
func (bgc buttonGroupSelect) Hidden(value bool) buttonGroupSelect {
	return bgc.set("hidden", value)
}

// HiddenOn 设置是否隐藏的表达式
func (bgc buttonGroupSelect) HiddenOn(value string) buttonGroupSelect {
	return bgc.set("hiddenOn", value)
}

// Hint 设置输入提示
func (bgc buttonGroupSelect) Hint(value string) buttonGroupSelect {
	return bgc.set("hint", value)
}

// Horizontal 设置水平布局具体的左右分配
func (bgc buttonGroupSelect) Horizontal(value string) buttonGroupSelect {
	return bgc.set("horizontal", value)
}

// Id 设置组件唯一 ID
func (bgc buttonGroupSelect) Id(value string) buttonGroupSelect {
	return bgc.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (bgc buttonGroupSelect) InitAutoFill(value string) buttonGroupSelect {
	return bgc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉取
func (bgc buttonGroupSelect) InitFetch(value bool) buttonGroupSelect {
	return bgc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (bgc buttonGroupSelect) InitFetchOn(value string) buttonGroupSelect {
	return bgc.set("initFetchOn", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (bgc buttonGroupSelect) Inline(value bool) buttonGroupSelect {
	return bgc.set("inline", value)
}

// InputClassName 设置 input className
func (bgc buttonGroupSelect) InputClassName(value string) buttonGroupSelect {
	return bgc.set("inputClassName", value)
}

// JoinValues 单选模式与多选模式的处理
func (bgc buttonGroupSelect) JoinValues(value bool) buttonGroupSelect {
	return bgc.set("joinValues", value)
}

// Label 设置描述标题
func (bgc buttonGroupSelect) Label(value string) buttonGroupSelect {
	return bgc.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (bgc buttonGroupSelect) LabelAlign(value string) buttonGroupSelect {
	return bgc.set("labelAlign", value)
}

// LabelClassName 设置 label className
func (bgc buttonGroupSelect) LabelClassName(value string) buttonGroupSelect {
	return bgc.set("labelClassName", value)
}

// LabelRemark 设置描述标题旁的小图标
func (bgc buttonGroupSelect) LabelRemark(value string) buttonGroupSelect {
	return bgc.set("labelRemark", value)
}

// LabelWidth 设置 label 宽度
func (bgc buttonGroupSelect) LabelWidth(value string) buttonGroupSelect {
	return bgc.set("labelWidth", value)
}

// Mode 设置当前表单项展示模式
func (bgc buttonGroupSelect) Mode(value string) buttonGroupSelect {
	return bgc.set("mode", value)
}

// Multiple 设置是否为多选模式
func (bgc buttonGroupSelect) Multiple(value bool) buttonGroupSelect {
	return bgc.set("multiple", value)
}

// Name 设置表单字段名
func (bgc buttonGroupSelect) Name(value string) buttonGroupSelect {
	return bgc.set("name", value)
}

// Value
func (bgc buttonGroupSelect) Value(value any) buttonGroupSelect {
	return bgc.set("value", value)
}

// OnChange 设置值变化时的回调函数
func (bgc buttonGroupSelect) OnChange(value string) buttonGroupSelect {
	return bgc.set("onChange", value)
}

// OnEvent 设置事件回调函数
func (bgc buttonGroupSelect) OnEvent(value any) buttonGroupSelect {
	return bgc.set("onEvent", value)
}

// OptionIcon 设置选项图标
func (bgc buttonGroupSelect) OptionIcon(value string) buttonGroupSelect {
	return bgc.set("optionIcon", value)
}

// OptionLabel 设置选项标签
func (bgc buttonGroupSelect) OptionLabel(value string) buttonGroupSelect {
	return bgc.set("optionLabel", value)
}

// OptionValue 设置选项值
func (bgc buttonGroupSelect) OptionValue(value string) buttonGroupSelect {
	return bgc.set("optionValue", value)
}

// Options 设置选项数组
func (bgc buttonGroupSelect) Options(value ...any) buttonGroupSelect {
	return bgc.set("options", value)
}

// Outline 设置是否为轮廓按钮
func (bgc buttonGroupSelect) Outline(value bool) buttonGroupSelect {
	return bgc.set("outline", value)
}

// Overlay 设置是否启用 Overlay
func (bgc buttonGroupSelect) Overlay(value bool) buttonGroupSelect {
	return bgc.set("overlay", value)
}

// Source 设置源 API
func (bgc buttonGroupSelect) Source(value string) buttonGroupSelect {
	return bgc.set("source", value)
}

// Size 设置控件尺寸
func (bgc buttonGroupSelect) Size(value string) buttonGroupSelect {
	return bgc.set("size", value)
}

// ValueField 设置值字段名
func (bgc buttonGroupSelect) ValueField(value string) buttonGroupSelect {
	return bgc.set("valueField", value)
}

// Validation 设置校验规则
func (bgc buttonGroupSelect) Validation(value string) buttonGroupSelect {
	return bgc.set("validation", value)
}

// Visible 设置是否可见
func (bgc buttonGroupSelect) Visible(value bool) buttonGroupSelect {
	return bgc.set("visible", value)
}

// VisibleOn 用表达式来配置是否显示
func (bgc buttonGroupSelect) VisibleOn(value string) buttonGroupSelect {
	return bgc.set("visibleOn", value)
}
