package comp

// Video 视频播放器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/video
//
// @version 6.7.0
type Video Schema

// NewVideo 创建一个新的 Video 实例
func NewVideo() Video {
	return Video{}.set("type", "video")
}

func (v Video) set(key string, value interface{}) Video {
	v[key] = value
	return v
}

// aspectRatio 视频比率 可选值: auto | 4:3 | 16:9
func (v Video) AspectRatio(value string) Video {
	return v.set("aspectRatio", value)
}

// autoPlay 是否自动播放
func (v Video) AutoPlay(value bool) Video {
	return v.set("autoPlay", value)
}

// className 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) ClassName(value string) Video {
	return v.set("className", value)
}

// columnsCount 如果显示切帧，通过此配置项可以控制每行显示多少帧
func (v Video) ColumnsCount(value string) Video {
	return v.set("columnsCount", value)
}

// disabled 是否禁用
func (v Video) Disabled(value bool) Video {
	return v.set("disabled", value)
}

// disabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (v Video) DisabledOn(value string) Video {
	return v.set("disabledOn", value)
}

// editorSetting 编辑器配置，运行时可以忽略
func (v Video) EditorSetting(value string) Video {
	return v.set("editorSetting", value)
}

// frames 设置后，可以显示切帧.点击帧的时候会将视频跳到对应时间。frames: {  '01:22': 'http://domain/xxx.jpg' }
func (v Video) Frames(value string) Video {
	return v.set("frames", value)
}

// framesClassName 配置帧列表容器className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) FramesClassName(value string) Video {
	return v.set("framesClassName", value)
}

// hidden 是否隐藏
func (v Video) Hidden(value bool) Video {
	return v.set("hidden", value)
}

// hiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (v Video) HiddenOn(value string) Video {
	return v.set("hiddenOn", value)
}

// id 组件唯一 id，主要用于日志采集
func (v Video) Id(value string) Video {
	return v.set("id", value)
}

// isLive 如果是实时的，请标记一下
func (v Video) IsLive(value bool) Video {
	return v.set("isLive", value)
}

// jumpBufferDuration 跳转到帧时，往前多少秒。
func (v Video) JumpBufferDuration(value string) Video {
	return v.set("jumpBufferDuration", value)
}

// jumpFrame 点击帧画面时是否跳转视频对应的点
func (v Video) JumpFrame(value bool) Video {
	return v.set("jumpFrame", value)
}

// loop 是否循环播放
func (v Video) Loop(value bool) Video {
	return v.set("loop", value)
}

// muted 是否初始静音
func (v Video) Muted(value bool) Video {
	return v.set("muted", value)
}

// onEvent 事件动作配置
func (v Video) OnEvent(value string) Video {
	return v.set("onEvent", value)
}

// playerClassName 配置播放器 className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) PlayerClassName(value string) Video {
	return v.set("playerClassName", value)
}

// poster 视频封面地址 (视频封面地址)
func (v Video) Poster(value string) Video {
	return v.set("poster", value)
}

// rates 视频速率
func (v Video) Rates(value string) Video {
	return v.set("rates", value)
}

// splitPoster 是否将视频和封面分开显示
func (v Video) SplitPoster(value bool) Video {
	return v.set("splitPoster", value)
}

// src 视频播放地址 (视频播放地址)
func (v Video) Src(value string) Video {
	return v.set("src", value)
}

// static 是否静态展示
func (v Video) Static(value bool) Video {
	return v.set("static", value)
}

// staticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) StaticClassName(value string) Video {
	return v.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) StaticInputClassName(value string) Video {
	return v.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (v Video) StaticLabelClassName(value string) Video {
	return v.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (v Video) StaticOn(value string) Video {
	return v.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (v Video) StaticPlaceholder(value string) Video {
	return v.set("staticPlaceholder", value)
}

// staticSchema
func (v Video) StaticSchema(value string) Video {
	return v.set("staticSchema", value)
}

// stopOnNextFrame 默认播放的时候到了下一帧会继续播放，同时高亮下一帧。 如果配置这个则会停止播放，等待用户选择新的区间再播放。
func (v Video) StopOnNextFrame(value bool) Video {
	return v.set("stopOnNextFrame", value)
}

// style 组件样式
func (v Video) Style(value string) Video {
	return v.set("style", value)
}

// testIdBuilder
func (v Video) TestIdBuilder(value string) Video {
	return v.set("testIdBuilder", value)
}

// testid
func (v Video) Testid(value string) Video {
	return v.set("testid", value)
}

// type 指定为视频类型
func (v Video) Type(value string) Video {
	return v.set("type", value)
}

// useMobileUI 可以组件级别用来关闭移动端样式
func (v Video) UseMobileUI(value bool) Video {
	return v.set("useMobileUI", value)
}

// videoType 视频类型如： video/x-flv
func (v Video) VideoType(value string) Video {
	return v.set("videoType", value)
}

// visible 是否显示
func (v Video) Visible(value bool) Video {
	return v.set("visible", value)
}

// visibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (v Video) VisibleOn(value string) Video {
	return v.set("visibleOn", value)
}
