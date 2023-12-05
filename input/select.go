package input

import (
	"github.com/manifoldco/promptui"
	"gitlab.com/tozd/go/errors"
)

type SelectArg func(o *wselect)

type wselect struct {
	label string
	items []string
	size  int
}

func (s *wselect) getSelectPrompt() promptui.Select {
	return promptui.Select{
		Label: s.label,
		Items: s.items,
		Size:  s.size,
	}
}

// NewSelectArg crea una nueva instancia de wselect
func Select(opts ...SelectArg) (int, string, errors.E) {
	s := &wselect{}
	for _, opt := range opts {
		opt(s)
	}

	sel := s.getSelectPrompt()
	idx, result, err := sel.Run()
	if err != nil {
		return 0, "", errors.Wrap(err, "input.Select")
	}

	return idx, result, nil
}

// WithLabel asigna label a wselect
func WithSelectLabel(label string) SelectArg {
	return func(s *wselect) {
		s.label = label
	}
}

// WithItems asigna items a wselect
func WithItems(items []string) SelectArg {
	return func(s *wselect) {
		s.items = items
	}
}

// WithSize asigna size a wselect
func WithSize(size int) SelectArg {
	return func(s *wselect) {
		s.size = size
	}
}
