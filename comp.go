package amisgo

import (
	"github.com/zrcoder/amisgo/comp"
	icomp "github.com/zrcoder/amisgo/internal/comp"
)

func (a *App) Action() comp.Action                           { return icomp.NewAction(a.mux) }
func (a *App) Alert() comp.Alert                             { return icomp.NewAlert() }
func (a *App) AnchorNav() comp.AnchorNav                     { return icomp.NewAnchorNav() }
func (a *App) App() comp.App                                 { return icomp.NewApp() }
func (a *App) Audio() comp.Audio                             { return icomp.NewAudio() }
func (a *App) Avatar() comp.Avatar                           { return icomp.NewAvatar() }
func (a *App) Amis() comp.Amis                               { return icomp.NewAmis() }
func (a *App) Barcode() comp.Barcode                         { return icomp.NewBarcode() }
func (a *App) Breadcrumb() comp.Breadcrumb                   { return icomp.NewBreadcrumb() }
func (a *App) Button() comp.Action                           { return icomp.NewButton(a.mux) }
func (a *App) ButtonGroup() comp.ButtonGroup                 { return icomp.NewButtonGroup() }
func (a *App) ButtonGroupSelect() comp.Select                { return icomp.NewButtonGroupSelect() }
func (a *App) ButtonToolbar() comp.ButtonToolbar             { return icomp.NewButtonToolbar() }
func (a *App) Calendar() comp.Calendar                       { return icomp.NewCalendar() }
func (a *App) Card() comp.Card                               { return icomp.NewCard() }
func (a *App) Cards() comp.Cards                             { return icomp.NewCards() }
func (a *App) Carousel() comp.Carousel                       { return icomp.NewCarousel() }
func (a *App) ChainedSelect() comp.ChainedSelect             { return icomp.NewChainedSelect() }
func (a *App) Chart() comp.Chart                             { return icomp.NewChart(a.mux) }
func (a *App) ChartRadios() comp.ChartRadios                 { return icomp.NewChartRadios() }
func (a *App) Checkbox() comp.Checkbox                       { return icomp.NewCheckbox() }
func (a *App) Checkboxes() comp.Checkboxes                   { return icomp.NewCheckboxes() }
func (a *App) Code() comp.Code                               { return icomp.NewCode() }
func (a *App) Collapse() comp.Collapse                       { return icomp.NewCollapse() }
func (a *App) CollapseGroup() comp.CollapseGroup             { return icomp.NewCollapseGroup() }
func (a *App) Color() comp.Color                             { return icomp.NewColor() }
func (a *App) Combo() comp.Combo                             { return icomp.NewCombo() }
func (a *App) ConditionBuilder() comp.ConditionBuilder       { return icomp.NewConditionBuilder() }
func (a *App) ConditionGroupValue() comp.ConditionGroupValue { return icomp.NewConditionGroupValue() }
func (a *App) Container() comp.Container                     { return icomp.NewContainer() }
func (a *App) Control() comp.Control                         { return icomp.NewControl() }
func (a *App) CopyButton() comp.CopyButton                   { return icomp.NewCopyButton() }
func (a *App) Crud() comp.Crud                               { return icomp.NewCrud(a.mux) }
func (a *App) CrudCards() comp.Crud                          { return icomp.NewCrudCards(a.mux) }
func (a *App) CrudList() comp.Crud                           { return icomp.NewCrudList(a.mux) }
func (a *App) CrudTable() comp.Crud                          { return icomp.NewCrudTable(a.mux) }
func (a *App) Custom() comp.Custom                           { return icomp.NewCustom() }
func (a *App) Date() comp.Date                               { return icomp.NewDate() }
func (a *App) DateRange() comp.DateRange                     { return icomp.NewDateRange() }
func (a *App) Dialog() comp.Dialog                           { return icomp.NewDialog() }
func (a *App) DialogAction() comp.DialogAction               { return icomp.NewDialogAction() }
func (a *App) DiffEditor() comp.DiffEditor                   { return icomp.NewDiffEditor() }
func (a *App) Divider() comp.Divider                         { return icomp.NewDivider() }
func (a *App) Drawer() comp.Drawer                           { return icomp.NewDrawer() }
func (a *App) DrawerAction() comp.DrawerAction               { return icomp.NewDrawerAction() }
func (a *App) DropdownButton() comp.DropdownButton           { return icomp.NewDropdownButton() }
func (a *App) Each() comp.Each                               { return icomp.NewEach() }
func (a *App) Editor() comp.Editor                           { return icomp.NewEditor() }
func (a *App) EmailAction() comp.EmailAction                 { return icomp.NewEmailAction() }
func (a *App) FieldSet() comp.FieldSet                       { return icomp.NewFieldSet() }
func (a *App) Flex() comp.Flex                               { return icomp.NewFlex() }
func (a *App) Form() comp.Form                               { return icomp.NewForm(a.mux) }
func (a *App) Formula() comp.Formula                         { return icomp.NewFormula() }
func (a *App) Grid() comp.Grid                               { return icomp.NewGrid() }
func (a *App) Grid2D() comp.Grid2d                           { return icomp.NewGrid2D() }
func (a *App) GridNav() comp.GridNav                         { return icomp.NewGridNav() }
func (a *App) Group() comp.Group                             { return icomp.NewGroup() }
func (a *App) HBox() comp.HBox                               { return icomp.NewHBox() }
func (a *App) Hidden() comp.Hidden                           { return icomp.NewHidden() }
func (a *App) Html() comp.Tpl                                { return icomp.NewHtml() }
func (a *App) Icon() comp.Icon                               { return icomp.NewIcon() }
func (a *App) IconPicker() comp.IconPicker                   { return icomp.NewIconPicker() }
func (a *App) Iframe() comp.Iframe                           { return icomp.NewIframe() }
func (a *App) Image() comp.Image                             { return icomp.NewImage() }
func (a *App) Images() comp.Images                           { return icomp.NewImages() }
func (a *App) InputArray() comp.InputArray                   { return icomp.NewInputArray() }
func (a *App) InputCity() comp.InputCity                     { return icomp.NewInputCity() }
func (a *App) InputColor() comp.InputColor                   { return icomp.NewInputColor() }
func (a *App) InputDate() comp.InputDate                     { return icomp.NewInputDate() }
func (a *App) InputDateRange() comp.InputDateRange           { return icomp.NewInputDateRange() }
func (a *App) InputDatetime() comp.InputDatetime             { return icomp.NewInputDatetime() }
func (a *App) InputDatetimeRange() comp.InputDatetimeRange   { return icomp.NewInputDatetimeRange() }
func (a *App) InputEmail() comp.InputText                    { return icomp.NewInputEmail() }
func (a *App) InputExcel() comp.InputExcel                   { return icomp.NewInputExcel() }
func (a *App) InputFile() comp.InputFile                     { return icomp.NewInputFile(a.mux) }
func (a *App) InputGroup() comp.InputGroup                   { return icomp.NewInputGroup() }
func (a *App) InputImage() comp.InputImage                   { return icomp.NewInputImage(a.mux) }
func (a *App) InputKV() comp.InputKV                         { return icomp.NewInputKV() }
func (a *App) InputKVS() comp.InputKVS                       { return icomp.NewInputKVS() }
func (a *App) InputMonth() comp.InputMonth                   { return icomp.NewInputMonth() }
func (a *App) InputMonthRange() comp.InputMonthRange         { return icomp.NewInputMonthRange() }
func (a *App) InputNumber() comp.InputNumber                 { return icomp.NewInputNumber() }
func (a *App) InputPassword() comp.InputText                 { return icomp.NewInputPassword() }
func (a *App) InputQuarter() comp.InputQuarter               { return icomp.NewInputQuarter() }
func (a *App) InputQuarterRange() comp.InputQuarterRange     { return icomp.NewInputQuarterRange() }
func (a *App) InputRange() comp.InputRange                   { return icomp.NewInputRange() }
func (a *App) InputRating() comp.InputRating                 { return icomp.NewInputRating() }
func (a *App) InputRepeat() comp.InputRepeat                 { return icomp.NewInputRepeat() }
func (a *App) InputRichText() comp.InputRichText             { return icomp.NewInputRichText() }
func (a *App) InputSignature() comp.InputSignature           { return icomp.NewInputSignature() }
func (a *App) InputSubForm() comp.InputSubForm               { return icomp.NewInputSubForm() }
func (a *App) InputTable() comp.InputTable                   { return icomp.NewInputTable() }
func (a *App) InputTag() comp.InputTag                       { return icomp.NewInputTag() }
func (a *App) InputText() comp.InputText                     { return icomp.NewInputText() }
func (a *App) InputTime() comp.InputTime                     { return icomp.NewInputTime() }
func (a *App) InputTimeRange() comp.InputTimeRange           { return icomp.NewInputTimeRange() }
func (a *App) InputTree() comp.InputTree                     { return icomp.NewInputTree() }
func (a *App) InputUrl() comp.InputText                      { return icomp.NewInputUrl() }
func (a *App) InputYear() comp.InputYear                     { return icomp.NewInputYear() }
func (a *App) InputYearRange() comp.InputYearRange           { return icomp.NewInputYearRange() }
func (a *App) Json() comp.Json                               { return icomp.NewJson() }
func (a *App) JsonEditor() comp.Editor                       { return icomp.NewJsonEditor() }
func (a *App) JsonSchemaEditor() comp.JsonSchemaEditor       { return icomp.NewJsonSchemaEditor() }
func (a *App) Link() comp.Link                               { return icomp.NewLink() }
func (a *App) LinkAction() comp.LinkAction                   { return icomp.NewLinkAction() }
func (a *App) List() comp.List                               { return icomp.NewList() }
func (a *App) ListSelect() comp.ListSelect                   { return icomp.NewListSelect() }
func (a *App) LocationPicker() comp.LocationPicker           { return icomp.NewLocationPicker() }
func (a *App) Log() comp.Log                                 { return icomp.NewLog() }
func (a *App) Mapping() comp.Mapping                         { return icomp.NewMapping() }
func (a *App) Markdown() comp.Markdown                       { return icomp.NewMarkdown() }
func (a *App) MatrixCheckboxes() comp.MatrixCheckboxes       { return icomp.NewMatrixCheckboxes() }
func (a *App) MultilineText() comp.MultilineText             { return icomp.NewMultilineText() }
func (a *App) Nav() comp.Nav                                 { return icomp.NewNav() }
func (a *App) NestedSelect() comp.NestedSelect               { return icomp.NewNestedSelect() }
func (a *App) NewTimeline() comp.Timeline                    { return icomp.NewTimeline() }
func (a *App) Operation() comp.Operation                     { return icomp.NewOperation() }
func (a *App) OtherAction() comp.OtherAction                 { return icomp.NewOtherAction() }
func (a *App) Page() comp.Page                               { return icomp.NewPage(a.mux) }
func (a *App) Pagination() comp.Pagination                   { return icomp.NewPagination() }
func (a *App) PaginationWrapper() comp.PaginationWrapper     { return icomp.NewPaginationWrapper() }
func (a *App) Panel() comp.Panel                             { return icomp.NewPanel() }
func (a *App) Password() comp.Password                       { return icomp.NewPassword() }
func (a *App) Picker() comp.Picker                           { return icomp.NewPicker() }
func (a *App) Plain() comp.Plain                             { return icomp.NewPlain() }
func (a *App) Portlet() comp.Portlet                         { return icomp.NewPortlet() }
func (a *App) Progress() comp.Progress                       { return icomp.NewProgress() }
func (a *App) Property() comp.Property                       { return icomp.NewProperty() }
func (a *App) QRCode() comp.QRCode                           { return icomp.NewQRCode() }
func (a *App) Radio() comp.Radio                             { return icomp.NewRadio() }
func (a *App) Radios() comp.Radios                           { return icomp.NewRadios() }
func (a *App) ReloadAction() comp.ReloadAction               { return icomp.NewReloadAction() }
func (a *App) Remark() comp.Remark                           { return icomp.NewRemark() }
func (a *App) SearchBox() comp.SearchBox                     { return icomp.NewSearchBox() }
func (a *App) Select() comp.Select                           { return icomp.NewSelect() }
func (a *App) Service() comp.Service                         { return icomp.NewService(a.mux) }
func (a *App) SparkLine() comp.SparkLine                     { return icomp.NewSparkLine() }
func (a *App) Spinner() comp.Spinner                         { return icomp.NewSpinner() }
func (a *App) Static() comp.Static                           { return icomp.NewStatic() }
func (a *App) Status() comp.Status                           { return icomp.NewStatus() }
func (a *App) Steps() comp.Steps                             { return icomp.NewSteps() }
func (a *App) SubmitAction() comp.Action                     { return icomp.NewSubmitAction(a.mux) }
func (a *App) SvgIcon() comp.SvgIcon                         { return icomp.NewSvgIcon() }
func (a *App) Switch() comp.Switch                           { return icomp.NewSwitch() }
func (a *App) SwitchContainer() comp.SwitchContainer         { return icomp.NewSwitchContainer() }
func (a *App) Table() comp.Table                             { return icomp.NewTable() }
func (a *App) Table2() comp.Table                            { return icomp.NewTable2() }
func (a *App) TableView() comp.TableView                     { return icomp.NewTableView() }
func (a *App) Tabs() comp.Tabs                               { return icomp.NewTabs() }
func (a *App) TabsTransfer() comp.TabsTransfer               { return icomp.NewTabsTransfer() }
func (a *App) TabsTransferPicker() comp.TabsTransferPicker   { return icomp.NewTabsTransferPicker() }
func (a *App) Tag() comp.Tag                                 { return icomp.NewTag() }
func (a *App) Tasks() comp.Tasks                             { return icomp.NewTasks() }
func (a *App) Text() comp.Text                               { return icomp.NewText() }
func (a *App) Textarea() comp.Textarea                       { return icomp.NewTextarea() }
func (a *App) ToastAction() comp.ToastAction                 { return icomp.NewToastAction() }
func (a *App) TooltipWrapper() comp.TooltipWrapper           { return icomp.NewTooltipWrapper() }
func (a *App) Tpl() comp.Tpl                                 { return icomp.NewTpl() }
func (a *App) Transfer() comp.Transfer                       { return icomp.NewTransfer() }
func (a *App) TransferPicker() comp.TransferPicker           { return icomp.NewTransferPicker() }
func (a *App) TreeSelect() comp.TreeSelect                   { return icomp.NewTreeSelect() }
func (a *App) UUID() comp.UUID                               { return icomp.NewUUID() }
func (a *App) UrlAction() comp.UrlAction                     { return icomp.NewUrlAction() }
func (a *App) UsersSelect() comp.UserSelect                  { return icomp.NewUsersSelect() }
func (a *App) VBox() comp.VBox                               { return icomp.NewVBox() }
func (a *App) Video() comp.Video                             { return icomp.NewVideo() }
func (a *App) WangEditor() comp.WangEditor                   { return icomp.NewWangEditor() }
func (a *App) Watermark() comp.Watermark                     { return icomp.NewWatermark() }
func (a *App) WebComponent() comp.WebComponent               { return icomp.NewWebComponent() }
func (a *App) Wizard() comp.Wizard                           { return icomp.NewWizard() }
func (a *App) WizardStep() comp.WizardStep                   { return icomp.NewWizardStep() }
func (a *App) Words() comp.Words                             { return icomp.NewWords() }
func (a *App) Wrapper() comp.Wrapper                         { return icomp.NewWrapper() }
func (a *App) Number() comp.Number                           { return icomp.NewNumber() }
func (a *App) Shape() comp.Shape                             { return icomp.NewShape() }
func (a *App) ThemeSelect() comp.Select                      { return icomp.NewSelect().Themes(a.mux, a.Conf.Templ) }
func (a *App) LocaleSelect() comp.Select                     { return icomp.NewSelect().Locales(a.mux, a.Conf.Templ) }
func (a *App) ThemeButtonGroupSelect() comp.Select {
	return icomp.NewButtonGroupSelect().Themes(a.mux, a.Conf.Templ)
}

