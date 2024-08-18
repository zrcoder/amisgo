package comp

// Root amis Page 渲染器。详情请见：https://aisuda.bce.baidu.com/amis/zh-CN/components/page
//
// @author  slowlyo
// @version 6.7.0
type Root Schema

// NewRoot 创建一个新的 Root 实例
func NewRoot() Root {
	return Root{}.set("type", "page")
}

func (r Root) set(key string, value interface{}) Root {
	r[key] = value
	return r
}

// Aside 边栏区域
func (r Root) Aside(value string) Root {
	return r.set("aside", value)
}

// AsideClassName 边栏区 css 类名
func (r Root) AsideClassName(value string) Root {
	return r.set("asideClassName", value)
}

// AsideMaxWidth 边栏最小宽度
func (r Root) AsideMaxWidth(value string) Root {
	return r.set("asideMaxWidth", value)
}

// AsideMinWidth 边栏最小宽度
func (r Root) AsideMinWidth(value string) Root {
	return r.set("asideMinWidth", value)
}

// AsideResizor 边栏是否允许拖动
func (r Root) AsideResizor(value bool) Root {
	return r.set("asideResizor", value)
}

// AsideSticky 边栏内容是否粘住，即不跟随滚动。
func (r Root) AsideSticky(value bool) Root {
	return r.set("asideSticky", value)
}

// Body 内容区域
func (r Root) Body(value ...interface{}) Root {
	return r.set("body", value)
}

// BodyClassName 内容区 css 类名
func (r Root) BodyClassName(value string) Root {
	return r.set("bodyClassName", value)
}

// ClassName 配置容器 className
func (r Root) ClassName(value string) Root {
	return r.set("className", value)
}

// Css 自定义页面级别样式表
func (r Root) Css(value string) Root {
	return r.set("css", value)
}

// CssVars css 变量
func (r Root) CssVars(value string) Root {
	return r.set("cssVars", value)
}

// Data 页面级别的初始数据
func (r Root) Data(value string) Root {
	return r.set("data", value)
}

// Definitions
func (r Root) Definitions(value string) Root {
	return r.set("definitions", value)
}

// Disabled 是否禁用
func (r Root) Disabled(value bool) Root {
	return r.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (r Root) DisabledOn(value string) Root {
	return r.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (r Root) EditorSetting(value string) Root {
	return r.set("editorSetting", value)
}

// HeaderClassName 配置 header 容器 className
func (r Root) HeaderClassName(value string) Root {
	return r.set("headerClassName", value)
}

// Hidden 是否隐藏
func (r Root) Hidden(value bool) Root {
	return r.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (r Root) HiddenOn(value string) Root {
	return r.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (r Root) Id(value string) Root {
	return r.set("id", value)
}

// InitApi 页面初始化的时候，可以设置一个 API 让其取拉取，发送数据会携带当前 data 数据（包含地址栏参数），获取得数据会合并到 data 中，供组件内使用。
func (r Root) InitApi(value string) Root {
	return r.set("initApi", value)
}

// InitFetch 是否默认就拉取？
func (r Root) InitFetch(value bool) Root {
	return r.set("initFetch", value)
}

// InitFetchOn 是否默认就拉取表达式
func (r Root) InitFetchOn(value string) Root {
	return r.set("initFetchOn", value)
}

// Interval 配置轮询间隔，配置后 initApi 将轮询加载。
func (r Root) Interval(value string) Root {
	return r.set("interval", value)
}

// LoadingConfig
func (r Root) LoadingConfig(value string) Root {
	return r.set("loadingConfig", value)
}

// Messages
func (r Root) Messages(value string) Root {
	return r.set("messages", value)
}

// MobileCSS 移动端下的样式表
func (r Root) MobileCSS(value string) Root {
	return r.set("mobileCSS", value)
}

// Name
func (r Root) Name(value string) Root {
	return r.set("name", value)
}

// OnEvent 事件动作配置
func (r Root) OnEvent(value string) Root {
	return r.set("onEvent", value)
}

// PullRefresh 下拉刷新配置
func (r Root) PullRefresh(value string) Root {
	return r.set("pullRefresh", value)
}

// Regions 默认不设置自动感觉内容来决定要不要展示这些区域 如果配置了，以配置为主。
func (r Root) Regions(value string) Root {
	return r.set("regions", value)
}

// Remark 页面描述, 标题旁边会出现个小图标，放上去会显示这个属性配置的内容。
func (r Root) Remark(value string) Root {
	return r.set("remark", value)
}

// ShowErrorMsg 是否显示错误信息，默认是显示的。
func (r Root) ShowErrorMsg(value bool) Root {
	return r.set("showErrorMsg", value)
}

// SilentPolling 是否要静默加载，也就是说不显示进度
func (r Root) SilentPolling(value bool) Root {
	return r.set("silentPolling", value)
}

// Static 是否静态展示
func (r Root) Static(value bool) Root {
	return r.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (r Root) StaticClassName(value string) Root {
	return r.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (r Root) StaticInputClassName(value string) Root {
	return r.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (r Root) StaticLabelClassName(value string) Root {
	return r.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (r Root) StaticOn(value string) Root {
	return r.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (r Root) StaticPlaceholder(value string) Root {
	return r.set("staticPlaceholder", value)
}

// StaticSchema
func (r Root) StaticSchema(value string) Root {
	return r.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件。
func (r Root) StopAutoRefreshWhen(value string) Root {
	return r.set("stopAutoRefreshWhen", value)
}

// Style 自定义样式
func (r Root) Style(value string) Root {
	return r.set("style", value)
}

// SubTitle 页面副标题
func (r Root) SubTitle(value string) Root {
	return r.set("subTitle", value)
}

// TestIdBuilder
func (r Root) TestIdBuilder(value string) Root {
	return r.set("testIdBuilder", value)
}

// Testid
func (r Root) Testid(value string) Root {
	return r.set("testid", value)
}

// Title 页面标题
func (r Root) Title(value string) Root {
	return r.set("title", value)
}

// Toolbar 页面顶部区域，当存在 title 时在右上角显示。
func (r Root) Toolbar(value string) Root {
	return r.set("toolbar", value)
}

// ToolbarClassName 配置 toolbar 容器 className
func (r Root) ToolbarClassName(value string) Root {
	return r.set("toolbarClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (r Root) UseMobileUI(value bool) Root {
	return r.set("useMobileUI", value)
}

// Visible 是否显示
func (r Root) Visible(value bool) Root {
	return r.set("visible", value)
}

// VisibleOn 是否显示表达式
func (r Root) VisibleOn(value string) Root {
	return r.set("visibleOn", value)
}
