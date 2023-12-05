package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/chzyer/readline"
	"github.com/wfrodriguez/console/output"
	"gitlab.com/tozd/go/errors"
)

const None = -1

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func drawRuler(s int) string {
	sbn := strings.Builder{}
	sbl := strings.Builder{}
	for i := 0; i < s; i = i + 10 {
		sbn.WriteString(fmt.Sprintf("%10d", i+10))
		sbl.WriteString("┄┄┄┄┄┄┄┄┄┼")
	}

	strn := sbn.String()
	strn = replaceAtIndex(strn, '0', 0)
	strl := sbl.String()
	strl = replaceAtIndex(strl, '├', 0)
	strl = replaceAtIndex(strl, '┤', utf8.RuneCountInString(strl)-1)

	return fmt.Sprintf("╔═══╗%s\n╚═══╝%s", strn, strl)
}

func TextArea(label string, ruleWidth int) ([]string, errors.E) {
	nline := 1
	var input []string
	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 fmt.Sprintf("┊%3d ", nline),
		HistorySearchFold:      false,
		DisableAutoSaveHistory: true,
	})
	if err != nil {
		return []string{}, errors.Wrap(err, "error al crear el lector de linea")
	}
	defer rl.Close()
	fmt.Println(output.Sprintf("<input>%s</>(%s.%[2]s para finalizar)", label, "<ENTER>"))
	if ruleWidth > 0 {
		fmt.Println(output.Sprintf("<extra>%s</>", drawRuler(ruleWidth)))
	}

	hasInput := true
	for hasInput {
		line, err := rl.Readline()
		if err != nil {
			return []string{}, errors.Wrap(err, "error al leer la linea")
		}
		line = strings.TrimSpace(line)
		if line == "." {
			hasInput = false
			continue
		}
		input = append(input, line)
		nline++
		rl.SetPrompt(fmt.Sprintf("┊%3d ", nline))
	}
	fmt.Println()

	return input, nil

}

func init() {
	output.AddStyle("input", output.NewStyle(output.WithForeground(output.New256Color(28, output.Foreground))))
	output.AddStyle("def", output.NewStyle(output.WithForeground(output.New256Color(178, output.Foreground))))
	output.AddStyle("extra", output.NewStyle(output.WithForeground(output.New256Color(122, output.Foreground))))
}

func Confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)

	for ; tries > 0; tries-- {
		fmt.Print(output.Sprintf("<input>%s </>[<def>y/n</>]: ", s))

		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	return false
}
