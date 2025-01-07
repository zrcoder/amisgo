package comp

// inputTable represents a table control

type inputTable Schema

// InputTable creates a new TableControl instance
func InputTable() inputTable {
	return inputTable{}.set("type", "input-table")
}

func (tc inputTable) set(key string, value any) inputTable {
	tc[key] = value
	return tc
}

// AddApi sets the add API
func (tc inputTable) AddApi(value string) inputTable {
	return tc.set("addApi", value)
}

// AddBtnIcon sets the add button icon
func (tc inputTable) AddBtnIcon(value string) inputTable {
	return tc.set("addBtnIcon", value)
}

// AddBtnLabel sets the add button label
func (tc inputTable) AddBtnLabel(value string) inputTable {
	return tc.set("addBtnLabel", value)
}

// Addable sets whether adding is allowed
func (tc inputTable) Addable(value bool) inputTable {
	return tc.set("addable", value)
}

// AffixFooter sets whether the footer is fixed
func (tc inputTable) AffixFooter(value bool) inputTable {
	return tc.set("affixFooter", value)
}

// AffixHeader sets whether the header is fixed
func (tc inputTable) AffixHeader(value bool) inputTable {
	return tc.set("affixHeader", value)
}

// AffixRow sets the footer summary row
func (tc inputTable) AffixRow(value string) inputTable {
	return tc.set("affixRow", value)
}

// AutoFill sets auto-fill when an option is selected
func (tc inputTable) AutoFill(value string) inputTable {
	return tc.set("autoFill", value)
}

// AutoFillHeight sets the table auto height
func (tc inputTable) AutoFillHeight(value string) inputTable {
	return tc.set("autoFillHeight", value)
}

// AutoGenerateFilter enables auto generation of filter form
func (tc inputTable) AutoGenerateFilter(value string) inputTable {
	return tc.set("autoGenerateFilter", value)
}

// CanAccessSuperData sets whether parent data can be accessed
func (tc inputTable) CanAccessSuperData(value bool) inputTable {
	return tc.set("canAccessSuperData", value)
}

// CancelBtnIcon sets the cancel button icon
func (tc inputTable) CancelBtnIcon(value string) inputTable {
	return tc.set("cancelBtnIcon", value)
}

// CancelBtnLabel sets the cancel button label
func (tc inputTable) CancelBtnLabel(value string) inputTable {
	return tc.set("cancelBtnLabel", value)
}

// ChildrenAddable sets whether child items can be added
func (tc inputTable) ChildrenAddable(value bool) inputTable {
	return tc.set("childrenAddable", value)
}

// ClassName sets the container CSS class name
func (tc inputTable) ClassName(value string) inputTable {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tc inputTable) ClearValueOnHidden(value bool) inputTable {
	return tc.set("clearValueOnHidden", value)
}

// Columns sets the table columns
func (tc inputTable) Columns(value ...any) inputTable {
	return tc.set("columns", value)
}

// ColumnsTogglable sets whether column toggling is enabled
func (tc inputTable) ColumnsTogglable(value bool) inputTable {
	return tc.set("columnsTogglable", value)
}

// CombineFromIndex sets the column merge start index
func (tc inputTable) CombineFromIndex(value string) inputTable {
	return tc.set("combineFromIndex", value)
}

// CombineNum sets the number of columns to merge
func (tc inputTable) CombineNum(value string) inputTable {
	return tc.set("combineNum", value)
}

