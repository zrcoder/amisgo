package comp

import "github.com/zrcoder/amisgo/schema"

type EventAction schema.Schema

func NewEventAction() EventAction {
	return EventAction{}
}

func NewEventActionToast() EventAction {
	return NewEventAction().ActionType("toast")
}

func NewEventActionDrawer(drawer ...any) EventAction {
	res := NewEventAction().ActionType("drawer")
	if len(drawer) > 0 {
		res.Drawer(drawer[0])
	}
	return res
}

func NewEventActionDialog(dialog ...any) EventAction {
	res := NewEventAction().ActionType("dialog")
	if len(dialog) > 0 {
		res.Dialog(dialog[0])
	}
	return res
}

// set sets a field value
func (la EventAction) set(key string, value any) EventAction {
	la[key] = value
	return la
}

// ActionType sets the action type
// Possible values:
// toast | ajax | dialog | closeDialog | drawer | closeDrawer | confirmDialog | alert
// url | link | goBack | refresh | copy | print | email | setEventData | wait
// setValue | reload | show | hidden | enabled | disabled | static | nonstatic
// broadcast | loop | break | continue | switch | parallel | validate | custom
func (la EventAction) ActionType(value string) EventAction {
	return la.set("actionType", value)
}

func (ea EventAction) Children(value ...EventAction) EventAction {
	return ea.set("children", value)
}

// Api sets the api when actionType is ajax
func (ea EventAction) Api(value any) EventAction {
	return ea.set("api", value)
}

// Drawer defines the drawer when the actionType is drawer
func (ea EventAction) Drawer(value any) EventAction {
	return ea.set("drawer", value)
}

// Dialog defines the dialog when the actionType is dialog
func (ea EventAction) Dialog(value any) EventAction {
	return ea.set("dialog", value)
}

// Script sets a custom script when actionType is custom
func (ea EventAction) Script(value string) EventAction {
	return ea.set("script", value)
}

// Args sets the arguments
func (la EventAction) Args(value any) EventAction {
	return la.set("args", value)
}

// ComponentID sets the component ID
func (la EventAction) ComponentID(value string) EventAction {
	return la.set("componentId", value)
}

// ComponentName sets the component name
func (la EventAction) ComponentName(value string) EventAction {
	return la.set("componentName", value)
}

// Value sets the value
func (ea EventAction) Value(value string) EventAction {
	return ea.set("value", value)
}

// ConfirmTitle sets the confirm dialog title
func (la EventAction) ConfirmTitle(value any) EventAction {
	return la.set("confirmTitle", value)
}

// Data sets the data
func (la EventAction) Data(value any) EventAction {
	return la.set("data", value)
}

// DataMergeMode sets the data merge mode
// Possible values: merge | override
func (la EventAction) DataMergeMode(value string) EventAction {
	return la.set("dataMergeMode", value)
}

// Description sets the description
func (la EventAction) Description(value string) EventAction {
	return la.set("description", value)
}

// ExecOn sets the execution condition
func (la EventAction) ExecOn(value string) EventAction {
	return la.set("execOn", value)
}

// Expression sets the execution condition as a boolean, expression, or ConditionBuilder
func (la EventAction) Expression(value any) EventAction {
	return la.set("expression", value)
}

// IgnoreError sets whether to ignore errors
func (la EventAction) IgnoreError(value bool) EventAction {
	return la.set("ignoreError", value)
}

// OutputVar sets the output variable
func (la EventAction) OutputVar(value string) EventAction {
	return la.set("outputVar", value)
}

// WaitForAction sets if wait the dialog/drawer action
func (ea EventAction) WaitForAction(value bool) EventAction {
	return ea.set("waitForAction", value)
}

// PreventDefault sets whether to prevent the default behavior as a boolean, expression, or ConditionBuilder
func (la EventAction) PreventDefault(value any) EventAction {
	return la.set("preventDefault", value)
}

// StopPropagation sets whether to stop event propagation as a boolean, expression, or ConditionBuilder
func (la EventAction) StopPropagation(value any) EventAction {
	return la.set("stopPropagation", value)
}
