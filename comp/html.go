package comp

// Html creates a template with type set to "html".
func Html() tpl {
	return Tpl().set("type", "html")
}
