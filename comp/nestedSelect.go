package comp

import "github.com/zrcoder/amisgo/model"

// nestedSelect represents a nested select component.

type nestedSelect model.Schema

func NestedSelect() nestedSelect {
	return make(nestedSelect).set("type", "nested-select")
}

// set sets a field value.
func (nsc nestedSelect) set(key string, value any) nestedSelect {
	nsc[key] = value
	return nsc
}

// AddApi sets the API for adding items.
func (nsc nestedSelect) AddApi(value string) nestedSelect {
	return nsc.set("addApi", value)
}

// AddControls sets the form items for adding.
func (nsc nestedSelect) AddControls(value string) nestedSelect {
	return nsc.set("addControls", value)
}

// AddDialog sets the dialog for adding items.
func (nsc nestedSelect) AddDialog(value string) nestedSelect {
	return nsc.set("addDialog", value)
}

// AutoFill sets auto-fill.
func (nsc nestedSelect) AutoFill(value string) nestedSelect {
	return nsc.set("autoFill", value)
}

// BorderMode sets the border mode.
func (nsc nestedSelect) BorderMode(value string) nestedSelect {
	return nsc.set("borderMode", value)
}

// Cascade sets whether parent and child are independent.
func (nsc nestedSelect) Cascade(value bool) nestedSelect {
	return nsc.set("cascade", value)
}

