package renderers

// Page amis Page 渲染器
type Page struct {
	BaseRenderer
}

// NewPage 创建一个新的 Page 实例
func NewPage() *Page {
	p := &Page{}
	p.set("type", "page")
	return p
}

// Aside 边栏区域
func (p *Page) Aside(value string) *Page {
	p.set("aside", value)
	return p
}

// AsideClassName 边栏区 css 类名
func (p *Page) AsideClassName(value string) *Page {
	p.set("asideClassName", value)
	return p
}

// AsideMaxWidth 边栏最小宽度
func (p *Page) AsideMaxWidth(value string) *Page {
	p.set("asideMaxWidth", value)
	return p
}

// AsideMinWidth 边栏最小宽度
func (p *Page) AsideMinWidth(value string) *Page {
	p.set("asideMinWidth", value)
	return p
}

// AsideResizor 边栏是否允许拖动
func (p *Page) AsideResizor(value bool) *Page {
	p.set("asideResizor", value)
	return p
}

// AsideSticky 边栏内容是否粘住
func (p *Page) AsideSticky(value bool) *Page {
	p.set("asideSticky", value)
	return p
}

// Body 内容区域
func (p *Page) Body(value string) *Page {
	p.set("body", value)
	return p
}

// BodyClassName 内容区 css 类名
func (p *Page) BodyClassName(value string) *Page {
	p.set("bodyClassName", value)
	return p
}

// ClassName 配置容器 className
func (p *Page) ClassName(value string) *Page {
	p.set("className", value)
	return p
}

// Css 自定义页面级别样式表
func (p *Page) Css(value string) *Page {
	p.set("css", value)
	return p
}

// CssVars css 变量
func (p *Page) CssVars(value string) *Page {
	p.set("cssVars", value)
	return p
}

// Data 页面级别的初始数据
func (p *Page) Data(value string) *Page {
	p.set("data", value)
	return p
}

// Definitions 定义
func (p *Page) Definitions(value string) *Page {
	p.set("definitions", value)
	return p
}

// Disabled 是否禁用
func (p *Page) Disabled(value bool) *Page {
	p.set("disabled", value)
	return p
}

// DisabledOn 是否禁用表达式
func (p *Page) DisabledOn(value string) *Page {
	p.set("disabledOn", value)
	return p
}

// EditorSetting 编辑器配置
func (p *Page) EditorSetting(value string) *Page {
	p.set("editorSetting", value)
	return p
}

// HeaderClassName 配置 header 容器 className
func (p *Page) HeaderClassName(value string) *Page {
	p.set("headerClassName", value)
	return p
}

// Hidden 是否隐藏
func (p *Page) Hidden(value bool) *Page {
	p.set("hidden", value)
	return p
}

// HiddenOn 是否隐藏表达式
func (p *Page) HiddenOn(value string) *Page {
	p.set("hiddenOn", value)
	return p
}

// Id 组件唯一 id
func (p *Page) Id(value string) *Page {
	p.set("id", value)
	return p
}

// InitApi 页面初始化 API
func (p *Page) InitApi(value string) *Page {
	p.set("initApi", value)
	return p
}

// InitFetch 是否默认就拉取
func (p *Page) InitFetch(value bool) *Page {
	p.set("initFetch", value)
	return p
}

// InitFetchOn 是否默认就拉取表达式
func (p *Page) InitFetchOn(value string) *Page {
	p.set("initFetchOn", value)
	return p
}

// Interval 配置轮询间隔
func (p *Page) Interval(value string) *Page {
	p.set("interval", value)
	return p
}

// LoadingConfig 加载配置
func (p *Page) LoadingConfig(value string) *Page {
	p.set("loadingConfig", value)
	return p
}

// Messages 消息文案配置
func (p *Page) Messages(value string) *Page {
	p.set("messages", value)
	return p
}

// MobileCSS 移动端下的样式表
func (p *Page) MobileCSS(value string) *Page {
	p.set("mobileCSS", value)
	return p
}

// Name 组件名字
func (p *Page) Name(value string) *Page {
	p.set("name", value)
	return p
}

// OnEvent 事件动作配置
func (p *Page) OnEvent(value string) *Page {
	p.set("onEvent", value)
	return p
}

// PullRefresh 下拉刷新配置
func (p *Page) PullRefresh(value string) *Page {
	p.set("pullRefresh", value)
	return p
}

// Regions 默认不设置自动感觉内容来决定要不要展示这些区域
func (p *Page) Regions(value string) *Page {
	p.set("regions", value)
	return p
}

// Remark 页面描述
func (p *Page) Remark(value string) *Page {
	p.set("remark", value)
	return p
}

// ShowErrorMsg 是否显示错误信息
func (p *Page) ShowErrorMsg(value bool) *Page {
	p.set("showErrorMsg", value)
	return p
}

// SilentPolling 是否要静默加载
func (p *Page) SilentPolling(value bool) *Page {
	p.set("silentPolling", value)
	return p
}

// Static 是否静态展示
func (p *Page) Static(value bool) *Page {
	p.set("static", value)
	return p
}

// StaticClassName 静态展示表单项类名
func (p *Page) StaticClassName(value string) *Page {
	p.set("staticClassName", value)
	return p
}

// StaticInputClassName 静态展示表单项Value类名
func (p *Page) StaticInputClassName(value string) *Page {
	p.set("staticInputClassName", value)
	return p
}

// StaticLabelClassName 静态展示表单项Label类名
func (p *Page) StaticLabelClassName(value string) *Page {
	p.set("staticLabelClassName", value)
	return p
}

// StaticOn 是否静态展示表达式
func (p *Page) StaticOn(value string) *Page {
	p.set("staticOn", value)
	return p
}

// StaticPlaceholder 静态展示空值占位
func (p *Page) StaticPlaceholder(value string) *Page {
	p.set("staticPlaceholder", value)
	return p
}

// StaticSchema 静态展示 schema
func (p *Page) StaticSchema(value string) *Page {
	p.set("staticSchema", value)
	return p
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (p *Page) StopAutoRefreshWhen(value string) *Page {
	p.set("stopAutoRefreshWhen", value)
	return p
}

// Style 自定义样式
func (p *Page) Style(value string) *Page {
	p.set("style", value)
	return p
}

// SubTitle 页面副标题
func (p *Page) SubTitle(value string) *Page {
	p.set("subTitle", value)
	return p
}

// TestIdBuilder 测试 ID 构建器
func (p *Page) TestIdBuilder(value string) *Page {
	p.set("testIdBuilder", value)
	return p
}

// Testid 测试 ID
func (p *Page) Testid(value string) *Page {
	p.set("testid", value)
	return p
}

// Title 页面标题
func (p *Page) Title(value string) *Page {
	p.set("title", value)
	return p
}

// Toolbar 页面顶部区域
func (p *Page) Toolbar(value string) *Page {
	p.set("toolbar", value)
	return p
}

// ToolbarClassName 配置 toolbar 容器 className
func (p *Page) ToolbarClassName(value string) *Page {
	p.set("toolbarClassName", value)
	return p
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (p *Page) UseMobileUI(value bool) *Page {
	p.set("useMobileUI", value)
	return p
}

// Visible 是否显示
func (p *Page) Visible(value bool) *Page {
	p.set("visible", value)
	return p
}

// VisibleOn 是否显示表达式
func (p *Page) VisibleOn(value string) *Page {
	p.set("visibleOn", value)
	return p
}
