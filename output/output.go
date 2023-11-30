package output

import (
	"fmt"
	"regexp"
	"strings"
)

var reColor *regexp.Regexp

func init() {
	reColor = regexp.MustCompile(`(<[^>]+>)`)
}

// Sprintf Formatea y retorna un string de igual forma que `fmt.Sprintf`, pero con colores a traves de etiquetas de
// colores encerrados en `<` y `>` y finalizando con la etiqueta `<reset>` o `</>`
func Sprintf(format string, a ...any) string { // {{{
	txt := format
	out := reColor.ReplaceAllStringFunc(
		txt,
		func(s string) string {
			if IsColor() {
				lowerKey := strings.ToLower(strings.Trim(s, " <>"))
				if value, ok := styleMap[lowerKey]; ok {
					return value.String()
				}
			}

			return ""
		},
	)

	return fmt.Sprintf(out, a...) + Reset
} // }}}
