package comp

type Button = VanillaAction

func NewButton() Button {
	return NewVanillaAction().set("type", "button")
}
