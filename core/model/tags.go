package model

type (
	Tags interface {
		Each(func(string))
		Exists(tag string) (ok bool)
	}

	tags map[string]bool
)

func (tags tags) Each(f func(string)) {
	for t, _ := range tags {
		f(t)
	}
}

func (tags tags) Exists(tag string) bool {
	return tags[tag]
}

func (tags tags) del(tag string) bool {
	if tags[tag] {
		delete(tags, tag)
		return true
	}
	return false
}

func (tags tags) set(tag string) bool {
	if tags[tag] {
		return false
	}
	tags[tag] = true
	return true
}


