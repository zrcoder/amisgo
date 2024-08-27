package comp

// qrCode
//
// @version 6.7.0
type qrCode schema

// QRCode 创建一个新的 QRCode 实例
func QRCode() qrCode {
	return qrCode{}.
		set("type", "qrcode")
}

func (q qrCode) set(key string, value any) qrCode {
	q[key] = value
	return q
}

// BackgroundColor 背景色
func (q qrCode) BackgroundColor(value string) qrCode {
	return q.set("backgroundColor", value)
}

// ClassName 容器 css 类名
func (q qrCode) ClassName(value string) qrCode {
	return q.set("className", value)
}

// CodeSize 二维码的宽高大小，默认 128
func (q qrCode) CodeSize(value int) qrCode {
	return q.set("codeSize", value)
}

// Disabled 是否禁用
func (q qrCode) Disabled(value bool) qrCode {
	return q.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (q qrCode) DisabledOn(value string) qrCode {
	return q.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (q qrCode) EditorSetting(value string) qrCode {
	return q.set("editorSetting", value)
}

// ForegroundColor 前景色
func (q qrCode) ForegroundColor(value string) qrCode {
	return q.set("foregroundColor", value)
}

// Hidden 是否隐藏
func (q qrCode) Hidden(value bool) qrCode {
	return q.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (q qrCode) HiddenOn(value string) qrCode {
	return q.set("hiddenOn", value)
}

// ID 组件唯一 id
func (q qrCode) ID(value string) qrCode {
	return q.set("id", value)
}

// ImageSettings 图片配置
func (q qrCode) ImageSettings(value string) qrCode {
	return q.set("imageSettings", value)
}

// Level 二维码复杂级别 有（'L' 'M' 'Q' 'H'）四种
func (q qrCode) Level(value string) qrCode {
	return q.set("level", value)
}

// Mode 渲染模式
func (q qrCode) Mode(value string) qrCode {
	return q.set("mode", value)
}

// Name 关联字段名
func (q qrCode) Name(value string) qrCode {
	return q.set("name", value)
}

// Value 原始数据, 即扫描二维码后得到的内容
func (q qrCode) Value(value string) qrCode {
	return q.set("value", value)
}

// OnEvent 事件动作配置
func (q qrCode) OnEvent(value any) qrCode {
	return q.set("onEvent", value)
}

// Placeholder 占位符
func (q qrCode) Placeholder(value string) qrCode {
	return q.set("placeholder", value)
}

// QrcodeClassName css 类名
func (q qrCode) QrcodeClassName(value string) qrCode {
	return q.set("qrcodeClassName", value)
}

// Static 是否静态展示
func (q qrCode) Static(value bool) qrCode {
	return q.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (q qrCode) StaticClassName(value string) qrCode {
	return q.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (q qrCode) StaticInputClassName(value string) qrCode {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (q qrCode) StaticLabelClassName(value string) qrCode {
	return q.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (q qrCode) StaticOn(value string) qrCode {
	return q.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (q qrCode) StaticPlaceholder(value string) qrCode {
	return q.set("staticPlaceholder", value)
}

// StaticSchema 静态展示的 Schema
func (q qrCode) StaticSchema(value string) qrCode {
	return q.set("staticSchema", value)
}

// Style 组件样式
func (q qrCode) Style(value any) qrCode {
	return q.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (q qrCode) TestIdBuilder(value string) qrCode {
	return q.set("testIdBuilder", value)
}

// TestID 测试 ID
func (q qrCode) TestID(value string) qrCode {
	return q.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (q qrCode) UseMobileUI(value bool) qrCode {
	return q.set("useMobileUI", value)
}

// Visible 是否显示
func (q qrCode) Visible(value bool) qrCode {
	return q.set("visible", value)
}

// VisibleOn 是否显示表达式
func (q qrCode) VisibleOn(value string) qrCode {
	return q.set("visibleOn", value)
}
