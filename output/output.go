package output

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/wfrodriguez/console"
)

var reColor *regexp.Regexp

func init() {
	reColor = regexp.MustCompile(`(<[^>]+>)`)

	reset := NewStyle(WithForeground(creset))
	AddStyle("reset", reset)
	AddStyle("/", reset)

	AddStyle("title", NewStyle(WithForeground(New256Color(69, Foreground))))
	AddStyle("subtitle", NewStyle(WithForeground(New256Color(60, Foreground))))
	AddStyle("note", NewStyle(WithForeground(New256Color(256, Foreground))))
	AddStyle("caution", NewStyle(WithForeground(New256Color(230, Foreground))))
	AddStyle("success", NewStyle(WithForeground(New256Color(41, Foreground))))
	AddStyle("info", NewStyle(WithForeground(New256Color(39, Foreground))))
	AddStyle("warning", NewStyle(WithForeground(New256Color(220, Foreground))))
	AddStyle("error", NewStyle(WithForeground(New256Color(196, Foreground))))
}

// Sprintf Formatea y retorna un string de igual forma que `fmt.Sprintf`, pero con colores a traves de etiquetas de
// colores encerrados en `<` y `>` y finalizando con la etiqueta `<reset>` o `</>`
func Sprintf(format string, a ...any) string {
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

	return fmt.Sprintf(out, a...)
}

// PrintSection Imprime una sección de texto en la consola
// con una línea de guiones
func PrintSection(styleName, title string, content ...string) {
	width, _ := console.GetConsoleSize()
	sb := strings.Builder{}
	line := strings.Repeat("─", width-(utf8.RuneCountInString(title)+3))
	sb.WriteString(fmt.Sprintf("<%s>┌┤%s├%s\n", styleName, title, line))
	for _, row := range content {
		sb.WriteString(fmt.Sprintf("│%s\n", row))
	}
	sb.WriteString("└───</>")
	fmt.Println(Sprintf(sb.String()))
}

// Title crea una cabecera estilo Setext de primer nivel
func Title(title string) {
	l := utf8.RuneCountInString(title)
	fmt.Println(Sprintf("<title>%s\n%s</>", title, strings.Repeat("=", l)))
}

// Subtitle crea una cabecera estilo Setext de segundo nivel
func Subtitle(title string) {
	l := utf8.RuneCountInString(title)
	fmt.Println(Sprintf("<subtitle>%s\n%s</>", title, strings.Repeat("-", l)))
}

// Note crea un bloque de texto de tipo nota
func Note(content ...string) {
	PrintSection("note", "Nota", content...)
}

// Caution crea un bloque de texto de tipo advertencia
func Caution(content ...string) {
	PrintSection("caution", "CAUTION!", content...)
}

// Success crea un bloque de texto de tipo éxito
func Success(content ...string) {
	PrintSection("success", "SUCCESS", content...)
}

// Info crea un bloque de texto de tipo información
func Info(content ...string) {
	PrintSection("info", "Información", content...)
}

// Warning crea un bloque de texto de tipo advertencia
func Warning(content ...string) {
	PrintSection("warning", "WARNING!!", content...)
}

// Error crea un bloque de texto de tipo error
func Error(content ...string) {
	PrintSection("error", "ERROR!!!", content...)
}
