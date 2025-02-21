package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/internal/template"
	"github.com/zrcoder/amisgo/schema"
)

// Select represents a select control configuration.
// Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/select
type Select schema.Schema

func NewSelect() Select {
	return Select{"type": "select"}
}

// NewButtonGroupSelect creates a new button group select instance
func NewButtonGroupSelect() Select {
	return Select{"type": "button-group-select"}
}

// Themes makes the select for theme selection.
// Note: Ensure that `conf.WithThemes` is called during app initialization to avoid a panic.
func (s Select) Themes(mux *http.ServeMux, templ *template.Templ) Select {
	const theme = "theme"
	if len(templ.Themes) == 0 {
		panic("ThemeSelect: conf.WithThemes must be called during app initialization")
	}
	url := servemux.ServeQuery(mux, func(m map[string]string) error {
		templ.UpdateTheme(m[theme])
		return nil
	}, theme)
	return s.Name(theme).Source(servemux.ServeData(mux, func() (any, error) {
		return schema.SuccessResponse(" ", schema.Schema{
			"value":   templ.CurrentTheme(),
			"options": templ.Themes,
		}), nil
	})).OnEvent(
		NewEvent().Change(
			NewEventActions(NewEventAction().ActionType("ajax").Api(url+"?"+theme+"=${event.data.value}"),
				NewEventAction().ActionType("refresh"),
			),
		),
	)
}

// Localesmakes the select for locale selection.
// Note: Ensure that `conf.WithLocales` is called during app initialization to avoid a panic.
func (s Select) Locales(mux *http.ServeMux, templ *template.Templ) Select {
	const locale = "locale"
	if len(templ.Locales) == 0 {
		panic("LocaleSelect: conf.WithLocales must be called during app initialization")
	}
	url := servemux.ServeQuery(mux, func(m map[string]string) error {
		templ.UpdateLocale(m[locale])
		return nil
	}, locale)
	return s.Name(locale).Source(servemux.ServeData(mux, func() (any, error) {
		return schema.SuccessResponse(" ", schema.Schema{
			"value":   templ.CurrentLocale(),
			"options": templ.Locales,
		}), nil
	})).OnEvent(
		NewEvent().Change(
			NewEventActions(NewEventAction().ActionType("ajax").Api(url+"?"+locale+"=${event.data.value}"),
				NewEventAction().ActionType("refresh"),
			),
		),
	)
}

// set sets the key-value pair in the selectControl instance
func (sc Select) set(key string, value any) Select {
	sc[key] = value
	return sc
}

// SelectMode sets the select mode. options: group, table, tree, chained, associated.
func (sc Select) SelectMode(value string) Select {
	return sc.set("selectMode", value)
}

// AddApi sets the API for adding
func (sc Select) AddApi(value string) Select {
	return sc.set("addApi", value)
}

// AddControls sets the form items for adding
func (sc Select) AddControls(value string) Select {
	return sc.set("addControls", value)
}

// AddDialog sets the settings for the add dialog
func (sc Select) AddDialog(value string) Select {
	return sc.set("addDialog", value)
}

// AutoCheckChildren sets whether to automatically select child nodes
func (sc Select) AutoCheckChildren(value string) Select {
	return sc.set("autoCheckChildren", value)
}

// AutoComplete sets the auto-complete API
func (sc Select) AutoComplete(value string) Select {
	return sc.set("autoComplete", value)
}

// AutoFill sets whether to automatically fill in the form when an option is selected
func (sc Select) AutoFill(value string) Select {
	return sc.set("autoFill", value)
}

// BorderMode sets the border mode. Optional values: full | half | none
func (sc Select) BorderMode(value string) Select {
	return sc.set("borderMode", value)
}

// CheckAll sets whether to allow full selection in multi-select mode
func (sc Select) CheckAll(value bool) Select {
	return sc.set("checkAll", value)
}

// CheckAllLabel sets the label for the full selection option in multi-select mode
func (sc Select) CheckAllLabel(value string) Select {
	return sc.set("checkAllLabel", value)
}

// ClassName sets the container CSS class name
func (sc Select) ClassName(value string) Select {
	return sc.set("className", value)
}

// ClearValueOnHidden sets whether to clear the form item value when it is hidden
func (sc Select) ClearValueOnHidden(value bool) Select {
	return sc.set("clearValueOnHidden", value)
}

// Clearable sets whether the form item can be cleared
func (sc Select) Clearable(value bool) Select {
	return sc.set("clearable", value)
}

// Columns sets the table column information when selectMode is table
func (sc Select) Columns(value ...any) Select {
	return sc.set("columns", value)
}

// Creatable sets whether new items can be created
func (sc Select) Creatable(value bool) Select {
	return sc.set("creatable", value)
}

// CreateBtnLabel sets the label for the create button
func (sc Select) CreateBtnLabel(value string) Select {
	return sc.set("createBtnLabel", value)
}

// DefaultCheckAll sets whether to default to fully select all values in multi-select mode
func (sc Select) DefaultCheckAll(value bool) Select {
	return sc.set("defaultCheckAll", value)
}

