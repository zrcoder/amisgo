package comp

// page 代表 amis page 渲染器

type page schema

// Page 创建一个新的 Page 实例
func Page() page {
	return page{}.set("type", "page")
}

// set 用于设置字段值
func (p page) set(key string, value any) page {
	p[key] = value
	return p
}

// Aside 边栏区域
func (p page) Aside(value ...any) page {
	return p.set("aside", value)
}

// AsideClassName 边栏区 css 类名
func (p page) AsideClassName(value string) page {
	return p.set("asideClassName", value)
}

// AsideMaxWidth 边栏最大宽度
func (p page) AsideMaxWidth(value int) page {
	return p.set("asideMaxWidth", value)
}

// AsideMinWidth 边栏最小宽度
func (p page) AsideMinWidth(value int) page {
	return p.set("asideMinWidth", value)
}

// AsideResizor 边栏是否允许拖动
func (p page) AsideResizor(value bool) page {
	return p.set("asideResizor", value)
}

// AsideSticky 边栏内容是否粘住，即不跟随滚动。
func (p page) AsideSticky(value bool) page {
	return p.set("asideSticky", value)
}

// Body 内容区域
func (p page) Body(value ...any) page {
	return p.set("body", value)
}

// BodyClassName 内容区 css 类名
func (p page) BodyClassName(value string) page {
	return p.set("bodyClassName", value)
}

// ClassName 配置容器 className
func (p page) ClassName(value string) page {
	return p.set("className", value)
}

// CSS 自定义页面级别样式表
func (p page) CSS(value any) page {
	return p.set("css", value)
}

// CSSVars css 变量
func (p page) CSSVars(value any) page {
	return p.set("cssVars", value)
}

// Data 页面级别的初始数据
func (p page) Data(value Data) page {
	return p.set("data", value)
}

// Definitions 配置定义
func (p page) Definitions(value string) page {
	return p.set("definitions", value)
}

// Disabled 是否禁用
func (p page) Disabled(value bool) page {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p page) DisabledOn(value string) page {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p page) EditorSetting(value string) page {
	return p.set("editorSetting", value)
}

// HeaderClassName 配置 header 容器 className
func (p page) HeaderClassName(value string) page {
	return p.set("headerClassName", value)
}

// Hidden 是否隐藏
func (p page) Hidden(value bool) page {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p page) HiddenOn(value string) page {
	return p.set("hiddenOn", value)
}

// ID 组件唯一 id
func (p page) ID(value string) page {
	return p.set("id", value)
}

// InitApi 页面初始化 API
func (p page) InitApi(value string) page {
	return p.set("initApi", value)
}

// InitData
func (p page) InitData(getter func() (any, error)) page {
	return p.InitApi(serveData(getter))
}

// InitFetch 是否默认就拉取
func (p page) InitFetch(value bool) page {
	return p.set("initFetch", value)
}

// InitFetchOn 是否默认就拉取表达式
func (p page) InitFetchOn(value string) page {
	return p.set("initFetchOn", value)
}

// Interval 刷新时间(ms, 最小 1000)
func (p page) Interval(value int) page {
	return p.set("interval", value)
}

// LoadingConfig 加载配置
func (p page) LoadingConfig(value string) page {
	return p.set("loadingConfig", value)
}

// Messages 消息文案配置
func (p page) Messages(value string) page {
	return p.set("messages", value)
}

// MobileCSS 移动端下的样式表
func (p page) MobileCSS(value any) page {
	return p.set("mobileCSS", value)
}

// Name 组件名字
func (p page) Name(value string) page {
	return p.set("name", value)
}

// OnEvent 事件动作配置
func (p page) OnEvent(value any) page {
	return p.set("onEvent", value)
}

// PullRefresh 下拉刷新配置
func (p page) PullRefresh(value pullRefresh) page {
	return p.set("pullRefresh", value)
}

// Regions 默认不设置自动感觉内容来决定要不要展示这些区域
func (p page) Regions(value string) page {
	return p.set("regions", value)
}

// Remark 页面描述 string or remark
func (p page) Remark(value any) page {
	return p.set("remark", value)
}

// ShowErrorMsg 是否显示错误信息
func (p page) ShowErrorMsg(value bool) page {
	return p.set("showErrorMsg", value)
}

// SilentPolling 是否要静默加载
func (p page) SilentPolling(value bool) page {
	return p.set("silentPolling", value)
}

// Static 是否静态展示
func (p page) Static(value bool) page {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p page) StaticClassName(value string) page {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (p page) StaticInputClassName(value string) page {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (p page) StaticLabelClassName(value string) page {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p page) StaticOn(value string) page {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p page) StaticPlaceholder(value string) page {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示模式的 schema
func (p page) StaticSchema(value string) page {
	return p.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (p page) StopAutoRefreshWhen(value string) page {
	return p.set("stopAutoRefreshWhen", value)
}

// Style 自定义样式
func (p page) Style(value any) page {
	return p.set("style", value)
}

// SubTitle 页面副标题
func (p page) SubTitle(value any) page {
	return p.set("subTitle", value)
}

// TestIdBuilder 自定义测试 ID 构建器
func (p page) TestIdBuilder(value string) page {
	return p.set("testIdBuilder", value)
}

// Testid 测试 ID
func (p page) Testid(value string) page {
	return p.set("testid", value)
}

// Title 页面标题
func (p page) Title(value any) page {
	return p.set("title", value)
}

// Toolbar 页面顶部区域，当存在 title 时在右上角显示。
func (p page) Toolbar(value ...any) page {
	return p.set("toolbar", value)
}

// ToolbarClassName 配置 toolbar 容器 className
func (p page) ToolbarClassName(value string) page {
	return p.set("toolbarClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p page) UseMobileUI(value bool) page {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p page) Visible(value bool) page {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p page) VisibleOn(value string) page {
	return p.set("visibleOn", value)
}
