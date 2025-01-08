package comp

import "github.com/zrcoder/amisgo/model"

// drawer represents a drawer component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/drawer
type drawer model.Schema

func Drawer() drawer {
	return make(drawer).set("type", "drawer")
}

func (d drawer) set(key string, value any) drawer {
	d[key] = value
	return d
}

// Actions sets the actions for the drawer. By default, confirm and cancel buttons are created automatically.
func (d drawer) Actions(value ...any) drawer {
	return d.set("actions", value)
}

// Body sets the content area of the drawer.
func (d drawer) Body(value ...any) drawer {
	return d.set("body", value)
}

// BodyClassName sets the class name for the body container.
func (d drawer) BodyClassName(value string) drawer {
	return d.set("bodyClassName", value)
}

// ClassName sets the class name for the outer container.
func (d drawer) ClassName(value string) drawer {
	return d.set("className", value)
}

// CloseOnEsc sets whether the drawer can be closed by pressing the ESC key.
func (d drawer) CloseOnEsc(value bool) drawer {
	return d.set("closeOnEsc", value)
}

// CloseOnOutside sets whether the drawer can be closed by clicking outside of it.
func (d drawer) CloseOnOutside(value bool) drawer {
	return d.set("closeOnOutside", value)
}

// Confirm sets whether the confirm button is automatically generated. This setting is ignored if custom buttons are configured.
func (d drawer) Confirm(value bool) drawer {
	return d.set("confirm", value)
}

// Data sets the data mapping for the drawer.
func (d drawer) Data(value model.Data) drawer {
	return d.set("data", value)
}

// Disabled sets whether the drawer is disabled.
func (d drawer) Disabled(value bool) drawer {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine whether the drawer is disabled.
func (d drawer) DisabledOn(value string) drawer {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, which can be ignored at runtime.
func (d drawer) EditorSetting(value string) drawer {
	return d.set("editorSetting", value)
}

// Footer sets the footer content of the drawer.
func (d drawer) Footer(value string) drawer {
	return d.set("footer", value)
}

// FooterClassName sets the class name for the footer container.
func (d drawer) FooterClassName(value string) drawer {
	return d.set("footerClassName", value)
}

// Header sets the header content of the drawer.
func (d drawer) Header(value string) drawer {
	return d.set("header", value)
}

// HeaderClassName sets the class name for the header container.
func (d drawer) HeaderClassName(value string) drawer {
	return d.set("headerClassName", value)
}

// Height sets the height of the drawer. This is effective when the position is top or bottom.
func (d drawer) Height(value any) drawer {
	return d.set("height", value)
}

// Hidden sets whether the drawer is hidden.
func (d drawer) Hidden(value bool) drawer {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine whether the drawer is hidden.
func (d drawer) HiddenOn(value string) drawer {
	return d.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly used for logging.
func (d drawer) ID(value string) drawer {
	return d.set("id", value)
}

// InputParams sets the input parameters for the drawer in JSONSchema format.
func (d drawer) InputParams(value string) drawer {
	return d.set("inputParams", value)
}

// Name sets the name of the component, which can be used for locating and component communication.
func (d drawer) Name(value string) drawer {
	return d.set("name", value)
}

// OnEvent sets the event action configuration for the drawer.
func (d drawer) OnEvent(value any) drawer {
	return d.set("onEvent", value)
}

// Overlay sets whether the overlay is displayed.
func (d drawer) Overlay(value bool) drawer {
	return d.set("overlay", value)
}

// Position sets the position from which the drawer pops out. Possible values: left, right, top, bottom.
func (d drawer) Position(value string) drawer {
	return d.set("position", value)
}

// Resizable sets whether the drawer size can be resized.
func (d drawer) Resizable(value bool) drawer {
	return d.set("resizable", value)
}

// ShowCloseButton sets whether the close button is displayed. When false, closeOnOutside is enabled by default.
func (d drawer) ShowCloseButton(value bool) drawer {
	return d.set("showCloseButton", value)
}

// ShowErrorMsg sets whether error messages are displayed.
func (d drawer) ShowErrorMsg(value bool) drawer {
	return d.set("showErrorMsg", value)
}

// Size sets the size of the drawer. Possible values: xs, sm, md, lg, xl.
func (d drawer) Size(value string) drawer {
	return d.set("size", value)
}

// Static sets whether the drawer is displayed statically.
func (d drawer) Static(value bool) drawer {
	return d.set("static", value)
}

// StaticClassName sets the class name for static display form items.
func (d drawer) StaticClassName(value string) drawer {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values.
func (d drawer) StaticInputClassName(value string) drawer {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels.
func (d drawer) StaticLabelClassName(value string) drawer {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine whether the drawer is displayed statically.
func (d drawer) StaticOn(value string) drawer {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display when the value is empty.
func (d drawer) StaticPlaceholder(value string) drawer {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (d drawer) StaticSchema(value string) drawer {
	return d.set("staticSchema", value)
}

// Style sets the style for the component.
func (d drawer) Style(value any) drawer {
	return d.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (d drawer) TestIdBuilder(value string) drawer {
	return d.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (d drawer) Testid(value string) drawer {
	return d.set("testid", value)
}

// Title sets the title of the drawer.
func (d drawer) Title(value any) drawer {
	return d.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level.
func (d drawer) UseMobileUI(value bool) drawer {
	return d.set("useMobileUI", value)
}

// Visible sets whether the drawer is visible.
func (d drawer) Visible(value bool) drawer {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine whether the drawer is visible.
func (d drawer) VisibleOn(value string) drawer {
	return d.set("visibleOn", value)
}

// Width sets the width of the drawer. This is effective when the position is left or right.
func (d drawer) Width(value any) drawer {
	return d.set("width", value)
}
