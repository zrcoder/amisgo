package renderers

// ButtonGroupControl 代表按钮组控件渲染器
type ButtonGroupControl struct {
	*BaseRenderer
}

// NewButtonGroupControl 创建一个新的 ButtonGroupControl 实例
func NewButtonGroupControl() *ButtonGroupControl {
	bgc := &ButtonGroupControl{BaseRenderer: NewBaseRenderer()}
	bgc.set("type", "button-group-select")
	return bgc
}

func (bgc *ButtonGroupControl) set(key string, value interface{}) *ButtonGroupControl {
	bgc.BaseRenderer.set(key, value)
	return bgc
}

// AddApi 设置添加时调用的接口
func (bgc *ButtonGroupControl) AddApi(value string) *ButtonGroupControl {
	return bgc.set("addApi", value)
}

// AddControls 设置新增时的表单项
func (bgc *ButtonGroupControl) AddControls(value string) *ButtonGroupControl {
	return bgc.set("addControls", value)
}

// AddDialog 设置控制新增弹框设置项
func (bgc *ButtonGroupControl) AddDialog(value string) *ButtonGroupControl {
	return bgc.set("addDialog", value)
}

// AutoFill 自动填充
func (bgc *ButtonGroupControl) AutoFill(value string) *ButtonGroupControl {
	return bgc.set("autoFill", value)
}

// BtnActiveClassName 设置按钮激活状态的类名
func (bgc *ButtonGroupControl) BtnActiveClassName(value string) *ButtonGroupControl {
	return bgc.set("btnActiveClassName", value)
}

// BtnActiveLevel 设置按钮选中的样式级别
func (bgc *ButtonGroupControl) BtnActiveLevel(value string) *ButtonGroupControl {
	return bgc.set("btnActiveLevel", value)
}

// BtnClassName 设置按钮的 CSS 类名
func (bgc *ButtonGroupControl) BtnClassName(value string) *ButtonGroupControl {
	return bgc.set("btnClassName", value)
}

// BtnLevel 设置按钮样式级别
func (bgc *ButtonGroupControl) BtnLevel(value string) *ButtonGroupControl {
	return bgc.set("btnLevel", value)
}

// Buttons 设置按钮集合
func (bgc *ButtonGroupControl) Buttons(value string) *ButtonGroupControl {
	return bgc.set("buttons", value)
}

