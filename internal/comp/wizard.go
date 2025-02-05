package comp

import "github.com/zrcoder/amisgo/model"

// Wizard form Wizard documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Wizard
type Wizard model.Schema

func NewWizard() Wizard {
	return Wizard{"type": "wizard"}
}

func (w Wizard) set(key string, value any) Wizard {
	w[key] = value
	return w
}

// ActionClassName sets the button className
func (w Wizard) ActionClassName(value string) Wizard {
	return w.set("actionClassName", value)
}

// ActionFinishLabel sets the finish button label
func (w Wizard) ActionFinishLabel(value string) Wizard {
	return w.set("actionFinishLabel", value)
}

// ActionNextLabel sets the next button label
func (w Wizard) ActionNextLabel(value string) Wizard {
	return w.set("actionNextLabel", value)
}

// ActionNextSaveLabel sets the next and save button label
func (w Wizard) ActionNextSaveLabel(value string) Wizard {
	return w.set("actionNextSaveLabel", value)
}

// ActionPrevLabel sets the previous button label
func (w Wizard) ActionPrevLabel(value string) Wizard {
	return w.set("actionPrevLabel", value)
}

// AffixFooter sets whether to fix the footer at the bottom
func (w Wizard) AffixFooter(value string) Wizard {
	return w.set("affixFooter", value)
}

// Api sets the API for saving data
func (w Wizard) Api(value string) Wizard {
	return w.set("api", value)
}

// BodyClassName sets the form area CSS class
func (w Wizard) BodyClassName(value string) Wizard {
	return w.set("bodyClassName", value)
}

// BulkSubmit sets whether to submit in bulk
func (w Wizard) BulkSubmit(value bool) Wizard {
	return w.set("bulkSubmit", value)
}

// ClassName sets the container CSS class
func (w Wizard) ClassName(value string) Wizard {
	return w.set("className", value)
}

// Disabled sets whether the wizard is disabled
func (w Wizard) Disabled(value bool) Wizard {
	return w.set("disabled", value)
}

// DisabledOn sets the expression to disable the wizard
func (w Wizard) DisabledOn(value string) Wizard {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (w Wizard) EditorSetting(value string) Wizard {
	return w.set("editorSetting", value)
}

// FooterClassName sets the footer CSS class
func (w Wizard) FooterClassName(value string) Wizard {
	return w.set("footerClassName", value)
}

// Hidden sets whether the wizard is hidden
func (w Wizard) Hidden(value bool) Wizard {
	return w.set("hidden", value)
}

// HiddenOn sets the expression to hide the wizard
func (w Wizard) HiddenOn(value string) Wizard {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID
func (w Wizard) ID(value string) Wizard {
	return w.set("id", value)
}

// InitApi sets the API for initial data
func (w Wizard) InitApi(value string) Wizard {
	return w.set("initApi", value)
}

// LoadingConfig sets the loading configuration
func (w Wizard) LoadingConfig(value string) Wizard {
	return w.set("loadingConfig", value)
}

// Mode sets the display mode (vertical or horizontal)
func (w Wizard) Mode(value string) Wizard {
	return w.set("mode", value)
}

// Name sets the component name
func (w Wizard) Name(value string) Wizard {
	return w.set("name", value)
}

// OnEvent sets the event action configuration
func (w Wizard) OnEvent(value any) Wizard {
	return w.set("onEvent", value)
}

// ReadOnly sets whether the wizard is read-only
func (w Wizard) ReadOnly(value bool) Wizard {
	return w.set("readOnly", value)
}

// Redirect sets the redirect URL after saving
func (w Wizard) Redirect(value string) Wizard {
	return w.set("redirect", value)
}

// Reload sets the reload action configuration
func (w Wizard) Reload(value string) Wizard {
	return w.set("reload", value)
}

// StartStep sets the starting step
func (w Wizard) StartStep(value string) Wizard {
	return w.set("startStep", value)
}

// Static sets whether the wizard is static
func (w Wizard) Static(value bool) Wizard {
	return w.set("static", value)
}

// StaticClassName sets the static display form item class
func (w Wizard) StaticClassName(value string) Wizard {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class
func (w Wizard) StaticInputClassName(value string) Wizard {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class
func (w Wizard) StaticLabelClassName(value string) Wizard {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (w Wizard) StaticOn(value string) Wizard {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (w Wizard) StaticPlaceholder(value string) Wizard {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w Wizard) StaticSchema(value string) Wizard {
	return w.set("staticSchema", value)
}

// StepClassName sets the step + body area CSS class
func (w Wizard) StepClassName(value string) Wizard {
	return w.set("stepClassName", value)
}

// Steps sets the steps
func (w Wizard) Steps(value string) Wizard {
	return w.set("steps", value)
}

// StepsClassName sets the steps area CSS class
func (w Wizard) StepsClassName(value string) Wizard {
	return w.set("stepsClassName", value)
}

// Style sets the component style
func (w Wizard) Style(value any) Wizard {
	return w.set("style", value)
}

// Target sets the target form or CRUD model name
func (w Wizard) Target(value string) Wizard {
	return w.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (w Wizard) TestIdBuilder(value string) Wizard {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w Wizard) Testid(value string) Wizard {
	return w.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (w Wizard) UseMobileUI(value bool) Wizard {
	return w.set("useMobileUI", value)
}

// Visible sets whether the wizard is visible
func (w Wizard) Visible(value bool) Wizard {
	return w.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (w Wizard) VisibleOn(value string) Wizard {
	return w.set("visibleOn", value)
}

// WrapWithPanel sets whether to wrap with a panel
func (w Wizard) WrapWithPanel(value bool) Wizard {
	return w.set("wrapWithPanel", value)
}
