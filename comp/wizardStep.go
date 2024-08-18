package comp

// WizardStep 表单向导的步骤
//
// @author  slowlyo
// @version 6.7.0
type WizardStep Schema

// NewWizardStep 创建一个新的 WizardStep 实例
func NewWizardStep() WizardStep {
	return WizardStep{}.set("type", "wizard-step")
}

func (ws WizardStep) set(key string, value interface{}) WizardStep {
	ws[key] = value
	return ws
}

// Actions 按钮集合，会固定在底部显示。
func (ws WizardStep) Actions(value string) WizardStep {
	return ws.set("actions", value)
}

// AffixFooter 是否固定底下的按钮在底部。
func (ws WizardStep) AffixFooter(value bool) WizardStep {
	return ws.set("affixFooter", value)
}

// Api Form 用来保存数据的 api。
func (ws WizardStep) Api(value string) WizardStep {
	return ws.set("api", value)
}

// AsyncApi 设置此属性后，表单提交发送保存接口后，还会继续轮询请求该接口，直到返回 finished 属性为 true 才 结束。
func (ws WizardStep) AsyncApi(value string) WizardStep {
	return ws.set("asyncApi", value)
}

// AutoFocus 是否自动将第一个表单元素聚焦。
func (ws WizardStep) AutoFocus(value bool) WizardStep {
	return ws.set("autoFocus", value)
}

// Body 表单项集合
func (ws WizardStep) Body(value string) WizardStep {
	return ws.set("body", value)
}

// CheckInterval 轮询请求的时间间隔，默认为 3秒。设置 asyncApi 才有效
func (ws WizardStep) CheckInterval(value string) WizardStep {
	return ws.set("checkInterval", value)
}

// ClassName 容器 css 类名
func (ws WizardStep) ClassName(value string) WizardStep {
	return ws.set("className", value)
}

// ClearAfterSubmit 提交后清空表单
func (ws WizardStep) ClearAfterSubmit(value bool) WizardStep {
	return ws.set("clearAfterSubmit", value)
}

// ClearPersistDataAfterSubmit 提交成功后清空本地缓存
func (ws WizardStep) ClearPersistDataAfterSubmit(value bool) WizardStep {
	return ws.set("clearPersistDataAfterSubmit", value)
}

// ColumnCount 表单项显示为几列
func (ws WizardStep) ColumnCount(value string) WizardStep {
	return ws.set("columnCount", value)
}

// Data
func (ws WizardStep) Data(value string) WizardStep {
	return ws.set("data", value)
}

// Debug 是否开启调试，开启后会在顶部实时显示表单项数据。
func (ws WizardStep) Debug(value bool) WizardStep {
	return ws.set("debug", value)
}

// DebugConfig Debug面板配置
func (ws WizardStep) DebugConfig(value string) WizardStep {
	return ws.set("debugConfig", value)
}

// Description 描述
func (ws WizardStep) Description(value string) WizardStep {
	return ws.set("description", value)
}

