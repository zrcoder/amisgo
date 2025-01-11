package comp

import "github.com/zrcoder/amisgo/model"

// Combo represents a composite input field component with multiple input capabilities
type combo model.Schema

// Combo creates a new Combo instance
func Combo() combo {
	return combo{"type": "combo"}
}

func (cc combo) set(key string, value any) combo {
	cc[key] = value
	return cc
}

// AddButtonClassName sets the CSS class name for the add button
func (cc combo) AddButtonClassName(value string) combo {
	return cc.set("addButtonClassName", value)
}

// AddButtonText sets the text displayed on the add button
func (cc combo) AddButtonText(value string) combo {
	return cc.set("addButtonText", value)
}

// Addable enables or disables the ability to add new items
func (cc combo) Addable(value bool) combo {
	return cc.set("addable", value)
}

// Addattop determines whether new items are added at the top of the list
func (cc combo) Addattop(value bool) combo {
	return cc.set("addattop", value)
}

// AutoFill configures automatic filling behavior
func (cc combo) AutoFill(value string) combo {
	return cc.set("autoFill", value)
}

// CanAccessSuperData enables or disables access to parent-level data
func (cc combo) CanAccessSuperData(value bool) combo {
	return cc.set("canAccessSuperData", value)
}

// ClassName sets the CSS class name for the container
func (cc combo) ClassName(value string) combo {
	return cc.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (cc combo) ClearValueOnHidden(value bool) combo {
	return cc.set("clearValueOnHidden", value)
}

// Conditions sets a schema that renders only when specific conditions are met
func (cc combo) Conditions(value string) combo {
	return cc.set("conditions", value)
}

// DeleteApi specifies the API endpoint for deletion operations
func (cc combo) DeleteApi(value string) combo {
	return cc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation message for deletion
func (cc combo) DeleteConfirmText(value string) combo {
	return cc.set("deleteConfirmText", value)
}

// Delimiter sets the separator character for values
func (cc combo) Delimiter(value string) combo {
	return cc.set("delimiter", value)
}

// Desc sets a short description
func (cc combo) Desc(value string) combo {
	return cc.set("desc", value)
}

// Description sets detailed description content
func (cc combo) Description(value string) combo {
	return cc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (cc combo) DescriptionClassName(value string) combo {
	return cc.set("descriptionClassName", value)
}

// Disabled enables or disables the entire component
func (cc combo) Disabled(value bool) combo {
	return cc.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (cc combo) DisabledOn(value string) combo {
	return cc.set("disabledOn", value)
}

// Draggable enables or disables drag-and-drop sorting
func (cc combo) Draggable(value bool) combo {
	return cc.set("draggable", value)
}

// DraggableTip sets the tooltip text for drag-and-drop functionality
func (cc combo) DraggableTip(value string) combo {
	return cc.set("draggableTip", value)
}

// EditorSetting configures editor-specific settings
func (cc combo) EditorSetting(value string) combo {
	return cc.set("editorSetting", value)
}

// ExtraName sets an additional field name
func (cc combo) ExtraName(value string) combo {
	return cc.set("extraName", value)
}

// Flat determines whether to flatten the result structure
func (cc combo) Flat(value bool) combo {
	return cc.set("flat", value)
}

// FormClassName sets the CSS class name for internal form groups
func (cc combo) FormClassName(value string) combo {
	return cc.set("formClassName", value)
}

// Hidden controls the overall visibility of the component
func (cc combo) Hidden(value bool) combo {
	return cc.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (cc combo) HiddenOn(value string) combo {
	return cc.set("hiddenOn", value)
}

// Hint sets an input hint or placeholder text
func (cc combo) Hint(value string) combo {
	return cc.set("hint", value)
}

// Horizontal configures horizontal layout distribution
func (cc combo) Horizontal(value string) combo {
	return cc.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (cc combo) ID(value string) combo {
	return cc.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (cc combo) InitAutoFill(value string) combo {
	return cc.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (cc combo) Inline(value bool) combo {
	return cc.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (cc combo) InputClassName(value string) combo {
	return cc.set("inputClassName", value)
}

// Items sets the child items for the array input
func (cc combo) Items(value ...any) combo {
	return cc.set("items", value)
}

// JoinValues determines whether to send flattened values using a delimiter
func (cc combo) JoinValues(value bool) combo {
	return cc.set("joinValues", value)
}

// Label sets the description title
func (cc combo) Label(value string) combo {
	return cc.set("label", value)
}

// LabelAlign sets the alignment for the description title
func (cc combo) LabelAlign(value string) combo {
	return cc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (cc combo) LabelClassName(value string) combo {
	return cc.set("labelClassName", value)
}

// LabelRemark adds a small icon tooltip next to the label
func (cc combo) LabelRemark(value string) combo {
	return cc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (cc combo) LabelWidth(value string) combo {
	return cc.set("labelWidth", value)
}

// LazyLoad enables or disables lazy loading
func (cc combo) LazyLoad(value bool) combo {
	return cc.set("lazyLoad", value)
}

// MaxLength sets the maximum number of items allowed
func (cc combo) MaxLength(value string) combo {
	return cc.set("maxLength", value)
}

// Messages configures custom validation or informational messages
func (cc combo) Messages(value string) combo {
	return cc.set("messages", value)
}

// MinLength sets the minimum number of items required
func (cc combo) MinLength(value string) combo {
	return cc.set("minLength", value)
}

// Mode sets the current display mode for the form item
func (cc combo) Mode(value string) combo {
	return cc.set("mode", value)
}

// MultiLine enables or disables multi-line input mode
func (cc combo) MultiLine(value bool) combo {
	return cc.set("multiLine", value)
}

// Multiple enables or disables multiple selection
func (cc combo) Multiple(value bool) combo {
	return cc.set("multiple", value)
}

// Name sets the field name for the component
func (cc combo) Name(value string) combo {
	return cc.set("name", value)
}

// NoBorder enables or disables border display
func (cc combo) NoBorder(value bool) combo {
	return cc.set("noBorder", value)
}

// Nullable determines whether null values are allowed
func (cc combo) Nullable(value bool) combo {
	return cc.set("nullable", value)
}

// OnEvent configures event-driven actions for the component
func (cc combo) OnEvent(value any) combo {
	return cc.set("onEvent", value)
}

// Placeholder sets the content displayed when no members exist
func (cc combo) Placeholder(value string) combo {
	return cc.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (cc combo) ReadOnly(value bool) combo {
	return cc.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (cc combo) ReadOnlyOn(value string) combo {
	return cc.set("readOnlyOn", value)
}

// Remark adds a small icon tooltip for additional information
func (cc combo) Remark(value string) combo {
	return cc.set("remark", value)
}

// Removable enables or disables the ability to remove items
func (cc combo) Removable(value bool) combo {
	return cc.set("removable", value)
}

// Required determines whether the field is mandatory
func (cc combo) Required(value bool) combo {
	return cc.set("required", value)
}

// Row sets the row configuration for form groups
func (cc combo) Row(value string) combo {
	return cc.set("row", value)
}

// SaveAs configures the save format
func (cc combo) SaveAs(value string) combo {
	return cc.set("saveAs", value)
}

// Size sets the component size
func (cc combo) Size(value string) combo {
	return cc.set("size", value)
}

// Source specifies the data source API endpoint
func (cc combo) Source(value string) combo {
	return cc.set("source", value)
}

// StaticClassName sets the CSS class name for static components
func (cc combo) StaticClassName(value string) combo {
	return cc.set("staticClassName", value)
}

// Striped enables or disables striped styling
func (cc combo) Striped(value bool) combo {
	return cc.set("striped", value)
}

// SubmitOnChange determines whether to trigger submission on value change
func (cc combo) SubmitOnChange(value bool) combo {
	return cc.set("submitOnChange", value)
}

// Tip sets an informational message
func (cc combo) Tip(value string) combo {
	return cc.set("tip", value)
}

// Value sets the component's value
func (cc combo) Value(value string) combo {
	return cc.set("value", value)
}

// Vertical configures vertical layout distribution
func (cc combo) Vertical(value string) combo {
	return cc.set("vertical", value)
}

// Visible controls the overall visibility of the component
func (cc combo) Visible(value bool) combo {
	return cc.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (cc combo) VisibleOn(value string) combo {
	return cc.set("visibleOn", value)
}

// Width sets the component's width
func (cc combo) Width(value string) combo {
	return cc.set("width", value)
}

// Wrap sets the CSS class name for the outer div
func (cc combo) Wrap(value string) combo {
	return cc.set("wrap", value)
}
