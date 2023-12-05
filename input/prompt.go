package input

import (
	"github.com/manifoldco/promptui"
	"gitlab.com/tozd/go/errors"
)

type PromptArg func(o *prompt)

type prompt struct {
	label        string
	defaultValue string
	validate     func(string) error
	mask         rune
	vimMode      bool
}

func (p prompt) getPrompt() promptui.Prompt {
	return promptui.Prompt{
		Label:     p.label,
		Default:   p.defaultValue,
		Validate:  p.validate,
		Mask:      p.mask,
		IsVimMode: p.vimMode,
	}
}

// NewPromptArg crea una nueva instancia de prompt (Input text, Password input) y devuelve el resultado
func Ask(opts ...PromptArg) (string, errors.E) {
	p := &prompt{}
	for _, opt := range opts {
		opt(p)
	}

	pmp := p.getPrompt()
	result, err := pmp.Run()
	if err != nil {
		return "", errors.Wrap(err, "input.Ask")
	}

	return result, nil
}

// WithLabel es el texto que aparece en la línea de comandos.
func WithLabel(label string) PromptArg {
	return func(p *prompt) {
		p.label = label
	}
}

// WithDefault es el valor inicial del prompt. Este valor se mostrará junto a la etiqueta de la consulta y el usuario podrá verlo
// o cambiarlo en función de las opciones.
func WithDefault(defaultValue string) PromptArg {
	return func(p *prompt) {
		p.defaultValue = defaultValue
	}
}

// WithValidate es una función opcional que se utiliza para validar el valor introducido.
func WithValidate(validate func(string) error) PromptArg {
	return func(p *prompt) {
		p.validate = validate
	}
}

// WithMask es una runa opcional que establece qué carácter mostrar en lugar de los caracteres introducidos. Esto permite ocultar
// información privada como contraseñas.
func WithMask(mask rune) PromptArg {
	return func(p *prompt) {
		p.mask = mask
	}
}

// WithVimMode habilita los movimientos tipo vi (hjkl) y la edición.
func WithVimMode() PromptArg {
	return func(p *prompt) {
		p.vimMode = true
	}
}
