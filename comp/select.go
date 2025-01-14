package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/internal/template"
	"github.com/zrcoder/amisgo/model"
)

// selectControl represents a select control configuration.
// Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/select
type selectControl model.Schema

// Select creates a new select instance
func Select() selectControl {
	return selectControl{"type": "select"}
}

// ButtonGroupSelect creates a new button group select instance
func ButtonGroupSelect() selectControl {
	return selectControl{"type": "button-group-select"}
}

// ThemeSelect creates a new select instance for theme selection.
// Note: Ensure that `conf.WithThemes` is called during app initialization to avoid a panic.
func ThemeSelect() selectControl {
	const theme = "theme"
	themes := template.GetThemes()
	if len(themes) == 0 {
		panic("ThemeSelect: conf.WithThemes must be called during app initialization")
	}
	url := servermux.ServeQuery(func(m map[string]string) error {
		template.SetTheme(m[theme])
		return nil
	}, theme)
	return Select().Name(theme).Source(servermux.ServeData(func() (any, error) {
		options := make([]any, len(themes))
		for i, v := range themes {
			label := v
			if v == "" {
				label = "default"
			}
			options[i] = Option().Label(label).Value(v)
		}
		return model.SuccessResponse(" ", model.Schema{
			"value":   template.GetTheme(),
			"options": options,
		}), nil
	})).OnEvent(
		model.Schema{
			"change": model.Schema{
				"actions": []any{
					EventAction().ActionType("ajax").Api(url + "?" + theme + "=${event.data.value}"),
					EventAction().ActionType("refresh"),
				},
			},
		},
	)
}

// ThemeButtonGroupSelect creates a new button group select instance for theme selection.
// Note: Ensure that `conf.WithThemes` is called during app initialization to avoid a panic.
func ThemeButtonGroupSelect() selectControl {
	return ThemeSelect().set("type", "button-group-select")
}

// set sets the key-value pair in the selectControl instance
func (sc selectControl) set(key string, value any) selectControl {
	sc[key] = value
	return sc
}

// SelectMode sets the select mode. options: group, table, tree, chained, associated.
func (sc selectControl) SelectMode(value string) selectControl {
	return sc.set("selectMode", value)
}

// AddApi sets the API for adding
func (sc selectControl) AddApi(value string) selectControl {
	return sc.set("addApi", value)
}

// AddControls sets the form items for adding
func (sc selectControl) AddControls(value string) selectControl {
	return sc.set("addControls", value)
}

// AddDialog sets the settings for the add dialog
func (sc selectControl) AddDialog(value string) selectControl {
	return sc.set("addDialog", value)
}

// AutoCheckChildren sets whether to automatically select child nodes
func (sc selectControl) AutoCheckChildren(value string) selectControl {
	return sc.set("autoCheckChildren", value)
}

// AutoComplete sets the auto-complete API
func (sc selectControl) AutoComplete(value string) selectControl {
	return sc.set("autoComplete", value)
}

// AutoFill sets whether to automatically fill in the form when an option is selected
func (sc selectControl) AutoFill(value string) selectControl {
	return sc.set("autoFill", value)
}

// BorderMode sets the border mode. Optional values: full | half | none
func (sc selectControl) BorderMode(value string) selectControl {
	return sc.set("borderMode", value)
}

// CheckAll sets whether to allow full selection in multi-select mode
func (sc selectControl) CheckAll(value bool) selectControl {
	return sc.set("checkAll", value)
}

// CheckAllLabel sets the label for the full selection option in multi-select mode
func (sc selectControl) CheckAllLabel(value string) selectControl {
	return sc.set("checkAllLabel", value)
}

