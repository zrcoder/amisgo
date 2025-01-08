package comp

import "github.com/zrcoder/amisgo/model"

// panel represents the amis panel renderer
type panel model.Schema

// Panel creates a new Panel instance
func Panel() panel {
	return panel{}.set("type", "panel")
}

// set sets a field value
func (p panel) set(key string, value any) panel {
	p[key] = value
	return p
}

// Actions sets the actions
func (p panel) Actions(value ...any) panel {
	return p.set("actions", value)
}

// ActionsClassName sets the actions class name
func (p panel) ActionsClassName(value string) panel {
	return p.set("actionsClassName", value)
}

// ActionsControlClassName sets the actions control class name
func (p panel) ActionsControlClassName(value string) panel {
	return p.set("actionsControlClassName", value)
}

// AffixFooter sets whether the footer is fixed
func (p panel) AffixFooter(value bool) panel {
	return p.set("affixFooter", value)
}

// Body sets the body content
func (p panel) Body(value ...any) panel {
	return p.set("body", value)
}

// BodyClassName sets the body container class name
func (p panel) BodyClassName(value string) panel {
	return p.set("bodyClassName", value)
}

// BodyControlClassName sets the body control class name
func (p panel) BodyControlClassName(value string) panel {
	return p.set("bodyControlClassName", value)
}

// ClassName sets the container CSS class name
func (p panel) ClassName(value string) panel {
	return p.set("className", value)
}

// Disabled sets whether the panel is disabled
func (p panel) Disabled(value bool) panel {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p panel) DisabledOn(value string) panel {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p panel) EditorSetting(value string) panel {
	return p.set("editorSetting", value)
}

// Footer sets the footer content
func (p panel) Footer(value ...any) panel {
	return p.set("footer", value)
}

// FooterClassName sets the footer container class name
func (p panel) FooterClassName(value string) panel {
	return p.set("footerClassName", value)
}

// FooterWrapClassName sets the footer and actions wrapper class name
func (p panel) FooterWrapClassName(value string) panel {
	return p.set("footerWrapClassName", value)
}

// Header sets the header content
func (p panel) Header(value ...any) panel {
	return p.set("header", value)
}

// HeaderClassName sets the header container class name
func (p panel) HeaderClassName(value string) panel {
	return p.set("headerClassName", value)
}

// HeaderControlClassName sets the header control class name
func (p panel) HeaderControlClassName(value string) panel {
	return p.set("headerControlClassName", value)
}

// Hidden sets whether the panel is hidden
func (p panel) Hidden(value bool) panel {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p panel) HiddenOn(value string) panel {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p panel) ID(value string) panel {
	return p.set("id", value)
}

// OnEvent sets the event configuration
func (p panel) OnEvent(value any) panel {
	return p.set("onEvent", value)
}

// Static sets whether the panel is static
func (p panel) Static(value bool) panel {
	return p.set("static", value)
}

// StaticClassName sets the static form item class name
func (p panel) StaticClassName(value string) panel {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static form item value class name
func (p panel) StaticInputClassName(value string) panel {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static form item label class name
func (p panel) StaticLabelClassName(value string) panel {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static expression
func (p panel) StaticOn(value string) panel {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (p panel) StaticPlaceholder(value string) panel {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (p panel) StaticSchema(value string) panel {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p panel) Style(value any) panel {
	return p.set("style", value)
}

// SubFormHorizontal sets the horizontal layout ratio for subforms
func (p panel) SubFormHorizontal(value string) panel {
	return p.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for subform items
func (p panel) SubFormMode(value string) panel {
	return p.set("subFormMode", value)
}

// TestIdBuilder sets the custom test ID builder
func (p panel) TestIdBuilder(value string) panel {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p panel) Testid(value string) panel {
	return p.set("testid", value)
}

// Title sets the panel title
func (p panel) Title(value ...any) panel {
	return p.set("title", value)
}

// UseMobileUI sets whether to use mobile UI
func (p panel) UseMobileUI(value bool) panel {
	return p.set("useMobileUI", value)
}

// Visible sets whether the panel is visible
func (p panel) Visible(value bool) panel {
	return p.set("visible", value)
}

// VisibleOn sets the visible expression
func (p panel) VisibleOn(value string) panel {
	return p.set("visibleOn", value)
}
