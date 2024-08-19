package comp

// audio 音频渲染器
type audio schema

func Audio() audio {
	return make(audio).set("type", "audio")
}

func (a audio) set(key string, value any) audio {
	a[key] = value
	return a
}

// AutoPlay 设置是否自动播放
func (a audio) AutoPlay(value bool) audio {
	return a.set("autoPlay", value)
}

// ClassName 设置容器 css 类名
func (a audio) ClassName(value string) audio {
	return a.set("className", value)
}

// Controls 配置控制器
func (a audio) Controls(value string) audio {
	return a.set("controls", value)
}

// Disabled 设置是否禁用
func (a audio) Disabled(value bool) audio {
	return a.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (a audio) DisabledOn(value string) audio {
	return a.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (a audio) EditorSetting(value string) audio {
	return a.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (a audio) Hidden(value bool) audio {
	return a.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (a audio) HiddenOn(value string) audio {
	return a.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (a audio) ID(value string) audio {
	return a.set("id", value)
}

// Inline 设置是否是内联模式
func (a audio) Inline(value bool) audio {
	return a.set("inline", value)
}

// Loop 设置是否循环播放
func (a audio) Loop(value bool) audio {
	return a.set("loop", value)
}

// OnEvent 设置事件动作配置
func (a audio) OnEvent(value any) audio {
	return a.set("onEvent", value)
}

// Rates 配置可选播放倍速
func (a audio) Rates(value string) audio {
	return a.set("rates", value)
}

// Src 设置视频播放地址
func (a audio) Src(value string) audio {
	return a.set("src", value)
}

// Static 设置是否静态展示
func (a audio) Static(value bool) audio {
	return a.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (a audio) StaticClassName(value string) audio {
	return a.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (a audio) StaticInputClassName(value string) audio {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (a audio) StaticLabelClassName(value string) audio {
	return a.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (a audio) StaticOn(value string) audio {
	return a.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (a audio) StaticPlaceholder(value string) audio {
	return a.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (a audio) StaticSchema(value string) audio {
	return a.set("staticSchema", value)
}

// Style 设置组件样式
func (a audio) Style(value any) audio {
	return a.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (a audio) TestIdBuilder(value string) audio {
	return a.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (a audio) Testid(value string) audio {
	return a.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (a audio) UseMobileUI(value bool) audio {
	return a.set("useMobileUI", value)
}

// Visible 设置是否显示
func (a audio) Visible(value bool) audio {
	return a.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (a audio) VisibleOn(value string) audio {
	return a.set("visibleOn", value)
}
