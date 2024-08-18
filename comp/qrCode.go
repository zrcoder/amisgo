package comp

// QRCode
//
// @version 6.7.0
type QRCode Schema

// NewQRCode 创建一个新的 QRCode 实例
func NewQRCode() QRCode {
	return QRCode{}.
		set("type", "qrcode")
}

func (q QRCode) set(key string, value interface{}) QRCode {
	q[key] = value
	return q
}

// BackgroundColor 背景色
func (q QRCode) BackgroundColor(value string) QRCode {
	return q.set("backgroundColor", value)
}

// ClassName 容器 css 类名
func (q QRCode) ClassName(value string) QRCode {
	return q.set("className", value)
}

// CodeSize 二维码的宽高大小，默认 128
func (q QRCode) CodeSize(value string) QRCode {
	return q.set("codeSize", value)
}

// Disabled 是否禁用
func (q QRCode) Disabled(value bool) QRCode {
	return q.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (q QRCode) DisabledOn(value string) QRCode {
	return q.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (q QRCode) EditorSetting(value string) QRCode {
	return q.set("editorSetting", value)
}

// ForegroundColor 前景色
func (q QRCode) ForegroundColor(value string) QRCode {
	return q.set("foregroundColor", value)
}

// Hidden 是否隐藏
func (q QRCode) Hidden(value bool) QRCode {
	return q.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (q QRCode) HiddenOn(value string) QRCode {
	return q.set("hiddenOn", value)
}

// ID 组件唯一 id
func (q QRCode) ID(value string) QRCode {
	return q.set("id", value)
}

// ImageSettings 图片配置
func (q QRCode) ImageSettings(value string) QRCode {
	return q.set("imageSettings", value)
}

// Level 二维码复杂级别
func (q QRCode) Level(value string) QRCode {
	return q.set("level", value)
}

// Mode 渲染模式
func (q QRCode) Mode(value string) QRCode {
	return q.set("mode", value)
}

// Name 关联字段名
func (q QRCode) Name(value string) QRCode {
	return q.set("name", value)
}

// OnEvent 事件动作配置
func (q QRCode) OnEvent(value string) QRCode {
	return q.set("onEvent", value)
}

// Placeholder 占位符
func (q QRCode) Placeholder(value string) QRCode {
	return q.set("placeholder", value)
}

// QrcodeClassName css 类名
func (q QRCode) QrcodeClassName(value string) QRCode {
	return q.set("qrcodeClassName", value)
}

// Static 是否静态展示
func (q QRCode) Static(value bool) QRCode {
	return q.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (q QRCode) StaticClassName(value string) QRCode {
	return q.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (q QRCode) StaticInputClassName(value string) QRCode {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (q QRCode) StaticLabelClassName(value string) QRCode {
	return q.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (q QRCode) StaticOn(value string) QRCode {
	return q.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (q QRCode) StaticPlaceholder(value string) QRCode {
	return q.set("staticPlaceholder", value)
}

// StaticSchema 静态展示的 Schema
func (q QRCode) StaticSchema(value string) QRCode {
	return q.set("staticSchema", value)
}

// Style 组件样式
func (q QRCode) Style(value string) QRCode {
	return q.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (q QRCode) TestIdBuilder(value string) QRCode {
	return q.set("testIdBuilder", value)
}

// TestID 测试 ID
func (q QRCode) TestID(value string) QRCode {
	return q.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (q QRCode) UseMobileUI(value bool) QRCode {
	return q.set("useMobileUI", value)
}

// Visible 是否显示
func (q QRCode) Visible(value bool) QRCode {
	return q.set("visible", value)
}

// VisibleOn 是否显示表达式
func (q QRCode) VisibleOn(value string) QRCode {
	return q.set("visibleOn", value)
}
