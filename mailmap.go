package hercules

import (
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func ParseMailmap(contents string) map[string]object.Signature {
  mm := map[string]object.Signature{}
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.LastIndex(line, ">") != len(line) - 1 {
			continue
		}
		ltp := strings.LastIndex(line, "<")
		fromEmail := line[ltp + 1:len(line) - 1]
		line = strings.TrimSpace(line[:ltp])
		gtp := strings.LastIndex(line, ">")
		fromName := ""
		if gtp != len(line) - 1 {
			fromName = strings.TrimSpace(line[gtp + 1:])
		}
		toEmail := ""
		if gtp > 0 {
			ltp = strings.LastIndex(line, "<")
			toEmail = line[ltp + 1:gtp]
			line = strings.TrimSpace(line[:ltp])
		}
		toName := line
		if fromEmail != "" {
			mm[fromEmail] = object.Signature{Name: toName, Email: toEmail}
		}
		if fromName != "" {
			mm[fromName] = object.Signature{Name: toName, Email: toEmail}
		}
	}
	return mm
}