func (a *App) LocaleButtonGroupSelect() comp.Select {
	return icomp.NewButtonGroupSelect().Locales(a.mux, a.Conf.Templ)
}

func (a *App) AnchorNavSection() comp.AnchorNavSection     { return icomp.NewAnchorNavSection() }
func (a *App) AutoFillHeight() comp.AutoFillHeight         { return icomp.NewAutoFillHeight() }
func (a *App) AutoGenerateFilter() comp.AutoGenerateFilter { return icomp.NewAutoGenerateFilter() }
func (a *App) Badge() comp.Badge                           { return icomp.NewBadge() }
func (a *App) ComboCondition() comp.ComboCondition         { return icomp.NewComboCondition() }
func (a *App) Component() comp.Component                   { return icomp.NewComponent() }
func (a *App) Expandable() comp.Expandable                 { return icomp.NewExpandable() }
func (a *App) FeedbackDialog() comp.FeedbackDialog         { return icomp.NewFeedbackDialog() }
func (a *App) GridColumn() comp.GridColumn                 { return icomp.NewGridColumn() }
func (a *App) HBoxColumn() comp.HBoxColumn                 { return icomp.NewHBoxColumn() }
func (a *App) IconChecked() comp.IconChecked               { return icomp.NewIconChecked() }
func (a *App) IconItem() comp.IconItem                     { return icomp.NewIconItem() }
func (a *App) ImageToolbarAction() comp.ImageToolbarAction { return icomp.NewImageToolbarAction() }
func (a *App) NavItem() comp.NavItem                       { return icomp.NewNavItem() }
func (a *App) NavOverflow() comp.NavOverflow               { return icomp.NewNavOverflow() }
func (a *App) NewTimelineItem() comp.TimelineItem          { return icomp.NewTimelineItem() }
func (a *App) PortletTab() comp.PortletTab                 { return icomp.NewPortletTab() }
func (a *App) QRCodeImageSettings() comp.QRCodeImageSettings {
	return icomp.NewQRCodeImageSettings()
}
func (a *App) RowSelection() comp.RowSelection { return icomp.NewRowSelection() }
func (a *App) RowSelectionOptions() comp.RowSelectionOptions {
	return icomp.NewRowSelectionOptions()
}
func (a *App) SchemaApi() comp.SchemaApi           { return icomp.NewSchemaApi() }
func (a *App) SchemaCopyable() comp.SchemaCopyable { return icomp.NewSchemaCopyable() }
func (a *App) SchemaMessage() comp.SchemaMessage   { return icomp.NewSchemaMessage() }
func (a *App) SchemaPopOver() comp.SchemaPopOver   { return icomp.NewSchemaPopOver() }
func (a *App) State() comp.State                   { return icomp.NewState() }
func (a *App) Step() comp.Step                     { return icomp.NewStep() }
func (a *App) Tr() comp.Tr                         { return icomp.NewTr() }
func (a *App) Td() comp.Td                         { return icomp.NewTd() }
func (a *App) Tcol() comp.Tcol                     { return icomp.NewTCol() }
func (a *App) TableColumn() comp.TableColumn       { return icomp.NewTableColumn() }
func (a *App) Toast() comp.Toast                   { return icomp.NewToast() }
func (a *App) ToastItem() comp.ToastItem           { return icomp.NewToastItem() }
func (a *App) ChartConfig() comp.ChartCfg          { return icomp.NewChartConfig() }
func (a *App) ChartSeries() comp.ChartSeri         { return icomp.NewChartSeries() }
func (a *App) ChartAxis() comp.ChartAxis           { return icomp.NewChartAxis() }
func (a *App) NewListItem() comp.ListItem          { return icomp.NewListItem() }
func (a *App) Api() comp.Api                       { return icomp.NewApi() }
func (a *App) BreadcrumbItem() comp.BreadcrumbItem { return icomp.NewBreadcrumbItem() }
func (a *App) Column(component ...any) comp.Column { return icomp.NewColumn(component...) }
func (a *App) Event() comp.Event                   { return icomp.NewEvent() }
func (a *App) EventActions(actions ...comp.EventAction) icomp.EventActions {
	return icomp.NewEventActions(actions...)
}
func (a *App) EventAction() comp.EventAction      { return icomp.NewEventAction() }
func (a *App) EventActionToast() comp.EventAction { return icomp.NewEventActionToast() }
func (a *App) EventActionDrawer(drawer ...comp.Drawer) icomp.EventAction {
	return icomp.NewEventActionDrawer(drawer...)
}

