package util

import (
	"unicode"
	"unicode/utf8"
)

func WordWrap(text string, lineWidth int) string { // {{{
	wrap := make([]byte, 0, len(text)+2*len(text)/lineWidth)
	eoLine := lineWidth
	inWord := false
	for i, j := 0, 0; ; {
		r, size := utf8.DecodeRuneInString(text[i:])
		if size == 0 && r == utf8.RuneError {
			r = ' '
		}
		if unicode.IsSpace(r) {
			if inWord {
				if i >= eoLine {
					wrap = append(wrap, '\n')
					eoLine = len(wrap) + lineWidth
				} else if len(wrap) > 0 {
					wrap = append(wrap, ' ')
				}
				wrap = append(wrap, text[j:i]...)
			}
			inWord = false
		} else if !inWord {
			inWord = true
			j = i
		}
		if size == 0 && r == ' ' {
			break
		}
		i += size
	}
	return string(wrap)
}  // }}}

// TernaryIf es una función que devuelve un valor dependiendo de una condición, similar al operador ternario `?:`
// para que funcione `T` debe ser del mismo tipo de dato.
// Esta función se usa con frecuencia como atajo para la instrucción if.
//
// Parámetros:
// - `condition`: Una expresión que se evalúa como true o false.
// - `trueVal`, `falseVal`: Expresion.o con valores de algún tipo.
// Resultado:
// - Si `condition` es true, el operador retorna el valor de `trueVal`; de lo contrario, devuelve el valor de `falseVal`.
func TernaryIf[T any](condition bool, trueVal, falseVal T) T {  // {{{
	if condition {
		return trueVal
	}

	return falseVal
}  // }}}