// Disabled 是否禁用
func (ws WizardStep) Disabled(value bool) WizardStep {
	return ws.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (ws WizardStep) DisabledOn(value string) WizardStep {
	return ws.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ws WizardStep) EditorSetting(value string) WizardStep {
	return ws.set("editorSetting", value)
}

// Feedback Form 也可以配置 feedback。
func (ws WizardStep) Feedback(value string) WizardStep {
	return ws.set("feedback", value)
}

// FieldSet
func (ws WizardStep) FieldSet(value string) WizardStep {
	return ws.set("fieldSet", value)
}

// FinishedField 如果决定结束的字段名不是 `finished` 请设置此属性，比如 `is_success`
func (ws WizardStep) FinishedField(value string) WizardStep {
	return ws.set("finishedField", value)
}

// Hidden 是否隐藏
func (ws WizardStep) Hidden(value bool) WizardStep {
	return ws.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (ws WizardStep) HiddenOn(value string) WizardStep {
	return ws.set("hiddenOn", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (ws WizardStep) Horizontal(value string) WizardStep {
	return ws.set("horizontal", value)
}

// Icon 图标
func (ws WizardStep) Icon(value string) WizardStep {
	return ws.set("icon", value)
}

// Id 组件唯一 id，主要用于日志采集
func (ws WizardStep) Id(value string) WizardStep {
	return ws.set("id", value)
}

// InitApi 用来初始化表单数据
func (ws WizardStep) InitApi(value string) WizardStep {
	return ws.set("initApi", value)
}

// InitAsyncApi Form 用来获取初始数据的 api
func (ws WizardStep) InitAsyncApi(value string) WizardStep {
	return ws.set("initAsyncApi", value)
}

// InitCheckInterval 设置了initAsyncApi以后，默认拉取的时间间隔
func (ws WizardStep) InitCheckInterval(value string) WizardStep {
	return ws.set("initCheckInterval", value)
}

// InitFetch 是否初始加载
func (ws WizardStep) InitFetch(value bool) WizardStep {
	return ws.set("initFetch", value)
}

// InitFetchOn 建议改成 api 的 sendOn 属性。
func (ws WizardStep) InitFetchOn(value string) WizardStep {
	return ws.set("initFetchOn", value)
}

// InitFinishedField 设置了initAsyncApi后，默认会从返回数据的data.finished来判断是否完成，也可以设置成其他的xxx
func (ws WizardStep) InitFinishedField(value string) WizardStep {
	return ws.set("initFinishedField", value)
}

// Interval 设置后将轮询调用 initApi
func (ws WizardStep) Interval(value string) WizardStep {
	return ws.set("interval", value)
}

// Jumpable 是否可直接跳转到该步骤
func (ws WizardStep) Jumpable(value bool) WizardStep {
	return ws.set("jumpable", value)
}

// JumpableOn 通过 JS 表达式来配置当前步骤可否被直接跳转到。
func (ws WizardStep) JumpableOn(value string) WizardStep {
	return ws.set("jumpableOn", value)
}

// Label
func (ws WizardStep) Label(value string) WizardStep {
	return ws.set("label", value)
}

// LabelAlign 表单label的对齐方式
func (ws WizardStep) LabelAlign(value string) WizardStep {
	return ws.set("labelAlign", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (ws WizardStep) LabelWidth(value string) WizardStep {
	return ws.set("labelWidth", value)
}

// Messages 消息文案配置
func (ws WizardStep) Messages(value string) WizardStep {
	return ws.set("messages", value)
}

// Mode 配置表单项默认的展示方式。
func (ws WizardStep) Mode(value string) WizardStep {
	return ws.set("mode", value)
}

// Name 组件名字
func (ws WizardStep) Name(value string) WizardStep {
	return ws.set("name", value)
}

// OnEvent 事件动作配置
func (ws WizardStep) OnEvent(value string) WizardStep {
	return ws.set("onEvent", value)
}

// PanelClassName 配置容器 panel className
func (ws WizardStep) PanelClassName(value string) WizardStep {
	return ws.set("panelClassName", value)
}

// PersistData 是否开启本地缓存
func (ws WizardStep) PersistData(value string) WizardStep {
	return ws.set("persistData", value)
}

// PersistDataKeys 开启本地缓存后限制保存哪些 key
func (ws WizardStep) PersistDataKeys(value string) WizardStep {
	return ws.set("persistDataKeys", value)
}

// PreventEnterSubmit 禁用回车提交
func (ws WizardStep) PreventEnterSubmit(value bool) WizardStep {
	return ws.set("preventEnterSubmit", value)
}

// PrimaryField 设置主键 id
func (ws WizardStep) PrimaryField(value string) WizardStep {
	return ws.set("primaryField", value)
}

// PromptPageLeave 页面离开提示
func (ws WizardStep) PromptPageLeave(value bool) WizardStep {
	return ws.set("promptPageLeave", value)
}

// PromptPageLeaveMessage 具体的提示信息
func (ws WizardStep) PromptPageLeaveMessage(value string) WizardStep {
	return ws.set("promptPageLeaveMessage", value)
}

// Redirect
func (ws WizardStep) Redirect(value string) WizardStep {
	return ws.set("redirect", value)
}

// Reload
func (ws WizardStep) Reload(value string) WizardStep {
	return ws.set("reload", value)
}

// ResetAfterSubmit 提交完后重置表单
func (ws WizardStep) ResetAfterSubmit(value bool) WizardStep {
	return ws.set("resetAfterSubmit", value)
}

// Rules 组合校验规则
func (ws WizardStep) Rules(value string) WizardStep {
	return ws.set("rules", value)
}

// SilentPolling 是否静默拉取
func (ws WizardStep) SilentPolling(value bool) WizardStep {
	return ws.set("silentPolling", value)
}

// Static 展示态时的className
func (ws WizardStep) Static(value bool) WizardStep {
	return ws.set("static", value)
}

// StaticClassName css类名
func (ws WizardStep) StaticClassName(value string) WizardStep {
	return ws.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (ws WizardStep) StaticInputClassName(value string) WizardStep {
	return ws.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (ws WizardStep) StaticLabelClassName(value string) WizardStep {
	return ws.set("staticLabelClassName", value)
}

// StaticOn 表达式，语法 `data.xxx > 5`。
func (ws WizardStep) StaticOn(value string) WizardStep {
	return ws.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ws WizardStep) StaticPlaceholder(value string) WizardStep {
	return ws.set("staticPlaceholder", value)
}

// StaticSchema
func (ws WizardStep) StaticSchema(value string) WizardStep {
	return ws.set("staticSchema", value)
}

// StopAutoRefreshWhen 配置停止轮询的条件
func (ws WizardStep) StopAutoRefreshWhen(value string) WizardStep {
	return ws.set("stopAutoRefreshWhen", value)
}

// Style 组件样式
func (ws WizardStep) Style(value string) WizardStep {
	return ws.set("style", value)
}

// SubTitle 子标题
func (ws WizardStep) SubTitle(value string) WizardStep {
	return ws.set("subTitle", value)
}

// SubmitOnChange 修改的时候是否直接提交表单。
func (ws WizardStep) SubmitOnChange(value bool) WizardStep {
	return ws.set("submitOnChange", value)
}

// SubmitOnInit 表单初始先提交一次
func (ws WizardStep) SubmitOnInit(value bool) WizardStep {
	return ws.set("submitOnInit", value)
}

// SubmitText 默认的提交按钮名称
func (ws WizardStep) SubmitText(value string) WizardStep {
	return ws.set("submitText", value)
}

// Tabs
func (ws WizardStep) Tabs(value string) WizardStep {
	return ws.set("tabs", value)
}

// Target 默认表单提交自己会通过发送 api 保存数据
func (ws WizardStep) Target(value string) WizardStep {
	return ws.set("target", value)
}

// TestIdBuilder
func (ws WizardStep) TestIdBuilder(value string) WizardStep {
	return ws.set("testIdBuilder", value)
}

// Testid
func (ws WizardStep) Testid(value string) WizardStep {
	return ws.set("testid", value)
}

// Title 表单标题
func (ws WizardStep) Title(value string) WizardStep {
	return ws.set("title", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ws WizardStep) UseMobileUI(value bool) WizardStep {
	return ws.set("useMobileUI", value)
}

// Value
func (ws WizardStep) Value(value string) WizardStep {
	return ws.set("value", value)
}

// Visible 是否显示
func (ws WizardStep) Visible(value bool) WizardStep {
	return ws.set("visible", value)
}

// VisibleOn 是否显示表达式
func (ws WizardStep) VisibleOn(value string) WizardStep {
	return ws.set("visibleOn", value)
}

// WrapWithPanel 是否用 panel 包裹起来
func (ws WizardStep) WrapWithPanel(value bool) WizardStep {
	return ws.set("wrapWithPanel", value)
}
