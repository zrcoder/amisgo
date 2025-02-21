package comp

import "github.com/zrcoder/amisgo/schema"

// Combo represents a composite input field component with multiple input capabilities
type Combo schema.Schema

func NewCombo() Combo {
	return Combo{"type": "combo"}
}

func (cc Combo) set(key string, value any) Combo {
	cc[key] = value
	return cc
}

// AddButtonClassName sets the CSS class name for the add button
func (cc Combo) AddButtonClassName(value string) Combo {
	return cc.set("addButtonClassName", value)
}

// AddButtonText sets the text displayed on the add button
func (cc Combo) AddButtonText(value string) Combo {
	return cc.set("addButtonText", value)
}

// Addable enables or disables the ability to add new items
func (cc Combo) Addable(value bool) Combo {
	return cc.set("addable", value)
}

// Addattop determines whether new items are added at the top of the list
func (cc Combo) Addattop(value bool) Combo {
	return cc.set("addattop", value)
}

// AutoFill configures automatic filling behavior
func (cc Combo) AutoFill(value string) Combo {
	return cc.set("autoFill", value)
}

// CanAccessSuperData enables or disables access to parent-level data
func (cc Combo) CanAccessSuperData(value bool) Combo {
	return cc.set("canAccessSuperData", value)
}

// ClassName sets the CSS class name for the container
func (cc Combo) ClassName(value string) Combo {
	return cc.set("className", value)
}

// ClearValueOnHidden determines whether to remove the form item's value when hidden
func (cc Combo) ClearValueOnHidden(value bool) Combo {
	return cc.set("clearValueOnHidden", value)
}

// Conditions sets a schema.Schema that renders only when specific conditions are met
func (cc Combo) Conditions(value string) Combo {
	return cc.set("conditions", value)
}

