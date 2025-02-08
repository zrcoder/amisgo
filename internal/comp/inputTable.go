package comp

import "github.com/zrcoder/amisgo/model"

// InputTable represents a table control
type InputTable model.Schema

func NewInputTable() InputTable {
	return InputTable{"type": "input-table"}
}

func (tc InputTable) set(key string, value any) InputTable {
	tc[key] = value
	return tc
}

// AddApi sets the add API
func (tc InputTable) AddApi(value string) InputTable {
	return tc.set("addApi", value)
}

// AddBtnIcon sets the add button icon
func (tc InputTable) AddBtnIcon(value string) InputTable {
	return tc.set("addBtnIcon", value)
}

// AddBtnLabel sets the add button label
func (tc InputTable) AddBtnLabel(value string) InputTable {
	return tc.set("addBtnLabel", value)
}

// Addable sets whether adding is allowed
func (tc InputTable) Addable(value bool) InputTable {
	return tc.set("addable", value)
}

// AffixFooter sets whether the footer is fixed
func (tc InputTable) AffixFooter(value bool) InputTable {
	return tc.set("affixFooter", value)
}

// AffixHeader sets whether the header is fixed
func (tc InputTable) AffixHeader(value bool) InputTable {
	return tc.set("affixHeader", value)
}

// AffixRow sets the footer summary row
func (tc InputTable) AffixRow(value string) InputTable {
	return tc.set("affixRow", value)
}

// AutoFill sets auto-fill when an option is selected
func (tc InputTable) AutoFill(value string) InputTable {
	return tc.set("autoFill", value)
}

// AutoFillHeight sets the table auto height
func (tc InputTable) AutoFillHeight(value string) InputTable {
	return tc.set("autoFillHeight", value)
}

// AutoGenerateFilter enables auto generation of filter form
func (tc InputTable) AutoGenerateFilter(value string) InputTable {
	return tc.set("autoGenerateFilter", value)
}

// CanAccessSuperData sets whether parent data can be accessed
func (tc InputTable) CanAccessSuperData(value bool) InputTable {
	return tc.set("canAccessSuperData", value)
}

// CancelBtnIcon sets the cancel button icon
func (tc InputTable) CancelBtnIcon(value string) InputTable {
	return tc.set("cancelBtnIcon", value)
}

// CancelBtnLabel sets the cancel button label
func (tc InputTable) CancelBtnLabel(value string) InputTable {
	return tc.set("cancelBtnLabel", value)
}

// ChildrenAddable sets whether child items can be added
func (tc InputTable) ChildrenAddable(value bool) InputTable {
	return tc.set("childrenAddable", value)
}

