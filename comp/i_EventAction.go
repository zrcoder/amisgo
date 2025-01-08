package comp

import "github.com/zrcoder/amisgo/model"

type MEventAction model.Schema

// EventAction creates a new EventAction instance
func EventAction() MEventAction {
	return MEventAction{}
}

// set sets a field value
func (la MEventAction) set(key string, value any) MEventAction {
	la[key] = value
	return la
}

// ActionType sets the action type
// Possible values:
// toast | ajax | dialog | closeDialog | drawer | closeDrawer | confirmDialog | alert
// url | link | goBack | refresh | copy | print | email | setEventData | wait
// setValue | reload | show | hidden | enabled | disabled | static | nonstatic
// broadcast | loop | break | continue | switch | parallel | custom
func (la MEventAction) ActionType(value string) MEventAction {
	return la.set("actionType", value)
}

// Drawer defines the drawer to pop up
func (ea MEventAction) Drawer(value any) MEventAction {
	return ea.set("drawer", value)
}

// Script sets a custom script when actionType is custom
func (ea MEventAction) Script(value string) MEventAction {
	return ea.set("script", value)
}

// Args sets the arguments
func (la MEventAction) Args(value model.Schema) MEventAction {
	return la.set("args", value)
}

// ComponentID sets the component ID
func (la MEventAction) ComponentID(value string) MEventAction {
	return la.set("componentId", value)
}

// ComponentName sets the component name
func (la MEventAction) ComponentName(value string) MEventAction {
	return la.set("componentName", value)
}

// Value sets the value
func (ea MEventAction) Value(value string) MEventAction {
	return ea.set("value", value)
}

// ConfirmTitle sets the confirm dialog title
func (la MEventAction) ConfirmTitle(value any) MEventAction {
	return la.set("confirmTitle", value)
}

// Data sets the data
func (la MEventAction) Data(value model.Data) MEventAction {
	return la.set("data", value)
}

// DataMergeMode sets the data merge mode
// Possible values: merge | override
func (la MEventAction) DataMergeMode(value string) MEventAction {
	return la.set("dataMergeMode", value)
}

// Description sets the description
func (la MEventAction) Description(value string) MEventAction {
	return la.set("description", value)
}

// ExecOn sets the execution condition
func (la MEventAction) ExecOn(value string) MEventAction {
	return la.set("execOn", value)
}

// Expression sets the execution condition as a boolean, expression, or ConditionBuilder
func (la MEventAction) Expression(value any) MEventAction {
	return la.set("expression", value)
}

// IgnoreError sets whether to ignore errors
func (la MEventAction) IgnoreError(value bool) MEventAction {
	return la.set("ignoreError", value)
}

// OutputVar sets the output variable
func (la MEventAction) OutputVar(value string) MEventAction {
	return la.set("outputVar", value)
}

// PreventDefault sets whether to prevent the default behavior as a boolean, expression, or ConditionBuilder
func (la MEventAction) PreventDefault(value any) MEventAction {
	return la.set("preventDefault", value)
}

// StopPropagation sets whether to stop event propagation as a boolean, expression, or ConditionBuilder
func (la MEventAction) StopPropagation(value any) MEventAction {
	return la.set("stopPropagation", value)
}
