package comp

import "github.com/zrcoder/amisgo/schema"

// NestedSelect represents a nested select component.
type NestedSelect schema.Schema

func NewNestedSelect() NestedSelect {
	return NestedSelect{"type": "nested-select"}
}

// set sets a field value.
func (nsc NestedSelect) set(key string, value any) NestedSelect {
	nsc[key] = value
	return nsc
}

// AddApi sets the API for adding items.
func (nsc NestedSelect) AddApi(value string) NestedSelect {
	return nsc.set("addApi", value)
}

// AddControls sets the form items for adding.
func (nsc NestedSelect) AddControls(value string) NestedSelect {
	return nsc.set("addControls", value)
}

// AddDialog sets the dialog for adding items.
func (nsc NestedSelect) AddDialog(value string) NestedSelect {
	return nsc.set("addDialog", value)
}

// AutoFill sets auto-fill.
func (nsc NestedSelect) AutoFill(value string) NestedSelect {
	return nsc.set("autoFill", value)
}

// BorderMode sets the border mode.
func (nsc NestedSelect) BorderMode(value string) NestedSelect {
	return nsc.set("borderMode", value)
}

// Cascade sets whether parent and child are independent.
func (nsc NestedSelect) Cascade(value bool) NestedSelect {
	return nsc.set("cascade", value)
}

