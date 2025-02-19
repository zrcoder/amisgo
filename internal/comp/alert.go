package comp

import "github.com/zrcoder/amisgo/schema"

// alert represents the alert renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/alert

type Alert schema.Schema

func NewAlert() Alert {
	return Alert{"type": "alert"}
}

func (a Alert) set(key string, value any) Alert {
	a[key] = value
	return a
}

// Actions sets the actions area
func (a Alert) Actions(value string) Alert {
	return a.set("actions", value)
}

// Body sets the content area
func (a Alert) Body(value ...any) Alert {
	return a.set("body", value)
}

// ClassName sets the container CSS class name
func (a Alert) ClassName(value string) Alert {
	return a.set("className", value)
}

// CloseButtonClassName sets the CSS class name for the close button
func (a Alert) CloseButtonClassName(value string) Alert {
	return a.set("closeButtonClassName", value)
}

// Disabled sets whether the alert is disabled
func (a Alert) Disabled(value bool) Alert {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine if the alert is disabled
func (a Alert) DisabledOn(value string) Alert {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (a Alert) EditorSetting(value string) Alert {
	return a.set("editorSetting", value)
}

// Hidden sets whether the alert is hidden
func (a Alert) Hidden(value bool) Alert {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine if the alert is hidden
func (a Alert) HiddenOn(value string) Alert {
	return a.set("hiddenOn", value)
}

// Icon sets the icon class name
func (a Alert) Icon(value string) Alert {
	return a.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (a Alert) IconClassName(value string) Alert {
	return a.set("iconClassName", value)
}

// ID sets the unique component ID, mainly used for logging
func (a Alert) ID(value string) Alert {
	return a.set("id", value)
}

// Level sets the alert type. Possible values: info | warning | success | danger
func (a Alert) Level(value string) Alert {
	return a.set("level", value)
}

// OnEvent sets the event action configuration
func (a Alert) OnEvent(value any) Alert {
	return a.set("onEvent", value)
}

// ShowCloseButton sets whether to show the close button
func (a Alert) ShowCloseButton(value bool) Alert {
	return a.set("showCloseButton", value)
}

// ShowIcon sets whether to show the icon
func (a Alert) ShowIcon(value bool) Alert {
	return a.set("showIcon", value)
}

// Static sets whether the alert is displayed statically
func (a Alert) Static(value bool) Alert {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (a Alert) StaticClassName(value string) Alert {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (a Alert) StaticInputClassName(value string) Alert {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (a Alert) StaticLabelClassName(value string) Alert {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the alert is displayed statically
func (a Alert) StaticOn(value string) Alert {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (a Alert) StaticPlaceholder(value string) Alert {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (a Alert) StaticSchema(value string) Alert {
	return a.set("staticSchema", value)
}

// Style sets the component style
func (a Alert) Style(value any) Alert {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (a Alert) TestIdBuilder(value string) Alert {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID
func (a Alert) Testid(value string) Alert {
	return a.set("testid", value)
}

// Title sets the alert title
func (a Alert) Title(value any) Alert {
	return a.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (a Alert) UseMobileUI(value bool) Alert {
	return a.set("useMobileUI", value)
}

// Visible sets whether the alert is visible
func (a Alert) Visible(value bool) Alert {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine if the alert is visible
func (a Alert) VisibleOn(value string) Alert {
	return a.set("visibleOn", value)
}
