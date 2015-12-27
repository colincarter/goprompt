package vcs

import "os"

type Vcs interface {
	Check()
	Name() string
	Path() string
	Found() bool
	Branch() string
	Modifications() string
	NewFiles() string
}

func vcsDir(folder string) bool {
	_, err := os.Stat(folder)
	return err == nil
}

// Start in folder and walk backwards
// until / is found
func findFolder(folder string) (foundFolder string, found bool) {
	foundFolder = ""
	found = false

	initialWd, err := os.Getwd()
	if err != nil {
		return "", false
	}

	// If we're already root then finish
	if initialWd == "/" {
		if vcsDir(folder) {
			return "/", true
		}

		_, err := os.Stat(folder)
		if err == nil {
			return "/", true
		}

		return "", false

	}

	for {
		_, err := os.Stat(folder)
		if err != nil {
			os.Chdir("..")

			wd, err := os.Getwd()
			if err != nil {
				found = false
			}

			if wd == "/" {
				break
			}

		} else {
			found = true
			break
		}
	}

	foundFolder, err = os.Getwd()

	os.Chdir(initialWd) // Change back to where we were

	return foundFolder, found
}