// ClassName sets the container CSS class name
func (tc InputTable) ClassName(value string) InputTable {
	return tc.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (tc InputTable) ClearValueOnHidden(value bool) InputTable {
	return tc.set("clearValueOnHidden", value)
}

// Columns sets the table columns
func (tc InputTable) Columns(value ...any) InputTable {
	return tc.set("columns", value)
}

// ColumnsTogglable sets whether column toggling is enabled
func (tc InputTable) ColumnsTogglable(value bool) InputTable {
	return tc.set("columnsTogglable", value)
}

// CombineFromIndex sets the column merge start index
func (tc InputTable) CombineFromIndex(value string) InputTable {
	return tc.set("combineFromIndex", value)
}

// CombineNum sets the number of columns to merge
func (tc InputTable) CombineNum(value string) InputTable {
	return tc.set("combineNum", value)
}

// ConfirmBtnIcon sets the confirm button icon
func (tc InputTable) ConfirmBtnIcon(value string) InputTable {
	return tc.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel sets the confirm button label
func (tc InputTable) ConfirmBtnLabel(value string) InputTable {
	return tc.set("confirmBtnLabel", value)
}

// CopyAddBtn sets whether the copy button is shown
func (tc InputTable) CopyAddBtn(value bool) InputTable {
	return tc.set("copyAddBtn", value)
}

// CopyBtnIcon sets the copy button icon
func (tc InputTable) CopyBtnIcon(value string) InputTable {
	return tc.set("copyBtnIcon", value)
}

// CopyBtnLabel sets the copy button label
func (tc InputTable) CopyBtnLabel(value string) InputTable {
	return tc.set("copyBtnLabel", value)
}

// CopyData sets the data mapping for copying
func (tc InputTable) CopyData(value string) InputTable {
	return tc.set("copyData", value)
}

// Copyable sets whether copying is allowed
func (tc InputTable) Copyable(value bool) InputTable {
	return tc.set("copyable", value)
}

// DeferApi sets the lazy load API
func (tc InputTable) DeferApi(value string) InputTable {
	return tc.set("deferApi", value)
}

// DeleteApi sets the delete API
func (tc InputTable) DeleteApi(value string) InputTable {
	return tc.set("deleteApi", value)
}

// DeleteBtnIcon sets the delete button icon
func (tc InputTable) DeleteBtnIcon(value string) InputTable {
	return tc.set("deleteBtnIcon", value)
}

// DeleteBtnLabel sets the delete button label
func (tc InputTable) DeleteBtnLabel(value string) InputTable {
	return tc.set("deleteBtnLabel", value)
}

// DeleteConfirmText sets the delete confirmation text
func (tc InputTable) DeleteConfirmText(value string) InputTable {
	return tc.set("deleteConfirmText", value)
}

// Desc sets the description
func (tc InputTable) Desc(value string) InputTable {
	return tc.set("desc", value)
}

// Description sets the description content
func (tc InputTable) Description(value string) InputTable {
	return tc.set("description", value)
}

// DescriptionClassName sets the description CSS class name
func (tc InputTable) DescriptionClassName(value string) InputTable {
	return tc.set("descriptionClassName", value)
}

// Disabled sets whether the table is disabled
func (tc InputTable) Disabled(value bool) InputTable {
	return tc.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (tc InputTable) DisabledOn(value string) InputTable {
	return tc.set("disabledOn", value)
}

// Draggable sets whether rows are draggable
func (tc InputTable) Draggable(value bool) InputTable {
	return tc.set("draggable", value)
}

// EditBtnIcon sets the edit button icon
func (tc InputTable) EditBtnIcon(value string) InputTable {
	return tc.set("editBtnIcon", value)
}

// EditBtnLabel sets the edit button label
func (tc InputTable) EditBtnLabel(value string) InputTable {
	return tc.set("editBtnLabel", value)
}

// Editable sets whether the table is editable
func (tc InputTable) Editable(value bool) InputTable {
	return tc.set("editable", value)
}

// EditorSetting sets the editor configuration
func (tc InputTable) EditorSetting(value string) InputTable {
	return tc.set("editorSetting", value)
}

// EnableStaticTransform sets whether static state switching is enabled
func (tc InputTable) EnableStaticTransform(value bool) InputTable {
	return tc.set("enableStaticTransform", value)
}

// ExtraName sets the extra field name
func (tc InputTable) ExtraName(value string) InputTable {
	return tc.set("extraName", value)
}

// Footable sets whether the footer display is enabled
func (tc InputTable) Footable(value bool) InputTable {
	return tc.set("footable", value)
}

// FooterAddBtn sets the footer add button configuration
func (tc InputTable) FooterAddBtn(value string) InputTable {
	return tc.set("footerAddBtn", value)
}

// FooterClassName sets the footer CSS class name
func (tc InputTable) FooterClassName(value string) InputTable {
	return tc.set("footerClassName", value)
}

// FooterTitle sets the footer title
func (tc InputTable) FooterTitle(value any) InputTable {
	return tc.set("footerTitle", value)
}

// FormApi sets the form API
func (tc InputTable) FormApi(value string) InputTable {
	return tc.set("formApi", value)
}

// HideQuickSave sets whether the quick save button is hidden
func (tc InputTable) HideQuickSave(value bool) InputTable {
	return tc.set("hideQuickSave", value)
}

// HoverAction sets the hover action configuration
func (tc InputTable) HoverAction(value string) InputTable {
	return tc.set("hoverAction", value)
}

// HoverTip sets the hover tip text
func (tc InputTable) HoverTip(value string) InputTable {
	return tc.set("hoverTip", value)
}

// Inline sets whether the edit box is inline
func (tc InputTable) Inline(value bool) InputTable {
	return tc.set("inline", value)
}

// IsPreview sets whether it is in preview mode
func (tc InputTable) IsPreview(value bool) InputTable {
	return tc.set("isPreview", value)
}

// ItemActionItem sets the list item action configuration
func (tc InputTable) ItemActionItem(value string) InputTable {
	return tc.set("itemActionItem", value)
}

// ItemActions sets the related action items
func (tc InputTable) ItemActions(value string) InputTable {
	return tc.set("itemActions", value)
}

// Label sets the label text
func (tc InputTable) Label(value string) InputTable {
	return tc.set("label", value)
}

// LabelClassName sets the label CSS class name
func (tc InputTable) LabelClassName(value string) InputTable {
	return tc.set("labelClassName", value)
}

// LabelRemark sets the label remark text
func (tc InputTable) LabelRemark(value string) InputTable {
	return tc.set("labelRemark", value)
}

// MaxLength sets the maximum length
func (tc InputTable) MaxLength(value string) InputTable {
	return tc.set("maxLength", value)
}

// Merge sets the merge configuration
func (tc InputTable) Merge(value string) InputTable {
	return tc.set("merge", value)
}

// MinLength sets the minimum length
func (tc InputTable) MinLength(value string) InputTable {
	return tc.set("minLength", value)
}

// MultiLine sets whether multi-line display is enabled
func (tc InputTable) MultiLine(value bool) InputTable {
	return tc.set("multiLine", value)
}

// Name sets the field name
func (tc InputTable) Name(value string) InputTable {
	return tc.set("name", value)
}

// NoEmpty sets whether to ignore empty fields
func (tc InputTable) NoEmpty(value bool) InputTable {
	return tc.set("noEmpty", value)
}

// NoBorder sets whether the component has no border
func (tc InputTable) NoBorder(value bool) InputTable {
	return tc.set("noBorder", value)
}

// NoIndex sets whether the index column is hidden
func (tc InputTable) NoIndex(value bool) InputTable {
	return tc.set("noIndex", value)
}

// NoResults sets the content displayed when there is no data
func (tc InputTable) NoResults(value string) InputTable {
	return tc.set("noResults", value)
}

// OnAction sets the callback function after an action is triggered
func (tc InputTable) OnAction(value string) InputTable {
	return tc.set("onAction", value)
}

// OnChange sets the callback function when the component value changes
func (tc InputTable) OnChange(value string) InputTable {
	return tc.set("onChange", value)
}

// OnConfirm sets the callback function when confirm is triggered
func (tc InputTable) OnConfirm(value string) InputTable {
	return tc.set("onConfirm", value)
}

// OnDelete sets the callback function when delete is triggered
func (tc InputTable) OnDelete(value string) InputTable {
	return tc.set("onDelete", value)
}

// OnEvent sets the event trigger configuration
func (tc InputTable) OnEvent(value any) InputTable {
	return tc.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (tc InputTable) Placeholder(value string) InputTable {
	return tc.set("placeholder", value)
}

// QuickSave sets whether quick save dialog is shown
func (tc InputTable) QuickSave(value bool) InputTable {
	return tc.set("quickSave", value)
}

// ReadOnly sets whether the table is read-only
func (tc InputTable) ReadOnly(value bool) InputTable {
	return tc.set("readOnly", value)
}

// Remarks sets the additional remarks
func (tc InputTable) Remarks(value string) InputTable {
	return tc.set("remarks", value)
}

// Remote sets the remote template configuration
func (tc InputTable) Remote(value string) InputTable {
	return tc.set("remote", value)
}

// RenderLabel sets the render label
func (tc InputTable) RenderLabel(value string) InputTable {
	return tc.set("renderLabel", value)
}

// SaveApi sets the save API
func (tc InputTable) SaveApi(value string) InputTable {
	return tc.set("saveApi", value)
}

// SearchApi sets the search API
func (tc InputTable) SearchApi(value string) InputTable {
	return tc.set("searchApi", value)
}

// Searchable sets whether the table is searchable
func (tc InputTable) Searchable(value bool) InputTable {
	return tc.set("searchable", value)
}

// Size sets the size (large, medium, small)
func (tc InputTable) Size(value string) InputTable {
	return tc.set("size", value)
}

// Source sets the remote interface configuration
func (tc InputTable) Source(value string) InputTable {
	return tc.set("source", value)
}

// Static sets the static data configuration
func (tc InputTable) Static(value string) InputTable {
	return tc.set("static", value)
}

// Status sets the status
func (tc InputTable) Status(value string) InputTable {
	return tc.set("status", value)
}

// SubForm sets the associated form
func (tc InputTable) SubForm(value string) InputTable {
	return tc.set("subForm", value)
}

// SubmitOnChange sets whether to submit on change
func (tc InputTable) SubmitOnChange(value bool) InputTable {
	return tc.set("submitOnChange", value)
}

// Tag sets the tag configuration
func (tc InputTable) Tag(value string) InputTable {
	return tc.set("tag", value)
}

// Union sets the form union configuration
func (tc InputTable) Union(value string) InputTable {
	return tc.set("union", value)
}

// Validate sets whether validation is enabled
func (tc InputTable) Validate(value bool) InputTable {
	return tc.set("validate", value)
}

// Value sets the default value
func (tc InputTable) Value(value string) InputTable {
	return tc.set("value", value)
}

// ValueType sets the value type
func (tc InputTable) ValueType(value string) InputTable {
	return tc.set("valueType", value)
}

// Width sets the control width
func (tc InputTable) Width(value string) InputTable {
	return tc.set("width", value)
}
