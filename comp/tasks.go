package comp

// Tasks 渲染器，格式说明 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tasks
//
// @author  slowlyo
// @version 6.7.0
type Tasks Schema

// NewTasks 创建一个新的 Tasks 实例
func NewTasks() Tasks {
	return Tasks{}.set("type", "tasks")
}

// set 用于设置属性值，返回更新后的 Tasks 实例
func (t Tasks) set(key string, value interface{}) Tasks {
	t[key] = value
	return t
}

// BtnClassName css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     }
func (t Tasks) BtnClassName(value string) Tasks {
	return t.set("btnClassName", value)
}

// BtnText 操作按钮文字
func (t Tasks) BtnText(value string) Tasks {
	return t.set("btnText", value)
}

// CanRetryStatusCode
func (t Tasks) CanRetryStatusCode(value string) Tasks {
	return t.set("canRetryStatusCode", value)
}

// CheckApi 用来获取任务状态的 API，当没有进行时任务时不会发送。 (用来获取任务状态的 API，当没有进行时任务时不会发送。)
func (t Tasks) CheckApi(value string) Tasks {
	return t.set("checkApi", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) ClassName(value string) Tasks {
	return t.set("className", value)
}

// Disabled 是否禁用
func (t Tasks) Disabled(value bool) Tasks {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t Tasks) DisabledOn(value string) Tasks {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t Tasks) EditorSetting(value string) Tasks {
	return t.set("editorSetting", value)
}

// ErrorStatusCode
func (t Tasks) ErrorStatusCode(value string) Tasks {
	return t.set("errorStatusCode", value)
}

// FinishStatusCode
func (t Tasks) FinishStatusCode(value string) Tasks {
	return t.set("finishStatusCode", value)
}

// Hidden 是否隐藏
func (t Tasks) Hidden(value bool) Tasks {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t Tasks) HiddenOn(value string) Tasks {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t Tasks) Id(value string) Tasks {
	return t.set("id", value)
}

// InitialStatusCode
func (t Tasks) InitialStatusCode(value string) Tasks {
	return t.set("initialStatusCode", value)
}

// Interval 当有任务进行中，会每隔一段时间再次检测，而时间间隔就是通过此项配置，默认 3s。
func (t Tasks) Interval(value string) Tasks {
	return t.set("interval", value)
}

// Items
func (t Tasks) Items(value string) Tasks {
	return t.set("items", value)
}

// LoadingConfig
func (t Tasks) LoadingConfig(value string) Tasks {
	return t.set("loadingConfig", value)
}

// LoadingStatusCode
func (t Tasks) LoadingStatusCode(value string) Tasks {
	return t.set("loadingStatusCode", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (t Tasks) Name(value string) Tasks {
	return t.set("name", value)
}

// OnEvent 事件动作配置
func (t Tasks) OnEvent(value string) Tasks {
	return t.set("onEvent", value)
}

// OperationLabel 操作列说明
func (t Tasks) OperationLabel(value string) Tasks {
	return t.set("operationLabel", value)
}

// ReSubmitApi 如果任务失败，且可以重试，提交的时候会使用此 API (如果任务失败，且可以重试，提交的时候会使用此 API)
func (t Tasks) ReSubmitApi(value string) Tasks {
	return t.set("reSubmitApi", value)
}

// ReadyStatusCode
func (t Tasks) ReadyStatusCode(value string) Tasks {
	return t.set("readyStatusCode", value)
}

// RemarkLabel 备注列说明
func (t Tasks) RemarkLabel(value string) Tasks {
	return t.set("remarkLabel", value)
}

// RetryBtnClassName 配置容器重试按钮 className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) RetryBtnClassName(value string) Tasks {
	return t.set("retryBtnClassName", value)
}

// RetryBtnText 重试操作按钮文字
func (t Tasks) RetryBtnText(value string) Tasks {
	return t.set("retryBtnText", value)
}

// Static 是否静态展示
func (t Tasks) Static(value bool) Tasks {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) StaticClassName(value string) Tasks {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) StaticInputClassName(value string) Tasks {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) StaticLabelClassName(value string) Tasks {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t Tasks) StaticOn(value string) Tasks {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Tasks) StaticPlaceholder(value string) Tasks {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t Tasks) StaticSchema(value string) Tasks {
	return t.set("staticSchema", value)
}

// StatusLabel 状态列说明
func (t Tasks) StatusLabel(value string) Tasks {
	return t.set("statusLabel", value)
}

// StatusLabelMap 状态显示对应的类名配置。
func (t Tasks) StatusLabelMap(value string) Tasks {
	return t.set("statusLabelMap", value)
}

// StatusTextMap 状态显示对应的文字显示配置。
func (t Tasks) StatusTextMap(value string) Tasks {
	return t.set("statusTextMap", value)
}

// Style 组件样式
func (t Tasks) Style(value string) Tasks {
	return t.set("style", value)
}

// SubmitApi 提交任务使用的 API (提交任务使用的 API)
func (t Tasks) SubmitApi(value string) Tasks {
	return t.set("submitApi", value)
}

// TableClassName 配置 table className (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (t Tasks) TableClassName(value string) Tasks {
	return t.set("tableClassName", value)
}

// TaskNameLabel 任务名称列说明
func (t Tasks) TaskNameLabel(value string) Tasks {
	return t.set("taskNameLabel", value)
}

// TestIdBuilder
func (t Tasks) TestIdBuilder(value string) Tasks {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Tasks) Testid(value string) Tasks {
	return t.set("testid", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Tasks) UseMobileUI(value bool) Tasks {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Tasks) Visible(value bool) Tasks {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t Tasks) VisibleOn(value string) Tasks {
	return t.set("visibleOn", value)
}
