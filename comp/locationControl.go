package comp

// LocationControl 选点组件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/location
//
// @version 6.7.0
type LocationControl Schema

// NewLocationControl 创建一个新的 LocationControl 实例
func NewLocationControl() LocationControl {
	return LocationControl{}.set("type", "location-picker")
}

func (lc LocationControl) set(key string, value interface{}) LocationControl {
	lc[key] = value
	return lc
}

// Ak 有的地图需要设置 ak 信息
func (lc LocationControl) Ak(value string) LocationControl {
	lc.set("ak", value)
	return lc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (lc LocationControl) AutoFill(value string) LocationControl {
	lc.set("autoFill", value)
	return lc
}

// AutoSelectCurrentLoc 是否自动选中当前地理位置
func (lc LocationControl) AutoSelectCurrentLoc(value bool) LocationControl {
	lc.set("autoSelectCurrentLoc", value)
	return lc
}

// ClassName 容器 css 类名
func (lc LocationControl) ClassName(value string) LocationControl {
	lc.set("className", value)
	return lc
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (lc LocationControl) ClearValueOnHidden(value bool) LocationControl {
	lc.set("clearValueOnHidden", value)
	return lc
}

// Desc 描述内容
func (lc LocationControl) Desc(value string) LocationControl {
	lc.set("desc", value)
	return lc
}

// Description 描述内容，支持 Html 片段
func (lc LocationControl) Description(value string) LocationControl {
	lc.set("description", value)
	return lc
}

// DescriptionClassName 配置描述上的 className
func (lc LocationControl) DescriptionClassName(value string) LocationControl {
	lc.set("descriptionClassName", value)
	return lc
}

// Disabled 是否禁用
func (lc LocationControl) Disabled(value bool) LocationControl {
	lc.set("disabled", value)
	return lc
}

// DisabledOn 是否禁用表达式
func (lc LocationControl) DisabledOn(value string) LocationControl {
	lc.set("disabledOn", value)
	return lc
}

// EditorSetting 编辑器配置，运行时可以忽略
func (lc LocationControl) EditorSetting(value string) LocationControl {
	lc.set("editorSetting", value)
	return lc
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (lc LocationControl) ExtraName(value string) LocationControl {
	lc.set("extraName", value)
	return lc
}

// GetLocationPlaceholder 开启只读模式后的占位提示
func (lc LocationControl) GetLocationPlaceholder(value string) LocationControl {
	lc.set("getLocationPlaceholder", value)
	return lc
}

// Hidden 是否隐藏
func (lc LocationControl) Hidden(value bool) LocationControl {
	lc.set("hidden", value)
	return lc
}

// HiddenOn 是否隐藏表达式
func (lc LocationControl) HiddenOn(value string) LocationControl {
	lc.set("hiddenOn", value)
	return lc
}

// Hint 输入提示，聚焦的时候显示
func (lc LocationControl) Hint(value string) LocationControl {
	lc.set("hint", value)
	return lc
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (lc LocationControl) Horizontal(value string) LocationControl {
	lc.set("horizontal", value)
	return lc
}

// Id 组件唯一 id，主要用于日志采集
func (lc LocationControl) Id(value string) LocationControl {
	lc.set("id", value)
	return lc
}

// InitAutoFill 初始化自动填充
func (lc LocationControl) InitAutoFill(value string) LocationControl {
	lc.set("initAutoFill", value)
	return lc
}

// Inline 表单 control 是否为 inline 模式
func (lc LocationControl) Inline(value bool) LocationControl {
	lc.set("inline", value)
	return lc
}

// InputClassName 配置 input className
func (lc LocationControl) InputClassName(value string) LocationControl {
	lc.set("inputClassName", value)
	return lc
}

// Label 描述标题
func (lc LocationControl) Label(value string) LocationControl {
	lc.set("label", value)
	return lc
}

// LabelAlign 描述标题对齐方式
func (lc LocationControl) LabelAlign(value string) LocationControl {
	lc.set("labelAlign", value)
	return lc
}

// LabelClassName 配置 label className
func (lc LocationControl) LabelClassName(value string) LocationControl {
	lc.set("labelClassName", value)
	return lc
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (lc LocationControl) LabelRemark(value string) LocationControl {
	lc.set("labelRemark", value)
	return lc
}

// LabelWidth label 自定义宽度，默认单位为px
func (lc LocationControl) LabelWidth(value string) LocationControl {
	lc.set("labelWidth", value)
	return lc
}

// Mode 配置当前表单项展示模式
func (lc LocationControl) Mode(value string) LocationControl {
	lc.set("mode", value)
	return lc
}

// Name 字段名，表单提交时的 key
func (lc LocationControl) Name(value string) LocationControl {
	lc.set("name", value)
	return lc
}

// OnEvent 事件动作配置
func (lc LocationControl) OnEvent(value string) LocationControl {
	lc.set("onEvent", value)
	return lc
}

// OnlySelectCurrentLoc 是否限制只能选中当前地理位置
func (lc LocationControl) OnlySelectCurrentLoc(value bool) LocationControl {
	lc.set("onlySelectCurrentLoc", value)
	return lc
}

// Placeholder 占位符
func (lc LocationControl) Placeholder(value string) LocationControl {
	lc.set("placeholder", value)
	return lc
}

// ReadOnly 是否只读
func (lc LocationControl) ReadOnly(value bool) LocationControl {
	lc.set("readOnly", value)
	return lc
}

// ReadOnlyOn 只读条件
func (lc LocationControl) ReadOnlyOn(value string) LocationControl {
	lc.set("readOnlyOn", value)
	return lc
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (lc LocationControl) Remark(value string) LocationControl {
	lc.set("remark", value)
	return lc
}

// Required 是否为必填
func (lc LocationControl) Required(value bool) LocationControl {
	lc.set("required", value)
	return lc
}

// Row 行数
func (lc LocationControl) Row(value string) LocationControl {
	lc.set("row", value)
	return lc
}

// SaveImmediately 是否立即保存
func (lc LocationControl) SaveImmediately(value bool) LocationControl {
	lc.set("saveImmediately", value)
	return lc
}

// Size 表单项大小
func (lc LocationControl) Size(value string) LocationControl {
	lc.set("size", value)
	return lc
}

// Static 是否静态展示
func (lc LocationControl) Static(value bool) LocationControl {
	lc.set("static", value)
	return lc
}

// StaticClassName 静态展示表单项类名
func (lc LocationControl) StaticClassName(value string) LocationControl {
	lc.set("staticClassName", value)
	return lc
}

// StaticInputClassName 静态展示表单项 Value 类名
func (lc LocationControl) StaticInputClassName(value string) LocationControl {
	lc.set("staticInputClassName", value)
	return lc
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (lc LocationControl) StaticLabelClassName(value string) LocationControl {
	lc.set("staticLabelClassName", value)
	return lc
}

// StaticOn 是否静态展示表达式
func (lc LocationControl) StaticOn(value string) LocationControl {
	lc.set("staticOn", value)
	return lc
}

// StaticPlaceholder 静态展示空值占位
func (lc LocationControl) StaticPlaceholder(value string) LocationControl {
	lc.set("staticPlaceholder", value)
	return lc
}

// StaticSchema 静态展示 schema
func (lc LocationControl) StaticSchema(value string) LocationControl {
	lc.set("staticSchema", value)
	return lc
}

// Style 组件样式
func (lc LocationControl) Style(value string) LocationControl {
	lc.set("style", value)
	return lc
}

// SubmitOnChange 当修改完的时候是否提交表单
func (lc LocationControl) SubmitOnChange(value bool) LocationControl {
	lc.set("submitOnChange", value)
	return lc
}

// TestIdBuilder 测试 ID 构建器
func (lc LocationControl) TestIdBuilder(value string) LocationControl {
	lc.set("testIdBuilder", value)
	return lc
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (lc LocationControl) UseMobileUI(value bool) LocationControl {
	lc.set("useMobileUI", value)
	return lc
}

// ValidateApi 远端校验表单项接口
func (lc LocationControl) ValidateApi(value string) LocationControl {
	lc.set("validateApi", value)
	return lc
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (lc LocationControl) ValidateOnChange(value bool) LocationControl {
	lc.set("validateOnChange", value)
	return lc
}

// ValidationErrors 验证失败的提示信息
func (lc LocationControl) ValidationErrors(value string) LocationControl {
	lc.set("validationErrors", value)
	return lc
}

// Validations 验证配置
func (lc LocationControl) Validations(value string) LocationControl {
	lc.set("validations", value)
	return lc
}

// Value 默认值
func (lc LocationControl) Value(value string) LocationControl {
	lc.set("value", value)
	return lc
}

// Vendor 选择地图类型
func (lc LocationControl) Vendor(value string) LocationControl {
	lc.set("vendor", value)
	return lc
}

// Visible 是否显示
func (lc LocationControl) Visible(value bool) LocationControl {
	lc.set("visible", value)
	return lc
}

// VisibleOn 是否显示表达式
func (lc LocationControl) VisibleOn(value string) LocationControl {
	lc.set("visibleOn", value)
	return lc
}

// Width 在 Table 中调整宽度
func (lc LocationControl) Width(value string) LocationControl {
	lc.set("width", value)
	return lc
}
