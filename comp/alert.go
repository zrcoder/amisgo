package comp

import "github.com/zrcoder/amisgo/model"

// alert represents the alert renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/alert

type alert model.Schema

// Alert creates a new Alert instance
func Alert() alert {
	return alert{"type": "alert"}
}

func (a alert) set(key string, value any) alert {
	a[key] = value
	return a
}

// Actions sets the actions area
func (a alert) Actions(value string) alert {
	return a.set("actions", value)
}

// Body sets the content area
func (a alert) Body(value ...any) alert {
	return a.set("body", value)
}

// ClassName sets the container CSS class name
func (a alert) ClassName(value string) alert {
	return a.set("className", value)
}

// CloseButtonClassName sets the CSS class name for the close button
func (a alert) CloseButtonClassName(value string) alert {
	return a.set("closeButtonClassName", value)
}

// Disabled sets whether the alert is disabled
func (a alert) Disabled(value bool) alert {
	return a.set("disabled", value)
}

// DisabledOn sets the expression to determine if the alert is disabled
func (a alert) DisabledOn(value string) alert {
	return a.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (a alert) EditorSetting(value string) alert {
	return a.set("editorSetting", value)
}

// Hidden sets whether the alert is hidden
func (a alert) Hidden(value bool) alert {
	return a.set("hidden", value)
}

// HiddenOn sets the expression to determine if the alert is hidden
func (a alert) HiddenOn(value string) alert {
	return a.set("hiddenOn", value)
}

// Icon sets the icon class name
func (a alert) Icon(value string) alert {
	return a.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (a alert) IconClassName(value string) alert {
	return a.set("iconClassName", value)
}

// ID sets the unique component ID, mainly used for logging
func (a alert) ID(value string) alert {
	return a.set("id", value)
}

// Level sets the alert type. Possible values: info | warning | success | danger
func (a alert) Level(value string) alert {
	return a.set("level", value)
}

// OnEvent sets the event action configuration
func (a alert) OnEvent(value any) alert {
	return a.set("onEvent", value)
}

// ShowCloseButton sets whether to show the close button
func (a alert) ShowCloseButton(value bool) alert {
	return a.set("showCloseButton", value)
}

// ShowIcon sets whether to show the icon
func (a alert) ShowIcon(value bool) alert {
	return a.set("showIcon", value)
}

// Static sets whether the alert is displayed statically
func (a alert) Static(value bool) alert {
	return a.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (a alert) StaticClassName(value string) alert {
	return a.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (a alert) StaticInputClassName(value string) alert {
	return a.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (a alert) StaticLabelClassName(value string) alert {
	return a.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the alert is displayed statically
func (a alert) StaticOn(value string) alert {
	return a.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (a alert) StaticPlaceholder(value string) alert {
	return a.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (a alert) StaticSchema(value string) alert {
	return a.set("staticSchema", value)
}

// Style sets the component style
func (a alert) Style(value any) alert {
	return a.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (a alert) TestIdBuilder(value string) alert {
	return a.set("testIdBuilder", value)
}

// Testid sets the test ID
func (a alert) Testid(value string) alert {
	return a.set("testid", value)
}

// Title sets the alert title
func (a alert) Title(value any) alert {
	return a.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (a alert) UseMobileUI(value bool) alert {
	return a.set("useMobileUI", value)
}

// Visible sets whether the alert is visible
func (a alert) Visible(value bool) alert {
	return a.set("visible", value)
}

// VisibleOn sets the expression to determine if the alert is visible
func (a alert) VisibleOn(value string) alert {
	return a.set("visibleOn", value)
}