// DeferApi sets the API for deferred loading
func (sc Select) DeferApi(value string) Select {
	return sc.set("deferApi", value)
}

// DeferField sets the deferred loading field
func (sc Select) DeferField(value string) Select {
	return sc.set("deferField", value)
}

// DeleteApi sets the API for deleting options
func (sc Select) DeleteApi(value string) Select {
	return sc.set("deleteApi", value)
}

// DeleteConfirmText sets the prompt text for deleting options
func (sc Select) DeleteConfirmText(value string) Select {
	return sc.set("deleteConfirmText", value)
}

// Delimiter sets the delimiter
func (sc Select) Delimiter(value string) Select {
	return sc.set("delimiter", value)
}

// Desc sets the description
func (sc Select) Desc(value string) Select {
	return sc.set("desc", value)
}

// Description sets the description content, supports HTML fragments
func (sc Select) Description(value string) Select {
	return sc.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (sc Select) DescriptionClassName(value string) Select {
	return sc.set("descriptionClassName", value)
}

// Disabled sets whether the select control is disabled
func (sc Select) Disabled(value bool) Select {
	return sc.set("disabled", value)
}

// DisabledOn sets the expression to disable the select control
func (sc Select) DisabledOn(value string) Select {
	return sc.set("disabledOn", value)
}

// EditApi sets the API for editing
func (sc Select) EditApi(value string) Select {
	return sc.set("editApi", value)
}

// EditControls sets the form items for editing options
func (sc Select) EditControls(value string) Select {
	return sc.set("editControls", value)
}

// EditDialog sets the settings for the edit dialog
func (sc Select) EditDialog(value string) Select {
	return sc.set("editDialog", value)
}

// Editable sets whether the options can be edited
func (sc Select) Editable(value bool) Select {
	return sc.set("editable", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (sc Select) EditorSetting(value string) Select {
	return sc.set("editorSetting", value)
}

// ExtraName sets the extra field name, can be used to flatten another value when it is a range component
func (sc Select) ExtraName(value string) Select {
	return sc.set("extraName", value)
}

// ExtractValue sets whether to wrap the value of the selected option value as an array and submit it as the value of the current form item
func (sc Select) ExtractValue(value bool) Select {
	return sc.set("extractValue", value)
}

// Hidden sets whether the select control is hidden
func (sc Select) Hidden(value bool) Select {
	return sc.set("hidden", value)
}

// HiddenOn sets the expression to hide the select control
func (sc Select) HiddenOn(value string) Select {
	return sc.set("hiddenOn", value)
}

// Hint sets the input hint, displayed when focused
func (sc Select) Hint(value string) Select {
	return sc.set("hint", value)
}

// Horizontal sets the specific left-right allocation when configured as horizontal layout
func (sc Select) Horizontal(value string) Select {
	return sc.set("horizontal", value)
}

// ID sets the unique component ID, mainly used for log collection
func (sc Select) ID(value string) Select {
	return sc.set("id", value)
}

// InitAutoFill sets the initial auto-fill
func (sc Select) InitAutoFill(value string) Select {
	return sc.set("initAutoFill", value)
}

// InitFetch configures the initial fetch for the source interface
func (sc Select) InitFetch(value bool) Select {
	return sc.set("initFetch", value)
}

// InitFetchOn configures whether to initially fetch the source interface with an expression
func (sc Select) InitFetchOn(value string) Select {
	return sc.set("initFetchOn", value)
}

// Inline sets whether the form control is in inline mode
func (sc Select) Inline(value bool) Select {
	return sc.set("inline", value)
}

// InputClassName sets the input class name
func (sc Select) InputClassName(value string) Select {
	return sc.set("inputClassName", value)
}

// ItemHeight sets the height of a single option, mainly used for virtual rendering
func (sc Select) ItemHeight(value string) Select {
	return sc.set("itemHeight", value)
}

// JoinValues sets whether to submit the value of the selected option value as an array when it is a single selection mode, or join the values of the selected multiple options with the delimiter when it is a multi-selection mode
func (sc Select) JoinValues(value bool) Select {
	return sc.set("joinValues", value)
}

// Label sets the label text
func (sc Select) Label(value string) Select {
	return sc.set("label", value)
}

// LabelClassName sets the class name for the label
func (sc Select) LabelClassName(value string) Select {
	return sc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (sc Select) LabelRemark(value string) Select {
	return sc.set("labelRemark", value)
}

// MatchObj
func (sc Select) MatchObj(value string) Select {
	return sc.set("matchObj", value)
}

// Multiple sets whether to support multiple selection
func (sc Select) Multiple(value bool) Select {
	return sc.set("multiple", value)
}

// Name sets the name of the form item
func (sc Select) Name(value string) Select {
	return sc.set("name", value)
}

// SelectFirst
func (sc Select) SelectFirst(value bool) Select {
	return sc.set("selectFirst", value)
}

// Mode sets the mode of the select control. Optional values: normal | inline | horizontal
func (sc Select) Mode(value string) Select {
	return sc.set("mode", value)
}

// NoneOptions sets the prompt text when the value provided by the autoComplete API is not supported by the select options themselves
func (sc Select) NoneOptions(value ...any) Select {
	return sc.set("noneOptions", value)
}

// OnChange sets the change callback
func (sc Select) OnChange(value string) Select {
	return sc.set("onChange", value)
}

// OnDialogConfirm sets the callback when the dialog confirm button is clicked
func (sc Select) OnDialogConfirm(value string) Select {
	return sc.set("onDialogConfirm", value)
}

// Options sets the options collection, mainly used to configure the drop-down list items
func (sc Select) Options(value ...any) Select {
	return sc.set("options", value)
}

// PopOverConfig sets the PopOver configuration
func (sc Select) PopOverConfig(value string) Select {
	return sc.set("popOverConfig", value)
}

// Reload sets whether to reload data when re-rendering
func (sc Select) Reload(value string) Select {
	return sc.set("reload", value)
}

// ResetValue sets the value when resetting
func (sc Select) ResetValue(value string) Select {
	return sc.set("resetValue", value)
}

// RootClassName sets the class name for the outer layer of the component
func (sc Select) RootClassName(value string) Select {
	return sc.set("rootClassName", value)
}

// Searchable sets whether to support search
func (sc Select) Searchable(value bool) Select {
	return sc.set("searchable", value)
}

// Source sets the data source
func (sc Select) Source(value string) Select {
	return sc.set("source", value)
}

// SourceClassName sets the additional class name for the source
func (sc Select) SourceClassName(value string) Select {
	return sc.set("sourceClassName", value)
}

// SourceApi sets the API for the source data
func (sc Select) SourceApi(value string) Select {
	return sc.set("sourceApi", value)
}

// SourceMessage sets the prompt message for the source data
func (sc Select) SourceMessage(value string) Select {
	return sc.set("sourceMessage", value)
}

// StaticClassName sets the class name
func (sc Select) StaticClassName(value string) Select {
	return sc.set("staticClassName", value)
}

// SubmitOnChange sets whether to automatically submit after the value changes
func (sc Select) SubmitOnChange(value bool) Select {
	return sc.set("submitOnChange", value)
}

// Type_ sets the component type
func (sc Select) Type_(value string) Select {
	return sc.set("type", value)
}

// Value sets the bound value
func (sc Select) Value(value any) Select {
	return sc.set("value", value)
}

// Visible sets the display condition
func (sc Select) Visible(value bool) Select {
	return sc.set("visible", value)
}

// VisibleOn sets the expression to determine if the component should be displayed
func (sc Select) VisibleOn(value string) Select {
	return sc.set("visibleOn", value)
}

// Virtual sets whether to use virtual rendering
func (sc Select) Virtual(value bool) Select {
	return sc.set("virtual", value)
}

// Width sets the component width
func (sc Select) Width(value string) Select {
	return sc.set("width", value)
}

// OnEvent sets the event listener
func (sc Select) OnEvent(value Event) Select {
	return sc.set("onEvent", value)
}

// BtnActiveClassName sets the CSS class name for button active state
func (sc Select) BtnActiveClassName(value string) Select {
	return sc.set("btnActiveClassName", value)
}

// BtnActiveLevel sets the style level for selected buttons
func (sc Select) BtnActiveLevel(value string) Select {
	return sc.set("btnActiveLevel", value)
}

// BtnClassName sets the CSS class name for buttons
func (sc Select) BtnClassName(value string) Select {
	return sc.set("btnClassName", value)
}

// BtnLevel sets the style level for buttons
func (sc Select) BtnLevel(value string) Select {
	return sc.set("btnLevel", value)
}

// Buttons configures the collection of buttons
func (sc Select) Buttons(value ...any) Select {
	return sc.set("buttons", value)
}

// LabelAlign sets the alignment for the description title
func (sc Select) LabelAlign(value string) Select {
	return sc.set("labelAlign", value)
}

// LabelWidth sets the width of the label
func (sc Select) LabelWidth(value string) Select {
	return sc.set("labelWidth", value)
}

// OptionIcon sets the option icon
func (sc Select) OptionIcon(value string) Select {
	return sc.set("optionIcon", value)
}

// OptionLabel sets the option label
func (sc Select) OptionLabel(value string) Select {
	return sc.set("optionLabel", value)
}

// OptionValue sets the option value
func (sc Select) OptionValue(value string) Select {
	return sc.set("optionValue", value)
}

// Outline enables or disables outline button style
func (sc Select) Outline(value bool) Select {
	return sc.set("outline", value)
}

// Overlay enables or disables Overlay
func (sc Select) Overlay(value bool) Select {
	return sc.set("overlay", value)
}

// Size sets the control size
func (sc Select) Size(value string) Select {
	return sc.set("size", value)
}

// ValueField sets the value field name
func (sc Select) ValueField(value string) Select {
	return sc.set("valueField", value)
}

// Validation sets the validation rules
func (sc Select) Validation(value string) Select {
	return sc.set("validation", value)
}

// NoResultsText sets the text when no result found
func (s Select) NoResultsText(value string) Select {
	return s.set("noResultsText", value)
}