// ClassName sets the container CSS class.
func (nsc nestedSelect) ClassName(value string) nestedSelect {
	return nsc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden.
func (nsc nestedSelect) ClearValueOnHidden(value bool) nestedSelect {
	return nsc.set("clearValueOnHidden", value)
}

// Clearable sets whether the field is clearable.
func (nsc nestedSelect) Clearable(value bool) nestedSelect {
	return nsc.set("clearable", value)
}

// Creatable sets whether new items can be created.
func (nsc nestedSelect) Creatable(value bool) nestedSelect {
	return nsc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button.
func (nsc nestedSelect) CreateBtnLabel(value string) nestedSelect {
	return nsc.set("createBtnLabel", value)
}

// DeferApi sets the API for deferred loading.
func (nsc nestedSelect) DeferApi(value string) nestedSelect {
	return nsc.set("deferApi", value)
}

// DeferField sets the field for lazy loading.
func (nsc nestedSelect) DeferField(value string) nestedSelect {
	return nsc.set("deferField", value)
}

// DeleteApi sets the API for deleting items.
func (nsc nestedSelect) DeleteApi(value string) nestedSelect {
	return nsc.set("deleteApi", value)
}

// DeleteConfirmText sets the confirmation text for deletion.
func (nsc nestedSelect) DeleteConfirmText(value string) nestedSelect {
	return nsc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter.
func (nsc nestedSelect) Delimiter(value string) nestedSelect {
	return nsc.set("delimiter", value)
}

// Desc sets the description.
func (nsc nestedSelect) Desc(value string) nestedSelect {
	return nsc.set("desc", value)
}

// Description sets the description content.
func (nsc nestedSelect) Description(value string) nestedSelect {
	return nsc.set("description", value)
}

// DescriptionClassName sets the CSS class for the description.
func (nsc nestedSelect) DescriptionClassName(value string) nestedSelect {
	return nsc.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (nsc nestedSelect) Disabled(value bool) nestedSelect {
	return nsc.set("disabled", value)
}

// DisabledOn sets the expression for disabling the field.
func (nsc nestedSelect) DisabledOn(value string) nestedSelect {
	return nsc.set("disabledOn", value)
}

// EditApi sets the API for editing items.
func (nsc nestedSelect) EditApi(value string) nestedSelect {
	return nsc.set("editApi", value)
}

// EditControls sets the form items for editing.
func (nsc nestedSelect) EditControls(value string) nestedSelect {
	return nsc.set("editControls", value)
}

// EditDialog sets the dialog for editing items.
func (nsc nestedSelect) EditDialog(value string) nestedSelect {
	return nsc.set("editDialog", value)
}

// Editable sets whether the field is editable.
func (nsc nestedSelect) Editable(value bool) nestedSelect {
	return nsc.set("editable", value)
}

// EditorSetting sets the editor configuration.
func (nsc nestedSelect) EditorSetting(value string) nestedSelect {
	return nsc.set("editorSetting", value)
}

// ExtraName sets the extra field name.
func (nsc nestedSelect) ExtraName(value string) nestedSelect {
	return nsc.set("extraName", value)
}

// ExtractValue sets whether to extract the value as an array.
func (nsc nestedSelect) ExtractValue(value bool) nestedSelect {
	return nsc.set("extractValue", value)
}

// Hidden sets whether the field is hidden.
func (nsc nestedSelect) Hidden(value bool) nestedSelect {
	return nsc.set("hidden", value)
}

// HiddenOn sets the expression for hiding the field.
func (nsc nestedSelect) HiddenOn(value string) nestedSelect {
	return nsc.set("hiddenOn", value)
}

// HideNodePathLabel sets whether to hide the ancestor node text.
func (nsc nestedSelect) HideNodePathLabel(value bool) nestedSelect {
	return nsc.set("hideNodePathLabel", value)
}

// Hint sets the input hint.
func (nsc nestedSelect) Hint(value string) nestedSelect {
	return nsc.set("hint", value)
}

// Horizontal sets the horizontal layout.
func (nsc nestedSelect) Horizontal(value string) nestedSelect {
	return nsc.set("horizontal", value)
}

// ID sets the unique component ID.
func (nsc nestedSelect) ID(value string) nestedSelect {
	return nsc.set("id", value)
}

// InitAutoFill sets the initial auto-fill.
func (nsc nestedSelect) InitAutoFill(value string) nestedSelect {
	return nsc.set("initAutoFill", value)
}

// InitFetch sets whether to fetch the source initially.
func (nsc nestedSelect) InitFetch(value bool) nestedSelect {
	return nsc.set("initFetch", value)
}

// InitFetchOn sets the expression for initial fetch.
func (nsc nestedSelect) InitFetchOn(value string) nestedSelect {
	return nsc.set("initFetchOn", value)
}

// Inline sets whether the control is inline.
func (nsc nestedSelect) Inline(value bool) nestedSelect {
	return nsc.set("inline", value)
}

// InputClassName sets the input CSS class.
func (nsc nestedSelect) InputClassName(value string) nestedSelect {
	return nsc.set("inputClassName", value)
}

// JoinValues sets whether to join values in single or multiple mode.
func (nsc nestedSelect) JoinValues(value bool) nestedSelect {
	return nsc.set("joinValues", value)
}

// Label sets the label.
func (nsc nestedSelect) Label(value string) nestedSelect {
	return nsc.set("label", value)
}

// LabelAlign sets the label alignment.
func (nsc nestedSelect) LabelAlign(value string) nestedSelect {
	return nsc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class.
func (nsc nestedSelect) LabelClassName(value string) nestedSelect {
	return nsc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (nsc nestedSelect) LabelRemark(value string) nestedSelect {
	return nsc.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (nsc nestedSelect) LabelWidth(value string) nestedSelect {
	return nsc.set("labelWidth", value)
}

// MaxTagCount sets the maximum number of tags to display.
func (nsc nestedSelect) MaxTagCount(value string) nestedSelect {
	return nsc.set("maxTagCount", value)
}

// MenuClassName sets the CSS class for the menu.
func (nsc nestedSelect) MenuClassName(value string) nestedSelect {
	return nsc.set("menuClassName", value)
}

// Mode sets the display mode.
func (nsc nestedSelect) Mode(value string) nestedSelect {
	return nsc.set("mode", value)
}

// Multiple sets whether the field is in multiple selection mode.
func (nsc nestedSelect) Multiple(value bool) nestedSelect {
	return nsc.set("multiple", value)
}

// Name sets the field name.
func (nsc nestedSelect) Name(value string) nestedSelect {
	return nsc.set("name", value)
}

// OnEvent sets the event actions.
func (nsc nestedSelect) OnEvent(value any) nestedSelect {
	return nsc.set("onEvent", value)
}

// OnlyChildren sets whether to include only child nodes when selecting parent.
func (nsc nestedSelect) OnlyChildren(value bool) nestedSelect {
	return nsc.set("onlyChildren", value)
}

// Options sets the options.
func (nsc nestedSelect) Options(value any) nestedSelect {
	return nsc.set("options", value)
}

// OptionTpl sets the option template.
func (nsc nestedSelect) OptionTpl(value string) nestedSelect {
	return nsc.set("optionTpl", value)
}

// Placeholder sets the placeholder text.
func (nsc nestedSelect) Placeholder(value string) nestedSelect {
	return nsc.set("placeholder", value)
}

// Remark sets the remark.
func (nsc nestedSelect) Remark(value string) nestedSelect {
	return nsc.set("remark", value)
}

// Render sets the render method.
func (nsc nestedSelect) Render(value string) nestedSelect {
	return nsc.set("render", value)
}

// SearchAble sets whether the field is searchable.
func (nsc nestedSelect) SearchAble(value bool) nestedSelect {
	return nsc.set("searchAble", value)
}

// ShowDesc sets whether to show the description.
func (nsc nestedSelect) ShowDesc(value bool) nestedSelect {
	return nsc.set("showDesc", value)
}

// Size sets the size.
func (nsc nestedSelect) Size(value string) nestedSelect {
	return nsc.set("size", value)
}

// Static sets whether the field is static.
func (nsc nestedSelect) Static(value bool) nestedSelect {
	return nsc.set("static", value)
}

// StaticLabel sets the static label content.
func (nsc nestedSelect) StaticLabel(value string) nestedSelect {
	return nsc.set("staticLabel", value)
}

// Style sets the style.
func (nsc nestedSelect) Style(value any) nestedSelect {
	return nsc.set("style", value)
}

// Tips sets the tips text.
func (nsc nestedSelect) Tips(value string) nestedSelect {
	return nsc.set("tips", value)
}

// Tree sets whether the child nodes are in tree structure.
func (nsc nestedSelect) Tree(value bool) nestedSelect {
	return nsc.set("tree", value)
}

// Value sets the selected value.
func (nsc nestedSelect) Value(value any) nestedSelect {
	return nsc.set("value", value)
}

// ValueField sets the value field.
func (nsc nestedSelect) ValueField(value string) nestedSelect {
	return nsc.set("valueField", value)
}

// Visible sets whether the field is visible.
func (nsc nestedSelect) Visible(value bool) nestedSelect {
	return nsc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (nsc nestedSelect) VisibleOn(value string) nestedSelect {
	return nsc.set("visibleOn", value)
}

// WithChildren sets whether to include child nodes.
func (nsc nestedSelect) WithChildren(value bool) nestedSelect {
	return nsc.set("withChildren", value)
}

// WithLabelChildren sets whether to show label child nodes.
func (nsc nestedSelect) WithLabelChildren(value bool) nestedSelect {
	return nsc.set("withLabelChildren", value)
}

// WithSelect sets whether to select nodes.
func (nsc nestedSelect) WithSelect(value bool) nestedSelect {
	return nsc.set("withSelect", value)
}
