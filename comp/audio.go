package comp

// Audio 音频渲染器
type Audio Schema

// NewAudio 创建一个新的 Audio 实例
func NewAudio() Audio {
	return make(Audio).set("type", "audio")
}

func (a Audio) set(key string, value interface{}) Audio {
	a[key] = value
	return a
}

// AutoPlay 设置是否自动播放
func (a Audio) AutoPlay(value bool) Audio {
	return a.set("autoPlay", value)
}

// ClassName 设置容器 css 类名
func (a Audio) ClassName(value string) Audio {
	return a.set("className", value)
}

// Controls 配置控制器
func (a Audio) Controls(value string) Audio {
	return a.set("controls", value)
}

// Disabled 设置是否禁用
func (a Audio) Disabled(value bool) Audio {
	return a.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (a Audio) DisabledOn(value string) Audio {
	return a.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (a Audio) EditorSetting(value string) Audio {
	return a.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (a Audio) Hidden(value bool) Audio {
	return a.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (a Audio) HiddenOn(value string) Audio {
	return a.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (a Audio) ID(value string) Audio {
	return a.set("id", value)
}

// Inline 设置是否是内联模式
func (a Audio) Inline(value bool) Audio {
	return a.set("inline", value)
}

// Loop 设置是否循环播放
func (a Audio) Loop(value bool) Audio {
	return a.set("loop", value)
}

// OnEvent 设置事件动作配置
func (a Audio) OnEvent(value string) Audio {
	return a.set("onEvent", value)
}

// Rates 配置可选播放倍速
func (a Audio) Rates(value string) Audio {
	return a.set("rates", value)
}

// Src 设置视频播放地址
func (a Audio) Src(value string) Audio {
	return a.set("src", value)
}

// Static 设置是否静态展示
func (a Audio) Static(value bool) Audio {
	return a.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (a Audio) StaticClassName(value string) Audio {
	return a.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (a Audio) StaticInputClassName(value string) Audio {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (a Audio) StaticLabelClassName(value string) Audio {
	return a.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (a Audio) StaticOn(value string) Audio {
	return a.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (a Audio) StaticPlaceholder(value string) Audio {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (a Audio) StaticSchema(value string) Audio {
	return a.set("staticSchema", value)
}

// Style 设置组件样式
func (a Audio) Style(value string) Audio {
	return a.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (a Audio) TestIdBuilder(value string) Audio {
	return a.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (a Audio) Testid(value string) Audio {
	return a.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (a Audio) UseMobileUI(value bool) Audio {
	return a.set("useMobileUI", value)
}

// Visible 设置是否显示
func (a Audio) Visible(value bool) Audio {
	return a.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (a Audio) VisibleOn(value string) Audio {
	return a.set("visibleOn", value)
}
