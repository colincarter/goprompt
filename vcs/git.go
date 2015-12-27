package vcs

import (
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
)

type Git struct {
	name   string
	found  bool
	folder string
}

func (g *Git) Check() {
	folder, found := findFolder(".git")

	g.name = "git"

	if found {
		g.found = true
		g.folder = folder
	}
}

func (g *Git) Name() string {
	return g.name
}

func (g *Git) Path() string {
	return g.folder
}

func (g *Git) Found() bool {
	return g.found
}

func (g *Git) Branch() string {
	filename := path.Join(g.folder, ".git", "HEAD")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	parts := strings.Split(string(data), "/")

	return strings.TrimSpace(parts[len(parts)-1])
}

func (g *Git) Modifications() string {
	cmd := exec.Command("git", "diff", "--no-ext-diff", "--quiet", "--exit-code")
	err := cmd.Run()
	if err != nil {
		return "+"
	}

	return ""
}

func (g *Git) NewFiles() string {
	cmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	out, err := cmd.Output()
	if err == nil && len(out) == 0 {
		return ""
	}

	return "?"
}