// ClassName sets the container CSS class name
func (sc selectControl) ClassName(value string) selectControl {
	return sc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the form item value when it is hidden
func (sc selectControl) ClearValueOnHidden(value bool) selectControl {
	return sc.set("clearValueOnHidden", value)
}

// Clearable sets whether the form item can be cleared
func (sc selectControl) Clearable(value bool) selectControl {
	return sc.set("clearable", value)
}

// Columns sets the table column information when selectMode is table
func (sc selectControl) Columns(value ...any) selectControl {
	return sc.set("columns", value)
}

// Creatable sets whether new items can be created
func (sc selectControl) Creatable(value bool) selectControl {
	return sc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (sc selectControl) CreateBtnLabel(value string) selectControl {
	return sc.set("createBtnLabel", value)
}

// DefaultCheckAll sets whether to default to fully select all values in multi-select mode
func (sc selectControl) DefaultCheckAll(value bool) selectControl {
	return sc.set("defaultCheckAll", value)
}

// DeferApi sets the API for deferred loading
func (sc selectControl) DeferApi(value string) selectControl {
	return sc.set("deferApi", value)
}

// DeferField sets the deferred loading field
func (sc selectControl) DeferField(value string) selectControl {
	return sc.set("deferField", value)
}

// DeleteApi sets the API for deleting options
func (sc selectControl) DeleteApi(value string) selectControl {
	return sc.set("deleteApi", value)
}

// DeleteConfirmText sets the prompt text for deleting options
func (sc selectControl) DeleteConfirmText(value string) selectControl {
	return sc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (sc selectControl) Delimiter(value string) selectControl {
	return sc.set("delimiter", value)
}

// Desc sets the description
func (sc selectControl) Desc(value string) selectControl {
	return sc.set("desc", value)
}

// Description sets the description content, supports HTML fragments
func (sc selectControl) Description(value string) selectControl {
	return sc.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (sc selectControl) DescriptionClassName(value string) selectControl {
	return sc.set("descriptionClassName", value)
}

// Disabled sets whether the select control is disabled
func (sc selectControl) Disabled(value bool) selectControl {
	return sc.set("disabled", value)
}

// DisabledOn sets the expression to disable the select control
func (sc selectControl) DisabledOn(value string) selectControl {
	return sc.set("disabledOn", value)
}

// EditApi sets the API for editing
func (sc selectControl) EditApi(value string) selectControl {
	return sc.set("editApi", value)
}

// EditControls sets the form items for editing options
func (sc selectControl) EditControls(value string) selectControl {
	return sc.set("editControls", value)
}

// EditDialog sets the settings for the edit dialog
func (sc selectControl) EditDialog(value string) selectControl {
	return sc.set("editDialog", value)
}

// Editable sets whether the options can be edited
func (sc selectControl) Editable(value bool) selectControl {
	return sc.set("editable", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (sc selectControl) EditorSetting(value string) selectControl {
	return sc.set("editorSetting", value)
}

// ExtraName sets the extra field name, can be used to flatten another value when it is a range component
func (sc selectControl) ExtraName(value string) selectControl {
	return sc.set("extraName", value)
}

// ExtractValue sets whether to wrap the value of the selected option value as an array and submit it as the value of the current form item
func (sc selectControl) ExtractValue(value bool) selectControl {
	return sc.set("extractValue", value)
}

// Hidden sets whether the select control is hidden
func (sc selectControl) Hidden(value bool) selectControl {
	return sc.set("hidden", value)
}

// HiddenOn sets the expression to hide the select control
func (sc selectControl) HiddenOn(value string) selectControl {
	return sc.set("hiddenOn", value)
}

// Hint sets the input hint, displayed when focused
func (sc selectControl) Hint(value string) selectControl {
	return sc.set("hint", value)
}

// Horizontal sets the specific left-right allocation when configured as horizontal layout
func (sc selectControl) Horizontal(value string) selectControl {
	return sc.set("horizontal", value)
}

// ID sets the unique component ID, mainly used for log collection
func (sc selectControl) ID(value string) selectControl {
	return sc.set("id", value)
}

// InitAutoFill sets the initial auto-fill
func (sc selectControl) InitAutoFill(value string) selectControl {
	return sc.set("initAutoFill", value)
}

// InitFetch configures the initial fetch for the source interface
func (sc selectControl) InitFetch(value bool) selectControl {
	return sc.set("initFetch", value)
}

// InitFetchOn configures whether to initially fetch the source interface with an expression
func (sc selectControl) InitFetchOn(value string) selectControl {
	return sc.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode
func (sc selectControl) Inline(value bool) selectControl {
	return sc.set("inline", value)
}

// InputClassName sets the input class name
func (sc selectControl) InputClassName(value string) selectControl {
	return sc.set("inputClassName", value)
}

// ItemHeight sets the height of a single option, mainly used for virtual rendering
func (sc selectControl) ItemHeight(value string) selectControl {
	return sc.set("itemHeight", value)
}

// JoinValues sets whether to submit the value of the selected option value as an array when it is a single selection mode, or join the values of the selected multiple options with the delimiter when it is a multi-selection mode
func (sc selectControl) JoinValues(value bool) selectControl {
	return sc.set("joinValues", value)
}

// Label sets the label text
func (sc selectControl) Label(value string) selectControl {
	return sc.set("label", value)
}

// LabelClassName sets the class name for the label
func (sc selectControl) LabelClassName(value string) selectControl {
	return sc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (sc selectControl) LabelRemark(value string) selectControl {
	return sc.set("labelRemark", value)
}

// MatchObj
func (sc selectControl) MatchObj(value string) selectControl {
	return sc.set("matchObj", value)
}

// Multiple sets whether to support multiple selection
func (sc selectControl) Multiple(value bool) selectControl {
	return sc.set("multiple", value)
}

// Name sets the name of the form item
func (sc selectControl) Name(value string) selectControl {
	return sc.set("name", value)
}

// SelectFirst
func (sc selectControl) SelectFirst(value bool) selectControl {
	return sc.set("selectFirst", value)
}

// Mode sets the mode of the select control. Optional values: normal | inline | horizontal
func (sc selectControl) Mode(value string) selectControl {
	return sc.set("mode", value)
}

// NoneOptions sets the prompt text when the value provided by the autoComplete API is not supported by the select options themselves
func (sc selectControl) NoneOptions(value ...any) selectControl {
	return sc.set("noneOptions", value)
}

// OnChange sets the change callback
func (sc selectControl) OnChange(value string) selectControl {
	return sc.set("onChange", value)
}

// OnDialogConfirm sets the callback when the dialog confirm button is clicked
func (sc selectControl) OnDialogConfirm(value string) selectControl {
	return sc.set("onDialogConfirm", value)
}

// Options sets the options collection, mainly used to configure the drop-down list items
func (sc selectControl) Options(value ...any) selectControl {
	return sc.set("options", value)
}

// PopOverConfig sets the PopOver configuration
func (sc selectControl) PopOverConfig(value string) selectControl {
	return sc.set("popOverConfig", value)
}

// Reload sets whether to reload data when re-rendering
func (sc selectControl) Reload(value string) selectControl {
	return sc.set("reload", value)
}

// ResetValue sets the value when resetting
func (sc selectControl) ResetValue(value string) selectControl {
	return sc.set("resetValue", value)
}

// RootClassName sets the class name for the outer layer of the component
func (sc selectControl) RootClassName(value string) selectControl {
	return sc.set("rootClassName", value)
}

// Searchable sets whether to support search
func (sc selectControl) Searchable(value bool) selectControl {
	return sc.set("searchable", value)
}

// Source sets the data source
func (sc selectControl) Source(value string) selectControl {
	return sc.set("source", value)
}

// SourceClassName sets the additional class name for the source
func (sc selectControl) SourceClassName(value string) selectControl {
	return sc.set("sourceClassName", value)
}

// SourceApi sets the API for the source data
func (sc selectControl) SourceApi(value string) selectControl {
	return sc.set("sourceApi", value)
}

// SourceMessage sets the prompt message for the source data
func (sc selectControl) SourceMessage(value string) selectControl {
	return sc.set("sourceMessage", value)
}

// StaticClassName sets the class name
func (sc selectControl) StaticClassName(value string) selectControl {
	return sc.set("staticClassName", value)
}

// SubmitOnChange sets whether to automatically submit after the value changes
func (sc selectControl) SubmitOnChange(value bool) selectControl {
	return sc.set("submitOnChange", value)
}

// Type_ sets the component type
func (sc selectControl) Type_(value string) selectControl {
	return sc.set("type", value)
}

// Value sets the bound value
func (sc selectControl) Value(value any) selectControl {
	return sc.set("value", value)
}

// Visible sets the display condition
func (sc selectControl) Visible(value bool) selectControl {
	return sc.set("visible", value)
}

// VisibleOn sets the expression to determine if the component should be displayed
func (sc selectControl) VisibleOn(value string) selectControl {
	return sc.set("visibleOn", value)
}

// Virtual sets whether to use virtual rendering
func (sc selectControl) Virtual(value bool) selectControl {
	return sc.set("virtual", value)
}

// Width sets the component width
func (sc selectControl) Width(value string) selectControl {
	return sc.set("width", value)
}

// OnEvent sets the event listener
func (sc selectControl) OnEvent(value any) selectControl {
	return sc.set("onEvent", value)
}

// BtnActiveClassName sets the CSS class name for button active state
func (sc selectControl) BtnActiveClassName(value string) selectControl {
	return sc.set("btnActiveClassName", value)
}

// BtnActiveLevel sets the style level for selected buttons
func (sc selectControl) BtnActiveLevel(value string) selectControl {
	return sc.set("btnActiveLevel", value)
}

// BtnClassName sets the CSS class name for buttons
func (sc selectControl) BtnClassName(value string) selectControl {
	return sc.set("btnClassName", value)
}

// BtnLevel sets the style level for buttons
func (sc selectControl) BtnLevel(value string) selectControl {
	return sc.set("btnLevel", value)
}

// Buttons configures the collection of buttons
func (sc selectControl) Buttons(value ...any) selectControl {
	return sc.set("buttons", value)
}

// LabelAlign sets the alignment for the description title
func (sc selectControl) LabelAlign(value string) selectControl {
	return sc.set("labelAlign", value)
}

// LabelWidth sets the width of the label
func (sc selectControl) LabelWidth(value string) selectControl {
	return sc.set("labelWidth", value)
}

// OptionIcon sets the option icon
func (sc selectControl) OptionIcon(value string) selectControl {
	return sc.set("optionIcon", value)
}

// OptionLabel sets the option label
func (sc selectControl) OptionLabel(value string) selectControl {
	return sc.set("optionLabel", value)
}

// OptionValue sets the option value
func (sc selectControl) OptionValue(value string) selectControl {
	return sc.set("optionValue", value)
}

// Outline enables or disables outline button style
func (sc selectControl) Outline(value bool) selectControl {
	return sc.set("outline", value)
}

// Overlay enables or disables Overlay
func (sc selectControl) Overlay(value bool) selectControl {
	return sc.set("overlay", value)
}

// Size sets the control size
func (sc selectControl) Size(value string) selectControl {
	return sc.set("size", value)
}

// ValueField sets the value field name
func (sc selectControl) ValueField(value string) selectControl {
	return sc.set("valueField", value)
}

// Validation sets the validation rules
func (sc selectControl) Validation(value string) selectControl {
	return sc.set("validation", value)
}