// ClassName sets the container CSS class.
func (nsc NestedSelect) ClassName(value string) NestedSelect {
	return nsc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden.
func (nsc NestedSelect) ClearValueOnHidden(value bool) NestedSelect {
	return nsc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (nsc NestedSelect) Clearable(value bool) NestedSelect {
	return nsc.set("clearable", value)
}

// Creatable sets whether new items can be created.
func (nsc NestedSelect) Creatable(value bool) NestedSelect {
	return nsc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (nsc NestedSelect) CreateBtnLabel(value string) NestedSelect {
	return nsc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (nsc NestedSelect) DeferApi(value string) NestedSelect {
	return nsc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (nsc NestedSelect) DeferField(value string) NestedSelect {
	return nsc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (nsc NestedSelect) DeleteApi(value string) NestedSelect {
	return nsc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (nsc NestedSelect) DeleteConfirmText(value string) NestedSelect {
	return nsc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter.
func (nsc NestedSelect) Delimiter(value string) NestedSelect {
	return nsc.set("delimiter", value)
}

// Desc sets the description.
func (nsc NestedSelect) Desc(value string) NestedSelect {
	return nsc.set("desc", value)
}

// Description sets the description content.
func (nsc NestedSelect) Description(value string) NestedSelect {
	return nsc.set("description", value)
}

// DescriptionClassName sets the CSS class for the description.
func (nsc NestedSelect) DescriptionClassName(value string) NestedSelect {
	return nsc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (nsc NestedSelect) Disabled(value bool) NestedSelect {
	return nsc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field.
func (nsc NestedSelect) DisabledOn(value string) NestedSelect {
	return nsc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (nsc NestedSelect) EditApi(value string) NestedSelect {
	return nsc.set("editApi", value)
}

// EditControls sets the form items for editing.
func (nsc NestedSelect) EditControls(value string) NestedSelect {
	return nsc.set("editControls", value)
}

// EditDialog sets the dialog for editing items.
func (nsc NestedSelect) EditDialog(value string) NestedSelect {
	return nsc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (nsc NestedSelect) Editable(value bool) NestedSelect {
	return nsc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (nsc NestedSelect) EditorSetting(value string) NestedSelect {
	return nsc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (nsc NestedSelect) ExtraName(value string) NestedSelect {
	return nsc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (nsc NestedSelect) ExtractValue(value bool) NestedSelect {
	return nsc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (nsc NestedSelect) Hidden(value bool) NestedSelect {
	return nsc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field.
func (nsc NestedSelect) HiddenOn(value string) NestedSelect {
	return nsc.set("hiddenOn", value)
}

// HideNodePathLabel sets whether to hide the ancestor node text.
func (nsc NestedSelect) HideNodePathLabel(value bool) NestedSelect {
	return nsc.set("hideNodePathLabel", value)
}

// Hint sets the input hint.
func (nsc NestedSelect) Hint(value string) NestedSelect {
	return nsc.set("hint", value)
}

// Horizontal sets the horizontal layout.
func (nsc NestedSelect) Horizontal(value string) NestedSelect {
	return nsc.set("horizontal", value)
}

// ID sets the unique component ID.
func (nsc NestedSelect) ID(value string) NestedSelect {
	return nsc.set("id", value)
}

// InitAutoFill sets the initial auto-fill.
func (nsc NestedSelect) InitAutoFill(value string) NestedSelect {
	return nsc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially.
func (nsc NestedSelect) InitFetch(value bool) NestedSelect {
	return nsc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch.
func (nsc NestedSelect) InitFetchOn(value string) NestedSelect {
	return nsc.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (nsc NestedSelect) Inline(value bool) NestedSelect {
	return nsc.set("inline", value)
}

// InputClassName sets the input CSS class.
func (nsc NestedSelect) InputClassName(value string) NestedSelect {
	return nsc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single or multiple mode.
func (nsc NestedSelect) JoinValues(value bool) NestedSelect {
	return nsc.set("joinValues", value)
}

// Label sets the label.
func (nsc NestedSelect) Label(value string) NestedSelect {
	return nsc.set("label", value)
}

// LabelAlign sets the label alignment.
func (nsc NestedSelect) LabelAlign(value string) NestedSelect {
	return nsc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class.
func (nsc NestedSelect) LabelClassName(value string) NestedSelect {
	return nsc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (nsc NestedSelect) LabelRemark(value string) NestedSelect {
	return nsc.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (nsc NestedSelect) LabelWidth(value string) NestedSelect {
	return nsc.set("labelWidth", value)
}

// MaxTagCount sets the maximum number of tags to display.
func (nsc NestedSelect) MaxTagCount(value string) NestedSelect {
	return nsc.set("maxTagCount", value)
}

// MenuClassName sets the CSS class for the menu.
func (nsc NestedSelect) MenuClassName(value string) NestedSelect {
	return nsc.set("menuClassName", value)
}

// Mode sets the display mode.
func (nsc NestedSelect) Mode(value string) NestedSelect {
	return nsc.set("mode", value)
}

// Multiple sets whether the field is in multiple selection mode.
func (nsc NestedSelect) Multiple(value bool) NestedSelect {
	return nsc.set("multiple", value)
}

// Name sets the field name.
func (nsc NestedSelect) Name(value string) NestedSelect {
	return nsc.set("name", value)
}

// OnEvent sets the event actions.
func (nsc NestedSelect) OnEvent(value any) NestedSelect {
	return nsc.set("onEvent", value)
}

// OnlyChildren sets whether to include only child nodes when selecting parent.
func (nsc NestedSelect) OnlyChildren(value bool) NestedSelect {
	return nsc.set("onlyChildren", value)
}

// Options sets the options.
func (nsc NestedSelect) Options(value any) NestedSelect {
	return nsc.set("options", value)
}

// OptionTpl sets the option template.
func (nsc NestedSelect) OptionTpl(value string) NestedSelect {
	return nsc.set("optionTpl", value)
}

// Placeholder sets the placeholder text.
func (nsc NestedSelect) Placeholder(value string) NestedSelect {
	return nsc.set("placeholder", value)
}

// Remark sets the remark.
func (nsc NestedSelect) Remark(value string) NestedSelect {
	return nsc.set("remark", value)
}

// Render sets the render method.
func (nsc NestedSelect) Render(value string) NestedSelect {
	return nsc.set("render", value)
}

// SearchAble sets whether the field is searchable.
func (nsc NestedSelect) SearchAble(value bool) NestedSelect {
	return nsc.set("searchAble", value)
}

// ShowDesc sets whether to show the description.
func (nsc NestedSelect) ShowDesc(value bool) NestedSelect {
	return nsc.set("showDesc", value)
}

// Size sets the size.
func (nsc NestedSelect) Size(value string) NestedSelect {
	return nsc.set("size", value)
}

// Static sets whether the field is static.
func (nsc NestedSelect) Static(value bool) NestedSelect {
	return nsc.set("static", value)
}

// StaticLabel sets the static label content.
func (nsc NestedSelect) StaticLabel(value string) NestedSelect {
	return nsc.set("staticLabel", value)
}

// Style sets the style.
func (nsc NestedSelect) Style(value any) NestedSelect {
	return nsc.set("style", value)
}

// Tips sets the tips text.
func (nsc NestedSelect) Tips(value string) NestedSelect {
	return nsc.set("tips", value)
}

// Tree sets whether the child nodes are in tree structure.
func (nsc NestedSelect) Tree(value bool) NestedSelect {
	return nsc.set("tree", value)
}

// Value sets the selected value.
func (nsc NestedSelect) Value(value any) NestedSelect {
	return nsc.set("value", value)
}

// ValueField sets the value field.
func (nsc NestedSelect) ValueField(value string) NestedSelect {
	return nsc.set("valueField", value)
}

// Visible sets whether the field is visible.
func (nsc NestedSelect) Visible(value bool) NestedSelect {
	return nsc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (nsc NestedSelect) VisibleOn(value string) NestedSelect {
	return nsc.set("visibleOn", value)
}

// WithChildren sets whether to include child nodes.
func (nsc NestedSelect) WithChildren(value bool) NestedSelect {
	return nsc.set("withChildren", value)
}

// WithLabelChildren sets whether to show label child nodes.
func (nsc NestedSelect) WithLabelChildren(value bool) NestedSelect {
	return nsc.set("withLabelChildren", value)
}

// WithSelect sets whether to select nodes.
func (nsc NestedSelect) WithSelect(value bool) NestedSelect {
	return nsc.set("withSelect", value)
}
