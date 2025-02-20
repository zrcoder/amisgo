package comp

import "github.com/zrcoder/amisgo/schema"

// Panel represents the amis Panel renderer
type Panel schema.Schema

func NewPanel() Panel {
	return Panel{"type": "panel"}
}

// set sets a field value
func (p Panel) set(key string, value any) Panel {
	p[key] = value
	return p
}

// Actions sets the actions
func (p Panel) Actions(value ...Action) Panel {
	return p.set("actions", value)
}

// ActionsClassName sets the actions class name
func (p Panel) ActionsClassName(value string) Panel {
	return p.set("actionsClassName", value)
}

// ActionsControlClassName sets the actions control class name
func (p Panel) ActionsControlClassName(value string) Panel {
	return p.set("actionsControlClassName", value)
}

// AffixFooter sets whether the footer is fixed
func (p Panel) AffixFooter(value bool) Panel {
	return p.set("affixFooter", value)
}

// Body sets the body content
func (p Panel) Body(value ...any) Panel {
	return p.set("body", value)
}

// BodyClassName sets the body container class name
func (p Panel) BodyClassName(value string) Panel {
	return p.set("bodyClassName", value)
}

// BodyControlClassName sets the body control class name
func (p Panel) BodyControlClassName(value string) Panel {
	return p.set("bodyControlClassName", value)
}

// ClassName sets the container CSS class name
func (p Panel) ClassName(value string) Panel {
	return p.set("className", value)
}

// Disabled sets whether the panel is disabled
func (p Panel) Disabled(value bool) Panel {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p Panel) DisabledOn(value string) Panel {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p Panel) EditorSetting(value string) Panel {
	return p.set("editorSetting", value)
}

// Footer sets the footer content
func (p Panel) Footer(value ...any) Panel {
	return p.set("footer", value)
}

// FooterClassName sets the footer container class name
func (p Panel) FooterClassName(value string) Panel {
	return p.set("footerClassName", value)
}

// FooterWrapClassName sets the footer and actions wrapper class name
func (p Panel) FooterWrapClassName(value string) Panel {
	return p.set("footerWrapClassName", value)
}

// Header sets the header content
func (p Panel) Header(value ...any) Panel {
	return p.set("header", value)
}

// HeaderClassName sets the header container class name
func (p Panel) HeaderClassName(value string) Panel {
	return p.set("headerClassName", value)
}

// HeaderControlClassName sets the header control class name
func (p Panel) HeaderControlClassName(value string) Panel {
	return p.set("headerControlClassName", value)
}

// Hidden sets whether the panel is hidden
func (p Panel) Hidden(value bool) Panel {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p Panel) HiddenOn(value string) Panel {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Panel) ID(value string) Panel {
	return p.set("id", value)
}

// OnEvent sets the event configuration
func (p Panel) OnEvent(value Event) Panel {
	return p.set("onEvent", value)
}

// Static sets whether the panel is static
func (p Panel) Static(value bool) Panel {
	return p.set("static", value)
}

// StaticClassName sets the static form item class name
func (p Panel) StaticClassName(value string) Panel {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static form item value class name
func (p Panel) StaticInputClassName(value string) Panel {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static form item label class name
func (p Panel) StaticLabelClassName(value string) Panel {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static expression
func (p Panel) StaticOn(value string) Panel {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (p Panel) StaticPlaceholder(value string) Panel {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (p Panel) StaticSchema(value string) Panel {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p Panel) Style(value any) Panel {
	return p.set("style", value)
}

// SubFormHorizontal sets the horizontal layout ratio for subforms
func (p Panel) SubFormHorizontal(value string) Panel {
	return p.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for subform items
func (p Panel) SubFormMode(value string) Panel {
	return p.set("subFormMode", value)
}

// TestIdBuilder sets the custom test ID builder
func (p Panel) TestIdBuilder(value string) Panel {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Panel) Testid(value string) Panel {
	return p.set("testid", value)
}

// Title sets the panel title
func (p Panel) Title(value ...any) Panel {
	return p.set("title", value)
}

// UseMobileUI sets whether to use mobile UI
func (p Panel) UseMobileUI(value bool) Panel {
	return p.set("useMobileUI", value)
}

// Visible sets whether the panel is visible
func (p Panel) Visible(value bool) Panel {
	return p.set("visible", value)
}

// VisibleOn sets the visible expression
func (p Panel) VisibleOn(value string) Panel {
	return p.set("visibleOn", value)
}