// ConfirmBtnIcon sets the confirm button icon
func (tc inputTable) ConfirmBtnIcon(value string) inputTable {
	return tc.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel sets the confirm button label
func (tc inputTable) ConfirmBtnLabel(value string) inputTable {
	return tc.set("confirmBtnLabel", value)
}

// CopyAddBtn sets whether the copy button is shown
func (tc inputTable) CopyAddBtn(value bool) inputTable {
	return tc.set("copyAddBtn", value)
}

// CopyBtnIcon sets the copy button icon
func (tc inputTable) CopyBtnIcon(value string) inputTable {
	return tc.set("copyBtnIcon", value)
}

// CopyBtnLabel sets the copy button label
func (tc inputTable) CopyBtnLabel(value string) inputTable {
	return tc.set("copyBtnLabel", value)
}

// CopyData sets the data mapping for copying
func (tc inputTable) CopyData(value string) inputTable {
	return tc.set("copyData", value)
}

// Copyable sets whether copying is allowed
func (tc inputTable) Copyable(value bool) inputTable {
	return tc.set("copyable", value)
}

// DeferApi sets the lazy load API
func (tc inputTable) DeferApi(value string) inputTable {
	return tc.set("deferApi", value)
}

// DeleteApi sets the delete API
func (tc inputTable) DeleteApi(value string) inputTable {
	return tc.set("deleteApi", value)
}

// DeleteBtnIcon sets the delete button icon
func (tc inputTable) DeleteBtnIcon(value string) inputTable {
	return tc.set("deleteBtnIcon", value)
}

// DeleteBtnLabel sets the delete button label
func (tc inputTable) DeleteBtnLabel(value string) inputTable {
	return tc.set("deleteBtnLabel", value)
}

// DeleteConfirmText sets the delete confirmation text
func (tc inputTable) DeleteConfirmText(value string) inputTable {
	return tc.set("deleteConfirmText", value)
}

// Desc sets the description
func (tc inputTable) Desc(value string) inputTable {
	return tc.set("desc", value)
}

// Description sets the description content
func (tc inputTable) Description(value string) inputTable {
	return tc.set("description", value)
}

// DescriptionClassName sets the description CSS class name
func (tc inputTable) DescriptionClassName(value string) inputTable {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the table is disabled
func (tc inputTable) Disabled(value bool) inputTable {
	return tc.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (tc inputTable) DisabledOn(value string) inputTable {
	return tc.set("disabledOn", value)
}

// Draggable sets whether rows are draggable
func (tc inputTable) Draggable(value bool) inputTable {
	return tc.set("draggable", value)
}

// EditBtnIcon sets the edit button icon
func (tc inputTable) EditBtnIcon(value string) inputTable {
	return tc.set("editBtnIcon", value)
}

// EditBtnLabel sets the edit button label
func (tc inputTable) EditBtnLabel(value string) inputTable {
	return tc.set("editBtnLabel", value)
}

// Editable sets whether the table is editable
func (tc inputTable) Editable(value bool) inputTable {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tc inputTable) EditorSetting(value string) inputTable {
	return tc.set("editorSetting", value)
}

// EnableStaticTransform sets whether static state switching is enabled
func (tc inputTable) EnableStaticTransform(value bool) inputTable {
	return tc.set("enableStaticTransform", value)
}

// ExtraName sets the extra field name
func (tc inputTable) ExtraName(value string) inputTable {
	return tc.set("extraName", value)
}

// Footable sets whether the footer display is enabled
func (tc inputTable) Footable(value bool) inputTable {
	return tc.set("footable", value)
}

// FooterAddBtn sets the footer add button configuration
func (tc inputTable) FooterAddBtn(value string) inputTable {
	return tc.set("footerAddBtn", value)
}

// FooterClassName sets the footer CSS class name
func (tc inputTable) FooterClassName(value string) inputTable {
	return tc.set("footerClassName", value)
}

// FooterTitle sets the footer title
func (tc inputTable) FooterTitle(value any) inputTable {
	return tc.set("footerTitle", value)
}

// FormApi sets the form API
func (tc inputTable) FormApi(value string) inputTable {
	return tc.set("formApi", value)
}

// HideQuickSave sets whether the quick save button is hidden
func (tc inputTable) HideQuickSave(value bool) inputTable {
	return tc.set("hideQuickSave", value)
}

// HoverAction sets the hover action configuration
func (tc inputTable) HoverAction(value string) inputTable {
	return tc.set("hoverAction", value)
}

// HoverTip sets the hover tip text
func (tc inputTable) HoverTip(value string) inputTable {
	return tc.set("hoverTip", value)
}

// Inline sets whether the edit box is inline
func (tc inputTable) Inline(value bool) inputTable {
	return tc.set("inline", value)
}

// IsPreview sets whether it is in preview mode
func (tc inputTable) IsPreview(value bool) inputTable {
	return tc.set("isPreview", value)
}

// ItemActionItem sets the list item action configuration
func (tc inputTable) ItemActionItem(value string) inputTable {
	return tc.set("itemActionItem", value)
}

// ItemActions sets the related action items
func (tc inputTable) ItemActions(value string) inputTable {
	return tc.set("itemActions", value)
}

// Label sets the label text
func (tc inputTable) Label(value string) inputTable {
	return tc.set("label", value)
}

// LabelClassName sets the label CSS class name
func (tc inputTable) LabelClassName(value string) inputTable {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the label remark text
func (tc inputTable) LabelRemark(value string) inputTable {
	return tc.set("labelRemark", value)
}

// MaxLength sets the maximum length
func (tc inputTable) MaxLength(value string) inputTable {
	return tc.set("maxLength", value)
}

// Merge sets the merge configuration
func (tc inputTable) Merge(value string) inputTable {
	return tc.set("merge", value)
}

// MinLength sets the minimum length
func (tc inputTable) MinLength(value string) inputTable {
	return tc.set("minLength", value)
}

// MultiLine sets whether multi-line display is enabled
func (tc inputTable) MultiLine(value bool) inputTable {
	return tc.set("multiLine", value)
}

// Name sets the field name
func (tc inputTable) Name(value string) inputTable {
	return tc.set("name", value)
}

// NoEmpty sets whether to ignore empty fields
func (tc inputTable) NoEmpty(value bool) inputTable {
	return tc.set("noEmpty", value)
}

// NoBorder sets whether the component has no border
func (tc inputTable) NoBorder(value bool) inputTable {
	return tc.set("noBorder", value)
}

// NoIndex sets whether the index column is hidden
func (tc inputTable) NoIndex(value bool) inputTable {
	return tc.set("noIndex", value)
}

// NoResults sets the content displayed when there is no data
func (tc inputTable) NoResults(value string) inputTable {
	return tc.set("noResults", value)
}

// OnAction sets the callback function after an action is triggered
func (tc inputTable) OnAction(value string) inputTable {
	return tc.set("onAction", value)
}

// OnChange sets the callback function when the component value changes
func (tc inputTable) OnChange(value string) inputTable {
	return tc.set("onChange", value)
}

// OnConfirm sets the callback function when confirm is triggered
func (tc inputTable) OnConfirm(value string) inputTable {
	return tc.set("onConfirm", value)
}

// OnDelete sets the callback function when delete is triggered
func (tc inputTable) OnDelete(value string) inputTable {
	return tc.set("onDelete", value)
}

// OnEvent sets the event trigger configuration
func (tc inputTable) OnEvent(value any) inputTable {
	return tc.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (tc inputTable) Placeholder(value string) inputTable {
	return tc.set("placeholder", value)
}

// QuickSave sets whether quick save dialog is shown
func (tc inputTable) QuickSave(value bool) inputTable {
	return tc.set("quickSave", value)
}

// ReadOnly sets whether the table is read-only
func (tc inputTable) ReadOnly(value bool) inputTable {
	return tc.set("readOnly", value)
}

// Remarks sets the additional remarks
func (tc inputTable) Remarks(value string) inputTable {
	return tc.set("remarks", value)
}

// Remote sets the remote template configuration
func (tc inputTable) Remote(value string) inputTable {
	return tc.set("remote", value)
}

// RenderLabel sets the render label
func (tc inputTable) RenderLabel(value string) inputTable {
	return tc.set("renderLabel", value)
}

// SaveApi sets the save API
func (tc inputTable) SaveApi(value string) inputTable {
	return tc.set("saveApi", value)
}

// SearchApi sets the search API
func (tc inputTable) SearchApi(value string) inputTable {
	return tc.set("searchApi", value)
}

// Searchable sets whether the table is searchable
func (tc inputTable) Searchable(value bool) inputTable {
	return tc.set("searchable", value)
}

// Size sets the size (large, medium, small)
func (tc inputTable) Size(value string) inputTable {
	return tc.set("size", value)
}

// Source sets the remote interface configuration
func (tc inputTable) Source(value string) inputTable {
	return tc.set("source", value)
}

// Static sets the static data configuration
func (tc inputTable) Static(value string) inputTable {
	return tc.set("static", value)
}

// Status sets the status
func (tc inputTable) Status(value string) inputTable {
	return tc.set("status", value)
}

// SubForm sets the associated form
func (tc inputTable) SubForm(value string) inputTable {
	return tc.set("subForm", value)
}

// SubmitOnChange sets whether to submit on change
func (tc inputTable) SubmitOnChange(value bool) inputTable {
	return tc.set("submitOnChange", value)
}

// Tag sets the tag configuration
func (tc inputTable) Tag(value string) inputTable {
	return tc.set("tag", value)
}

// Union sets the form union configuration
func (tc inputTable) Union(value string) inputTable {
	return tc.set("union", value)
}

// Validate sets whether validation is enabled
func (tc inputTable) Validate(value bool) inputTable {
	return tc.set("validate", value)
}

// Value sets the default value
func (tc inputTable) Value(value string) inputTable {
	return tc.set("value", value)
}

// ValueType sets the value type
func (tc inputTable) ValueType(value string) inputTable {
	return tc.set("valueType", value)
}

// Width sets the control width
func (tc inputTable) Width(value string) inputTable {
	return tc.set("width", value)
}