func (a *App) EventActionDialog(dialog ...comp.Dialog) icomp.EventAction {
	return icomp.NewEventActionDialog(dialog...)
}
func (a *App) EventActionArgs() comp.EventActionArgs   { return icomp.NewEventActionArgs() }
func (a *App) GridItem(component ...any) comp.GridItem { return icomp.NewGridItem(component...) }
func (a *App) Horizontal() comp.Horizontal             { return icomp.NewHorizontal() }
func (a *App) ListBodyField() comp.ListBodyField       { return icomp.NewListBodyField() }
func (a *App) NavLink() comp.NavLink                   { return icomp.NewNavLink() }
func (a *App) Option() comp.Option                     { return icomp.NewOption() }
func (a *App) Options() comp.Options                   { return icomp.NewOptions() }
func (a *App) PageItem() comp.PageItem                 { return icomp.NewPageItem() }
func (a *App) PropertyItem() comp.PropertyItem         { return icomp.NewPropertyItem() }
func (a *App) PullRefresh() comp.PullRefresh           { return icomp.NewPullRefresh() }
func (a *App) Rule() comp.Rule                         { return icomp.NewRule() }
func (a *App) Tab() comp.Tab                           { return icomp.NewTab() }
func (a *App) CarouselOption() comp.CarouselOption     { return icomp.NewCarouselOption() }
