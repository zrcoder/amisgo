package theme

// Available themes
const (
	Default = ""
	Cxd     = "cxd"
	Antd    = "antd"
	Ang     = "ang"
	Dark    = "dark"
)

// Theme represents a theme
type Theme struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
