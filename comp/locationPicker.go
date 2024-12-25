package comp

// locationPicker 选点组件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location-picker

type locationPicker Schema

// LocationPicker 创建一个新的 LocationPicker 实例
func LocationPicker() locationPicker {
	return locationPicker{}.set("type", "location-picker")
}

func (lc locationPicker) set(key string, value any) locationPicker {
	lc[key] = value
	return lc
}

// Ak 有的地图需要设置 ak 信息
func (lc locationPicker) Ak(value string) locationPicker {
	lc.set("ak", value)
	return lc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (lc locationPicker) AutoFill(value string) locationPicker {
	lc.set("autoFill", value)
	return lc
}

// AutoSelectCurrentLoc 是否自动选中当前地理位置
func (lc locationPicker) AutoSelectCurrentLoc(value bool) locationPicker {
	lc.set("autoSelectCurrentLoc", value)
	return lc
}

// ClassName 容器 css 类名
func (lc locationPicker) ClassName(value string) locationPicker {
	lc.set("className", value)
	return lc
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (lc locationPicker) ClearValueOnHidden(value bool) locationPicker {
	lc.set("clearValueOnHidden", value)
	return lc
}

// Desc 描述内容
func (lc locationPicker) Desc(value string) locationPicker {
	lc.set("desc", value)
	return lc
}

// Description 描述内容，支持 Html 片段
func (lc locationPicker) Description(value string) locationPicker {
	lc.set("description", value)
	return lc
}

// DescriptionClassName 配置描述上的 className
func (lc locationPicker) DescriptionClassName(value string) locationPicker {
	lc.set("descriptionClassName", value)
	return lc
}

// Disabled 是否禁用
func (lc locationPicker) Disabled(value bool) locationPicker {
	lc.set("disabled", value)
	return lc
}

// DisabledOn 是否禁用表达式
func (lc locationPicker) DisabledOn(value string) locationPicker {
	lc.set("disabledOn", value)
	return lc
}

// EditorSetting 编辑器配置，运行时可以忽略
func (lc locationPicker) EditorSetting(value string) locationPicker {
	lc.set("editorSetting", value)
	return lc
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (lc locationPicker) ExtraName(value string) locationPicker {
	lc.set("extraName", value)
	return lc
}

// GetLocationPlaceholder 开启只读模式后的占位提示
func (lc locationPicker) GetLocationPlaceholder(value string) locationPicker {
	lc.set("getLocationPlaceholder", value)
	return lc
}

// Hidden 是否隐藏
func (lc locationPicker) Hidden(value bool) locationPicker {
	lc.set("hidden", value)
	return lc
}

// HiddenOn 是否隐藏表达式
func (lc locationPicker) HiddenOn(value string) locationPicker {
	lc.set("hiddenOn", value)
	return lc
}

// Hint 输入提示，聚焦的时候显示
func (lc locationPicker) Hint(value string) locationPicker {
	lc.set("hint", value)
	return lc
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (lc locationPicker) Horizontal(value string) locationPicker {
	lc.set("horizontal", value)
	return lc
}

// Id 组件唯一 id，主要用于日志采集
func (lc locationPicker) ID(value string) locationPicker {
	lc.set("id", value)
	return lc
}

// InitAutoFill 初始化自动填充
func (lc locationPicker) InitAutoFill(value string) locationPicker {
	lc.set("initAutoFill", value)
	return lc
}

// Inline 表单 control 是否为 inline 模式
func (lc locationPicker) Inline(value bool) locationPicker {
	lc.set("inline", value)
	return lc
}

// InputClassName 配置 input className
func (lc locationPicker) InputClassName(value string) locationPicker {
	lc.set("inputClassName", value)
	return lc
}

// Label 描述标题
func (lc locationPicker) Label(value string) locationPicker {
	lc.set("label", value)
	return lc
}

// LabelAlign 描述标题对齐方式
func (lc locationPicker) LabelAlign(value string) locationPicker {
	lc.set("labelAlign", value)
	return lc
}

// LabelClassName 配置 label className
func (lc locationPicker) LabelClassName(value string) locationPicker {
	lc.set("labelClassName", value)
	return lc
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (lc locationPicker) LabelRemark(value string) locationPicker {
	lc.set("labelRemark", value)
	return lc
}

// LabelWidth label 自定义宽度，默认单位为px
func (lc locationPicker) LabelWidth(value string) locationPicker {
	lc.set("labelWidth", value)
	return lc
}

// Mode 配置当前表单项展示模式
func (lc locationPicker) Mode(value string) locationPicker {
	lc.set("mode", value)
	return lc
}

// Name 字段名，表单提交时的 key
func (lc locationPicker) Name(value string) locationPicker {
	lc.set("name", value)
	return lc
}

// OnEvent 事件动作配置
func (lc locationPicker) OnEvent(value any) locationPicker {
	lc.set("onEvent", value)
	return lc
}

// OnlySelectCurrentLoc 是否限制只能选中当前地理位置
func (lc locationPicker) OnlySelectCurrentLoc(value bool) locationPicker {
	lc.set("onlySelectCurrentLoc", value)
	return lc
}

// Placeholder 占位符
func (lc locationPicker) Placeholder(value string) locationPicker {
	lc.set("placeholder", value)
	return lc
}

// ReadOnly 是否只读
func (lc locationPicker) ReadOnly(value bool) locationPicker {
	lc.set("readOnly", value)
	return lc
}

// ReadOnlyOn 只读条件
func (lc locationPicker) ReadOnlyOn(value string) locationPicker {
	lc.set("readOnlyOn", value)
	return lc
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (lc locationPicker) Remark(value string) locationPicker {
	lc.set("remark", value)
	return lc
}

// Required 是否为必填
func (lc locationPicker) Required(value bool) locationPicker {
	lc.set("required", value)
	return lc
}

// Row 行数
func (lc locationPicker) Row(value string) locationPicker {
	lc.set("row", value)
	return lc
}

// SaveImmediately 是否立即保存
func (lc locationPicker) SaveImmediately(value bool) locationPicker {
	lc.set("saveImmediately", value)
	return lc
}

// Size 表单项大小
func (lc locationPicker) Size(value string) locationPicker {
	lc.set("size", value)
	return lc
}

// Static 是否静态展示
func (lc locationPicker) Static(value bool) locationPicker {
	lc.set("static", value)
	return lc
}

// StaticClassName 静态展示表单项类名
func (lc locationPicker) StaticClassName(value string) locationPicker {
	lc.set("staticClassName", value)
	return lc
}

// StaticInputClassName 静态展示表单项 Value 类名
func (lc locationPicker) StaticInputClassName(value string) locationPicker {
	lc.set("staticInputClassName", value)
	return lc
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (lc locationPicker) StaticLabelClassName(value string) locationPicker {
	lc.set("staticLabelClassName", value)
	return lc
}

// StaticOn 是否静态展示表达式
func (lc locationPicker) StaticOn(value string) locationPicker {
	lc.set("staticOn", value)
	return lc
}

// StaticPlaceholder 静态展示空值占位
func (lc locationPicker) StaticPlaceholder(value string) locationPicker {
	lc.set("staticPlaceholder", value)
	return lc
}

// StaticSchema 静态展示 schema
func (lc locationPicker) StaticSchema(value string) locationPicker {
	lc.set("staticSchema", value)
	return lc
}

// Style 组件样式
func (lc locationPicker) Style(value any) locationPicker {
	lc.set("style", value)
	return lc
}

// SubmitOnChange 当修改完的时候是否提交表单
func (lc locationPicker) SubmitOnChange(value bool) locationPicker {
	lc.set("submitOnChange", value)
	return lc
}

// TestIdBuilder 测试 ID 构建器
func (lc locationPicker) TestIdBuilder(value string) locationPicker {
	lc.set("testIdBuilder", value)
	return lc
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (lc locationPicker) UseMobileUI(value bool) locationPicker {
	lc.set("useMobileUI", value)
	return lc
}

// ValidateApi 远端校验表单项接口
func (lc locationPicker) ValidateApi(value string) locationPicker {
	lc.set("validateApi", value)
	return lc
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (lc locationPicker) ValidateOnChange(value bool) locationPicker {
	lc.set("validateOnChange", value)
	return lc
}

// ValidationErrors 验证失败的提示信息
func (lc locationPicker) ValidationErrors(value string) locationPicker {
	lc.set("validationErrors", value)
	return lc
}

// Validations 验证配置
func (lc locationPicker) Validations(value string) locationPicker {
	lc.set("validations", value)
	return lc
}

// Value 默认值
func (lc locationPicker) Value(value string) locationPicker {
	lc.set("value", value)
	return lc
}

// Vendor 选择地图类型
func (lc locationPicker) Vendor(value string) locationPicker {
	lc.set("vendor", value)
	return lc
}

// Visible 是否显示
func (lc locationPicker) Visible(value bool) locationPicker {
	lc.set("visible", value)
	return lc
}

// VisibleOn 是否显示表达式
func (lc locationPicker) VisibleOn(value string) locationPicker {
	lc.set("visibleOn", value)
	return lc
}

// Width 在 Table 中调整宽度
func (lc locationPicker) Width(value string) locationPicker {
	lc.set("width", value)
	return lc
}
