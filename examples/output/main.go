package main

import (
	"github.com/wfrodriguez/console/output"
)

const dump = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus accumsan nisi non libero cursus porttitor. Quisque euismod dictum magna ut venenatis. Nulla dignissim elit accumsan sodales viverra. Vestibulum ante ipsum primis in faucibus orci"

func main() { // {{{
	var arrDump = []string{
		dump,
		dump,
	}

	output.Title("Console output test")
	output.Subtitle("Some samples about output")
	output.Error(dump)
	output.Info(dump)
	output.Warning(dump)
	output.Note(arrDump...)
	output.Caution(dump, dump, dump)
	output.Success(dump)
} // }}}