// DeleteApi specifies the API endpoint for deletion operations
func (cc Combo) DeleteApi(value string) Combo {
	return cc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation message for deletion
func (cc Combo) DeleteConfirmText(value string) Combo {
	return cc.set("deleteConfirmText", value)
}

// Delimiter sets the separator character for values
func (cc Combo) Delimiter(value string) Combo {
	return cc.set("delimiter", value)
}

// Desc sets a short description
func (cc Combo) Desc(value string) Combo {
	return cc.set("desc", value)
}

// Description sets detailed description content
func (cc Combo) Description(value string) Combo {
	return cc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (cc Combo) DescriptionClassName(value string) Combo {
	return cc.set("descriptionClassName", value)
}

// Disabled enables or disables the entire component
func (cc Combo) Disabled(value bool) Combo {
	return cc.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (cc Combo) DisabledOn(value string) Combo {
	return cc.set("disabledOn", value)
}

// Draggable enables or disables drag-and-drop sorting
func (cc Combo) Draggable(value bool) Combo {
	return cc.set("draggable", value)
}

// DraggableTip sets the tooltip text for drag-and-drop functionality
func (cc Combo) DraggableTip(value string) Combo {
	return cc.set("draggableTip", value)
}

// EditorSetting configures editor-specific settings
func (cc Combo) EditorSetting(value string) Combo {
	return cc.set("editorSetting", value)
}

// ExtraName sets an additional field name
func (cc Combo) ExtraName(value string) Combo {
	return cc.set("extraName", value)
}

// Flat determines whether to flatten the result structure
func (cc Combo) Flat(value bool) Combo {
	return cc.set("flat", value)
}

// FormClassName sets the CSS class name for internal form groups
func (cc Combo) FormClassName(value string) Combo {
	return cc.set("formClassName", value)
}

// Hidden controls the overall visibility of the component
func (cc Combo) Hidden(value bool) Combo {
	return cc.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (cc Combo) HiddenOn(value string) Combo {
	return cc.set("hiddenOn", value)
}

// Hint sets an input hint or placeholder text
func (cc Combo) Hint(value string) Combo {
	return cc.set("hint", value)
}

// Horizontal configures horizontal layout distribution
func (cc Combo) Horizontal(value string) Combo {
	return cc.set("horizontal", value)
}

// ID sets a unique identifier for the component
func (cc Combo) ID(value string) Combo {
	return cc.set("id", value)
}

// InitAutoFill configures initial automatic filling behavior
func (cc Combo) InitAutoFill(value string) Combo {
	return cc.set("initAutoFill", value)
}

// Inline enables or disables inline form control mode
func (cc Combo) Inline(value bool) Combo {
	return cc.set("inline", value)
}

// InputClassName sets the CSS class name for input elements
func (cc Combo) InputClassName(value string) Combo {
	return cc.set("inputClassName", value)
}

// Items sets the child items for the array input
func (cc Combo) Items(value ...any) Combo {
	return cc.set("items", value)
}

// JoinValues determines whether to send flattened values using a delimiter
func (cc Combo) JoinValues(value bool) Combo {
	return cc.set("joinValues", value)
}

// Label sets the description title
func (cc Combo) Label(value string) Combo {
	return cc.set("label", value)
}

// LabelAlign sets the alignment for the description title
func (cc Combo) LabelAlign(value string) Combo {
	return cc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (cc Combo) LabelClassName(value string) Combo {
	return cc.set("labelClassName", value)
}

// LabelRemark adds a small icon tooltip next to the label
func (cc Combo) LabelRemark(value string) Combo {
	return cc.set("labelRemark", value)
}

// LabelWidth sets a custom width for the label
func (cc Combo) LabelWidth(value string) Combo {
	return cc.set("labelWidth", value)
}

// LazyLoad enables or disables lazy loading
func (cc Combo) LazyLoad(value bool) Combo {
	return cc.set("lazyLoad", value)
}

// MaxLength sets the maximum number of items allowed
func (cc Combo) MaxLength(value string) Combo {
	return cc.set("maxLength", value)
}

// Messages configures custom validation or informational messages
func (cc Combo) Messages(value string) Combo {
	return cc.set("messages", value)
}

// MinLength sets the minimum number of items required
func (cc Combo) MinLength(value string) Combo {
	return cc.set("minLength", value)
}

// Mode sets the current display mode for the form item
func (cc Combo) Mode(value string) Combo {
	return cc.set("mode", value)
}

// MultiLine enables or disables multi-line input mode
func (cc Combo) MultiLine(value bool) Combo {
	return cc.set("multiLine", value)
}

// Multiple enables or disables multiple selection
func (cc Combo) Multiple(value bool) Combo {
	return cc.set("multiple", value)
}

// Name sets the field name for the component
func (cc Combo) Name(value string) Combo {
	return cc.set("name", value)
}

// NoBorder enables or disables border display
func (cc Combo) NoBorder(value bool) Combo {
	return cc.set("noBorder", value)
}

// Nullable determines whether null values are allowed
func (cc Combo) Nullable(value bool) Combo {
	return cc.set("nullable", value)
}

// OnEvent configures event-driven actions for the component
func (cc Combo) OnEvent(value Event) Combo {
	return cc.set("onEvent", value)
}

// Placeholder sets the content displayed when no members exist
func (cc Combo) Placeholder(value string) Combo {
	return cc.set("placeholder", value)
}

// ReadOnly enables or disables read-only mode
func (cc Combo) ReadOnly(value bool) Combo {
	return cc.set("readOnly", value)
}

// ReadOnlyOn sets a conditional expression for read-only state
func (cc Combo) ReadOnlyOn(value string) Combo {
	return cc.set("readOnlyOn", value)
}

// Remark adds a small icon tooltip for additional information
func (cc Combo) Remark(value string) Combo {
	return cc.set("remark", value)
}

// Removable enables or disables the ability to remove items
func (cc Combo) Removable(value bool) Combo {
	return cc.set("removable", value)
}

// Required determines whether the field is mandatory
func (cc Combo) Required(value bool) Combo {
	return cc.set("required", value)
}

// Row sets the row configuration for form groups
func (cc Combo) Row(value string) Combo {
	return cc.set("row", value)
}

// SaveAs configures the save format
func (cc Combo) SaveAs(value string) Combo {
	return cc.set("saveAs", value)
}

// Size sets the component size
func (cc Combo) Size(value string) Combo {
	return cc.set("size", value)
}

// Source specifies the data source API endpoint
func (cc Combo) Source(value string) Combo {
	return cc.set("source", value)
}

// StaticClassName sets the CSS class name for static components
func (cc Combo) StaticClassName(value string) Combo {
	return cc.set("staticClassName", value)
}

// Striped enables or disables striped styling
func (cc Combo) Striped(value bool) Combo {
	return cc.set("striped", value)
}

// SubmitOnChange determines whether to trigger submission on value change
func (cc Combo) SubmitOnChange(value bool) Combo {
	return cc.set("submitOnChange", value)
}

// Tip sets an informational message
func (cc Combo) Tip(value string) Combo {
	return cc.set("tip", value)
}

// Value sets the component's value
func (cc Combo) Value(value string) Combo {
	return cc.set("value", value)
}

// Vertical configures vertical layout distribution
func (cc Combo) Vertical(value string) Combo {
	return cc.set("vertical", value)
}

// Visible controls the overall visibility of the component
func (cc Combo) Visible(value bool) Combo {
	return cc.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (cc Combo) VisibleOn(value string) Combo {
	return cc.set("visibleOn", value)
}

// Width sets the component's width
func (cc Combo) Width(value string) Combo {
	return cc.set("width", value)
}

// Wrap sets the CSS class name for the outer div
func (cc Combo) Wrap(value string) Combo {
	return cc.set("wrap", value)
}
