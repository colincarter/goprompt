package main

import (
	"flag"
	"fmt"

	"bytes"

	"github.com/colincarter/goprompt/vcs"
)

type vcsParts struct {
	found         bool
	name          string
	branch        string
	modifications string
	newFiles      string
}

func main() {
	outputFormat := flag.String("f", "[t:bnm]", "The format string")

	flag.Parse()

	vcses := []vcs.Vcs{new(vcs.Git), new(vcs.Hg)}
	parts := findVcs(vcses)

	if parts.found {
		fmt.Println(outputVcsDetails(parts, *outputFormat))
	}
}

func outputVcsDetails(parts *vcsParts, formatString string) string {
	var outFormatString bytes.Buffer

	for n := 0; n < len(formatString); n++ {
		switch formatString[n] {
		case 't':
			outFormatString.WriteString(parts.name)

		case 'b':
			outFormatString.WriteString(parts.branch)

		case 'n':
			outFormatString.WriteString(parts.newFiles)

		case 'm':
			outFormatString.WriteString(parts.modifications)

		default:
			outFormatString.WriteString(string(formatString[n]))
		}
	}

	return outFormatString.String()
}

func findVcs(vcses []vcs.Vcs) (parts *vcsParts) {
	for _, v := range vcses {
		v.Check()

		if v.Found() {
			return &vcsParts{
				found:         true,
				name:          v.Name(),
				branch:        v.Branch(),
				modifications: v.Modifications(),
				newFiles:      v.NewFiles(),
			}
		}
	}

	return &vcsParts{}
}
