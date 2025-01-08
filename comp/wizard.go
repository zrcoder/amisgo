package comp

import "github.com/zrcoder/amisgo/model"

// wizard form wizard documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/wizard

type wizard model.Schema

// Wizard creates a new Wizard instance
func Wizard() wizard {
	return wizard{}.set("type", "wizard")
}

func (w wizard) set(key string, value any) wizard {
	w[key] = value
	return w
}

// ActionClassName sets the button className
func (w wizard) ActionClassName(value string) wizard {
	return w.set("actionClassName", value)
}

// ActionFinishLabel sets the finish button label
func (w wizard) ActionFinishLabel(value string) wizard {
	return w.set("actionFinishLabel", value)
}

// ActionNextLabel sets the next button label
func (w wizard) ActionNextLabel(value string) wizard {
	return w.set("actionNextLabel", value)
}

// ActionNextSaveLabel sets the next and save button label
func (w wizard) ActionNextSaveLabel(value string) wizard {
	return w.set("actionNextSaveLabel", value)
}

// ActionPrevLabel sets the previous button label
func (w wizard) ActionPrevLabel(value string) wizard {
	return w.set("actionPrevLabel", value)
}

// AffixFooter sets whether to fix the footer at the bottom
func (w wizard) AffixFooter(value string) wizard {
	return w.set("affixFooter", value)
}

// Api sets the API for saving data
func (w wizard) Api(value string) wizard {
	return w.set("api", value)
}

// BodyClassName sets the form area CSS class
func (w wizard) BodyClassName(value string) wizard {
	return w.set("bodyClassName", value)
}

// BulkSubmit sets whether to submit in bulk
func (w wizard) BulkSubmit(value bool) wizard {
	return w.set("bulkSubmit", value)
}

// ClassName sets the container CSS class
func (w wizard) ClassName(value string) wizard {
	return w.set("className", value)
}

// Disabled sets whether the wizard is disabled
func (w wizard) Disabled(value bool) wizard {
	return w.set("disabled", value)
}

// DisabledOn sets the expression to disable the wizard
func (w wizard) DisabledOn(value string) wizard {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (w wizard) EditorSetting(value string) wizard {
	return w.set("editorSetting", value)
}

// FooterClassName sets the footer CSS class
func (w wizard) FooterClassName(value string) wizard {
	return w.set("footerClassName", value)
}

// Hidden sets whether the wizard is hidden
func (w wizard) Hidden(value bool) wizard {
	return w.set("hidden", value)
}

// HiddenOn sets the expression to hide the wizard
func (w wizard) HiddenOn(value string) wizard {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID
func (w wizard) ID(value string) wizard {
	return w.set("id", value)
}

// InitApi sets the API for initial data
func (w wizard) InitApi(value string) wizard {
	return w.set("initApi", value)
}

// LoadingConfig sets the loading configuration
func (w wizard) LoadingConfig(value string) wizard {
	return w.set("loadingConfig", value)
}

// Mode sets the display mode (vertical or horizontal)
func (w wizard) Mode(value string) wizard {
	return w.set("mode", value)
}

// Name sets the component name
func (w wizard) Name(value string) wizard {
	return w.set("name", value)
}

// OnEvent sets the event action configuration
func (w wizard) OnEvent(value any) wizard {
	return w.set("onEvent", value)
}

// ReadOnly sets whether the wizard is read-only
func (w wizard) ReadOnly(value bool) wizard {
	return w.set("readOnly", value)
}

// Redirect sets the redirect URL after saving
func (w wizard) Redirect(value string) wizard {
	return w.set("redirect", value)
}

// Reload sets the reload action configuration
func (w wizard) Reload(value string) wizard {
	return w.set("reload", value)
}

// StartStep sets the starting step
func (w wizard) StartStep(value string) wizard {
	return w.set("startStep", value)
}

// Static sets whether the wizard is static
func (w wizard) Static(value bool) wizard {
	return w.set("static", value)
}

// StaticClassName sets the static display form item class
func (w wizard) StaticClassName(value string) wizard {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class
func (w wizard) StaticInputClassName(value string) wizard {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class
func (w wizard) StaticLabelClassName(value string) wizard {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (w wizard) StaticOn(value string) wizard {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (w wizard) StaticPlaceholder(value string) wizard {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w wizard) StaticSchema(value string) wizard {
	return w.set("staticSchema", value)
}

// StepClassName sets the step + body area CSS class
func (w wizard) StepClassName(value string) wizard {
	return w.set("stepClassName", value)
}

// Steps sets the steps
func (w wizard) Steps(value string) wizard {
	return w.set("steps", value)
}

// StepsClassName sets the steps area CSS class
func (w wizard) StepsClassName(value string) wizard {
	return w.set("stepsClassName", value)
}

// Style sets the component style
func (w wizard) Style(value any) wizard {
	return w.set("style", value)
}

// Target sets the target form or CRUD model name
func (w wizard) Target(value string) wizard {
	return w.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (w wizard) TestIdBuilder(value string) wizard {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w wizard) Testid(value string) wizard {
	return w.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (w wizard) UseMobileUI(value bool) wizard {
	return w.set("useMobileUI", value)
}

// Visible sets whether the wizard is visible
func (w wizard) Visible(value bool) wizard {
	return w.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (w wizard) VisibleOn(value string) wizard {
	return w.set("visibleOn", value)
}

// WrapWithPanel sets whether to wrap with a panel
func (w wizard) WrapWithPanel(value bool) wizard {
	return w.set("wrapWithPanel", value)
}
