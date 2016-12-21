package gollapse

import "regexp"

type Gollapser struct {
	Match   *regexp.Regexp
	Replace string
}

func NewGollapser(match, replace string) Gollapser {
	m := regexp.MustCompile(match)
	return Gollapser{m, replace}
}

func (g *Gollapser) Collapse(path string) (string, bool) {
	if g.Match.MatchString(path) {
		return g.Match.ReplaceAllString(path, g.Replace), true
	}
	return "", false
}

type Gollapsers []Gollapser

func (gs Gollapsers) Collapse(path string) (string, bool) {
	for _, g := range gs {
		if replaced, ok := g.Collapse(path); ok {
			return replaced, ok
		}
	}
	return path, false
}