// ClassName 设置容器的 CSS 类名
func (bgc *ButtonGroupControl) ClassName(value string) *ButtonGroupControl {
	return bgc.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除掉该表单项值
func (bgc *ButtonGroupControl) ClearValueOnHidden(value bool) *ButtonGroupControl {
	return bgc.set("clearValueOnHidden", value)
}

// Clearable 设置是否可清除
func (bgc *ButtonGroupControl) Clearable(value bool) *ButtonGroupControl {
	return bgc.set("clearable", value)
}

// Creatable 设置是否可以新增
func (bgc *ButtonGroupControl) Creatable(value bool) *ButtonGroupControl {
	return bgc.set("creatable", value)
}

// CreateBtnLabel 设置新增文字
func (bgc *ButtonGroupControl) CreateBtnLabel(value string) *ButtonGroupControl {
	return bgc.set("createBtnLabel", value)
}

// DeferApi 设置延时加载的 API
func (bgc *ButtonGroupControl) DeferApi(value string) *ButtonGroupControl {
	return bgc.set("deferApi", value)
}

// DeferField 设置懒加载字段
func (bgc *ButtonGroupControl) DeferField(value string) *ButtonGroupControl {
	return bgc.set("deferField", value)
}

// DeleteApi 设置选项删除 API
func (bgc *ButtonGroupControl) DeleteApi(value string) *ButtonGroupControl {
	return bgc.set("deleteApi", value)
}

// DeleteConfirmText 设置选项删除提示文字
func (bgc *ButtonGroupControl) DeleteConfirmText(value string) *ButtonGroupControl {
	return bgc.set("deleteConfirmText", value)
}

// Delimiter 设置分割符
func (bgc *ButtonGroupControl) Delimiter(value string) *ButtonGroupControl {
	return bgc.set("delimiter", value)
}

// Desc 设置描述
func (bgc *ButtonGroupControl) Desc(value string) *ButtonGroupControl {
	return bgc.set("desc", value)
}

// Description 设置描述内容
func (bgc *ButtonGroupControl) Description(value string) *ButtonGroupControl {
	return bgc.set("description", value)
}

// DescriptionClassName 设置描述上的 className
func (bgc *ButtonGroupControl) DescriptionClassName(value string) *ButtonGroupControl {
	return bgc.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (bgc *ButtonGroupControl) Disabled(value bool) *ButtonGroupControl {
	return bgc.set("disabled", value)
}

// DisabledOn 设置通过 JS 表达式来配置禁用状态
func (bgc *ButtonGroupControl) DisabledOn(value string) *ButtonGroupControl {
	return bgc.set("disabledOn", value)
}

// EditApi 设置编辑时调用的 API
func (bgc *ButtonGroupControl) EditApi(value string) *ButtonGroupControl {
	return bgc.set("editApi", value)
}

// EditControls 设置选项修改的表单项
func (bgc *ButtonGroupControl) EditControls(value string) *ButtonGroupControl {
	return bgc.set("editControls", value)
}

// EditDialog 设置控制编辑弹框设置项
func (bgc *ButtonGroupControl) EditDialog(value string) *ButtonGroupControl {
	return bgc.set("editDialog", value)
}

// Editable 设置是否可以编辑
func (bgc *ButtonGroupControl) Editable(value bool) *ButtonGroupControl {
	return bgc.set("editable", value)
}

// EditorSetting 设置编辑器配置
func (bgc *ButtonGroupControl) EditorSetting(value string) *ButtonGroupControl {
	return bgc.set("editorSetting", value)
}

// ExtraName 设置额外的字段名
func (bgc *ButtonGroupControl) ExtraName(value string) *ButtonGroupControl {
	return bgc.set("extraName", value)
}

// ExtractValue 开启后将选中的选项值封装为数组
func (bgc *ButtonGroupControl) ExtractValue(value bool) *ButtonGroupControl {
	return bgc.set("extractValue", value)
}

// Hidden 设置是否隐藏
func (bgc *ButtonGroupControl) Hidden(value bool) *ButtonGroupControl {
	return bgc.set("hidden", value)
}

// HiddenOn 设置是否隐藏的表达式
func (bgc *ButtonGroupControl) HiddenOn(value string) *ButtonGroupControl {
	return bgc.set("hiddenOn", value)
}

// Hint 设置输入提示
func (bgc *ButtonGroupControl) Hint(value string) *ButtonGroupControl {
	return bgc.set("hint", value)
}

// Horizontal 设置水平布局具体的左右分配
func (bgc *ButtonGroupControl) Horizontal(value string) *ButtonGroupControl {
	return bgc.set("horizontal", value)
}

// Id 设置组件唯一 ID
func (bgc *ButtonGroupControl) Id(value string) *ButtonGroupControl {
	return bgc.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (bgc *ButtonGroupControl) InitAutoFill(value string) *ButtonGroupControl {
	return bgc.set("initAutoFill", value)
}

// InitFetch 配置 source 接口初始拉取
func (bgc *ButtonGroupControl) InitFetch(value bool) *ButtonGroupControl {
	return bgc.set("initFetch", value)
}

// InitFetchOn 用表达式来配置 source 接口初始要不要拉取
func (bgc *ButtonGroupControl) InitFetchOn(value string) *ButtonGroupControl {
	return bgc.set("initFetchOn", value)
}

// Inline 设置表单 control 是否为 inline 模式
func (bgc *ButtonGroupControl) Inline(value bool) *ButtonGroupControl {
	return bgc.set("inline", value)
}

// InputClassName 设置 input className
func (bgc *ButtonGroupControl) InputClassName(value string) *ButtonGroupControl {
	return bgc.set("inputClassName", value)
}

// JoinValues 单选模式与多选模式的处理
func (bgc *ButtonGroupControl) JoinValues(value bool) *ButtonGroupControl {
	return bgc.set("joinValues", value)
}

// Label 设置描述标题
func (bgc *ButtonGroupControl) Label(value string) *ButtonGroupControl {
	return bgc.set("label", value)
}

// LabelAlign 设置描述标题对齐方式
func (bgc *ButtonGroupControl) LabelAlign(value string) *ButtonGroupControl {
	return bgc.set("labelAlign", value)
}

// LabelClassName 设置 label className
func (bgc *ButtonGroupControl) LabelClassName(value string) *ButtonGroupControl {
	return bgc.set("labelClassName", value)
}

// LabelRemark 设置描述标题旁的小图标
func (bgc *ButtonGroupControl) LabelRemark(value string) *ButtonGroupControl {
	return bgc.set("labelRemark", value)
}

// LabelWidth 设置 label 宽度
func (bgc *ButtonGroupControl) LabelWidth(value string) *ButtonGroupControl {
	return bgc.set("labelWidth", value)
}

// Mode 设置当前表单项展示模式
func (bgc *ButtonGroupControl) Mode(value string) *ButtonGroupControl {
	return bgc.set("mode", value)
}

// Multiple 设置是否为多选模式
func (bgc *ButtonGroupControl) Multiple(value bool) *ButtonGroupControl {
	return bgc.set("multiple", value)
}

// Name 设置表单字段名
func (bgc *ButtonGroupControl) Name(value string) *ButtonGroupControl {
	return bgc.set("name", value)
}

// OnChange 设置值变化时的回调函数
func (bgc *ButtonGroupControl) OnChange(value string) *ButtonGroupControl {
	return bgc.set("onChange", value)
}

// OnEvent 设置事件回调函数
func (bgc *ButtonGroupControl) OnEvent(value string) *ButtonGroupControl {
	return bgc.set("onEvent", value)
}

// OptionIcon 设置选项图标
func (bgc *ButtonGroupControl) OptionIcon(value string) *ButtonGroupControl {
	return bgc.set("optionIcon", value)
}

// OptionLabel 设置选项标签
func (bgc *ButtonGroupControl) OptionLabel(value string) *ButtonGroupControl {
	return bgc.set("optionLabel", value)
}

// OptionValue 设置选项值
func (bgc *ButtonGroupControl) OptionValue(value string) *ButtonGroupControl {
	return bgc.set("optionValue", value)
}

// Options 设置选项数组
func (bgc *ButtonGroupControl) Options(value string) *ButtonGroupControl {
	return bgc.set("options", value)
}

// Outline 设置是否为轮廓按钮
func (bgc *ButtonGroupControl) Outline(value bool) *ButtonGroupControl {
	return bgc.set("outline", value)
}

// Overlay 设置是否启用 Overlay
func (bgc *ButtonGroupControl) Overlay(value bool) *ButtonGroupControl {
	return bgc.set("overlay", value)
}

// Source 设置源 API
func (bgc *ButtonGroupControl) Source(value string) *ButtonGroupControl {
	return bgc.set("source", value)
}

// Size 设置控件尺寸
func (bgc *ButtonGroupControl) Size(value string) *ButtonGroupControl {
	return bgc.set("size", value)
}

// ValueField 设置值字段名
func (bgc *ButtonGroupControl) ValueField(value string) *ButtonGroupControl {
	return bgc.set("valueField", value)
}

// Validation 设置校验规则
func (bgc *ButtonGroupControl) Validation(value string) *ButtonGroupControl {
	return bgc.set("validation", value)
}

// Visible 设置是否可见
func (bgc *ButtonGroupControl) Visible(value bool) *ButtonGroupControl {
	return bgc.set("visible", value)
}

// VisibleOn 用表达式来配置是否显示
func (bgc *ButtonGroupControl) VisibleOn(value string) *ButtonGroupControl {
	return bgc.set("visibleOn", value)
}
