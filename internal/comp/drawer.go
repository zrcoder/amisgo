package comp

import "github.com/zrcoder/amisgo/model"

// Drawer represents a Drawer component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Drawer
type Drawer model.Schema

func NewDrawer() Drawer {
	return Drawer{"type": "drawer"}
}

func (d Drawer) set(key string, value any) Drawer {
	d[key] = value
	return d
}

// Actions sets the actions for the drawer. By default, confirm and cancel buttons are created automatically.
func (d Drawer) Actions(value ...any) Drawer {
	return d.set("actions", value)
}

// Body sets the content area of the drawer.
func (d Drawer) Body(value ...any) Drawer {
	return d.set("body", value)
}

// BodyClassName sets the class name for the body container.
func (d Drawer) BodyClassName(value string) Drawer {
	return d.set("bodyClassName", value)
}

// ClassName sets the class name for the outer container.
func (d Drawer) ClassName(value string) Drawer {
	return d.set("className", value)
}

// CloseOnEsc sets whether the drawer can be closed by pressing the ESC key.
func (d Drawer) CloseOnEsc(value bool) Drawer {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside sets whether the drawer can be closed by clicking outside of it.
func (d Drawer) CloseOnOutside(value bool) Drawer {
	return d.set("closeOnOutside", value)
}

// Confirm sets whether the confirm button is automatically generated. This setting is ignored if custom buttons are configured.
func (d Drawer) Confirm(value bool) Drawer {
	return d.set("confirm", value)
}

// Data sets the data mapping for the drawer.
func (d Drawer) Data(value any) Drawer {
	return d.set("data", value)
}

// Disabled sets whether the drawer is disabled.
func (d Drawer) Disabled(value bool) Drawer {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the drawer is disabled.
func (d Drawer) DisabledOn(value string) Drawer {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (d Drawer) EditorSetting(value string) Drawer {
	return d.set("editorSetting", value)
}

// Footer sets the footer content of the drawer.
func (d Drawer) Footer(value string) Drawer {
	return d.set("footer", value)
}

// FooterClassName sets the class name for the footer container.
func (d Drawer) FooterClassName(value string) Drawer {
	return d.set("footerClassName", value)
}

// Header sets the header content of the drawer.
func (d Drawer) Header(value string) Drawer {
	return d.set("header", value)
}

// HeaderClassName sets the class name for the header container.
func (d Drawer) HeaderClassName(value string) Drawer {
	return d.set("headerClassName", value)
}

// Height sets the height of the drawer. This is effective when the position is top or bottom.
func (d Drawer) Height(value any) Drawer {
	return d.set("height", value)
}

// Hidden sets whether the drawer is hidden.
func (d Drawer) Hidden(value bool) Drawer {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the drawer is hidden.
func (d Drawer) HiddenOn(value string) Drawer {
	return d.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly used for logging.
func (d Drawer) ID(value string) Drawer {
	return d.set("id", value)
}

// InputParams sets the input parameters for the drawer in JSONSchema format.
func (d Drawer) InputParams(value string) Drawer {
	return d.set("inputParams", value)
}

// Name sets the name of the component, which can be used for locating and component communication.
func (d Drawer) Name(value string) Drawer {
	return d.set("name", value)
}

// OnEvent sets the event action configuration for the drawer.
func (d Drawer) OnEvent(value any) Drawer {
	return d.set("onEvent", value)
}

// Overlay sets whether the overlay is displayed.
func (d Drawer) Overlay(value bool) Drawer {
	return d.set("overlay", value)
}

// Position sets the position from which the drawer pops out. Possible values: left, right, top, bottom.
func (d Drawer) Position(value string) Drawer {
	return d.set("position", value)
}

// Resizable sets whether the drawer size can be resized.
func (d Drawer) Resizable(value bool) Drawer {
	return d.set("resizable", value)
}

// ShowCloseButton sets whether the close button is displayed. When false, closeOnOutside is enabled by default.
func (d Drawer) ShowCloseButton(value bool) Drawer {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg sets whether error messages are displayed.
func (d Drawer) ShowErrorMsg(value bool) Drawer {
	return d.set("showErrorMsg", value)
}

// Size sets the size of the drawer. Possible values: xs, sm, md, lg, xl.
func (d Drawer) Size(value string) Drawer {
	return d.set("size", value)
}

// Static sets whether the drawer is displayed statically.
func (d Drawer) Static(value bool) Drawer {
	return d.set("static", value)
}

// StaticClassName sets the class name for static display form items.
func (d Drawer) StaticClassName(value string) Drawer {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values.
func (d Drawer) StaticInputClassName(value string) Drawer {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels.
func (d Drawer) StaticLabelClassName(value string) Drawer {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the drawer is displayed statically.
func (d Drawer) StaticOn(value string) Drawer {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty.
func (d Drawer) StaticPlaceholder(value string) Drawer {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (d Drawer) StaticSchema(value string) Drawer {
	return d.set("staticSchema", value)
}

// Style sets the style for the component.
func (d Drawer) Style(value any) Drawer {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (d Drawer) TestIdBuilder(value string) Drawer {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (d Drawer) Testid(value string) Drawer {
	return d.set("testid", value)
}

// Title sets the title of the drawer.
func (d Drawer) Title(value any) Drawer {
	return d.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level.
func (d Drawer) UseMobileUI(value bool) Drawer {
	return d.set("useMobileUI", value)
}

// Visible sets whether the drawer is visible.
func (d Drawer) Visible(value bool) Drawer {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine whether the drawer is visible.
func (d Drawer) VisibleOn(value string) Drawer {
	return d.set("visibleOn", value)
}

// Width sets the width of the drawer. This is effective when the position is left or right.
func (d Drawer) Width(value any) Drawer {
	return d.set("width", value)
}
