package comp

import "github.com/zrcoder/amisgo/schema"

type CarouselOption schema.Schema

func NewCarouselOption() CarouselOption {
	return CarouselOption{}
}

func (co CarouselOption) set(key string, value any) CarouselOption {
	co[key] = value
	return co
}

func (co CarouselOption) Image(value string) CarouselOption {
	return co.set("image", value)
}

func (co CarouselOption) Href(value string) CarouselOption {
	return co.set("href", value)
}

func (co CarouselOption) ImageClassName(value string) CarouselOption {
	return co.set("imageClassName", value)
}

func (co CarouselOption) Title(value string) CarouselOption {
	return co.set("title", value)
}

func (co CarouselOption) TitleClassName(value string) CarouselOption {
	return co.set("titleClassName", value)
}

func (co CarouselOption) Description(value string) CarouselOption {
	return co.set("description", value)
}

func (co CarouselOption) DescriptionClassName(value string) CarouselOption {
	return co.set("descriptionClassName", value)
}

func (co CarouselOption) Html(value string) CarouselOption {
	return co.set("html", value)
}
