package mvc

type (
	Page struct {
		Head Head
		Body Body
	}
	Head struct {
		Title       string
		StyleSheets []string
		Scripts     []string
	}
	Body struct {
		ViewPath string
	}
)

func (h *Head) AddStyleSheet(url string) {
	h.StyleSheets = append(h.StyleSheets, url)
}

func (h *Head) AddScript(url string) {
	h.Scripts = append(h.Scripts, url)
}
