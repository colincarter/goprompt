package vcs

type Hg struct {
	name   string
	found  bool
	folder string
}

func (h *Hg) Check() {
	folder, found := findFolder(".hg")

	h.name = "mercurial"

	if found {
		h.found = true
		h.folder = folder
	}
}

func (h *Hg) Name() string {
	return h.name
}

func (h *Hg) Path() string {
	return h.folder
}

func (h *Hg) Found() bool {
	return h.found
}

func (h *Hg) Branch() string {
	return ""
}

func (h *Hg) Modifications() string {
	return ""
}

func (h *Hg) NewFiles() string {
	return ""
}
